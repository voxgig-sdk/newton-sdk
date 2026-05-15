
import { test, describe } from 'node:test'
import { equal } from 'node:assert'


import { NewtonSDK } from '..'


describe('exists', async () => {

  test('test-mode', async () => {
    const testsdk = await NewtonSDK.test()
    equal(null !== testsdk, true)
  })

})
