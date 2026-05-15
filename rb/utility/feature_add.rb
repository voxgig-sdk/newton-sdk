# Newton SDK utility: feature_add
module NewtonUtilities
  FeatureAdd = ->(ctx, f) {
    ctx.client.features << f
  }
end
