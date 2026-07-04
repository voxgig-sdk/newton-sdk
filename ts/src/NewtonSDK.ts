// Newton Ts SDK

import { AbsEntity } from './entity/AbsEntity'
import { ArccoEntity } from './entity/ArccoEntity'
import { ArcsinEntity } from './entity/ArcsinEntity'
import { ArctanEntity } from './entity/ArctanEntity'
import { AreaEntity } from './entity/AreaEntity'
import { CosEntity } from './entity/CosEntity'
import { DeriveEntity } from './entity/DeriveEntity'
import { FactorEntity } from './entity/FactorEntity'
import { IntegrateEntity } from './entity/IntegrateEntity'
import { LogEntity } from './entity/LogEntity'
import { SimplifyEntity } from './entity/SimplifyEntity'
import { SinEntity } from './entity/SinEntity'
import { TanEntity } from './entity/TanEntity'
import { TangentEntity } from './entity/TangentEntity'
import { ZeroEntity } from './entity/ZeroEntity'

export type * from './NewtonTypes'


import { inspect } from 'node:util'

import type { Context, Feature } from './types'

import { config } from './Config'
import { NewtonEntityBase } from './NewtonEntityBase'
import { Utility } from './utility/Utility'


import { BaseFeature } from './feature/base/BaseFeature'


const stdutil = new Utility()


class NewtonSDK {
  _mode: string = 'live'
  _options: any
  _utility = new Utility()
  _features: Feature[]
  _rootctx: Context

  constructor(options?: any) {

    this._rootctx = this._utility.makeContext({
      client: this,
      utility: this._utility,
      config,
      options,
      shared: new WeakMap()
    })

    this._options = this._utility.makeOptions(this._rootctx)

    const struct = this._utility.struct
    const getpath = struct.getpath
    const items = struct.items

    if (true === getpath(this._options.feature, 'test.active')) {
      this._mode = 'test'
    }

    this._rootctx.options = this._options

    this._features = []

    const featureAdd = this._utility.featureAdd
    const featureInit = this._utility.featureInit

    items(this._options.feature, (fitem: [string, any]) => {
      const fname = fitem[0]
      const fopts = fitem[1]
      if (fopts.active) {
        featureAdd(this._rootctx, this._rootctx.config.makeFeature(fname))
      }
    })

    if (null != this._options.extend) {
      for (let f of this._options.extend) {
        featureAdd(this._rootctx, f)
      }
    }

    for (let f of this._features) {
      featureInit(this._rootctx, f)
    }

    const featureHook = this._utility.featureHook
    featureHook(this._rootctx, 'PostConstruct')
  }


  options() {
    return this._utility.struct.clone(this._options)
  }


  utility() {
    return this._utility.struct.clone(this._utility)
  }


  async prepare(fetchargs?: any) {
    const utility = this._utility
    const struct = utility.struct
    const clone = struct.clone

    const {
      makeContext,
      makeFetchDef,
      prepareHeaders,
      prepareAuth,
    } = utility

    fetchargs = fetchargs || {}

    let ctx: Context = makeContext({
      opname: 'prepare',
      ctrl: fetchargs.ctrl || {},
    }, this._rootctx)

    const options = this._options

    // Build spec directly from SDK options + user-provided fetch args.
    const spec: any = {
      base: options.base,
      prefix: options.prefix,
      suffix: options.suffix,
      path: fetchargs.path || '',
      method: fetchargs.method || 'GET',
      params: fetchargs.params || {},
      query: fetchargs.query || {},
      headers: prepareHeaders(ctx),
      body: fetchargs.body,
      step: 'start',
    }

    ctx.spec = spec

    // Merge user-provided headers over SDK defaults.
    if (fetchargs.headers) {
      const uheaders = fetchargs.headers
      for (let key in uheaders) {
        spec.headers[key] = uheaders[key]
      }
    }

    // Apply SDK auth (apikey, auth prefix, etc.)
    const authResult = prepareAuth(ctx)
    if (authResult instanceof Error) {
      return authResult
    }

    return makeFetchDef(ctx)
  }


  async direct(fetchargs?: any) {
    const utility = this._utility
    const fetcher = utility.fetcher
    const makeContext = utility.makeContext

    const fetchdef = await this.prepare(fetchargs)
    if (fetchdef instanceof Error) {
      return fetchdef
    }

    let ctx: Context = makeContext({
      opname: 'direct',
      ctrl: (fetchargs || {}).ctrl || {},
    }, this._rootctx)

    try {
      const fetched = await fetcher(ctx, fetchdef.url, fetchdef)

      if (null == fetched) {
        return { ok: false, err: ctx.error('direct_no_response', 'response: undefined') }
      }
      else if (fetched instanceof Error) {
        return { ok: false, err: fetched }
      }

      const status = fetched.status

      // No body responses (204 No Content, 304 Not Modified) and explicit
      // zero content-length must skip JSON parsing — fetched.json() would
      // throw `Unexpected end of JSON input` on an empty body.
      const headers = fetched.headers
      const contentLength = headers && 'function' === typeof headers.get
        ? headers.get('content-length')
        : (headers || {})['content-length']
      const noBody = 204 === status || 304 === status || '0' === String(contentLength)

      let json: any = undefined
      if (!noBody) {
        try {
          json = 'function' === typeof fetched.json ? await fetched.json() : fetched.json
        }
        catch (parseErr) {
          // Body wasn't valid JSON — surface the raw response rather than
          // throwing. data stays undefined; callers can inspect status/headers.
          json = undefined
        }
      }

      return {
        ok: status >= 200 && status < 300,
        status,
        headers: fetched.headers,
        data: json,
      }
    }
    catch (err: any) {
      return { ok: false, err }
    }
  }



  _abs?: AbsEntity

  // Idiomatic facade: `client.abs.list()` / `client.abs.load({ id })`.
  get abs(): AbsEntity {
    return (this._abs ??= new AbsEntity(this, undefined))
  }

  /** @deprecated Use `client.abs` instead. */
  Abs(data?: any) {
    const self = this
    return new AbsEntity(self,data)
  }


  _arcco?: ArccoEntity

  // Idiomatic facade: `client.arcco.list()` / `client.arcco.load({ id })`.
  get arcco(): ArccoEntity {
    return (this._arcco ??= new ArccoEntity(this, undefined))
  }

  /** @deprecated Use `client.arcco` instead. */
  Arcco(data?: any) {
    const self = this
    return new ArccoEntity(self,data)
  }


  _arcsin?: ArcsinEntity

  // Idiomatic facade: `client.arcsin.list()` / `client.arcsin.load({ id })`.
  get arcsin(): ArcsinEntity {
    return (this._arcsin ??= new ArcsinEntity(this, undefined))
  }

  /** @deprecated Use `client.arcsin` instead. */
  Arcsin(data?: any) {
    const self = this
    return new ArcsinEntity(self,data)
  }


  _arctan?: ArctanEntity

  // Idiomatic facade: `client.arctan.list()` / `client.arctan.load({ id })`.
  get arctan(): ArctanEntity {
    return (this._arctan ??= new ArctanEntity(this, undefined))
  }

  /** @deprecated Use `client.arctan` instead. */
  Arctan(data?: any) {
    const self = this
    return new ArctanEntity(self,data)
  }


  _area?: AreaEntity

  // Idiomatic facade: `client.area.list()` / `client.area.load({ id })`.
  get area(): AreaEntity {
    return (this._area ??= new AreaEntity(this, undefined))
  }

  /** @deprecated Use `client.area` instead. */
  Area(data?: any) {
    const self = this
    return new AreaEntity(self,data)
  }


  _cos?: CosEntity

  // Idiomatic facade: `client.cos.list()` / `client.cos.load({ id })`.
  get cos(): CosEntity {
    return (this._cos ??= new CosEntity(this, undefined))
  }

  /** @deprecated Use `client.cos` instead. */
  Cos(data?: any) {
    const self = this
    return new CosEntity(self,data)
  }


  _derive?: DeriveEntity

  // Idiomatic facade: `client.derive.list()` / `client.derive.load({ id })`.
  get derive(): DeriveEntity {
    return (this._derive ??= new DeriveEntity(this, undefined))
  }

  /** @deprecated Use `client.derive` instead. */
  Derive(data?: any) {
    const self = this
    return new DeriveEntity(self,data)
  }


  _factor?: FactorEntity

  // Idiomatic facade: `client.factor.list()` / `client.factor.load({ id })`.
  get factor(): FactorEntity {
    return (this._factor ??= new FactorEntity(this, undefined))
  }

  /** @deprecated Use `client.factor` instead. */
  Factor(data?: any) {
    const self = this
    return new FactorEntity(self,data)
  }


  _integrate?: IntegrateEntity

  // Idiomatic facade: `client.integrate.list()` / `client.integrate.load({ id })`.
  get integrate(): IntegrateEntity {
    return (this._integrate ??= new IntegrateEntity(this, undefined))
  }

  /** @deprecated Use `client.integrate` instead. */
  Integrate(data?: any) {
    const self = this
    return new IntegrateEntity(self,data)
  }


  _log?: LogEntity

  // Idiomatic facade: `client.log.list()` / `client.log.load({ id })`.
  get log(): LogEntity {
    return (this._log ??= new LogEntity(this, undefined))
  }

  /** @deprecated Use `client.log` instead. */
  Log(data?: any) {
    const self = this
    return new LogEntity(self,data)
  }


  _simplify?: SimplifyEntity

  // Idiomatic facade: `client.simplify.list()` / `client.simplify.load({ id })`.
  get simplify(): SimplifyEntity {
    return (this._simplify ??= new SimplifyEntity(this, undefined))
  }

  /** @deprecated Use `client.simplify` instead. */
  Simplify(data?: any) {
    const self = this
    return new SimplifyEntity(self,data)
  }


  _sin?: SinEntity

  // Idiomatic facade: `client.sin.list()` / `client.sin.load({ id })`.
  get sin(): SinEntity {
    return (this._sin ??= new SinEntity(this, undefined))
  }

  /** @deprecated Use `client.sin` instead. */
  Sin(data?: any) {
    const self = this
    return new SinEntity(self,data)
  }


  _tan?: TanEntity

  // Idiomatic facade: `client.tan.list()` / `client.tan.load({ id })`.
  get tan(): TanEntity {
    return (this._tan ??= new TanEntity(this, undefined))
  }

  /** @deprecated Use `client.tan` instead. */
  Tan(data?: any) {
    const self = this
    return new TanEntity(self,data)
  }


  _tangent?: TangentEntity

  // Idiomatic facade: `client.tangent.list()` / `client.tangent.load({ id })`.
  get tangent(): TangentEntity {
    return (this._tangent ??= new TangentEntity(this, undefined))
  }

  /** @deprecated Use `client.tangent` instead. */
  Tangent(data?: any) {
    const self = this
    return new TangentEntity(self,data)
  }


  _zero?: ZeroEntity

  // Idiomatic facade: `client.zero.list()` / `client.zero.load({ id })`.
  get zero(): ZeroEntity {
    return (this._zero ??= new ZeroEntity(this, undefined))
  }

  /** @deprecated Use `client.zero` instead. */
  Zero(data?: any) {
    const self = this
    return new ZeroEntity(self,data)
  }




  static test(testoptsarg?: any, sdkoptsarg?: any) {
    const struct = stdutil.struct
    const setpath = struct.setpath
    const getdef = struct.getdef
    const clone = struct.clone
    const setprop = struct.setprop

    const sdkopts = getdef(clone(sdkoptsarg), {})
    const testopts = getdef(clone(testoptsarg), {})
    setprop(testopts, 'active', true)
    setpath(sdkopts, 'feature.test', testopts)

    const testsdk = new NewtonSDK(sdkopts)
    testsdk._mode = 'test'

    return testsdk
  }


  tester(testopts?: any, sdkopts?: any) {
    return NewtonSDK.test(testopts, sdkopts)
  }


  toJSON() {
    return { name: 'Newton' }
  }

  toString() {
    return 'Newton ' + this._utility.struct.jsonify(this.toJSON())
  }

  [inspect.custom]() {
    return this.toString()
  }

}




const SDK = NewtonSDK


export {
  stdutil,

  BaseFeature,
  NewtonEntityBase,

  NewtonSDK,
  SDK,
}


