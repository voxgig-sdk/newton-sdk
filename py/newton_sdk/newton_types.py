# Typed models for the Newton SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.
#
# These are TypedDicts, not dataclasses: the SDK ops return/accept plain dicts
# at runtime, and a TypedDict IS a dict shape, so the types match the runtime.
# Optional (req:false) keys are modelled as TypedDict key-optionality
# (total=False), split into a required base + total=False subclass when a type
# has both required and optional keys.

from __future__ import annotations

from typing import TypedDict, Any


class AbsRequired(TypedDict):
    expression: str
    operation: str
    result: str


class Abs(AbsRequired, total=False):
    id: str


class AbsLoadMatch(TypedDict):
    id: str


class ArccoRequired(TypedDict):
    expression: str
    operation: str
    result: str


class Arcco(ArccoRequired, total=False):
    id: str


class ArccoLoadMatch(TypedDict):
    id: str


class ArcsinRequired(TypedDict):
    expression: str
    operation: str
    result: str


class Arcsin(ArcsinRequired, total=False):
    id: str


class ArcsinLoadMatch(TypedDict):
    id: str


class ArctanRequired(TypedDict):
    expression: str
    operation: str
    result: str


class Arctan(ArctanRequired, total=False):
    id: str


class ArctanLoadMatch(TypedDict):
    id: str


class AreaRequired(TypedDict):
    expression: str
    operation: str
    result: str


class Area(AreaRequired, total=False):
    id: str


class AreaLoadMatch(TypedDict):
    id: str


class CosRequired(TypedDict):
    expression: str
    operation: str
    result: str


class Cos(CosRequired, total=False):
    id: str


class CosLoadMatch(TypedDict):
    id: str


class DeriveRequired(TypedDict):
    expression: str
    operation: str
    result: str


class Derive(DeriveRequired, total=False):
    id: str


class DeriveLoadMatch(TypedDict):
    id: str


class FactorRequired(TypedDict):
    expression: str
    operation: str
    result: str


class Factor(FactorRequired, total=False):
    id: str


class FactorLoadMatch(TypedDict):
    id: str


class IntegrateRequired(TypedDict):
    expression: str
    operation: str
    result: str


class Integrate(IntegrateRequired, total=False):
    id: str


class IntegrateLoadMatch(TypedDict):
    id: str


class LogRequired(TypedDict):
    expression: str
    operation: str
    result: str


class Log(LogRequired, total=False):
    id: str


class LogLoadMatch(TypedDict):
    id: str


class SimplifyRequired(TypedDict):
    expression: str
    operation: str
    result: str


class Simplify(SimplifyRequired, total=False):
    id: str


class SimplifyLoadMatch(TypedDict):
    id: str


class SinRequired(TypedDict):
    expression: str
    operation: str
    result: str


class Sin(SinRequired, total=False):
    id: str


class SinLoadMatch(TypedDict):
    id: str


class TanRequired(TypedDict):
    expression: str
    operation: str
    result: str


class Tan(TanRequired, total=False):
    id: str


class TanLoadMatch(TypedDict):
    id: str


class TangentRequired(TypedDict):
    expression: str
    operation: str
    result: str


class Tangent(TangentRequired, total=False):
    id: str


class TangentLoadMatch(TypedDict):
    id: str


class ZeroRequired(TypedDict):
    expression: str
    operation: str
    result: str


class Zero(ZeroRequired, total=False):
    id: str


class ZeroLoadMatch(TypedDict):
    id: str
