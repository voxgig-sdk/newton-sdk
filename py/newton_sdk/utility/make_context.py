# Newton SDK utility: make_context

from newton_sdk.core.context import NewtonContext


def make_context_util(ctxmap, basectx):
    return NewtonContext(ctxmap, basectx)
