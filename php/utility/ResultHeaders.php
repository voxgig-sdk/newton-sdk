<?php
declare(strict_types=1);

// Newton SDK utility: result_headers

class NewtonResultHeaders
{
    public static function call(NewtonContext $ctx): ?NewtonResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result) {
            if ($response && is_array($response->headers)) {
                $result->headers = $response->headers;
            } else {
                $result->headers = [];
            }
        }
        return $result;
    }
}
