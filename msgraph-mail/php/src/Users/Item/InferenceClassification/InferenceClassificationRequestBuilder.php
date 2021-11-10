<?php

namespace Microsoft\Graph\Users\Item\InferenceClassification;

use Microsoft\Graph\Models\Microsoft\Graph\InferenceClassification;
use Microsoft\Graph\Users\Item\InferenceClassification\Overrides\Item\InferenceClassificationOverrideRequestBuilder;
use Microsoft\Graph\Users\Item\InferenceClassification\Overrides\OverridesRequestBuilder;
use Microsoft\Kiota\Abstractions\HttpMethod;
use Microsoft\Kiota\Abstractions\MiddlewareOption;
use Microsoft\Kiota\Abstractions\RequestAdapter;
use Microsoft\Kiota\Abstractions\RequestInformation;
use Microsoft\Kiota\Abstractions\ResponseHandler;
use Microsoft\Kiota\Abstractions\Serialization\Parsable;

class InferenceClassificationRequestBuilder 
{
    public function overrides(): OverridesRequestBuilder {
        return new OverridesRequestBuilder($this->pathParameters, $this->requestAdapter);
    }
    
    /** @var array|null $pathParameters Path parameters for the request */
    private array $pathParameters;
    
    /** @var RequestAdapter|null $requestAdapter The request adapter to use to execute the requests. */
    private RequestAdapter $requestAdapter;
    
    /** @var string|null $urlTemplate Url template to use to build the URL for the current request builder */
    private ?string $urlTemplate;
    
    /**
     * Instantiates a new InferenceClassificationRequestBuilder and sets the default values.
     * @param array $pathParameters Path parameters for the request
     * @param RequestAdapter $requestAdapter The request adapter to use to execute the requests.
    */
    public function __construct(array $pathParameters, RequestAdapter $requestAdapter) {
        $this->urlTemplate = '{+baseurl}/users/{user_id}/inferenceClassification{?select,expand}';
        $this->requestAdapter = $requestAdapter;
        $this->pathParameters = $pathParameters;
    }

    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @param array|null $headers Request headers
     * @param array|null $options Request options
     * @return RequestInformation
    */
    public function createDeleteRequestInformation(?array $headers, ?array $options): RequestInformation {
        $requestInfo = new RequestInformation();
        $requestInfo->urlTemplate = $this->urlTemplate;
        $requestInfo->pathParameters = $this->pathParameters;
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
        $requestInfo->urlTemplate = $this->urlTemplate;
        $requestInfo->pathParameters = $this->pathParameters;
        $requestInfo->httpMethod = HttpMethod::GET;
        $requestInfo->setHeadersFromRawObject($headers);
        $requestInfo->setQueryStringParametersFromRawObject($queryParameters);
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
        $requestInfo = new RequestInformation();
        $requestInfo->urlTemplate = $this->urlTemplate;
        $requestInfo->pathParameters = $this->pathParameters;
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
        $requestInfo = $this->createDeleteRequestInformation($headers, $options);
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
        $requestInfo = $this->createGetRequestInformation($queryParameters, $headers, $options);
    }

    /**
     * Gets an item from the Microsoft\Graph.users.item.inferenceClassification.overrides.item collection
     * @param String $id Unique identifier of the item
     * @return InferenceClassificationOverrideRequestBuilder
    */
    public function overridesById(String $id): InferenceClassificationOverrideRequestBuilder {
        $urlTplParams = $this->pathParameters;
        $urlTplParams['inferenceClassificationOverride_id'] = $id;
        return new InferenceClassificationOverrideRequestBuilder($urlTplParams, $this->requestAdapter);
    }

    /**
     * Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
     * @param InferenceClassification $body 
     * @param array|null $headers Request headers
     * @param array|null $options Request options
     * @param ResponseHandler|null $responseHandler Response handler to use in place of the default response handling provided by the core service
    */
    public function patch(InferenceClassification $body, ?array $headers, ?array $options, ?ResponseHandler $responseHandler): void {
        $requestInfo = $this->createPatchRequestInformation($body, $headers, $options);
    }

}
