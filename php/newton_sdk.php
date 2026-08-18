<?php
declare(strict_types=1);

// Newton SDK

require_once __DIR__ . '/utility/struct/Struct.php';
require_once __DIR__ . '/core/UtilityType.php';
require_once __DIR__ . '/core/Spec.php';
require_once __DIR__ . '/core/Helpers.php';

// Load utility registration
require_once __DIR__ . '/utility/Register.php';

// Load config and features
require_once __DIR__ . '/config.php';
require_once __DIR__ . '/feature/BaseFeature.php';
require_once __DIR__ . '/features.php';

use Voxgig\Struct\Struct;

// Features record diagnostic state on the client as dynamic properties
// (_retry, _cache, _metrics, ...); allow them explicitly (PHP 8.2+
// deprecates implicit dynamic properties).
#[\AllowDynamicProperties]
class NewtonSDK
{
    public string $mode;
    public array $features;
    public ?array $options;

    private $_utility;
    private $_rootctx;

    public function __construct(array $options = [])
    {
        $this->mode = "live";
        $this->features = [];
        $this->options = null;

        $utility = new NewtonUtility();
        $this->_utility = $utility;

        $config = NewtonConfig::shared_config();

        $this->_rootctx = ($utility->make_context)([
            "client" => $this,
            "utility" => $utility,
            "config" => $config,
            "options" => $options ?? [],
            "shared" => [],
        ], null);

        $this->options = ($utility->make_options)($this->_rootctx);

        if (Struct::getpath($this->options, "feature.test.active") === true) {
            $this->mode = "test";
        }

        $this->_rootctx->options = $this->options;

        // Add features in the resolved order (make_options puts an explicit
        // list order first, else defaults to test-first). Ordering matters: the
        // `test` feature installs the base mock transport and the transport
        // features (retry/cache/netsim/proxy/ratelimit) wrap whatever is
        // current, so `test` must be added before them to sit at the base.
        $feature_opts = NewtonHelpers::to_map(Struct::getprop($this->options, "feature"));
        if ($feature_opts) {
            $featureorder = Struct::getpath($this->options, "__derived__.featureorder");
            if (is_array($featureorder)) {
                foreach ($featureorder as $fname) {
                    $fopts = NewtonHelpers::to_map($feature_opts[$fname] ?? null);
                    if ($fopts && isset($fopts["active"]) && $fopts["active"] === true) {
                        ($utility->feature_add)($this->_rootctx, NewtonFeatures::make_feature($fname));
                    }
                }
            }
        }

        // Add extension features.
        $extend_val = Struct::getprop($this->options, "extend");
        if (is_array($extend_val)) {
            foreach ($extend_val as $f) {
                if (is_object($f) && method_exists($f, 'get_name')) {
                    ($utility->feature_add)($this->_rootctx, $f);
                }
            }
        }

        // Initialize features.
        foreach ($this->features as $f) {
            ($utility->feature_init)($this->_rootctx, $f);
        }

        ($utility->feature_hook)($this->_rootctx, "PostConstruct");
    }

    public function options_map(): array
    {
        $out = Struct::clone($this->options);
        return is_array($out) ? $out : [];
    }

    public function get_utility()
    {
        return NewtonUtility::copy($this->_utility);
    }

    public function get_root_ctx()
    {
        return $this->_rootctx;
    }

    public function prepare(array $fetchargs = []): mixed
    {
        $utility = $this->_utility;
        $fetchargs = $fetchargs ?? [];

        $ctrl = NewtonHelpers::to_map(Struct::getprop($fetchargs, "ctrl")) ?? [];

        $ctx = ($utility->make_context)([
            "opname" => "prepare",
            "ctrl" => $ctrl,
        ], $this->_rootctx);

        $opts = $this->options;
        $path = Struct::getprop($fetchargs, "path") ?? "";
        $path = is_string($path) ? $path : "";
        $method_val = Struct::getprop($fetchargs, "method") ?? "GET";
        $method_val = is_string($method_val) ? $method_val : "GET";
        $params = NewtonHelpers::to_map(Struct::getprop($fetchargs, "params")) ?? [];
        $query = NewtonHelpers::to_map(Struct::getprop($fetchargs, "query")) ?? [];
        $headers = ($utility->prepare_headers)($ctx);

        $base = Struct::getprop($opts, "base") ?? "";
        $base = is_string($base) ? $base : "";
        $prefix = Struct::getprop($opts, "prefix") ?? "";
        $prefix = is_string($prefix) ? $prefix : "";
        $suffix = Struct::getprop($opts, "suffix") ?? "";
        $suffix = is_string($suffix) ? $suffix : "";

        $ctx->spec = new NewtonSpec([
            "base" => $base, "prefix" => $prefix, "suffix" => $suffix,
            "path" => $path, "method" => $method_val,
            "params" => $params, "query" => $query, "headers" => $headers,
            "body" => Struct::getprop($fetchargs, "body"),
            "step" => "start",
        ]);

        // Merge user-provided headers.
        $uh = Struct::getprop($fetchargs, "headers");
        if (is_array($uh)) {
            foreach ($uh as $k => $v) {
                $ctx->spec->headers[$k] = $v;
            }
        }

        [$_, $err] = ($utility->prepare_auth)($ctx);
        if ($err) {
            return ($utility->make_error)($ctx, $err);
        }

        [$fetchdef, $fd_err] = ($utility->make_fetch_def)($ctx);
        if ($fd_err) {
            return ($utility->make_error)($ctx, $fd_err);
        }
        return $fetchdef;
    }

    // Raw endpoint access is operator-controllable, like every entity op.
    // Blocking it means denying BOTH the 'direct' and 'graphql' tokens,
    // since either one reaches the same endpoint.
    public function direct(array $fetchargs = []): mixed
    {
        if (!$this->op_allowed("direct")) {
            return $this->op_denied("direct");
        }

        return $this->raw_request($fetchargs);
    }

    // Is this raw-access op permitted by the SDK's allow.op option?
    private function op_allowed(string $op): bool
    {
        $allow_op = Struct::getpath($this->options, "allow.op");
        return is_string($allow_op) && str_contains($allow_op, $op);
    }

    private function op_denied(string $op): array
    {
        $allow_op = Struct::getpath($this->options, "allow.op");
        return [
            "ok" => false,
            "err" => new NewtonError($op . "_allow",
                "NewtonSDK: " . $op . ": operation not allowed by" .
                " SDK option allow.op value: \"" . (string)$allow_op . "\""),
        ];
    }

    // Ungated request path shared by direct and graphql, each of which
    // checks its own allow.op token first. Private, rather than a flag on
    // fetchargs: a caller-supplied marker would let anyone opt straight back
    // out of the gate by passing it.
    private function raw_request(array $fetchargs = []): mixed
    {
        $utility = $this->_utility;

        // direct() is the raw-HTTP escape hatch: it never throws, it returns
        // an {ok, err, ...} dict. prepare() now raises on error, so catch it
        // and surface the failure through the dict instead.
        try {
            $fetchdef = $this->prepare($fetchargs);
        } catch (\Throwable $err) {
            return ["ok" => false, "err" => $err];
        }

        $fetchargs = $fetchargs ?? [];
        $ctrl = NewtonHelpers::to_map(Struct::getprop($fetchargs, "ctrl")) ?? [];

        $ctx = ($utility->make_context)([
            "opname" => "direct",
            "ctrl" => $ctrl,
        ], $this->_rootctx);

        $url = $fetchdef["url"] ?? "";
        [$fetched, $fetch_err] = ($utility->fetcher)($ctx, $url, $fetchdef);

        if ($fetch_err) {
            return ["ok" => false, "err" => $fetch_err];
        }

        if ($fetched === null) {
            return [
                "ok" => false,
                "err" => $ctx->make_error("direct_no_response", "response: undefined"),
            ];
        }

        if (is_array($fetched)) {
            $status = NewtonHelpers::to_int(Struct::getprop($fetched, "status"));
            $headers = Struct::getprop($fetched, "headers") ?? [];

            // No-body responses (204, 304) and explicit zero content-length
            // must skip JSON parsing — calling json() on an empty body errors.
            $content_length = is_array($headers) ? ($headers["content-length"] ?? null) : null;
            $no_body = $status === 204 || $status === 304 || (string)$content_length === "0";

            $json_data = null;
            if (!$no_body) {
                $jf = Struct::getprop($fetched, "json");
                if (is_callable($jf)) {
                    try {
                        $json_data = $jf();
                    } catch (\Throwable $e) {
                        // Non-JSON body — leave data null but keep status/ok.
                        $json_data = null;
                    }
                }
            }

            return [
                "ok" => $status >= 200 && $status < 300,
                "status" => $status,
                "headers" => Struct::getprop($fetched, "headers"),
                "data" => $json_data,
            ];
        }

        return [
            "ok" => false,
            "err" => $ctx->make_error("direct_invalid", "invalid response type"),
        ];
    }

    // Raw GraphQL access: the pressure valve that makes the generated
    // surface's deliberate omissions (per-call selection sets, typed filter
    // builders, batching, subscriptions) livable — the whole schema stays
    // reachable.
    //
    // Thin wrapper over the same prepare/fetch path direct uses, with the
    // one thing raw direct cannot do for GraphQL: a GraphQL failure rides
    // HTTP 200 as a top-level `errors` array, so status alone would report
    // a failed query as ok.
    //
    // NOTE: like direct, this bypasses the feature pipeline — no retry,
    // ratelimit or paging features apply.
    public function graphql(string $query, ?array $variables = null, ?array $ctrl = null): mixed
    {
        if (!$this->op_allowed("graphql")) {
            return $this->op_denied("graphql");
        }

        $res = $this->raw_request([
            "method" => "POST",
            "headers" => ["content-type" => "application/json"],
            "body" => ["query" => $query, "variables" => $variables ?? []],
            "ctrl" => $ctrl ?? [],
        ]);

        if (!is_array($res)) {
            return $res;
        }

        // Errors are read BEFORE any status check: a GraphQL parse or
        // validation failure comes back as HTTP 400 carrying the standard
        // { errors: [...] } body, and the raw path represents a non-2xx as
        // ok:false with no err — so returning early on status would discard
        // the server's own diagnostics, which are the only useful part of
        // that response.
        $errors = Struct::getpath($res, "data.errors");

        if (is_array($errors) && 0 < count($errors)) {
            $first = is_array($errors[0]) ? $errors[0] : [];
            $msg = $first["message"] ?? "";
            if (!is_string($msg) || "" === $msg) {
                $msg = "graphql error";
            }
            $res["ok"] = false;
            $res["err"] = new NewtonError("graphql_error",
                "NewtonSDK: graphql: " . $msg);
            $res["graphql"] = $errors;
        }

        return $res;
    }


    private $_abs = null;

    // Canonical facade: $client->Abs()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->abs()
    // resolves here too.
    public function Abs($data = null)
    {
        require_once __DIR__ . '/entity/abs_entity.php';
        if ($data === null) {
            if ($this->_abs === null) {
                $this->_abs = new AbsEntity($this, null);
            }
            return $this->_abs;
        }
        return new AbsEntity($this, $data);
    }


    private $_arcco = null;

    // Canonical facade: $client->Arcco()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->arcco()
    // resolves here too.
    public function Arcco($data = null)
    {
        require_once __DIR__ . '/entity/arcco_entity.php';
        if ($data === null) {
            if ($this->_arcco === null) {
                $this->_arcco = new ArccoEntity($this, null);
            }
            return $this->_arcco;
        }
        return new ArccoEntity($this, $data);
    }


    private $_arcsin = null;

    // Canonical facade: $client->Arcsin()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->arcsin()
    // resolves here too.
    public function Arcsin($data = null)
    {
        require_once __DIR__ . '/entity/arcsin_entity.php';
        if ($data === null) {
            if ($this->_arcsin === null) {
                $this->_arcsin = new ArcsinEntity($this, null);
            }
            return $this->_arcsin;
        }
        return new ArcsinEntity($this, $data);
    }


    private $_arctan = null;

    // Canonical facade: $client->Arctan()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->arctan()
    // resolves here too.
    public function Arctan($data = null)
    {
        require_once __DIR__ . '/entity/arctan_entity.php';
        if ($data === null) {
            if ($this->_arctan === null) {
                $this->_arctan = new ArctanEntity($this, null);
            }
            return $this->_arctan;
        }
        return new ArctanEntity($this, $data);
    }


    private $_area = null;

    // Canonical facade: $client->Area()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->area()
    // resolves here too.
    public function Area($data = null)
    {
        require_once __DIR__ . '/entity/area_entity.php';
        if ($data === null) {
            if ($this->_area === null) {
                $this->_area = new AreaEntity($this, null);
            }
            return $this->_area;
        }
        return new AreaEntity($this, $data);
    }


    private $_cos = null;

    // Canonical facade: $client->Cos()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->cos()
    // resolves here too.
    public function Cos($data = null)
    {
        require_once __DIR__ . '/entity/cos_entity.php';
        if ($data === null) {
            if ($this->_cos === null) {
                $this->_cos = new CosEntity($this, null);
            }
            return $this->_cos;
        }
        return new CosEntity($this, $data);
    }


    private $_derive = null;

    // Canonical facade: $client->Derive()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->derive()
    // resolves here too.
    public function Derive($data = null)
    {
        require_once __DIR__ . '/entity/derive_entity.php';
        if ($data === null) {
            if ($this->_derive === null) {
                $this->_derive = new DeriveEntity($this, null);
            }
            return $this->_derive;
        }
        return new DeriveEntity($this, $data);
    }


    private $_factor = null;

    // Canonical facade: $client->Factor()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->factor()
    // resolves here too.
    public function Factor($data = null)
    {
        require_once __DIR__ . '/entity/factor_entity.php';
        if ($data === null) {
            if ($this->_factor === null) {
                $this->_factor = new FactorEntity($this, null);
            }
            return $this->_factor;
        }
        return new FactorEntity($this, $data);
    }


    private $_integrate = null;

    // Canonical facade: $client->Integrate()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->integrate()
    // resolves here too.
    public function Integrate($data = null)
    {
        require_once __DIR__ . '/entity/integrate_entity.php';
        if ($data === null) {
            if ($this->_integrate === null) {
                $this->_integrate = new IntegrateEntity($this, null);
            }
            return $this->_integrate;
        }
        return new IntegrateEntity($this, $data);
    }


    private $_log = null;

    // Canonical facade: $client->Log()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->log()
    // resolves here too.
    public function Log($data = null)
    {
        require_once __DIR__ . '/entity/log_entity.php';
        if ($data === null) {
            if ($this->_log === null) {
                $this->_log = new LogEntity($this, null);
            }
            return $this->_log;
        }
        return new LogEntity($this, $data);
    }


    private $_simplify = null;

    // Canonical facade: $client->Simplify()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->simplify()
    // resolves here too.
    public function Simplify($data = null)
    {
        require_once __DIR__ . '/entity/simplify_entity.php';
        if ($data === null) {
            if ($this->_simplify === null) {
                $this->_simplify = new SimplifyEntity($this, null);
            }
            return $this->_simplify;
        }
        return new SimplifyEntity($this, $data);
    }


    private $_sin = null;

    // Canonical facade: $client->Sin()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->sin()
    // resolves here too.
    public function Sin($data = null)
    {
        require_once __DIR__ . '/entity/sin_entity.php';
        if ($data === null) {
            if ($this->_sin === null) {
                $this->_sin = new SinEntity($this, null);
            }
            return $this->_sin;
        }
        return new SinEntity($this, $data);
    }


    private $_tan = null;

    // Canonical facade: $client->Tan()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->tan()
    // resolves here too.
    public function Tan($data = null)
    {
        require_once __DIR__ . '/entity/tan_entity.php';
        if ($data === null) {
            if ($this->_tan === null) {
                $this->_tan = new TanEntity($this, null);
            }
            return $this->_tan;
        }
        return new TanEntity($this, $data);
    }


    private $_tangent = null;

    // Canonical facade: $client->Tangent()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->tangent()
    // resolves here too.
    public function Tangent($data = null)
    {
        require_once __DIR__ . '/entity/tangent_entity.php';
        if ($data === null) {
            if ($this->_tangent === null) {
                $this->_tangent = new TangentEntity($this, null);
            }
            return $this->_tangent;
        }
        return new TangentEntity($this, $data);
    }


    private $_zero = null;

    // Canonical facade: $client->Zero()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->zero()
    // resolves here too.
    public function Zero($data = null)
    {
        require_once __DIR__ . '/entity/zero_entity.php';
        if ($data === null) {
            if ($this->_zero === null) {
                $this->_zero = new ZeroEntity($this, null);
            }
            return $this->_zero;
        }
        return new ZeroEntity($this, $data);
    }



    public static function test(?array $testopts = null, ?array $sdkopts = null): self
    {
        $sdkopts = $sdkopts ?? [];
        $sdkopts = Struct::clone($sdkopts);
        $sdkopts = is_array($sdkopts) ? $sdkopts : [];

        $testopts = $testopts ?? [];
        $testopts = Struct::clone($testopts);
        $testopts = is_array($testopts) ? $testopts : [];
        $testopts["active"] = true;

        if (!isset($sdkopts["feature"])) {
            $sdkopts["feature"] = [];
        }
        $sdkopts["feature"]["test"] = $testopts;

        $sdk = new NewtonSDK($sdkopts);
        $sdk->mode = "test";
        return $sdk;
    }
}
