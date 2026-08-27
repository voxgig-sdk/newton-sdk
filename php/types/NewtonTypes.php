<?php
declare(strict_types=1);

// Typed models for the Newton SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
//
// These are documentation-grade value objects (PHP 8 typed properties),
// registered on the composer classmap autoload. The SDK boundary exchanges
// assoc-arrays; these classes name the shapes for tooling and typed callers.

/** Abs entity data model. */
class Abs
{
    public string $expression;
    public ?string $id = null;
    public string $operation;
    public string $result;
}

/** Request payload for Abs#load. */
class AbsLoadMatch
{
    public string $id;
}

/** Arcco entity data model. */
class Arcco
{
    public string $expression;
    public ?string $id = null;
    public string $operation;
    public string $result;
}

/** Request payload for Arcco#load. */
class ArccoLoadMatch
{
    public string $id;
}

/** Arcsin entity data model. */
class Arcsin
{
    public string $expression;
    public ?string $id = null;
    public string $operation;
    public string $result;
}

/** Request payload for Arcsin#load. */
class ArcsinLoadMatch
{
    public string $id;
}

/** Arctan entity data model. */
class Arctan
{
    public string $expression;
    public ?string $id = null;
    public string $operation;
    public string $result;
}

/** Request payload for Arctan#load. */
class ArctanLoadMatch
{
    public string $id;
}

/** Area entity data model. */
class Area
{
    public string $expression;
    public ?string $id = null;
    public string $operation;
    public string $result;
}

/** Request payload for Area#load. */
class AreaLoadMatch
{
    public string $id;
}

/** Cos entity data model. */
class Cos
{
    public string $expression;
    public ?string $id = null;
    public string $operation;
    public string $result;
}

/** Request payload for Cos#load. */
class CosLoadMatch
{
    public string $id;
}

/** Derive entity data model. */
class Derive
{
    public string $expression;
    public ?string $id = null;
    public string $operation;
    public string $result;
}

/** Request payload for Derive#load. */
class DeriveLoadMatch
{
    public string $id;
}

/** Factor entity data model. */
class Factor
{
    public string $expression;
    public ?string $id = null;
    public string $operation;
    public string $result;
}

/** Request payload for Factor#load. */
class FactorLoadMatch
{
    public string $id;
}

/** Integrate entity data model. */
class Integrate
{
    public string $expression;
    public ?string $id = null;
    public string $operation;
    public string $result;
}

/** Request payload for Integrate#load. */
class IntegrateLoadMatch
{
    public string $id;
}

/** Log entity data model. */
class Log
{
    public string $expression;
    public ?string $id = null;
    public string $operation;
    public string $result;
}

/** Request payload for Log#load. */
class LogLoadMatch
{
    public string $id;
}

/** Simplify entity data model. */
class Simplify
{
    public string $expression;
    public ?string $id = null;
    public string $operation;
    public string $result;
}

/** Request payload for Simplify#load. */
class SimplifyLoadMatch
{
    public string $id;
}

/** Sin entity data model. */
class Sin
{
    public string $expression;
    public ?string $id = null;
    public string $operation;
    public string $result;
}

/** Request payload for Sin#load. */
class SinLoadMatch
{
    public string $id;
}

/** Tan entity data model. */
class Tan
{
    public string $expression;
    public ?string $id = null;
    public string $operation;
    public string $result;
}

/** Request payload for Tan#load. */
class TanLoadMatch
{
    public string $id;
}

/** Tangent entity data model. */
class Tangent
{
    public string $expression;
    public ?string $id = null;
    public string $operation;
    public string $result;
}

/** Request payload for Tangent#load. */
class TangentLoadMatch
{
    public string $id;
}

/** Zero entity data model. */
class Zero
{
    public string $expression;
    public ?string $id = null;
    public string $operation;
    public string $result;
}

/** Request payload for Zero#load. */
class ZeroLoadMatch
{
    public string $id;
}

