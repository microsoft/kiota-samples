<?php

namespace Microsoft\Graph\Users\Messages\Content;

use Microsoft\Kiota\Abstractions\HttpCore;
use Microsoft\Kiota\Abstractions\HttpMethod;
use Microsoft\Kiota\Abstractions\MiddlewareOption;
use Microsoft\Kiota\Abstractions\QueryParametersBase;
use Microsoft\Kiota\Abstractions\RequestInformation;
use Microsoft\Kiota\Abstractions\ResponseHandler;
use Microsoft\Kiota\Abstractions\Serialization\ParseNode;
use Psr\Http\Message\StreamInterface;

class ContentRequestBuilder 
{
    /** @var string $currentPath Current path for the request */
    private string $currentPath;
    
    private HttpCoreInterface $httpCore;
    
    /** @var bool $isRawUrl Whether the current path is a raw URL */
    private bool $isRawUrl;
    
    /** @var string $pathSegment Path segment to use to build the URL for the current request builder */
    private string $pathSegment;
    
    /**
     * Instantiates a new ContentRequestBuilder and sets the default values.
     * @param ?string $currentPath $currentPath Current path for the request
     * @param bool|null $isRawUrl $isRawUrl Whether the current path is a raw URL
     * @param ?HttpCore $httpCore $httpCore The http core service to use to execute the requests.
    */
    public function __construct(?string $currentPath, bool $isRawUrl, ?HttpCore $httpCore) {
        $this->pathSegment = '/$value';
        $this->httpCore = $httpCore;
        $this->currentPath = $currentPath;
    }

    /**
     * Get media content for the navigation property messages from users
     * @param array|null $headers $h Request headers
     * @param array|null $options $o Request options for HTTP middlewares
     * @return RequestInformation
    */
    public function createGetRequestInformation(array $headers, array $options): RequestInformation {
    }

    /**
     * Update media content for the navigation property messages in users
     * @param ?StreamInterface $body $body Binary request body
     * @param array|null $headers $h Request headers
     * @param array|null $options $o Request options for HTTP middlewares
     * @return RequestInformation
    */
    public function createPutRequestInformation(?StreamInterface $body, array $headers, array $options): RequestInformation {
    }

    /**
     * Get media content for the navigation property messages from users
     * @param array|null $headers $h Request headers
     * @param array|null $options $o Request options for HTTP middlewares
     * @param ResponseHandlerInterface|null $responseHandler $responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return StreamInterface|null
    */
    public function get(array $headers, array $options, ResponseHandlerInterface $responseHandler): ?StreamInterface {
    }

    /**
     * Update media content for the navigation property messages in users
     * @param ?StreamInterface $body $body Binary request body
     * @param array|null $headers $h Request headers
     * @param array|null $options $o Request options for HTTP middlewares
     * @param ResponseHandlerInterface|null $responseHandler $responseHandler Response handler to use in place of the default response handling provided by the core service
    */
    public function put(?StreamInterface $body, array $headers, array $options, ResponseHandlerInterface $responseHandler): void {
    }

}
