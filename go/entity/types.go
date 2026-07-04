// Typed models for the Newton SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
package entity

import "encoding/json"

// Abs is the typed data model for the abs entity.
type Abs struct {
	Expression string `json:"expression"`
	Operation string `json:"operation"`
	Result string `json:"result"`
}

// AbsLoadMatch is the typed request payload for Abs.LoadTyped.
type AbsLoadMatch struct {
	Id string `json:"id"`
}

// Arcco is the typed data model for the arcco entity.
type Arcco struct {
	Expression string `json:"expression"`
	Operation string `json:"operation"`
	Result string `json:"result"`
}

// ArccoLoadMatch is the typed request payload for Arcco.LoadTyped.
type ArccoLoadMatch struct {
	Id string `json:"id"`
}

// Arcsin is the typed data model for the arcsin entity.
type Arcsin struct {
	Expression string `json:"expression"`
	Operation string `json:"operation"`
	Result string `json:"result"`
}

// ArcsinLoadMatch is the typed request payload for Arcsin.LoadTyped.
type ArcsinLoadMatch struct {
	Id string `json:"id"`
}

// Arctan is the typed data model for the arctan entity.
type Arctan struct {
	Expression string `json:"expression"`
	Operation string `json:"operation"`
	Result string `json:"result"`
}

// ArctanLoadMatch is the typed request payload for Arctan.LoadTyped.
type ArctanLoadMatch struct {
	Id string `json:"id"`
}

// Area is the typed data model for the area entity.
type Area struct {
	Expression string `json:"expression"`
	Operation string `json:"operation"`
	Result string `json:"result"`
}

// AreaLoadMatch is the typed request payload for Area.LoadTyped.
type AreaLoadMatch struct {
	Id string `json:"id"`
}

// Cos is the typed data model for the cos entity.
type Cos struct {
	Expression string `json:"expression"`
	Operation string `json:"operation"`
	Result string `json:"result"`
}

// CosLoadMatch is the typed request payload for Cos.LoadTyped.
type CosLoadMatch struct {
	Id string `json:"id"`
}

// Derive is the typed data model for the derive entity.
type Derive struct {
	Expression string `json:"expression"`
	Operation string `json:"operation"`
	Result string `json:"result"`
}

// DeriveLoadMatch is the typed request payload for Derive.LoadTyped.
type DeriveLoadMatch struct {
	Id string `json:"id"`
}

// Factor is the typed data model for the factor entity.
type Factor struct {
	Expression string `json:"expression"`
	Operation string `json:"operation"`
	Result string `json:"result"`
}

// FactorLoadMatch is the typed request payload for Factor.LoadTyped.
type FactorLoadMatch struct {
	Id string `json:"id"`
}

// Integrate is the typed data model for the integrate entity.
type Integrate struct {
	Expression string `json:"expression"`
	Operation string `json:"operation"`
	Result string `json:"result"`
}

// IntegrateLoadMatch is the typed request payload for Integrate.LoadTyped.
type IntegrateLoadMatch struct {
	Id string `json:"id"`
}

// Log is the typed data model for the log entity.
type Log struct {
	Expression string `json:"expression"`
	Operation string `json:"operation"`
	Result string `json:"result"`
}

// LogLoadMatch is the typed request payload for Log.LoadTyped.
type LogLoadMatch struct {
	Id string `json:"id"`
}

// Simplify is the typed data model for the simplify entity.
type Simplify struct {
	Expression string `json:"expression"`
	Operation string `json:"operation"`
	Result string `json:"result"`
}

// SimplifyLoadMatch is the typed request payload for Simplify.LoadTyped.
type SimplifyLoadMatch struct {
	Id string `json:"id"`
}

// Sin is the typed data model for the sin entity.
type Sin struct {
	Expression string `json:"expression"`
	Operation string `json:"operation"`
	Result string `json:"result"`
}

// SinLoadMatch is the typed request payload for Sin.LoadTyped.
type SinLoadMatch struct {
	Id string `json:"id"`
}

// Tan is the typed data model for the tan entity.
type Tan struct {
	Expression string `json:"expression"`
	Operation string `json:"operation"`
	Result string `json:"result"`
}

// TanLoadMatch is the typed request payload for Tan.LoadTyped.
type TanLoadMatch struct {
	Id string `json:"id"`
}

// Tangent is the typed data model for the tangent entity.
type Tangent struct {
	Expression string `json:"expression"`
	Operation string `json:"operation"`
	Result string `json:"result"`
}

// TangentLoadMatch is the typed request payload for Tangent.LoadTyped.
type TangentLoadMatch struct {
	Id string `json:"id"`
}

// Zero is the typed data model for the zero entity.
type Zero struct {
	Expression string `json:"expression"`
	Operation string `json:"operation"`
	Result string `json:"result"`
}

// ZeroLoadMatch is the typed request payload for Zero.LoadTyped.
type ZeroLoadMatch struct {
	Id string `json:"id"`
}

// asMap turns a typed request/data struct into the map[string]any the
// runtime op pipeline consumes, honouring the json tags above.
func asMap(v any) map[string]any {
	out := map[string]any{}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedFrom decodes a runtime value (a map[string]any produced by the op
// pipeline) into a typed model T via a JSON round-trip. On any error it
// returns the zero value of T; the op's own (value, error) tuple carries the
// real error.
func typedFrom[T any](v any) T {
	var out T
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedSliceFrom decodes a runtime list value ([]any of maps) into a typed
// slice []T via a JSON round-trip, for list ops.
func typedSliceFrom[T any](v any) []T {
	var out []T
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}
