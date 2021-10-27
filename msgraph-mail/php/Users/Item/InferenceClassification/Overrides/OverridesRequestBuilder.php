<?php

namespace Microsoft\Graph\Users\Item\InferenceClassification\Overrides;

use Microsoft\Graph\Models\Microsoft\Graph\InferenceClassificationOverride;
use Microsoft\Graph\Users\Item\InferenceClassification\Overrides\Item\InferenceClassificationOverrideRequestBuilder;
use Microsoft\Kiota\Abstractions\HttpMethod;
use Microsoft\Kiota\Abstractions\MiddlewareOption;
use Microsoft\Kiota\Abstractions\RequestAdapter;
use Microsoft\Kiota\Abstractions\RequestInformation;
use Microsoft\Kiota\Abstractions\ResponseHandler;
use Microsoft\Kiota\Abstractions\Serialization\Parsable;

class OverridesRequestBuilder 
{
    /** @var string $currentPath Current path for the request */
    private string $currentPath;
    
    /** @var bool $isRawUrl Whether the current path is a raw URL */
    private bool $isRawUrl;
    
    /** @var string $pathSegment Path segment to use to build the URL for the current request builder */
    private string $pathSegment;
    
    private RequestAdapter $requestAdapter;
    
    /**
     * Instantiates a new OverridesRequestBuilder and sets the default values.
     * @param string $currentPath $currentPath Current path for the request
     * @param bool|null $isRawUrl $isRawUrl Whether the current path is a raw URL
     * @param IRequestAdapter $$requestAdapter $requestAdapter The http core service to use to execute the requests.
    */
    public function __construct(string $currentPath, ?bool $isRawUrl, IRequestAdapter $$requestAdapter) {
        $this->pathSegment = '/overrides';
        $this->requestAdapter = $requestAdapter;
        $this->currentPath = $currentPath;
    }

    /**
     * A set of overrides for a user to always classify messages from specific senders in certain ways: focused, or other. Read-only. Nullable.
     * @param GetQueryParameters|null $q $q Request query parameters
     * @param array|null $ $h Request headers
     * @param array|null $options $o Request options
     * @return RequestInformation
    */
    public function createGetRequestInformation(?GetQueryParameters $q, ?array $, ?array $options): RequestInformation {
    }

    /**
     * A set of overrides for a user to always classify messages from specific senders in certain ways: focused, or other. Read-only. Nullable.
     * @param InferenceClassificationOverride $body $body 
     * @param array|null $ $h Request headers
     * @param array|null $options $o Request options
     * @return RequestInformation
    */
    public function createPostRequestInformation(InferenceClassificationOverride $body, ?array $, ?array $options): RequestInformation {
    }

    /**
     * A set of overrides for a user to always classify messages from specific senders in certain ways: focused, or other. Read-only. Nullable.
     * @param GetQueryParameters|null $q $q Request query parameters
     * @param array|null $ $h Request headers
     * @param array|null $options $o Request options
     * @param IResponseHandler|null $$responseHandler $responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return OverridesResponse|null
    */
    public function get(?GetQueryParameters $q, ?array $, ?array $options, ?IResponseHandler $$responseHandler): ?OverridesResponse {
    }

    /**
     * A set of overrides for a user to always classify messages from specific senders in certain ways: focused, or other. Read-only. Nullable.
     * @param InferenceClassificationOverride $body $body 
     * @param array|null $ $h Request headers
     * @param array|null $options $o Request options
     * @param IResponseHandler|null $$responseHandler $responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return InferenceClassificationOverride|null
    */
    public function post(InferenceClassificationOverride $body, ?array $, ?array $options, ?IResponseHandler $$responseHandler): ?InferenceClassificationOverride {
    }

}
