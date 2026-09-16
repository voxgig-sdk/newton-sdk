# Newton SDK feature factory

require_relative 'feature/base_feature'
require_relative 'feature/ratelimit_feature'
require_relative 'feature/retry_feature'
require_relative 'feature/test_feature'
require_relative 'feature/timeout_feature'


module NewtonFeatures
  def self.make_feature(name)
    case name
    when "base"
      NewtonBaseFeature.new
    when "ratelimit"
      NewtonRatelimitFeature.new
    when "retry"
      NewtonRetryFeature.new
    when "test"
      NewtonTestFeature.new
    when "timeout"
      NewtonTimeoutFeature.new
    else
      NewtonBaseFeature.new
    end
  end
end
