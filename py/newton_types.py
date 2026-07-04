# Typed models for the Newton SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.

from __future__ import annotations

from dataclasses import dataclass
from typing import Optional, Any


@dataclass
class Abs:
    expression: str
    operation: str
    result: str


@dataclass
class AbsLoadMatch:
    id: str


@dataclass
class Arcco:
    expression: str
    operation: str
    result: str


@dataclass
class ArccoLoadMatch:
    id: str


@dataclass
class Arcsin:
    expression: str
    operation: str
    result: str


@dataclass
class ArcsinLoadMatch:
    id: str


@dataclass
class Arctan:
    expression: str
    operation: str
    result: str


@dataclass
class ArctanLoadMatch:
    id: str


@dataclass
class Area:
    expression: str
    operation: str
    result: str


@dataclass
class AreaLoadMatch:
    id: str


@dataclass
class Cos:
    expression: str
    operation: str
    result: str


@dataclass
class CosLoadMatch:
    id: str


@dataclass
class Derive:
    expression: str
    operation: str
    result: str


@dataclass
class DeriveLoadMatch:
    id: str


@dataclass
class Factor:
    expression: str
    operation: str
    result: str


@dataclass
class FactorLoadMatch:
    id: str


@dataclass
class Integrate:
    expression: str
    operation: str
    result: str


@dataclass
class IntegrateLoadMatch:
    id: str


@dataclass
class Log:
    expression: str
    operation: str
    result: str


@dataclass
class LogLoadMatch:
    id: str


@dataclass
class Simplify:
    expression: str
    operation: str
    result: str


@dataclass
class SimplifyLoadMatch:
    id: str


@dataclass
class Sin:
    expression: str
    operation: str
    result: str


@dataclass
class SinLoadMatch:
    id: str


@dataclass
class Tan:
    expression: str
    operation: str
    result: str


@dataclass
class TanLoadMatch:
    id: str


@dataclass
class Tangent:
    expression: str
    operation: str
    result: str


@dataclass
class TangentLoadMatch:
    id: str


@dataclass
class Zero:
    expression: str
    operation: str
    result: str


@dataclass
class ZeroLoadMatch:
    id: str

