<?php
declare(strict_types=1);

// Newton SDK utility: feature_add

class NewtonFeatureAdd
{
    public static function call(NewtonContext $ctx, mixed $f): void
    {
        $ctx->client->features[] = $f;
    }
}
