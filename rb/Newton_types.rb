# frozen_string_literal: true

# Typed models for the Newton SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Member types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Ruby types are unenforced; these YARD
# annotations document the shapes. Do not edit by hand.

# Abs entity data model.
#
# @!attribute [rw] expression
#   @return [String]
#
# @!attribute [rw] operation
#   @return [String]
#
# @!attribute [rw] result
#   @return [String]
Abs = Struct.new(
  :expression,
  :operation,
  :result,
  keyword_init: true
)

# Request payload for Abs#load.
#
# @!attribute [rw] id
#   @return [String]
AbsLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Arcco entity data model.
#
# @!attribute [rw] expression
#   @return [String]
#
# @!attribute [rw] operation
#   @return [String]
#
# @!attribute [rw] result
#   @return [String]
Arcco = Struct.new(
  :expression,
  :operation,
  :result,
  keyword_init: true
)

# Request payload for Arcco#load.
#
# @!attribute [rw] id
#   @return [String]
ArccoLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Arcsin entity data model.
#
# @!attribute [rw] expression
#   @return [String]
#
# @!attribute [rw] operation
#   @return [String]
#
# @!attribute [rw] result
#   @return [String]
Arcsin = Struct.new(
  :expression,
  :operation,
  :result,
  keyword_init: true
)

# Request payload for Arcsin#load.
#
# @!attribute [rw] id
#   @return [String]
ArcsinLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Arctan entity data model.
#
# @!attribute [rw] expression
#   @return [String]
#
# @!attribute [rw] operation
#   @return [String]
#
# @!attribute [rw] result
#   @return [String]
Arctan = Struct.new(
  :expression,
  :operation,
  :result,
  keyword_init: true
)

# Request payload for Arctan#load.
#
# @!attribute [rw] id
#   @return [String]
ArctanLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Area entity data model.
#
# @!attribute [rw] expression
#   @return [String]
#
# @!attribute [rw] operation
#   @return [String]
#
# @!attribute [rw] result
#   @return [String]
Area = Struct.new(
  :expression,
  :operation,
  :result,
  keyword_init: true
)

# Request payload for Area#load.
#
# @!attribute [rw] id
#   @return [String]
AreaLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Cos entity data model.
#
# @!attribute [rw] expression
#   @return [String]
#
# @!attribute [rw] operation
#   @return [String]
#
# @!attribute [rw] result
#   @return [String]
Cos = Struct.new(
  :expression,
  :operation,
  :result,
  keyword_init: true
)

# Request payload for Cos#load.
#
# @!attribute [rw] id
#   @return [String]
CosLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Derive entity data model.
#
# @!attribute [rw] expression
#   @return [String]
#
# @!attribute [rw] operation
#   @return [String]
#
# @!attribute [rw] result
#   @return [String]
Derive = Struct.new(
  :expression,
  :operation,
  :result,
  keyword_init: true
)

# Request payload for Derive#load.
#
# @!attribute [rw] id
#   @return [String]
DeriveLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Factor entity data model.
#
# @!attribute [rw] expression
#   @return [String]
#
# @!attribute [rw] operation
#   @return [String]
#
# @!attribute [rw] result
#   @return [String]
Factor = Struct.new(
  :expression,
  :operation,
  :result,
  keyword_init: true
)

# Request payload for Factor#load.
#
# @!attribute [rw] id
#   @return [String]
FactorLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Integrate entity data model.
#
# @!attribute [rw] expression
#   @return [String]
#
# @!attribute [rw] operation
#   @return [String]
#
# @!attribute [rw] result
#   @return [String]
Integrate = Struct.new(
  :expression,
  :operation,
  :result,
  keyword_init: true
)

# Request payload for Integrate#load.
#
# @!attribute [rw] id
#   @return [String]
IntegrateLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Log entity data model.
#
# @!attribute [rw] expression
#   @return [String]
#
# @!attribute [rw] operation
#   @return [String]
#
# @!attribute [rw] result
#   @return [String]
Log = Struct.new(
  :expression,
  :operation,
  :result,
  keyword_init: true
)

# Request payload for Log#load.
#
# @!attribute [rw] id
#   @return [String]
LogLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Simplify entity data model.
#
# @!attribute [rw] expression
#   @return [String]
#
# @!attribute [rw] operation
#   @return [String]
#
# @!attribute [rw] result
#   @return [String]
Simplify = Struct.new(
  :expression,
  :operation,
  :result,
  keyword_init: true
)

# Request payload for Simplify#load.
#
# @!attribute [rw] id
#   @return [String]
SimplifyLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Sin entity data model.
#
# @!attribute [rw] expression
#   @return [String]
#
# @!attribute [rw] operation
#   @return [String]
#
# @!attribute [rw] result
#   @return [String]
Sin = Struct.new(
  :expression,
  :operation,
  :result,
  keyword_init: true
)

# Request payload for Sin#load.
#
# @!attribute [rw] id
#   @return [String]
SinLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Tan entity data model.
#
# @!attribute [rw] expression
#   @return [String]
#
# @!attribute [rw] operation
#   @return [String]
#
# @!attribute [rw] result
#   @return [String]
Tan = Struct.new(
  :expression,
  :operation,
  :result,
  keyword_init: true
)

# Request payload for Tan#load.
#
# @!attribute [rw] id
#   @return [String]
TanLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Tangent entity data model.
#
# @!attribute [rw] expression
#   @return [String]
#
# @!attribute [rw] operation
#   @return [String]
#
# @!attribute [rw] result
#   @return [String]
Tangent = Struct.new(
  :expression,
  :operation,
  :result,
  keyword_init: true
)

# Request payload for Tangent#load.
#
# @!attribute [rw] id
#   @return [String]
TangentLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Zero entity data model.
#
# @!attribute [rw] expression
#   @return [String]
#
# @!attribute [rw] operation
#   @return [String]
#
# @!attribute [rw] result
#   @return [String]
Zero = Struct.new(
  :expression,
  :operation,
  :result,
  keyword_init: true
)

# Request payload for Zero#load.
#
# @!attribute [rw] id
#   @return [String]
ZeroLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

