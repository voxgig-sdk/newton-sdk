
import { Context } from './Context'


class NewtonError extends Error {

  isNewtonError = true

  sdk = 'Newton'

  code: string
  ctx: Context

  constructor(code: string, msg: string, ctx: Context) {
    super(msg)
    this.code = code
    this.ctx = ctx
  }

}

export {
  NewtonError
}

