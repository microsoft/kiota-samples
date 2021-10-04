<?php

namespace Microsoft\Graph\Users\MailFolders\ChildFolders\Item;

use Microsoft\Graph\Users\MailFolder;
use Microsoft\Kiota\Abstractions\HttpCore;
use Microsoft\Kiota\Abstractions\HttpMethod;
use Microsoft\Kiota\Abstractions\MiddlewareOption;
use Microsoft\Kiota\Abstractions\QueryParametersBase;
use Microsoft\Kiota\Abstractions\RequestInformation;
use Microsoft\Kiota\Abstractions\ResponseHandler;
use Microsoft\Kiota\Abstractions\Serialization\ParseNode;

class MailFolderRequestBuilder 
{
    /** @var string $currentPath Current path for the request */
    private string $currentPath;
    
    private HttpCoreInterface $httpCore;
    
    /** @var bool $isRawUrl Whether the current path is a raw URL */
    private bool $isRawUrl;
    
    /** @var string $pathSegment Path segment to use to build the URL for the current request builder */
    private string $pathSegment;
    
    /**
     * Instantiates a new MailFolderRequestBuilder and sets the default values.
     * @param ?string $currentPath $currentPath Current path for the request
     * @param bool|null $isRawUrl $isRawUrl Whether the current path is a raw URL
     * @param ?HttpCore $httpCore $httpCore The http core service to use to execute the requests.
    */
    public function __construct(?string $currentPath, bool $isRawUrl, ?HttpCore $httpCore) {
        $this->pathSegment = '';
        $this->httpCore = $httpCore;
        $this->currentPath = $currentPath;
    }

    /**
     * The collection of child folders in the mailFolder.
     * @param array|null $headers $h Request headers
     * @param array|null $options $o Request options for HTTP middlewares
     * @return RequestInformation
    */
    public function createDeleteRequestInformation(array $headers, array $options): RequestInformation {
    }

    /**
     * The collection of child folders in the mailFolder.
     * @param GetQueryParameters|null $q $q Request query parameters
     * @param array|null $headers $h Request headers
     * @param array|null $options $o Request options for HTTP middlewares
     * @return RequestInformation
    */
    public function createGetRequestInformation(GetQueryParameters $q, array $headers, array $options): RequestInformation {
    }

    /**
     * The collection of child folders in the mailFolder.
     * @param ?MailFolder $body $body 
     * @param array|null $headers $h Request headers
     * @param array|null $options $o Request options for HTTP middlewares
     * @return RequestInformation
    */
    public function createPatchRequestInformation(?MailFolder $body, array $headers, array $options): RequestInformation {
    }

    /**
     * The collection of child folders in the mailFolder.
     * @param array|null $headers $h Request headers
     * @param array|null $options $o Request options for HTTP middlewares
     * @param ResponseHandlerInterface|null $responseHandler $responseHandler Response handler to use in place of the default response handling provided by the core service
    */
    public function delete(array $headers, array $options, ResponseHandlerInterface $responseHandler): void {
    }

    /**
     * The collection of child folders in the mailFolder.
     * @param GetQueryParameters|null $q $q Request query parameters
     * @param array|null $headers $h Request headers
     * @param array|null $options $o Request options for HTTP middlewares
     * @param ResponseHandlerInterface|null $responseHandler $responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return MailFolder|null
    */
    public function get(GetQueryParameters $q, array $headers, array $options, ResponseHandlerInterface $responseHandler): ?MailFolder {
    }

    /**
     * The collection of child folders in the mailFolder.
     * @param ?MailFolder $body $body 
     * @param array|null $headers $h Request headers
     * @param array|null $options $o Request options for HTTP middlewares
     * @param ResponseHandlerInterface|null $responseHandler $responseHandler Response handler to use in place of the default response handling provided by the core service
    */
    public function patch(?MailFolder $body, array $headers, array $options, ResponseHandlerInterface $responseHandler): void {
    }

}
