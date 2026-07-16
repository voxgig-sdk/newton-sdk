package core

import (
	"fmt"

	vs "github.com/voxgig-sdk/newton-sdk/go/utility/struct"
)

type NewtonSDK struct {
	Mode     string
	options  map[string]any
	utility  *Utility
	Features []Feature
	rootctx  *Context
}

func NewNewtonSDK(options map[string]any) *NewtonSDK {
	sdk := &NewtonSDK{
		Mode:     "live",
		Features: []Feature{},
	}

	sdk.utility = NewUtility()

	config := MakeConfig()

	sdk.rootctx = sdk.utility.MakeContext(map[string]any{
		"client":  sdk,
		"utility": sdk.utility,
		"config":  config,
		"options": options,
		"shared":  map[string]any{},
	}, nil)

	sdk.options = sdk.utility.MakeOptions(sdk.rootctx)

	if vs.GetPath([]any{"feature", "test", "active"}, sdk.options) == true {
		sdk.Mode = "test"
	}

	sdk.rootctx.Options = sdk.options

	// Add features in the resolved order (MakeOptions puts an explicit array
	// order first, else defaults to test-first). Ordering matters: the `test`
	// feature installs the base mock transport and the transport features
	// (retry/cache/netsim/proxy/ratelimit) wrap whatever is current, so `test`
	// must be added before them to sit at the base of the chain.
	featureOpts := ToMapAny(vs.GetProp(sdk.options, "feature"))
	if featureOpts != nil {
		if fo, ok := vs.GetPath([]any{"__derived__", "featureorder"}, sdk.options).([]any); ok {
			for _, n := range fo {
				fname, _ := n.(string)
				fopts := ToMapAny(featureOpts[fname])
				if fopts != nil {
					if active, ok := fopts["active"]; ok {
						if ab, ok := active.(bool); ok && ab {
							sdk.utility.FeatureAdd(sdk.rootctx, makeFeature(fname))
						}
					}
				}
			}
		}
	}

	// Add extension features.
	if extend := vs.GetProp(sdk.options, "extend"); extend != nil {
		if extList, ok := extend.([]any); ok {
			for _, f := range extList {
				if feat, ok := f.(Feature); ok {
					sdk.utility.FeatureAdd(sdk.rootctx, feat)
				}
			}
		}
	}

	// Initialize features.
	for _, f := range sdk.Features {
		sdk.utility.FeatureInit(sdk.rootctx, f)
	}

	sdk.utility.FeatureHook(sdk.rootctx, "PostConstruct")

	return sdk
}

func (sdk *NewtonSDK) OptionsMap() map[string]any {
	out := vs.Clone(sdk.options)
	if om, ok := out.(map[string]any); ok {
		return om
	}
	return map[string]any{}
}

func (sdk *NewtonSDK) GetUtility() *Utility {
	return CopyUtility(sdk.utility)
}

func (sdk *NewtonSDK) GetRootCtx() *Context {
	return sdk.rootctx
}

func (sdk *NewtonSDK) Prepare(fetchargs map[string]any) (map[string]any, error) {
	utility := sdk.utility

	if fetchargs == nil {
		fetchargs = map[string]any{}
	}

	var ctrl map[string]any
	if c := vs.GetProp(fetchargs, "ctrl"); c != nil {
		if cm, ok := c.(map[string]any); ok {
			ctrl = cm
		}
	}
	if ctrl == nil {
		ctrl = map[string]any{}
	}

	ctx := utility.MakeContext(map[string]any{
		"opname": "prepare",
		"ctrl":   ctrl,
	}, sdk.rootctx)

	options := sdk.options

	path, _ := vs.GetProp(fetchargs, "path").(string)
	method, _ := vs.GetProp(fetchargs, "method").(string)
	if method == "" {
		method = "GET"
	}

	params := ToMapAny(vs.GetProp(fetchargs, "params"))
	if params == nil {
		params = map[string]any{}
	}
	query := ToMapAny(vs.GetProp(fetchargs, "query"))
	if query == nil {
		query = map[string]any{}
	}

	headers := utility.PrepareHeaders(ctx)

	base, _ := vs.GetProp(options, "base").(string)
	prefix, _ := vs.GetProp(options, "prefix").(string)
	suffix, _ := vs.GetProp(options, "suffix").(string)

	ctx.Spec = NewSpec(map[string]any{
		"base":    base,
		"prefix":  prefix,
		"suffix":  suffix,
		"path":    path,
		"method":  method,
		"params":  params,
		"query":   query,
		"headers": headers,
		"body":    vs.GetProp(fetchargs, "body"),
		"step":    "start",
	})

	// Merge user-provided headers.
	if uh := vs.GetProp(fetchargs, "headers"); uh != nil {
		if uhm, ok := uh.(map[string]any); ok {
			for k, v := range uhm {
				ctx.Spec.Headers[k] = v
			}
		}
	}

	_, err := utility.PrepareAuth(ctx)
	if err != nil {
		return nil, err
	}

	return utility.MakeFetchDef(ctx)
}

func (sdk *NewtonSDK) Direct(fetchargs map[string]any) (map[string]any, error) {
	utility := sdk.utility

	fetchdef, err := sdk.Prepare(fetchargs)
	if err != nil {
		return map[string]any{"ok": false, "err": err}, nil
	}

	if fetchargs == nil {
		fetchargs = map[string]any{}
	}

	var ctrl map[string]any
	if c := vs.GetProp(fetchargs, "ctrl"); c != nil {
		if cm, ok := c.(map[string]any); ok {
			ctrl = cm
		}
	}
	if ctrl == nil {
		ctrl = map[string]any{}
	}

	ctx := utility.MakeContext(map[string]any{
		"opname": "direct",
		"ctrl":   ctrl,
	}, sdk.rootctx)

	url, _ := fetchdef["url"].(string)
	fetched, fetchErr := utility.Fetcher(ctx, url, fetchdef)

	if fetchErr != nil {
		return map[string]any{"ok": false, "err": fetchErr}, nil
	}

	if fetched == nil {
		return map[string]any{
			"ok":  false,
			"err": ctx.MakeError("direct_no_response", "response: undefined"),
		}, nil
	}

	if fm, ok := fetched.(map[string]any); ok {
		status := ToInt(vs.GetProp(fm, "status"))
		headers := vs.GetProp(fm, "headers")

		// No-body responses (204, 304) and explicit zero content-length
		// must skip JSON parsing — calling json() on an empty body errors.
		var contentLength string
		if hm, ok := headers.(map[string]any); ok {
			if cl, ok := hm["content-length"]; ok {
				contentLength = fmt.Sprintf("%v", cl)
			}
		}
		noBody := status == 204 || status == 304 || contentLength == "0"

		var jsonData any
		if !noBody {
			if jf := vs.GetProp(fm, "json"); jf != nil {
				if f, ok := jf.(func() any); ok {
					// f() returns nil on parse error in our fetcher.
					jsonData = f()
				}
			}
		}

		return map[string]any{
			"ok":      status >= 200 && status < 300,
			"status":  status,
			"headers": headers,
			"data":    jsonData,
		}, nil
	}

	return map[string]any{"ok": false, "err": ctx.MakeError("direct_invalid", "invalid response type")}, nil
}


// Abs returns a Abs entity bound to this client.
// Idiomatic usage: client.Abs(nil).List(nil, nil) or
// client.Abs(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NewtonSDK) Abs(data map[string]any) NewtonEntity {
	return NewAbsEntityFunc(sdk, data)
}


// Arcco returns a Arcco entity bound to this client.
// Idiomatic usage: client.Arcco(nil).List(nil, nil) or
// client.Arcco(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NewtonSDK) Arcco(data map[string]any) NewtonEntity {
	return NewArccoEntityFunc(sdk, data)
}


// Arcsin returns a Arcsin entity bound to this client.
// Idiomatic usage: client.Arcsin(nil).List(nil, nil) or
// client.Arcsin(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NewtonSDK) Arcsin(data map[string]any) NewtonEntity {
	return NewArcsinEntityFunc(sdk, data)
}


// Arctan returns a Arctan entity bound to this client.
// Idiomatic usage: client.Arctan(nil).List(nil, nil) or
// client.Arctan(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NewtonSDK) Arctan(data map[string]any) NewtonEntity {
	return NewArctanEntityFunc(sdk, data)
}


// Area returns a Area entity bound to this client.
// Idiomatic usage: client.Area(nil).List(nil, nil) or
// client.Area(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NewtonSDK) Area(data map[string]any) NewtonEntity {
	return NewAreaEntityFunc(sdk, data)
}


// Cos returns a Cos entity bound to this client.
// Idiomatic usage: client.Cos(nil).List(nil, nil) or
// client.Cos(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NewtonSDK) Cos(data map[string]any) NewtonEntity {
	return NewCosEntityFunc(sdk, data)
}


// Derive returns a Derive entity bound to this client.
// Idiomatic usage: client.Derive(nil).List(nil, nil) or
// client.Derive(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NewtonSDK) Derive(data map[string]any) NewtonEntity {
	return NewDeriveEntityFunc(sdk, data)
}


// Factor returns a Factor entity bound to this client.
// Idiomatic usage: client.Factor(nil).List(nil, nil) or
// client.Factor(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NewtonSDK) Factor(data map[string]any) NewtonEntity {
	return NewFactorEntityFunc(sdk, data)
}


// Integrate returns a Integrate entity bound to this client.
// Idiomatic usage: client.Integrate(nil).List(nil, nil) or
// client.Integrate(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NewtonSDK) Integrate(data map[string]any) NewtonEntity {
	return NewIntegrateEntityFunc(sdk, data)
}


// Log returns a Log entity bound to this client.
// Idiomatic usage: client.Log(nil).List(nil, nil) or
// client.Log(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NewtonSDK) Log(data map[string]any) NewtonEntity {
	return NewLogEntityFunc(sdk, data)
}


// Simplify returns a Simplify entity bound to this client.
// Idiomatic usage: client.Simplify(nil).List(nil, nil) or
// client.Simplify(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NewtonSDK) Simplify(data map[string]any) NewtonEntity {
	return NewSimplifyEntityFunc(sdk, data)
}


// Sin returns a Sin entity bound to this client.
// Idiomatic usage: client.Sin(nil).List(nil, nil) or
// client.Sin(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NewtonSDK) Sin(data map[string]any) NewtonEntity {
	return NewSinEntityFunc(sdk, data)
}


// Tan returns a Tan entity bound to this client.
// Idiomatic usage: client.Tan(nil).List(nil, nil) or
// client.Tan(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NewtonSDK) Tan(data map[string]any) NewtonEntity {
	return NewTanEntityFunc(sdk, data)
}


// Tangent returns a Tangent entity bound to this client.
// Idiomatic usage: client.Tangent(nil).List(nil, nil) or
// client.Tangent(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NewtonSDK) Tangent(data map[string]any) NewtonEntity {
	return NewTangentEntityFunc(sdk, data)
}


// Zero returns a Zero entity bound to this client.
// Idiomatic usage: client.Zero(nil).List(nil, nil) or
// client.Zero(nil).Load(map[string]any{"id": ...}, nil).
func (sdk *NewtonSDK) Zero(data map[string]any) NewtonEntity {
	return NewZeroEntityFunc(sdk, data)
}



func TestSDK(testopts map[string]any, sdkopts map[string]any) *NewtonSDK {
	if sdkopts == nil {
		sdkopts = map[string]any{}
	}
	sdkopts = vs.Clone(sdkopts).(map[string]any)

	if testopts == nil {
		testopts = map[string]any{}
	}
	testopts = vs.Clone(testopts).(map[string]any)
	testopts["active"] = true

	vs.SetPath(sdkopts, []any{"feature", "test"}, testopts)

	sdk := NewNewtonSDK(sdkopts)
	sdk.Mode = "test"

	return sdk
}
