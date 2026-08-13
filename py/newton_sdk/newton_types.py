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


class Abs(TypedDict):
    expression: str
    operation: str
    result: str


class AbsLoadMatch(TypedDict):
    id: str


class Arcco(TypedDict):
    expression: str
    operation: str
    result: str


class ArccoLoadMatch(TypedDict):
    id: str


class Arcsin(TypedDict):
    expression: str
    operation: str
    result: str


class ArcsinLoadMatch(TypedDict):
    id: str


class Arctan(TypedDict):
    expression: str
    operation: str
    result: str


class ArctanLoadMatch(TypedDict):
    id: str


class Area(TypedDict):
    expression: str
    operation: str
    result: str


class AreaLoadMatch(TypedDict):
    id: str


class Cos(TypedDict):
    expression: str
    operation: str
    result: str


class CosLoadMatch(TypedDict):
    id: str


class Derive(TypedDict):
    expression: str
    operation: str
    result: str


class DeriveLoadMatch(TypedDict):
    id: str


class Factor(TypedDict):
    expression: str
    operation: str
    result: str


class FactorLoadMatch(TypedDict):
    id: str


class Integrate(TypedDict):
    expression: str
    operation: str
    result: str


class IntegrateLoadMatch(TypedDict):
    id: str


class Log(TypedDict):
    expression: str
    operation: str
    result: str


class LogLoadMatch(TypedDict):
    id: str


class Simplify(TypedDict):
    expression: str
    operation: str
    result: str


class SimplifyLoadMatch(TypedDict):
    id: str


class Sin(TypedDict):
    expression: str
    operation: str
    result: str


class SinLoadMatch(TypedDict):
    id: str


class Tan(TypedDict):
    expression: str
    operation: str
    result: str


class TanLoadMatch(TypedDict):
    id: str


class Tangent(TypedDict):
    expression: str
    operation: str
    result: str


class TangentLoadMatch(TypedDict):
    id: str


class Zero(TypedDict):
    expression: str
    operation: str
    result: str


class ZeroLoadMatch(TypedDict):
    id: str
