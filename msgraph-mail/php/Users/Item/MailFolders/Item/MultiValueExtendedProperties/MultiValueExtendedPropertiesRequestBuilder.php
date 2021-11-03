<?php

namespace Microsoft\Graph\Users\Item\MailFolders\Item\MultiValueExtendedProperties;

use Microsoft\Graph\Models\Microsoft\Graph\MultiValueLegacyExtendedProperty;
use Microsoft\Graph\Users\Item\MailFolders\Item\MultiValueExtendedProperties\Item\MultiValueLegacyExtendedPropertyRequestBuilder;
use Microsoft\Kiota\Abstractions\HttpMethod;
use Microsoft\Kiota\Abstractions\MiddlewareOption;
use Microsoft\Kiota\Abstractions\RequestAdapter;
use Microsoft\Kiota\Abstractions\RequestInformation;
use Microsoft\Kiota\Abstractions\ResponseHandler;
use Microsoft\Kiota\Abstractions\Serialization\Parsable;

class MultiValueExtendedPropertiesRequestBuilder 
{
    /** @var array $pathParameters Path parameters for the request */
    private array $pathParameters;
    
    /** @var IRequestAdapter $requestAdapter The request adapter to use to execute the requests. */
    private RequestAdapter $requestAdapter;
    
    /** @var string $urlTemplate Url template to use to build the URL for the current request builder */
    private string $urlTemplate;
    
    /**
     * Instantiates a new MultiValueExtendedPropertiesRequestBuilder and sets the default values.
     * @param array $pathParameters Path parameters for the request
     * @param RequestAdapter $requestAdapter The request adapter to use to execute the requests.
    */
    public function __construct(array $pathParameters, RequestAdapter $requestAdapter) {
        if (is_null($pathParameters)) {
            throw new \Exception('$pathParameters cannot be null');
        }
        if (is_null($requestAdapter)) {
            throw new \Exception('$requestAdapter cannot be null');
        }
        $this->urlTemplate = 'https://graph.microsoft.com/v1.0/users/{user_id}/mailFolders/{mailFolder_id}/multiValueExtendedProperties{?top,skip,search,filter,count,orderby,select,expand}';
        $this->requestAdapter = $requestAdapter;
        $this->pathParameters = $pathParameters;
    }

    /**
     * The collection of multi-value extended properties defined for the mailFolder. Read-only. Nullable.
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
        $requestInfo.setHeadersFromRawObject($headers);
        $requestInfo->setQueryStringParametersFromRawObject($queryString);
        $requestInfo->addRequestOptions(...$options);
        return $requestInfo;
    }

    /**
     * The collection of multi-value extended properties defined for the mailFolder. Read-only. Nullable.
     * @param MultiValueLegacyExtendedProperty $body 
     * @param array|null $headers Request headers
     * @param array|null $options Request options
     * @return RequestInformation
    */
    public function createPostRequestInformation(MultiValueLegacyExtendedProperty $body, ?array $headers, ?array $options): RequestInformation {
        if (is_null($body)) {
            throw new \Exception('$body cannot be null');
        }
        $requestInfo = new RequestInformation();
        $requestInfo->urlTemplate = $this->urlTemplate;
        $requestInfo->pathParameters = $this->pathParameters;
        $requestInfo->httpMethod = HttpMethod::POST;
        $requestInfo.setHeadersFromRawObject($headers);
        $requestInfo->setContentFromParsable(this.requestAdapter, "application/json", $body);
        $requestInfo->addRequestOptions(...$options);
        return $requestInfo;
    }

    /**
     * The collection of multi-value extended properties defined for the mailFolder. Read-only. Nullable.
     * @param GetQueryParameters|null $queryParameters Request query parameters
     * @param array|null $headers Request headers
     * @param array|null $options Request options
     * @param ResponseHandler|null $responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return MultiValueExtendedPropertiesResponse|null
    */
    public function get(?GetQueryParameters $queryParameters, ?array $headers, ?array $options, ?ResponseHandler $responseHandler): ?MultiValueExtendedPropertiesResponse {
    }

    /**
     * The collection of multi-value extended properties defined for the mailFolder. Read-only. Nullable.
     * @param MultiValueLegacyExtendedProperty $body 
     * @param array|null $headers Request headers
     * @param array|null $options Request options
     * @param ResponseHandler|null $responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return MultiValueLegacyExtendedProperty|null
    */
    public function post(MultiValueLegacyExtendedProperty $body, ?array $headers, ?array $options, ?ResponseHandler $responseHandler): ?MultiValueLegacyExtendedProperty {
        if (is_null($body)) {
            throw new \Exception('$body cannot be null');
        }
    }

}
