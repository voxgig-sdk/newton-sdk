-- Newton SDK error

local NewtonError = {}
NewtonError.__index = NewtonError


function NewtonError.new(code, msg, ctx)
  local self = setmetatable({}, NewtonError)
  self.is_sdk_error = true
  self.sdk = "Newton"
  self.code = code or ""
  self.msg = msg or ""
  self.ctx = ctx
  self.result = nil
  self.spec = nil
  return self
end


function NewtonError:error()
  return self.msg
end


function NewtonError:__tostring()
  return self.msg
end


return NewtonError
