<?php
declare(strict_types=1);

// Newton SDK feature factory

require_once __DIR__ . '/feature/BaseFeature.php';
require_once __DIR__ . '/feature/TestFeature.php';


class NewtonFeatures
{
    public static function make_feature(string $name)
    {
        switch ($name) {
            case "base":
                return new NewtonBaseFeature();
            case "test":
                return new NewtonTestFeature();
            default:
                return new NewtonBaseFeature();
        }
    }
}
