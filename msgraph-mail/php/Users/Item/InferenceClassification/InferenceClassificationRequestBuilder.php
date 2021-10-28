<?php

namespace Microsoft\Graph\Users\Item\InferenceClassification;

use Microsoft\Graph\Models\Microsoft\Graph\InferenceClassification;
use Microsoft\Graph\Users\Item\InferenceClassification\Overrides\OverridesRequestBuilder;
use Microsoft\Kiota\Abstractions\HttpMethod;
use Microsoft\Kiota\Abstractions\MiddlewareOption;
use Microsoft\Kiota\Abstractions\RequestAdapter;
use Microsoft\Kiota\Abstractions\RequestInformation;
use Microsoft\Kiota\Abstractions\ResponseHandler;
use Microsoft\Kiota\Abstractions\Serialization\Parsable;

class InferenceClassificationRequestBuilder 
{
    /** @var string $currentPath Current path for the request */
    private string $currentPath;
    
    /** @var bool $isRawUrl Whether the current path is a raw URL */
    private bool $isRawUrl;
    
    public function overrides(): OverridesRequestBuilder {
        return new OverridesRequestBuilder($this->currentPath . $this->pathSegment, $this->requestAdapter);
    }
    
    /** @var string $pathSegment Path segment to use to build the URL for the current request builder */
    private string $pathSegment;
    
    private RequestAdapter $requestAdapter;
    
    /**
     * Instantiates a new InferenceClassificationRequestBuilder and sets the default values.
     * @param string $currentPath Current path for the request
     * @param bool|null $isRawUrl Whether the current path is a raw URL
     * @param RequestAdapter $requestAdapter The http core service to use to execute the requests.
    */
    public function __construct(string $currentPath, ?bool $isRawUrl, RequestAdapter $requestAdapter) {
        if (is_null($currentPath)) {
            throw new \Exception('$currentPath cannot be null');
        }
        if (is_null($requestAdapter)) {
            throw new \Exception('$requestAdapter cannot be null');
        }
        $this->pathSegment = '/inferenceClassification';
        $this->requestAdapter = $requestAdapter;
        $this->currentPath = $currentPath;
    }

    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @param array|null $headers Request headers
     * @param array|null $options Request options
     * @return RequestInformation
    */
    public function createDeleteRequestInformation(?array $headers, ?array $options): RequestInformation {
        $requestInfo = new RequestInformation();
        $requestInfo->setUri($this->currentPath, $this->pathSegment, $this->isRawUrl);
        $requestInfo->httpMethod = HttpMethod::DELETE;
        $requestInfo->setHeadersFromRawObject($headers);
        $requestInfo->addRequestOptions(...$options);
        return $requestInfo;
    }

    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @param GetQueryParameters|null $queryParameters Request query parameters
     * @param array|null $headers Request headers
     * @param array|null $options Request options
     * @return RequestInformation
    */
    public function createGetRequestInformation(?GetQueryParameters $queryParameters, ?array $headers, ?array $options): RequestInformation {
        $requestInfo = new RequestInformation();
        $requestInfo->setUri($this->currentPath, $this->pathSegment, $this->isRawUrl);
        $requestInfo->httpMethod = HttpMethod::GET;
        $requestInfo->setHeadersFromRawObject($headers);
        $requestInfo->setQueryStringParametersFromRawObject($queryStringParams);
        $requestInfo->addRequestOptions(...$options);
        return $requestInfo;
    }

    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @param InferenceClassification $body 
     * @param array|null $headers Request headers
     * @param array|null $options Request options
     * @return RequestInformation
    */
    public function createPatchRequestInformation(InferenceClassification $body, ?array $headers, ?array $options): RequestInformation {
        if (is_null($body)) {
            throw new \Exception('$body cannot be null');
        }
        $requestInfo = new RequestInformation();
        $requestInfo->setUri($this->currentPath, $this->pathSegment, $this->isRawUrl);
        $requestInfo->httpMethod = HttpMethod::PATCH;
        $requestInfo->setHeadersFromRawObject($headers);
        $requestInfo->setContentFromParsable($this->requestAdapter, "application/json", $body);
        $requestInfo->addRequestOptions(...$options);
        return $requestInfo;
    }

    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @param array|null $headers Request headers
     * @param array|null $options Request options
     * @param ResponseHandler|null $responseHandler Response handler to use in place of the default response handling provided by the core service
    */
    public function delete(?array $headers, ?array $options, ?ResponseHandler $responseHandler): void {
    }

    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @param GetQueryParameters|null $queryParameters Request query parameters
     * @param array|null $headers Request headers
     * @param array|null $options Request options
     * @param ResponseHandler|null $responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return InferenceClassification|null
    */
    public function get(?GetQueryParameters $queryParameters, ?array $headers, ?array $options, ?ResponseHandler $responseHandler): ?InferenceClassification {
    }

    /**
     * Gets an item from the Microsoft\Graph.users.item.inferenceClassification.overrides.item collection
     * @param String $id Unique identifier of the item
     * @return InferenceClassificationOverrideRequestBuilder
    */
    public function overridesById(String $id): InferenceClassificationOverrideRequestBuilder {
        if (is_null($id)) {
            throw new \Exception('$id cannot be null');
        }
    }

    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @param InferenceClassification $body 
     * @param array|null $headers Request headers
     * @param array|null $options Request options
     * @param ResponseHandler|null $responseHandler Response handler to use in place of the default response handling provided by the core service
    */
    public function patch(InferenceClassification $body, ?array $headers, ?array $options, ?ResponseHandler $responseHandler): void {
        if (is_null($body)) {
            throw new \Exception('$body cannot be null');
        }
    }

}
