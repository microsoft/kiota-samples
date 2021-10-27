<?php

namespace Microsoft\Graph\Users\Item\MailFolders\Item\SingleValueExtendedProperties;

use Microsoft\Graph\Models\Microsoft\Graph\SingleValueLegacyExtendedProperty;
use Microsoft\Graph\Users\Item\MailFolders\Item\SingleValueExtendedProperties\Item\SingleValueLegacyExtendedPropertyRequestBuilder;
use Microsoft\Kiota\Abstractions\HttpMethod;
use Microsoft\Kiota\Abstractions\MiddlewareOption;
use Microsoft\Kiota\Abstractions\RequestAdapter;
use Microsoft\Kiota\Abstractions\RequestInformation;
use Microsoft\Kiota\Abstractions\ResponseHandler;
use Microsoft\Kiota\Abstractions\Serialization\Parsable;

class SingleValueExtendedPropertiesRequestBuilder 
{
    /** @var string $currentPath Current path for the request */
    private string $currentPath;
    
    /** @var bool $isRawUrl Whether the current path is a raw URL */
    private bool $isRawUrl;
    
    /** @var string $pathSegment Path segment to use to build the URL for the current request builder */
    private string $pathSegment;
    
    private RequestAdapter $requestAdapter;
    
    /**
     * Instantiates a new SingleValueExtendedPropertiesRequestBuilder and sets the default values.
     * @param string $currentPath $currentPath Current path for the request
     * @param bool|null $isRawUrl $isRawUrl Whether the current path is a raw URL
     * @param IRequestAdapter $$requestAdapter $requestAdapter The http core service to use to execute the requests.
    */
    public function __construct(string $currentPath, ?bool $isRawUrl, IRequestAdapter $$requestAdapter) {
        $this->pathSegment = '/singleValueExtendedProperties';
        $this->requestAdapter = $requestAdapter;
        $this->currentPath = $currentPath;
    }

    /**
     * The collection of single-value extended properties defined for the mailFolder. Read-only. Nullable.
     * @param GetQueryParameters|null $q $q Request query parameters
     * @param array|null $ $h Request headers
     * @param array|null $options $o Request options
     * @return RequestInformation
    */
    public function createGetRequestInformation(?GetQueryParameters $q, ?array $, ?array $options): RequestInformation {
    }

    /**
     * The collection of single-value extended properties defined for the mailFolder. Read-only. Nullable.
     * @param SingleValueLegacyExtendedProperty $body $body 
     * @param array|null $ $h Request headers
     * @param array|null $options $o Request options
     * @return RequestInformation
    */
    public function createPostRequestInformation(SingleValueLegacyExtendedProperty $body, ?array $, ?array $options): RequestInformation {
    }

    /**
     * The collection of single-value extended properties defined for the mailFolder. Read-only. Nullable.
     * @param GetQueryParameters|null $q $q Request query parameters
     * @param array|null $ $h Request headers
     * @param array|null $options $o Request options
     * @param IResponseHandler|null $$responseHandler $responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return SingleValueExtendedPropertiesResponse|null
    */
    public function get(?GetQueryParameters $q, ?array $, ?array $options, ?IResponseHandler $$responseHandler): ?SingleValueExtendedPropertiesResponse {
    }

    /**
     * The collection of single-value extended properties defined for the mailFolder. Read-only. Nullable.
     * @param SingleValueLegacyExtendedProperty $body $body 
     * @param array|null $ $h Request headers
     * @param array|null $options $o Request options
     * @param IResponseHandler|null $$responseHandler $responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return SingleValueLegacyExtendedProperty|null
    */
    public function post(SingleValueLegacyExtendedProperty $body, ?array $, ?array $options, ?IResponseHandler $$responseHandler): ?SingleValueLegacyExtendedProperty {
    }

}
