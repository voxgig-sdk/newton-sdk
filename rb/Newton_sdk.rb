# Newton SDK

require_relative 'utility/struct/voxgig_struct'
require_relative 'core/utility_type'
require_relative 'core/spec'
require_relative 'core/helpers'

# Load utility registration
require_relative 'utility/register'

# Load config and features
require_relative 'config'
require_relative 'feature/base_feature'
require_relative 'features'

# Load typed models (Struct value objects).
require_relative 'Newton_types'


class NewtonSDK
  attr_accessor :mode, :features, :options

  def initialize(options = {})
    @mode = "live"
    @features = []
    @options = nil

    utility = NewtonUtility.new
    @_utility = utility

    config = NewtonConfig.make_config

    @_rootctx = utility.make_context.call({
      "client" => self,
      "utility" => utility,
      "config" => config,
      "options" => options || {},
      "shared" => {},
    }, nil)

    @options = utility.make_options.call(@_rootctx)

    if VoxgigStruct.getpath(@options, "feature.test.active") == true
      @mode = "test"
    end

    @_rootctx.options = @options

    # Add features from config.
    feature_opts = NewtonHelpers.to_map(VoxgigStruct.getprop(@options, "feature"))
    if feature_opts
      items = VoxgigStruct.items(feature_opts)
      if items
        items.each do |item|
          fname = item[0]
          fopts = NewtonHelpers.to_map(item[1])
          if fopts && fopts["active"] == true
            utility.feature_add.call(@_rootctx, NewtonFeatures.make_feature(fname))
          end
        end
      end
    end

    # Add extension features.
    extend_val = VoxgigStruct.getprop(@options, "extend")
    if extend_val.is_a?(Array)
      extend_val.each do |f|
        if f.respond_to?(:get_name)
          utility.feature_add.call(@_rootctx, f)
        end
      end
    end

    # Initialize features.
    @features.each do |f|
      utility.feature_init.call(@_rootctx, f)
    end

    utility.feature_hook.call(@_rootctx, "PostConstruct")
  end

  def options_map
    out = VoxgigStruct.clone(@options)
    out.is_a?(Hash) ? out : {}
  end

  def get_utility
    NewtonUtility.copy(@_utility)
  end

  def get_root_ctx
    @_rootctx
  end

  def prepare(fetchargs = {})
    utility = @_utility
    fetchargs ||= {}

    ctrl = NewtonHelpers.to_map(VoxgigStruct.getprop(fetchargs, "ctrl")) || {}

    ctx = utility.make_context.call({
      "opname" => "prepare",
      "ctrl" => ctrl,
    }, @_rootctx)

    opts = @options
    path = VoxgigStruct.getprop(fetchargs, "path") || ""
    path = "" unless path.is_a?(String)
    method_val = VoxgigStruct.getprop(fetchargs, "method") || "GET"
    method_val = "GET" unless method_val.is_a?(String)
    params = NewtonHelpers.to_map(VoxgigStruct.getprop(fetchargs, "params")) || {}
    query = NewtonHelpers.to_map(VoxgigStruct.getprop(fetchargs, "query")) || {}
    headers = utility.prepare_headers.call(ctx)

    base = VoxgigStruct.getprop(opts, "base") || ""
    base = "" unless base.is_a?(String)
    prefix = VoxgigStruct.getprop(opts, "prefix") || ""
    prefix = "" unless prefix.is_a?(String)
    suffix = VoxgigStruct.getprop(opts, "suffix") || ""
    suffix = "" unless suffix.is_a?(String)

    ctx.spec = NewtonSpec.new({
      "base" => base, "prefix" => prefix, "suffix" => suffix,
      "path" => path, "method" => method_val,
      "params" => params, "query" => query, "headers" => headers,
      "body" => VoxgigStruct.getprop(fetchargs, "body"),
      "step" => "start",
    })

    # Merge user-provided headers.
    uh = VoxgigStruct.getprop(fetchargs, "headers")
    if uh.is_a?(Hash)
      uh.each { |k, v| ctx.spec.headers[k] = v }
    end

    _, err = utility.prepare_auth.call(ctx)
    raise err if err

    utility.make_fetch_def.call(ctx)
  end

  def direct(fetchargs = {})
    utility = @_utility

    # direct() is the raw-HTTP escape hatch: it always returns a result hash
    # ({ "ok" => ..., ... }) and never raises. prepare() raises on error, so
    # trap that and surface it in the hash.
    begin
      fetchdef = prepare(fetchargs)
    rescue NewtonError => err
      return { "ok" => false, "err" => err }
    end

    fetchargs ||= {}
    ctrl = NewtonHelpers.to_map(VoxgigStruct.getprop(fetchargs, "ctrl")) || {}

    ctx = utility.make_context.call({
      "opname" => "direct",
      "ctrl" => ctrl,
    }, @_rootctx)

    url = fetchdef["url"] || ""
    fetched, fetch_err = utility.fetcher.call(ctx, url, fetchdef)

    return { "ok" => false, "err" => fetch_err } if fetch_err

    if fetched.nil?
      return {
        "ok" => false,
        "err" => ctx.make_error("direct_no_response", "response: undefined"),
      }
    end

    if fetched.is_a?(Hash)
      status = NewtonHelpers.to_int(VoxgigStruct.getprop(fetched, "status"))
      headers = VoxgigStruct.getprop(fetched, "headers") || {}

      # No-body responses (204, 304) and explicit zero content-length must
      # skip JSON parsing — calling json() on an empty body errors.
      content_length = headers.is_a?(Hash) ? headers["content-length"] : nil
      no_body = status == 204 || status == 304 || content_length.to_s == "0"

      json_data = nil
      unless no_body
        jf = VoxgigStruct.getprop(fetched, "json")
        if jf.is_a?(Proc)
          begin
            json_data = jf.call
          rescue StandardError
            # Non-JSON body — leave data nil, keep status/headers.
            json_data = nil
          end
        end
      end

      return {
        "ok" => status >= 200 && status < 300,
        "status" => status,
        "headers" => headers,
        "data" => json_data,
      }
    end

    return {
      "ok" => false,
      "err" => ctx.make_error("direct_invalid", "invalid response type"),
    }
  end


  # Idiomatic facade: client.abs.list / client.abs.load({ "id" => ... })
  def abs
    require_relative 'entity/abs_entity'
    @abs ||= AbsEntity.new(self, nil)
  end

  # Deprecated: use client.abs instead.
  def Abs(data = nil)
    require_relative 'entity/abs_entity'
    AbsEntity.new(self, data)
  end


  # Idiomatic facade: client.arcco.list / client.arcco.load({ "id" => ... })
  def arcco
    require_relative 'entity/arcco_entity'
    @arcco ||= ArccoEntity.new(self, nil)
  end

  # Deprecated: use client.arcco instead.
  def Arcco(data = nil)
    require_relative 'entity/arcco_entity'
    ArccoEntity.new(self, data)
  end


  # Idiomatic facade: client.arcsin.list / client.arcsin.load({ "id" => ... })
  def arcsin
    require_relative 'entity/arcsin_entity'
    @arcsin ||= ArcsinEntity.new(self, nil)
  end

  # Deprecated: use client.arcsin instead.
  def Arcsin(data = nil)
    require_relative 'entity/arcsin_entity'
    ArcsinEntity.new(self, data)
  end


  # Idiomatic facade: client.arctan.list / client.arctan.load({ "id" => ... })
  def arctan
    require_relative 'entity/arctan_entity'
    @arctan ||= ArctanEntity.new(self, nil)
  end

  # Deprecated: use client.arctan instead.
  def Arctan(data = nil)
    require_relative 'entity/arctan_entity'
    ArctanEntity.new(self, data)
  end


  # Idiomatic facade: client.area.list / client.area.load({ "id" => ... })
  def area
    require_relative 'entity/area_entity'
    @area ||= AreaEntity.new(self, nil)
  end

  # Deprecated: use client.area instead.
  def Area(data = nil)
    require_relative 'entity/area_entity'
    AreaEntity.new(self, data)
  end


  # Idiomatic facade: client.cos.list / client.cos.load({ "id" => ... })
  def cos
    require_relative 'entity/cos_entity'
    @cos ||= CosEntity.new(self, nil)
  end

  # Deprecated: use client.cos instead.
  def Cos(data = nil)
    require_relative 'entity/cos_entity'
    CosEntity.new(self, data)
  end


  # Idiomatic facade: client.derive.list / client.derive.load({ "id" => ... })
  def derive
    require_relative 'entity/derive_entity'
    @derive ||= DeriveEntity.new(self, nil)
  end

  # Deprecated: use client.derive instead.
  def Derive(data = nil)
    require_relative 'entity/derive_entity'
    DeriveEntity.new(self, data)
  end


  # Idiomatic facade: client.factor.list / client.factor.load({ "id" => ... })
  def factor
    require_relative 'entity/factor_entity'
    @factor ||= FactorEntity.new(self, nil)
  end

  # Deprecated: use client.factor instead.
  def Factor(data = nil)
    require_relative 'entity/factor_entity'
    FactorEntity.new(self, data)
  end


  # Idiomatic facade: client.integrate.list / client.integrate.load({ "id" => ... })
  def integrate
    require_relative 'entity/integrate_entity'
    @integrate ||= IntegrateEntity.new(self, nil)
  end

  # Deprecated: use client.integrate instead.
  def Integrate(data = nil)
    require_relative 'entity/integrate_entity'
    IntegrateEntity.new(self, data)
  end


  # Idiomatic facade: client.log.list / client.log.load({ "id" => ... })
  def log
    require_relative 'entity/log_entity'
    @log ||= LogEntity.new(self, nil)
  end

  # Deprecated: use client.log instead.
  def Log(data = nil)
    require_relative 'entity/log_entity'
    LogEntity.new(self, data)
  end


  # Idiomatic facade: client.simplify.list / client.simplify.load({ "id" => ... })
  def simplify
    require_relative 'entity/simplify_entity'
    @simplify ||= SimplifyEntity.new(self, nil)
  end

  # Deprecated: use client.simplify instead.
  def Simplify(data = nil)
    require_relative 'entity/simplify_entity'
    SimplifyEntity.new(self, data)
  end


  # Idiomatic facade: client.sin.list / client.sin.load({ "id" => ... })
  def sin
    require_relative 'entity/sin_entity'
    @sin ||= SinEntity.new(self, nil)
  end

  # Deprecated: use client.sin instead.
  def Sin(data = nil)
    require_relative 'entity/sin_entity'
    SinEntity.new(self, data)
  end


  # Idiomatic facade: client.tan.list / client.tan.load({ "id" => ... })
  def tan
    require_relative 'entity/tan_entity'
    @tan ||= TanEntity.new(self, nil)
  end

  # Deprecated: use client.tan instead.
  def Tan(data = nil)
    require_relative 'entity/tan_entity'
    TanEntity.new(self, data)
  end


  # Idiomatic facade: client.tangent.list / client.tangent.load({ "id" => ... })
  def tangent
    require_relative 'entity/tangent_entity'
    @tangent ||= TangentEntity.new(self, nil)
  end

  # Deprecated: use client.tangent instead.
  def Tangent(data = nil)
    require_relative 'entity/tangent_entity'
    TangentEntity.new(self, data)
  end


  # Idiomatic facade: client.zero.list / client.zero.load({ "id" => ... })
  def zero
    require_relative 'entity/zero_entity'
    @zero ||= ZeroEntity.new(self, nil)
  end

  # Deprecated: use client.zero instead.
  def Zero(data = nil)
    require_relative 'entity/zero_entity'
    ZeroEntity.new(self, data)
  end



  def self.test(testopts = nil, sdkopts = nil)
    sdkopts = sdkopts || {}
    sdkopts = VoxgigStruct.clone(sdkopts)
    sdkopts = {} unless sdkopts.is_a?(Hash)

    testopts = testopts || {}
    testopts = VoxgigStruct.clone(testopts)
    testopts = {} unless testopts.is_a?(Hash)
    testopts["active"] = true

    VoxgigStruct.setpath(sdkopts, "feature.test", testopts)

    sdk = NewtonSDK.new(sdkopts)
    sdk.mode = "test"
    sdk
  end
end
