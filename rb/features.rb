# Newton SDK feature factory

require_relative 'feature/base_feature'
require_relative 'feature/test_feature'


module NewtonFeatures
  def self.make_feature(name)
    case name
    when "base"
      NewtonBaseFeature.new
    when "test"
      NewtonTestFeature.new
    else
      NewtonBaseFeature.new
    end
  end
end
