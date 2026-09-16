

import Path from 'node:path'
import * as Fs from 'node:fs'

import { test, describe, afterEach } from 'node:test'
import assert from 'node:assert'
import { createLiveTransport } from '../../live-runner'
import { runLiveEntity } from '../../live-entity'


import { NewtonSDK, BaseFeature, stdutil } from '../../..'

import {
  envOverride,
  liveClientOptions,
  liveDelay,
  loadEnvLocal,
  makeCtrl,
  makeMatch,
  makeReqdata,
  makeStepData,
  makeValid,
  maybeSkipControl,
} from '../../utility'


// AFTER the imports on purpose: TypeScript hoists `import` above any
// statement in the emitted CommonJS, so a loader placed above them would
// run only after every imported module had already been evaluated - and
// anything reading process.env at module scope would miss these values.
loadEnvLocal(__dirname + '/../../../.env.local')


describe('CosEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when NEWTON_TEST_LIVE=TRUE.
  afterEach(liveDelay('NEWTON_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = NewtonSDK.test()
    const ent = testsdk.Cos()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.NEWTON_TEST_LIVE
    for (const op of ['load']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'cos.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":[{"active":true,"name":"expression","req":true,"short":"The mathematical expression that was processed","type":"`$STRING`","index$":0},{"active":true,"name":"id","req":false,"type":"`$STRING`","index$":1},{"active":true,"name":"operation","req":true,"short":"The mathematical operation that was performed","type":"`$STRING`","index$":2},{"active":true,"name":"result","req":true,"short":"The result of the mathematical operation","type":"`$STRING`","index$":3}],"id":{"field":"id","name":"id"},"name":"cos","op":{"load":{"input":"data","name":"load","points":[{"active":true,"args":{"params":[{"active":true,"example":"pi","kind":"param","name":"id","orig":"expression","reqd":true,"type":"`$STRING`","index$":0}]},"contract":{"id":"GET /cos/{expression}","json":"{\"operationId\":\"cosine\",\"parameters\":[{\"description\":\"URL-encoded mathematical expression for cosine calculation (e.g., pi)\",\"example\":\"pi\",\"in\":\"path\",\"name\":\"expression\",\"required\":true,\"schema\":{\"type\":\"string\"}}],\"protocol\":\"http\",\"responses\":{\"200\":{\"content\":{\"application/json\":{\"example\":{\"expression\":\"pi\",\"operation\":\"cos\",\"result\":\"-1\"},\"schema\":{\"properties\":{\"expression\":{\"description\":\"The mathematical expression that was processed\",\"example\":\"x^2\",\"type\":\"string\"},\"operation\":{\"description\":\"The mathematical operation that was performed\",\"example\":\"derive\",\"type\":\"string\"},\"result\":{\"description\":\"The result of the mathematical operation\",\"example\":\"2 x\",\"type\":\"string\"}},\"required\":[\"operation\",\"expression\",\"result\"],\"type\":\"object\"}}},\"description\":\"Successful operation\"}},\"securitySource\":\"unspecified\"}","source":"openapi3","version":1},"kind":"http","method":"GET","orig":"/cos/{expression}","rename":{"param":{"expression":"id"}},"segments":[{"lit":"cos"},{"var":"id"}],"select":{"exist":["id"]},"transform":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"load"}},"relations":{"ancestors":[]},"key$":"cos","name__orig":"cos","Name":"Cos","name_":"cos","name-":"cos","NAME":"COS","index$":5}, {"active":true,"entity":"cos","key$":"BasicCosFlow","kind":"basic","name":"BasicCosFlow","param":{},"step":[{"active":true,"data":{},"input":{"ref":"cos_ref01","srcdatavar":"cos_ref01_data","suffix":"_dt0"},"match":{"id":"cos01"},"op":"load","spec":[],"valid":[{"apply":"TextFieldMark","def":{"mark":"Mark01-cos_ref01"}}],"index$":0}]}, 'Cos')
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let cos_ref01_data = Object.values(setup.data.existing.cos)[0] as any

    // LOAD
    const cos_ref01_ent = client.Cos()
    const cos_ref01_match_dt0: any = {}
    cos_ref01_match_dt0.id = cos_ref01_data.id
    const cos_ref01_data_dt0 = (await cos_ref01_ent.load(cos_ref01_match_dt0)).data()
    assert(cos_ref01_data_dt0.id === cos_ref01_data.id)


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/cos/CosTestData.json')

  // TODO: file ready util needed?
  const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8')

  // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
  const entityData = JSON.parse(entityDataSource)

  options.entity = entityData.existing

  let client = NewtonSDK.test(options, extra)
  const struct = client.utility().struct
  const merge = struct.merge
  const transform = struct.transform

  let idmap = transform(
    ['cos01','cos02','cos03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'NEWTON_TEST_COS_ENTID': idmap,
    'NEWTON_TEST_LIVE': 'FALSE',
    'NEWTON_TEST_EXPLAIN': 'FALSE',
  })

  idmap = env['NEWTON_TEST_COS_ENTID']

  const live = 'TRUE' === env.NEWTON_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['NEWTON_TEST_COS_ENTID']
    idmap = rawIds && rawIds.trim() ? JSON.parse(rawIds) : {}
    if (!idmap || Array.isArray(idmap) || typeof idmap !== 'object') {
      throw new Error('Live ENTID must be a JSON object')
    }
    client = new NewtonSDK(merge([
      // FIRST, so the generated fields below win: sdk-test-control.json's
      // test.client.options adds to the live client, it does not redirect it.
      liveClientOptions(),
      {
      },
      // 'extra || {}', not a bare 'extra': struct.merge returns UNDEFINED when the
      // last entry is undefined, and basicSetup is normally called with no
      // argument at all - so a bare 'extra' silently discarded the apikey
      // and server values above and handed the SDK undefined. Harmless
      // while there was nothing in that object; not harmless now.
      extra || {},
      { system: { fetch: transport.fetch } }
    ]))
  }

  const setup = {
    idmap,
    env,
    options,
    client,
    struct,
    data: entityData,
    explain: 'TRUE' === env.NEWTON_TEST_EXPLAIN,
    live,
    transport,
    now: Date.now(),
  }

  return setup
}
  
