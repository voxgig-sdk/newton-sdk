# Newton SDK utility: make_context
require_relative '../core/context'
module NewtonUtilities
  MakeContext = ->(ctxmap, basectx) {
    NewtonContext.new(ctxmap, basectx)
  }
end
