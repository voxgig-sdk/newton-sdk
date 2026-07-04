# Newton SDK

from utility.voxgig_struct import voxgig_struct as vs
from core.utility_type import NewtonUtility
from core.spec import NewtonSpec
from core import helpers

# Load utility registration (populates Utility._registrar)
from utility import register

# Load features
from feature.base_feature import NewtonBaseFeature
from features import _make_feature


class NewtonSDK:

    def __init__(self, options=None):
        self.mode = "live"
        self.features = []
        self.options = None

        utility = NewtonUtility()
        self._utility = utility

        from config import make_config
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

        # Add features from config.
        feature_opts = helpers.to_map(vs.getprop(self.options, "feature"))
        if feature_opts is not None:
            feature_items = vs.items(feature_opts)
            if feature_items is not None:
                for item in feature_items:
                    fname = item[0]
                    fopts = helpers.to_map(item[1])
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

    def direct(self, fetchargs=None):
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


    @property
    def abs(self):
        """Idiomatic facade: client.abs.list() / client.abs.load({"id": ...})."""
        from entity.abs_entity import AbsEntity
        cached = getattr(self, "_abs", None)
        if cached is None:
            cached = AbsEntity(self, None)
            self._abs = cached
        return cached

    def Abs(self, data=None):
        # Deprecated: use client.abs instead.
        from entity.abs_entity import AbsEntity
        return AbsEntity(self, data)


    @property
    def arcco(self):
        """Idiomatic facade: client.arcco.list() / client.arcco.load({"id": ...})."""
        from entity.arcco_entity import ArccoEntity
        cached = getattr(self, "_arcco", None)
        if cached is None:
            cached = ArccoEntity(self, None)
            self._arcco = cached
        return cached

    def Arcco(self, data=None):
        # Deprecated: use client.arcco instead.
        from entity.arcco_entity import ArccoEntity
        return ArccoEntity(self, data)


    @property
    def arcsin(self):
        """Idiomatic facade: client.arcsin.list() / client.arcsin.load({"id": ...})."""
        from entity.arcsin_entity import ArcsinEntity
        cached = getattr(self, "_arcsin", None)
        if cached is None:
            cached = ArcsinEntity(self, None)
            self._arcsin = cached
        return cached

    def Arcsin(self, data=None):
        # Deprecated: use client.arcsin instead.
        from entity.arcsin_entity import ArcsinEntity
        return ArcsinEntity(self, data)


    @property
    def arctan(self):
        """Idiomatic facade: client.arctan.list() / client.arctan.load({"id": ...})."""
        from entity.arctan_entity import ArctanEntity
        cached = getattr(self, "_arctan", None)
        if cached is None:
            cached = ArctanEntity(self, None)
            self._arctan = cached
        return cached

    def Arctan(self, data=None):
        # Deprecated: use client.arctan instead.
        from entity.arctan_entity import ArctanEntity
        return ArctanEntity(self, data)


    @property
    def area(self):
        """Idiomatic facade: client.area.list() / client.area.load({"id": ...})."""
        from entity.area_entity import AreaEntity
        cached = getattr(self, "_area", None)
        if cached is None:
            cached = AreaEntity(self, None)
            self._area = cached
        return cached

    def Area(self, data=None):
        # Deprecated: use client.area instead.
        from entity.area_entity import AreaEntity
        return AreaEntity(self, data)


    @property
    def cos(self):
        """Idiomatic facade: client.cos.list() / client.cos.load({"id": ...})."""
        from entity.cos_entity import CosEntity
        cached = getattr(self, "_cos", None)
        if cached is None:
            cached = CosEntity(self, None)
            self._cos = cached
        return cached

    def Cos(self, data=None):
        # Deprecated: use client.cos instead.
        from entity.cos_entity import CosEntity
        return CosEntity(self, data)


    @property
    def derive(self):
        """Idiomatic facade: client.derive.list() / client.derive.load({"id": ...})."""
        from entity.derive_entity import DeriveEntity
        cached = getattr(self, "_derive", None)
        if cached is None:
            cached = DeriveEntity(self, None)
            self._derive = cached
        return cached

    def Derive(self, data=None):
        # Deprecated: use client.derive instead.
        from entity.derive_entity import DeriveEntity
        return DeriveEntity(self, data)


    @property
    def factor(self):
        """Idiomatic facade: client.factor.list() / client.factor.load({"id": ...})."""
        from entity.factor_entity import FactorEntity
        cached = getattr(self, "_factor", None)
        if cached is None:
            cached = FactorEntity(self, None)
            self._factor = cached
        return cached

    def Factor(self, data=None):
        # Deprecated: use client.factor instead.
        from entity.factor_entity import FactorEntity
        return FactorEntity(self, data)


    @property
    def integrate(self):
        """Idiomatic facade: client.integrate.list() / client.integrate.load({"id": ...})."""
        from entity.integrate_entity import IntegrateEntity
        cached = getattr(self, "_integrate", None)
        if cached is None:
            cached = IntegrateEntity(self, None)
            self._integrate = cached
        return cached

    def Integrate(self, data=None):
        # Deprecated: use client.integrate instead.
        from entity.integrate_entity import IntegrateEntity
        return IntegrateEntity(self, data)


    @property
    def log(self):
        """Idiomatic facade: client.log.list() / client.log.load({"id": ...})."""
        from entity.log_entity import LogEntity
        cached = getattr(self, "_log", None)
        if cached is None:
            cached = LogEntity(self, None)
            self._log = cached
        return cached

    def Log(self, data=None):
        # Deprecated: use client.log instead.
        from entity.log_entity import LogEntity
        return LogEntity(self, data)


    @property
    def simplify(self):
        """Idiomatic facade: client.simplify.list() / client.simplify.load({"id": ...})."""
        from entity.simplify_entity import SimplifyEntity
        cached = getattr(self, "_simplify", None)
        if cached is None:
            cached = SimplifyEntity(self, None)
            self._simplify = cached
        return cached

    def Simplify(self, data=None):
        # Deprecated: use client.simplify instead.
        from entity.simplify_entity import SimplifyEntity
        return SimplifyEntity(self, data)


    @property
    def sin(self):
        """Idiomatic facade: client.sin.list() / client.sin.load({"id": ...})."""
        from entity.sin_entity import SinEntity
        cached = getattr(self, "_sin", None)
        if cached is None:
            cached = SinEntity(self, None)
            self._sin = cached
        return cached

    def Sin(self, data=None):
        # Deprecated: use client.sin instead.
        from entity.sin_entity import SinEntity
        return SinEntity(self, data)


    @property
    def tan(self):
        """Idiomatic facade: client.tan.list() / client.tan.load({"id": ...})."""
        from entity.tan_entity import TanEntity
        cached = getattr(self, "_tan", None)
        if cached is None:
            cached = TanEntity(self, None)
            self._tan = cached
        return cached

    def Tan(self, data=None):
        # Deprecated: use client.tan instead.
        from entity.tan_entity import TanEntity
        return TanEntity(self, data)


    @property
    def tangent(self):
        """Idiomatic facade: client.tangent.list() / client.tangent.load({"id": ...})."""
        from entity.tangent_entity import TangentEntity
        cached = getattr(self, "_tangent", None)
        if cached is None:
            cached = TangentEntity(self, None)
            self._tangent = cached
        return cached

    def Tangent(self, data=None):
        # Deprecated: use client.tangent instead.
        from entity.tangent_entity import TangentEntity
        return TangentEntity(self, data)


    @property
    def zero(self):
        """Idiomatic facade: client.zero.list() / client.zero.load({"id": ...})."""
        from entity.zero_entity import ZeroEntity
        cached = getattr(self, "_zero", None)
        if cached is None:
            cached = ZeroEntity(self, None)
            self._zero = cached
        return cached

    def Zero(self, data=None):
        # Deprecated: use client.zero instead.
        from entity.zero_entity import ZeroEntity
        return ZeroEntity(self, data)



    @classmethod
    def test(cls, testopts=None, sdkopts=None):
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
