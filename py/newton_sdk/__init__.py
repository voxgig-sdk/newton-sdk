# Newton SDK

from newton_sdk.utility.voxgig_struct import voxgig_struct as vs
from newton_sdk.core.utility_type import NewtonUtility
from newton_sdk.core.spec import NewtonSpec
from newton_sdk.core import helpers

# Load utility registration (populates Utility._registrar)
from newton_sdk.utility import register

# Load features
from newton_sdk.feature.base_feature import NewtonBaseFeature
from newton_sdk.features import _make_feature


class NewtonSDK:

    def __init__(self, options=None):
        self.mode = "live"
        self.features = []
        self.options = None

        utility = NewtonUtility()
        self._utility = utility

        from newton_sdk.config import make_config
        config = make_config()

        self._rootctx = utility.make_context({
            "client": self,
            "utility": utility,
            "config": config,
            "options": options if options is not None else {},
            "shared": {},
        }, None)

        self.options = utility.make_options(self._rootctx)

        if vs.getpath(self.options, "feature.test.active") is True:
            self.mode = "test"

        self._rootctx.options = self.options

        # Add features in the resolved order (make_options puts an explicit
        # list order first, else defaults to test-first). Ordering matters: the
        # `test` feature installs the base mock transport and the transport
        # features (retry/cache/netsim/proxy/ratelimit) wrap whatever is
        # current, so `test` must be added before them to sit at the base.
        feature_opts = helpers.to_map(vs.getprop(self.options, "feature"))
        if feature_opts is not None:
            featureorder = vs.getpath(self.options, "__derived__.featureorder")
            if isinstance(featureorder, list):
                for fname in featureorder:
                    fopts = helpers.to_map(feature_opts.get(fname))
                    if fopts is not None and fopts.get("active") is True:
                        utility.feature_add(self._rootctx, _make_feature(fname))

        # Add extension features.
        extend = vs.getprop(self.options, "extend")
        if isinstance(extend, list):
            for f in extend:
                if isinstance(f, dict) or (hasattr(f, "get_name") and callable(f.get_name)):
                    utility.feature_add(self._rootctx, f)

        # Initialize features.
        for f in self.features:
            utility.feature_init(self._rootctx, f)

        utility.feature_hook(self._rootctx, "PostConstruct")

        # #BuildFeatures

    def options_map(self):
        out = vs.clone(self.options)
        if isinstance(out, dict):
            return out
        return {}

    def get_utility(self):
        return NewtonUtility.copy(self._utility)

    def get_root_ctx(self):
        return self._rootctx

    def prepare(self, fetchargs=None):
        utility = self._utility

        if fetchargs is None:
            fetchargs = {}

        ctrl = helpers.to_map(vs.getprop(fetchargs, "ctrl"))
        if ctrl is None:
            ctrl = {}

        ctx = utility.make_context({
            "opname": "prepare",
            "ctrl": ctrl,
        }, self._rootctx)

        options = self.options

        path = vs.getprop(fetchargs, "path") or ""
        if not isinstance(path, str):
            path = ""

        method = vs.getprop(fetchargs, "method") or "GET"
        if not isinstance(method, str):
            method = "GET"

        params = helpers.to_map(vs.getprop(fetchargs, "params"))
        if params is None:
            params = {}
        query = helpers.to_map(vs.getprop(fetchargs, "query"))
        if query is None:
            query = {}

        headers = utility.prepare_headers(ctx)

        base = vs.getprop(options, "base") or ""
        if not isinstance(base, str):
            base = ""
        prefix = vs.getprop(options, "prefix") or ""
        if not isinstance(prefix, str):
            prefix = ""
        suffix = vs.getprop(options, "suffix") or ""
        if not isinstance(suffix, str):
            suffix = ""

        ctx.spec = NewtonSpec({
            "base": base,
            "prefix": prefix,
            "suffix": suffix,
            "path": path,
            "method": method,
            "params": params,
            "query": query,
            "headers": headers,
            "body": vs.getprop(fetchargs, "body"),
            "step": "start",
        })

        # Merge user-provided headers.
        uh = vs.getprop(fetchargs, "headers")
        if isinstance(uh, dict):
            for k, v in uh.items():
                ctx.spec.headers[k] = v

        _, err = utility.prepare_auth(ctx)
        if err is not None:
            raise err

        fetchdef, err = utility.make_fetch_def(ctx)
        if err is not None:
            raise err

        return fetchdef

    # Raw endpoint access is operator-controllable, like every entity op.
    # Blocking it means denying BOTH the 'direct' and 'graphql' tokens, since
    # either one reaches the same endpoint.
    def direct(self, fetchargs=None):
        if not self._op_allowed("direct"):
            return self._op_denied("direct")

        return self._raw_request(fetchargs)

    # Is this raw-access op permitted by the SDK's allow.op option?
    def _op_allowed(self, op):
        allow_op = vs.getpath(self.options, "allow.op")
        return isinstance(allow_op, str) and op in allow_op

    def _op_denied(self, op):
        allow_op = vs.getpath(self.options, "allow.op")
        return {
            "ok": False,
            "err": Exception(
                "NewtonSDK: " + op + ": operation not allowed by"
                ' SDK option allow.op value: "' + str(allow_op) + '"'),
        }

    # Ungated request path shared by direct and graphql, each of which checks
    # its own allow.op token first. Private, rather than a flag on fetchargs:
    # a caller-supplied marker would let anyone opt straight back out of the
    # gate by passing it.
    def _raw_request(self, fetchargs=None):
        utility = self._utility

        try:
            fetchdef = self.prepare(fetchargs)
        except Exception as err:
            # direct() is the raw-HTTP escape hatch: it never raises, it
            # returns a result object callers branch on via result["ok"].
            return {"ok": False, "err": err}

        if fetchargs is None:
            fetchargs = {}
        ctrl = helpers.to_map(vs.getprop(fetchargs, "ctrl"))
        if ctrl is None:
            ctrl = {}

        ctx = utility.make_context({
            "opname": "direct",
            "ctrl": ctrl,
        }, self._rootctx)

        url = fetchdef.get("url", "")
        fetched, fetch_err = utility.fetcher(ctx, url, fetchdef)

        if fetch_err is not None:
            return {"ok": False, "err": fetch_err}

        if fetched is None:
            return {
                "ok": False,
                "err": ctx.make_error("direct_no_response", "response: undefined"),
            }

        if isinstance(fetched, dict):
            status = helpers.to_int(vs.getprop(fetched, "status"))
            headers = vs.getprop(fetched, "headers") or {}

            # No-body responses (204, 304) and explicit zero content-length
            # must skip JSON parsing — calling json() on an empty body raises.
            content_length = None
            if isinstance(headers, dict):
                content_length = headers.get("content-length")
            no_body = status in (204, 304) or str(content_length) == "0"

            json_data = None
            if not no_body:
                jf = vs.getprop(fetched, "json")
                if callable(jf):
                    try:
                        json_data = jf()
                    except Exception:
                        # Non-JSON body (e.g. text/plain, text/html). Surface
                        # status + headers but leave data as None.
                        json_data = None

            return {
                "ok": status >= 200 and status < 300,
                "status": status,
                "headers": headers,
                "data": json_data,
            }

        return {
            "ok": False,
            "err": ctx.make_error("direct_invalid", "invalid response type"),
        }

    # Raw GraphQL access: the pressure valve that makes the generated
    # surface's deliberate omissions (per-call selection sets, typed filter
    # builders, batching, subscriptions) livable — the whole schema stays
    # reachable.
    #
    # Thin wrapper over the same prepare/fetch path direct uses, with the one
    # thing raw direct cannot do for GraphQL: a GraphQL failure rides HTTP 200
    # as a top-level `errors` array, so status alone would report a failed
    # query as ok.
    #
    # NOTE: like direct, this bypasses the feature pipeline — no retry,
    # ratelimit or paging features apply.
    def graphql(self, query, variables=None, ctrl=None):
        if not self._op_allowed("graphql"):
            return self._op_denied("graphql")

        res = self._raw_request({
            "method": "POST",
            "headers": {"content-type": "application/json"},
            "body": {"query": query, "variables": variables or {}},
            "ctrl": ctrl or {},
        })

        # Errors are read BEFORE any status check: a GraphQL parse or
        # validation failure comes back as HTTP 400 carrying the standard
        # { errors: [...] } body, and the raw path represents a non-2xx as
        # ok:False with no err — so returning early on status would discard
        # the server's own diagnostics, which are the only useful part of
        # that response.
        errors = vs.getpath(res, "data.errors")

        if isinstance(errors, list) and 0 < len(errors):
            first = errors[0] if isinstance(errors[0], dict) else {}
            msg = first.get("message") or "graphql error"
            res["ok"] = False
            res["err"] = Exception("NewtonSDK: graphql: " + str(msg))
            res["graphql"] = errors

        return res


    def Abs(self, data=None) -> "AbsEntity":
        """Entity factory: client.Abs().list() / client.Abs().load({"id": ...})."""
        from newton_sdk.entity.abs_entity import AbsEntity
        return AbsEntity(self, data)


    def Arcco(self, data=None) -> "ArccoEntity":
        """Entity factory: client.Arcco().list() / client.Arcco().load({"id": ...})."""
        from newton_sdk.entity.arcco_entity import ArccoEntity
        return ArccoEntity(self, data)


    def Arcsin(self, data=None) -> "ArcsinEntity":
        """Entity factory: client.Arcsin().list() / client.Arcsin().load({"id": ...})."""
        from newton_sdk.entity.arcsin_entity import ArcsinEntity
        return ArcsinEntity(self, data)


    def Arctan(self, data=None) -> "ArctanEntity":
        """Entity factory: client.Arctan().list() / client.Arctan().load({"id": ...})."""
        from newton_sdk.entity.arctan_entity import ArctanEntity
        return ArctanEntity(self, data)


    def Area(self, data=None) -> "AreaEntity":
        """Entity factory: client.Area().list() / client.Area().load({"id": ...})."""
        from newton_sdk.entity.area_entity import AreaEntity
        return AreaEntity(self, data)


    def Cos(self, data=None) -> "CosEntity":
        """Entity factory: client.Cos().list() / client.Cos().load({"id": ...})."""
        from newton_sdk.entity.cos_entity import CosEntity
        return CosEntity(self, data)


    def Derive(self, data=None) -> "DeriveEntity":
        """Entity factory: client.Derive().list() / client.Derive().load({"id": ...})."""
        from newton_sdk.entity.derive_entity import DeriveEntity
        return DeriveEntity(self, data)


    def Factor(self, data=None) -> "FactorEntity":
        """Entity factory: client.Factor().list() / client.Factor().load({"id": ...})."""
        from newton_sdk.entity.factor_entity import FactorEntity
        return FactorEntity(self, data)


    def Integrate(self, data=None) -> "IntegrateEntity":
        """Entity factory: client.Integrate().list() / client.Integrate().load({"id": ...})."""
        from newton_sdk.entity.integrate_entity import IntegrateEntity
        return IntegrateEntity(self, data)


    def Log(self, data=None) -> "LogEntity":
        """Entity factory: client.Log().list() / client.Log().load({"id": ...})."""
        from newton_sdk.entity.log_entity import LogEntity
        return LogEntity(self, data)


    def Simplify(self, data=None) -> "SimplifyEntity":
        """Entity factory: client.Simplify().list() / client.Simplify().load({"id": ...})."""
        from newton_sdk.entity.simplify_entity import SimplifyEntity
        return SimplifyEntity(self, data)


    def Sin(self, data=None) -> "SinEntity":
        """Entity factory: client.Sin().list() / client.Sin().load({"id": ...})."""
        from newton_sdk.entity.sin_entity import SinEntity
        return SinEntity(self, data)


    def Tan(self, data=None) -> "TanEntity":
        """Entity factory: client.Tan().list() / client.Tan().load({"id": ...})."""
        from newton_sdk.entity.tan_entity import TanEntity
        return TanEntity(self, data)


    def Tangent(self, data=None) -> "TangentEntity":
        """Entity factory: client.Tangent().list() / client.Tangent().load({"id": ...})."""
        from newton_sdk.entity.tangent_entity import TangentEntity
        return TangentEntity(self, data)


    def Zero(self, data=None) -> "ZeroEntity":
        """Entity factory: client.Zero().list() / client.Zero().load({"id": ...})."""
        from newton_sdk.entity.zero_entity import ZeroEntity
        return ZeroEntity(self, data)



    @classmethod
    def test(cls, testopts=None, sdkopts=None) -> "NewtonSDK":
        if sdkopts is None:
            sdkopts = {}
        sdkopts = vs.clone(sdkopts)
        if not isinstance(sdkopts, dict):
            sdkopts = {}

        if testopts is None:
            testopts = {}
        testopts = vs.clone(testopts)
        if not isinstance(testopts, dict):
            testopts = {}
        testopts["active"] = True

        vs.setpath(sdkopts, "feature.test", testopts)

        sdk = cls(sdkopts)
        sdk.mode = "test"

        return sdk


from typing import TYPE_CHECKING

if TYPE_CHECKING:
    from newton_sdk.entity.abs_entity import AbsEntity
    from newton_sdk.entity.arcco_entity import ArccoEntity
    from newton_sdk.entity.arcsin_entity import ArcsinEntity
    from newton_sdk.entity.arctan_entity import ArctanEntity
    from newton_sdk.entity.area_entity import AreaEntity
    from newton_sdk.entity.cos_entity import CosEntity
    from newton_sdk.entity.derive_entity import DeriveEntity
    from newton_sdk.entity.factor_entity import FactorEntity
    from newton_sdk.entity.integrate_entity import IntegrateEntity
    from newton_sdk.entity.log_entity import LogEntity
    from newton_sdk.entity.simplify_entity import SimplifyEntity
    from newton_sdk.entity.sin_entity import SinEntity
    from newton_sdk.entity.tan_entity import TanEntity
    from newton_sdk.entity.tangent_entity import TangentEntity
    from newton_sdk.entity.zero_entity import ZeroEntity
