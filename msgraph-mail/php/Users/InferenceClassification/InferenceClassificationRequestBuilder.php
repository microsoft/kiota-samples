<?php

namespace Microsoft\Graph\Users\InferenceClassification;

use Microsoft\Graph\Users\InferenceClassification\Overrides\OverridesRequestBuilder;
use Microsoft\Kiota\Abstractions\HttpCore;
use Microsoft\Kiota\Abstractions\HttpMethod;
use Microsoft\Kiota\Abstractions\MiddlewareOption;
use Microsoft\Kiota\Abstractions\QueryParametersBase;
use Microsoft\Kiota\Abstractions\RequestInformation;
use Microsoft\Kiota\Abstractions\ResponseHandler;
use Microsoft\Kiota\Abstractions\Serialization\ParseNode;

class InferenceClassificationRequestBuilder 
{
    /** @var string $currentPath Current path for the request */
    private string $currentPath;
    
    private HttpCoreInterface $httpCore;
    
    /** @var bool $isRawUrl Whether the current path is a raw URL */
    private bool $isRawUrl;
    
    public function overrides(): OverridesRequestBuilder {
        return new OverridesRequestBuilder($this->currentPath . $this->pathSegment, $this->httpCore);
    }
    
    /** @var string $pathSegment Path segment to use to build the URL for the current request builder */
    private string $pathSegment;
    
    /**
     * Instantiates a new InferenceClassificationRequestBuilder and sets the default values.
     * @param ?string $currentPath $currentPath Current path for the request
     * @param bool|null $isRawUrl $isRawUrl Whether the current path is a raw URL
     * @param ?HttpCore $httpCore $httpCore The http core service to use to execute the requests.
    */
    public function __construct(?string $currentPath, bool $isRawUrl, ?HttpCore $httpCore) {
        $this->pathSegment = '/inferenceClassification';
        $this->httpCore = $httpCore;
        $this->currentPath = $currentPath;
    }

    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @param array|null $headers $h Request headers
     * @param array|null $options $o Request options for HTTP middlewares
     * @return RequestInformation
    */
    public function createDeleteRequestInformation(array $headers, array $options): RequestInformation {
    }

    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @param GetQueryParameters|null $q $q Request query parameters
     * @param array|null $headers $h Request headers
     * @param array|null $options $o Request options for HTTP middlewares
     * @return RequestInformation
    */
    public function createGetRequestInformation(GetQueryParameters $q, array $headers, array $options): RequestInformation {
    }

    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @param ?InferenceClassification $body $body 
     * @param array|null $headers $h Request headers
     * @param array|null $options $o Request options for HTTP middlewares
     * @return RequestInformation
    */
    public function createPatchRequestInformation(?InferenceClassification $body, array $headers, array $options): RequestInformation {
    }

    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @param array|null $headers $h Request headers
     * @param array|null $options $o Request options for HTTP middlewares
     * @param ResponseHandlerInterface|null $responseHandler $responseHandler Response handler to use in place of the default response handling provided by the core service
    */
    public function delete(array $headers, array $options, ResponseHandlerInterface $responseHandler): void {
    }

    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @param GetQueryParameters|null $q $q Request query parameters
     * @param array|null $headers $h Request headers
     * @param array|null $options $o Request options for HTTP middlewares
     * @param ResponseHandlerInterface|null $responseHandler $responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return InferenceClassification|null
    */
    public function get(GetQueryParameters $q, array $headers, array $options, ResponseHandlerInterface $responseHandler): ?InferenceClassification {
    }

    /**
     * Gets an item from the Microsoft\Graph.users.inferenceClassification.overrides collection
     * @param ?String $id $id Unique identifier of the item
     * @return InferenceClassificationOverrideRequestBuilder
    */
    public function overridesById(?String $id): InferenceClassificationOverrideRequestBuilder {
    }

    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @param ?InferenceClassification $body $body 
     * @param array|null $headers $h Request headers
     * @param array|null $options $o Request options for HTTP middlewares
     * @param ResponseHandlerInterface|null $responseHandler $responseHandler Response handler to use in place of the default response handling provided by the core service
    */
    public function patch(?InferenceClassification $body, array $headers, array $options, ResponseHandlerInterface $responseHandler): void {
    }

}
