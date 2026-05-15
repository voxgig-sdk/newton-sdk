# Newton SDK exists test

require "minitest/autorun"
require_relative "../Newton_sdk"

class ExistsTest < Minitest::Test
  def test_create_test_sdk
    testsdk = NewtonSDK.test(nil, nil)
    assert !testsdk.nil?
  end
end
