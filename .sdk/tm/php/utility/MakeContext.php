<?php
declare(strict_types=1);

// Newton SDK utility: make_context

require_once __DIR__ . '/../core/Context.php';

class NewtonMakeContext
{
    public static function call(array $ctxmap, ?NewtonContext $basectx): NewtonContext
    {
        return new NewtonContext($ctxmap, $basectx);
    }
}
