<?php
declare(strict_types=1);

// Derive entity test

require_once __DIR__ . '/../newton_sdk.php';
require_once __DIR__ . '/Runner.php';

use PHPUnit\Framework\TestCase;
use Voxgig\Struct\Struct as Vs;

class DeriveEntityTest extends TestCase
{
    public function test_create_instance(): void
    {
        $testsdk = NewtonSDK::test(null, null);
        $ent = $testsdk->Derive(null);
        $this->assertNotNull($ent);
    }

    public function test_basic_flow(): void
    {
        $setup = derive_basic_setup(null);
        // Per-op sdk-test-control.json skip.
        $_live = !empty($setup["live"]);
        foreach (["load"] as $_op) {
            [$_shouldSkip, $_reason] = Runner::is_control_skipped("entityOp", "derive." . $_op, $_live ? "live" : "unit");
            if ($_shouldSkip) {
                $this->markTestSkipped($_reason ?? "skipped via sdk-test-control.json");
                return;
            }
        }
        // The basic flow consumes synthetic IDs from the fixture. In live mode
        // without an *_ENTID env override, those IDs hit the live API and 4xx.
        if (!empty($setup["synthetic_only"])) {
            $this->markTestSkipped("live entity test uses synthetic IDs from fixture — set NEWTON_TEST_DERIVE_ENTID JSON to run live");
            return;
        }
        $client = $setup["client"];

        // Bootstrap entity data from existing test data.
        $derive_ref01_data_raw = Vs::items(Helpers::to_map(
            Vs::getpath($setup["data"], "existing.derive")));
        $derive_ref01_data = null;
        if (count($derive_ref01_data_raw) > 0) {
            $derive_ref01_data = Helpers::to_map($derive_ref01_data_raw[0][1]);
        }

        // LOAD
        $derive_ref01_ent = $client->Derive(null);
        $derive_ref01_match_dt0 = [
            "id" => $derive_ref01_data["id"],
        ];
        $derive_ref01_data_dt0_loaded = $derive_ref01_ent->load($derive_ref01_match_dt0, null);
        $derive_ref01_data_dt0_load_result = Helpers::to_map(is_object($derive_ref01_data_dt0_loaded) && method_exists($derive_ref01_data_dt0_loaded, 'data_get') ? $derive_ref01_data_dt0_loaded->data_get() : $derive_ref01_data_dt0_loaded);
        $this->assertNotNull($derive_ref01_data_dt0_load_result);
        $this->assertEquals($derive_ref01_data_dt0_load_result["id"], $derive_ref01_data["id"]);

    }
}

function derive_basic_setup($extra)
{
    Runner::load_env_local();

    $entity_data_file = __DIR__ . '/../../.sdk/test/entity/derive/DeriveTestData.json';
    $entity_data_source = file_get_contents($entity_data_file);
    $entity_data = json_decode($entity_data_source, true);

    $options = [];
    $options["entity"] = $entity_data["existing"];

    $client = NewtonSDK::test($options, $extra);

    // Generate idmap.
    $idmap = [];
    foreach (["derive01", "derive02", "derive03"] as $k) {
        $idmap[$k] = strtoupper($k);
    }

    // Detect ENTID env override before envOverride consumes it. When live
    // mode is on without a real override, the basic test runs against synthetic
    // IDs from the fixture and 4xx's. Surface this so the test can skip.
    $entid_env_raw = getenv("NEWTON_TEST_DERIVE_ENTID");
    $idmap_overridden = $entid_env_raw !== false && str_starts_with(trim($entid_env_raw), "{");

    $env = Runner::env_override([
        "NEWTON_TEST_DERIVE_ENTID" => $idmap,
        "NEWTON_TEST_LIVE" => "FALSE",
        "NEWTON_TEST_EXPLAIN" => "FALSE",
    ]);

    $idmap_resolved = Helpers::to_map(
        $env["NEWTON_TEST_DERIVE_ENTID"]);
    if ($idmap_resolved === null) {
        $idmap_resolved = Helpers::to_map($idmap);
    }

    if ($env["NEWTON_TEST_LIVE"] === "TRUE") {
        $merged_opts = Vs::merge([
            [
            ],
            $extra ?? [],
        ]);
        $client = new NewtonSDK(Helpers::to_map($merged_opts));
    }

    $live = $env["NEWTON_TEST_LIVE"] === "TRUE";
    return [
        "client" => $client,
        "data" => $entity_data,
        "idmap" => $idmap_resolved,
        "env" => $env,
        "explain" => $env["NEWTON_TEST_EXPLAIN"] === "TRUE",
        "live" => $live,
        "synthetic_only" => $live && !$idmap_overridden,
        "now" => (int)(microtime(true) * 1000),
    ];
}
