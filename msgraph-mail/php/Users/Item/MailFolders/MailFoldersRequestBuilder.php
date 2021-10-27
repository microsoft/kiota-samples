<?php

namespace Microsoft\Graph\Users\Item\MailFolders;

use Microsoft\Graph\Models\Microsoft\Graph\MailFolder;
use Microsoft\Graph\Users\Item\MailFolders\Item\MailFolderRequestBuilder;
use Microsoft\Kiota\Abstractions\HttpMethod;
use Microsoft\Kiota\Abstractions\MiddlewareOption;
use Microsoft\Kiota\Abstractions\RequestAdapter;
use Microsoft\Kiota\Abstractions\RequestInformation;
use Microsoft\Kiota\Abstractions\ResponseHandler;
use Microsoft\Kiota\Abstractions\Serialization\Parsable;

class MailFoldersRequestBuilder 
{
    /** @var string $currentPath Current path for the request */
    private string $currentPath;
    
    /** @var bool $isRawUrl Whether the current path is a raw URL */
    private bool $isRawUrl;
    
    /** @var string $pathSegment Path segment to use to build the URL for the current request builder */
    private string $pathSegment;
    
    private RequestAdapter $requestAdapter;
    
    /**
     * Instantiates a new MailFoldersRequestBuilder and sets the default values.
     * @param string $currentPath $currentPath Current path for the request
     * @param bool|null $isRawUrl $isRawUrl Whether the current path is a raw URL
     * @param IRequestAdapter $$requestAdapter $requestAdapter The http core service to use to execute the requests.
    */
    public function __construct(string $currentPath, ?bool $isRawUrl, IRequestAdapter $$requestAdapter) {
        $this->pathSegment = '/mailFolders';
        $this->requestAdapter = $requestAdapter;
        $this->currentPath = $currentPath;
    }

    /**
     * The user's mail folders. Read-only. Nullable.
     * @param GetQueryParameters|null $q $q Request query parameters
     * @param array|null $ $h Request headers
     * @param array|null $options $o Request options
     * @return RequestInformation
    */
    public function createGetRequestInformation(?GetQueryParameters $q, ?array $, ?array $options): RequestInformation {
    }

    /**
     * The user's mail folders. Read-only. Nullable.
     * @param MailFolder $body $body 
     * @param array|null $ $h Request headers
     * @param array|null $options $o Request options
     * @return RequestInformation
    */
    public function createPostRequestInformation(MailFolder $body, ?array $, ?array $options): RequestInformation {
    }

    /**
     * The user's mail folders. Read-only. Nullable.
     * @param GetQueryParameters|null $q $q Request query parameters
     * @param array|null $ $h Request headers
     * @param array|null $options $o Request options
     * @param IResponseHandler|null $$responseHandler $responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return MailFoldersResponse|null
    */
    public function get(?GetQueryParameters $q, ?array $, ?array $options, ?IResponseHandler $$responseHandler): ?MailFoldersResponse {
    }

    /**
     * The user's mail folders. Read-only. Nullable.
     * @param MailFolder $body $body 
     * @param array|null $ $h Request headers
     * @param array|null $options $o Request options
     * @param IResponseHandler|null $$responseHandler $responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return MailFolder|null
    */
    public function post(MailFolder $body, ?array $, ?array $options, ?IResponseHandler $$responseHandler): ?MailFolder {
    }

}
