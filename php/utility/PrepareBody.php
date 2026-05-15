<?php
declare(strict_types=1);

// Newton SDK utility: prepare_body

class NewtonPrepareBody
{
    public static function call(NewtonContext $ctx): mixed
    {
        if ($ctx->op->input === 'data') {
            return ($ctx->utility->transform_request)($ctx);
        }
        return null;
    }
}
