<?php
declare(strict_types=1);

// Newton SDK utility: result_body

class NewtonResultBody
{
    public static function call(NewtonContext $ctx): ?NewtonResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result && $response && $response->json_func && $response->body) {
            $result->body = ($response->json_func)();
        }
        return $result;
    }
}
