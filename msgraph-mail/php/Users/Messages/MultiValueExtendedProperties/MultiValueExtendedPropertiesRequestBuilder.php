<?php

namespace Microsoft\Graph\Users\Messages\MultiValueExtendedProperties;

use Microsoft\Graph\Users\Messages\MultiValueExtendedProperties\Item\MultiValueLegacyExtendedPropertyRequestBuilder;
use Microsoft\Graph\Users\MultiValueLegacyExtendedProperty;
use Microsoft\Kiota\Abstractions\HttpCore;
use Microsoft\Kiota\Abstractions\HttpMethod;
use Microsoft\Kiota\Abstractions\MiddlewareOption;
use Microsoft\Kiota\Abstractions\QueryParametersBase;
use Microsoft\Kiota\Abstractions\RequestInformation;
use Microsoft\Kiota\Abstractions\ResponseHandler;
use Microsoft\Kiota\Abstractions\Serialization\ParseNode;

class MultiValueExtendedPropertiesRequestBuilder 
{
    /** @var string $currentPath Current path for the request */
    private string $currentPath;
    
    private HttpCoreInterface $httpCore;
    
    /** @var bool $isRawUrl Whether the current path is a raw URL */
    private bool $isRawUrl;
    
    /** @var string $pathSegment Path segment to use to build the URL for the current request builder */
    private string $pathSegment;
    
    /**
     * Instantiates a new MultiValueExtendedPropertiesRequestBuilder and sets the default values.
     * @param ?string $currentPath $currentPath Current path for the request
     * @param bool|null $isRawUrl $isRawUrl Whether the current path is a raw URL
     * @param ?HttpCore $httpCore $httpCore The http core service to use to execute the requests.
    */
    public function __construct(?string $currentPath, bool $isRawUrl, ?HttpCore $httpCore) {
        $this->pathSegment = '/multiValueExtendedProperties';
        $this->httpCore = $httpCore;
        $this->currentPath = $currentPath;
    }

    /**
     * The collection of multi-value extended properties defined for the message. Nullable.
     * @param GetQueryParameters|null $q $q Request query parameters
     * @param array|null $headers $h Request headers
     * @param array|null $options $o Request options for HTTP middlewares
     * @return RequestInformation
    */
    public function createGetRequestInformation(GetQueryParameters $q, array $headers, array $options): RequestInformation {
    }

    /**
     * The collection of multi-value extended properties defined for the message. Nullable.
     * @param ?MultiValueLegacyExtendedProperty $body $body 
     * @param array|null $headers $h Request headers
     * @param array|null $options $o Request options for HTTP middlewares
     * @return RequestInformation
    */
    public function createPostRequestInformation(?MultiValueLegacyExtendedProperty $body, array $headers, array $options): RequestInformation {
    }

    /**
     * The collection of multi-value extended properties defined for the message. Nullable.
     * @param GetQueryParameters|null $q $q Request query parameters
     * @param array|null $headers $h Request headers
     * @param array|null $options $o Request options for HTTP middlewares
     * @param ResponseHandlerInterface|null $responseHandler $responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return MultiValueExtendedPropertiesResponse|null
    */
    public function get(GetQueryParameters $q, array $headers, array $options, ResponseHandlerInterface $responseHandler): ?MultiValueExtendedPropertiesResponse {
    }

    /**
     * The collection of multi-value extended properties defined for the message. Nullable.
     * @param ?MultiValueLegacyExtendedProperty $body $body 
     * @param array|null $headers $h Request headers
     * @param array|null $options $o Request options for HTTP middlewares
     * @param ResponseHandlerInterface|null $responseHandler $responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return MultiValueLegacyExtendedProperty|null
    */
    public function post(?MultiValueLegacyExtendedProperty $body, array $headers, array $options, ResponseHandlerInterface $responseHandler): ?MultiValueLegacyExtendedProperty {
    }

}
