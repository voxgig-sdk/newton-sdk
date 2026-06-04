# Sin entity test

require "minitest/autorun"
require "json"
require_relative "../Newton_sdk"
require_relative "runner"

class SinEntityTest < Minitest::Test
  def test_create_instance
    testsdk = NewtonSDK.test(nil, nil)
    ent = testsdk.Sin(nil)
    assert !ent.nil?
  end

  def test_basic_flow
    setup = sin_basic_setup(nil)
    # Per-op sdk-test-control.json skip.
    _live = setup[:live] || false
    ["load"].each do |_op|
      _should_skip, _reason = Runner.is_control_skipped("entityOp", "sin." + _op, _live ? "live" : "unit")
      if _should_skip
        skip(_reason || "skipped via sdk-test-control.json")
        return
      end
    end
    # The basic flow consumes synthetic IDs from the fixture. In live mode
    # without an *_ENTID env override, those IDs hit the live API and 4xx.
    if setup[:synthetic_only]
      skip "live entity test uses synthetic IDs from fixture — set NEWTON_TEST_SIN_ENTID JSON to run live"
      return
    end
    client = setup[:client]

    # Bootstrap entity data from existing test data.
    sin_ref01_data_raw = Vs.items(Helpers.to_map(
      Vs.getpath(setup[:data], "existing.sin")))
    sin_ref01_data = nil
    if sin_ref01_data_raw.length > 0
      sin_ref01_data = Helpers.to_map(sin_ref01_data_raw[0][1])
    end

    # LOAD
    sin_ref01_ent = client.Sin(nil)
    sin_ref01_match_dt0 = {}
    sin_ref01_data_dt0_loaded, err = sin_ref01_ent.load(sin_ref01_match_dt0, nil)
    assert_nil err
    assert !sin_ref01_data_dt0_loaded.nil?

  end
end

def sin_basic_setup(extra)
  Runner.load_env_local

  entity_data_file = File.join(__dir__, "..", "..", ".sdk", "test", "entity", "sin", "SinTestData.json")
  entity_data_source = File.read(entity_data_file)
  entity_data = JSON.parse(entity_data_source)

  options = {}
  options["entity"] = entity_data["existing"]

  client = NewtonSDK.test(options, extra)

  # Generate idmap via transform.
  idmap = Vs.transform(
    ["sin01", "sin02", "sin03"],
    {
      "`$PACK`" => ["", {
        "`$KEY`" => "`$COPY`",
        "`$VAL`" => ["`$FORMAT`", "upper", "`$COPY`"],
      }],
    }
  )

  # Detect ENTID env override before envOverride consumes it. When live
  # mode is on without a real override, the basic test runs against synthetic
  # IDs from the fixture and 4xx's. Surface this so the test can skip.
  entid_env_raw = ENV["NEWTON_TEST_SIN_ENTID"]
  idmap_overridden = !entid_env_raw.nil? && entid_env_raw.strip.start_with?("{")

  env = Runner.env_override({
    "NEWTON_TEST_SIN_ENTID" => idmap,
    "NEWTON_TEST_LIVE" => "FALSE",
    "NEWTON_TEST_EXPLAIN" => "FALSE",
  })

  idmap_resolved = Helpers.to_map(
    env["NEWTON_TEST_SIN_ENTID"])
  if idmap_resolved.nil?
    idmap_resolved = Helpers.to_map(idmap)
  end

  if env["NEWTON_TEST_LIVE"] == "TRUE"
    merged_opts = Vs.merge([
      {
      },
      extra || {},
    ])
    client = NewtonSDK.new(Helpers.to_map(merged_opts))
  end

  live = env["NEWTON_TEST_LIVE"] == "TRUE"
  {
    client: client,
    data: entity_data,
    idmap: idmap_resolved,
    env: env,
    explain: env["NEWTON_TEST_EXPLAIN"] == "TRUE",
    live: live,
    synthetic_only: live && !idmap_overridden,
    now: (Time.now.to_f * 1000).to_i,
  }
end
