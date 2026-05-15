# ProjectName SDK exists test

import pytest
from newton_sdk import NewtonSDK


class TestExists:

    def test_should_create_test_sdk(self):
        testsdk = NewtonSDK.test(None, None)
        assert testsdk is not None
