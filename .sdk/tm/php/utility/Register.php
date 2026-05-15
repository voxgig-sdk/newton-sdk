<?php
declare(strict_types=1);

// Newton SDK utility registration

require_once __DIR__ . '/../core/UtilityType.php';
require_once __DIR__ . '/Clean.php';
require_once __DIR__ . '/Done.php';
require_once __DIR__ . '/MakeError.php';
require_once __DIR__ . '/FeatureAdd.php';
require_once __DIR__ . '/FeatureHook.php';
require_once __DIR__ . '/FeatureInit.php';
require_once __DIR__ . '/Fetcher.php';
require_once __DIR__ . '/MakeFetchDef.php';
require_once __DIR__ . '/MakeContext.php';
require_once __DIR__ . '/MakeOptions.php';
require_once __DIR__ . '/MakeRequest.php';
require_once __DIR__ . '/MakeResponse.php';
require_once __DIR__ . '/MakeResult.php';
require_once __DIR__ . '/MakePoint.php';
require_once __DIR__ . '/MakeSpec.php';
require_once __DIR__ . '/MakeUrl.php';
require_once __DIR__ . '/Param.php';
require_once __DIR__ . '/PrepareAuth.php';
require_once __DIR__ . '/PrepareBody.php';
require_once __DIR__ . '/PrepareHeaders.php';
require_once __DIR__ . '/PrepareMethod.php';
require_once __DIR__ . '/PrepareParams.php';
require_once __DIR__ . '/PreparePath.php';
require_once __DIR__ . '/PrepareQuery.php';
require_once __DIR__ . '/ResultBasic.php';
require_once __DIR__ . '/ResultBody.php';
require_once __DIR__ . '/ResultHeaders.php';
require_once __DIR__ . '/TransformRequest.php';
require_once __DIR__ . '/TransformResponse.php';

NewtonUtility::setRegistrar(function (NewtonUtility $u): void {
    $u->clean = [NewtonClean::class, 'call'];
    $u->done = [NewtonDone::class, 'call'];
    $u->make_error = [NewtonMakeError::class, 'call'];
    $u->feature_add = [NewtonFeatureAdd::class, 'call'];
    $u->feature_hook = [NewtonFeatureHook::class, 'call'];
    $u->feature_init = [NewtonFeatureInit::class, 'call'];
    $u->fetcher = [NewtonFetcher::class, 'call'];
    $u->make_fetch_def = [NewtonMakeFetchDef::class, 'call'];
    $u->make_context = [NewtonMakeContext::class, 'call'];
    $u->make_options = [NewtonMakeOptions::class, 'call'];
    $u->make_request = [NewtonMakeRequest::class, 'call'];
    $u->make_response = [NewtonMakeResponse::class, 'call'];
    $u->make_result = [NewtonMakeResult::class, 'call'];
    $u->make_point = [NewtonMakePoint::class, 'call'];
    $u->make_spec = [NewtonMakeSpec::class, 'call'];
    $u->make_url = [NewtonMakeUrl::class, 'call'];
    $u->param = [NewtonParam::class, 'call'];
    $u->prepare_auth = [NewtonPrepareAuth::class, 'call'];
    $u->prepare_body = [NewtonPrepareBody::class, 'call'];
    $u->prepare_headers = [NewtonPrepareHeaders::class, 'call'];
    $u->prepare_method = [NewtonPrepareMethod::class, 'call'];
    $u->prepare_params = [NewtonPrepareParams::class, 'call'];
    $u->prepare_path = [NewtonPreparePath::class, 'call'];
    $u->prepare_query = [NewtonPrepareQuery::class, 'call'];
    $u->result_basic = [NewtonResultBasic::class, 'call'];
    $u->result_body = [NewtonResultBody::class, 'call'];
    $u->result_headers = [NewtonResultHeaders::class, 'call'];
    $u->transform_request = [NewtonTransformRequest::class, 'call'];
    $u->transform_response = [NewtonTransformResponse::class, 'call'];
});
