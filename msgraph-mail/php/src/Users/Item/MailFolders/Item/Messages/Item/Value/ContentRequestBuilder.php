<?php

namespace Microsoft\Graph\Users\Item\MailFolders\Item\Messages\Item\Value;

use Microsoft\Kiota\Abstractions\HttpMethod;
use Microsoft\Kiota\Abstractions\MiddlewareOption;
use Microsoft\Kiota\Abstractions\RequestAdapter;
use Microsoft\Kiota\Abstractions\RequestInformation;
use Microsoft\Kiota\Abstractions\ResponseHandler;
use Microsoft\Kiota\Abstractions\Serialization\Parsable;
use Psr\Http\Message\StreamInterface;

class ContentRequestBuilder 
{
    /** @var array|null $pathParameters Path parameters for the request */
    private array $pathParameters;
    
    /** @var RequestAdapter|null $requestAdapter The request adapter to use to execute the requests. */
    private RequestAdapter $requestAdapter;
    
    /** @var string|null $urlTemplate Url template to use to build the URL for the current request builder */
    private ?string $urlTemplate;
    
    /**
     * Instantiates a new ContentRequestBuilder and sets the default values.
     * @param array $pathParameters Path parameters for the request
     * @param RequestAdapter $requestAdapter The request adapter to use to execute the requests.
    */
    public function __construct(array $pathParameters, RequestAdapter $requestAdapter) {
        $this->urlTemplate = '{+baseurl}/users/{user_id}/mailFolders/{mailFolder_id}/messages/{message_id}/$value';
        $this->requestAdapter = $requestAdapter;
        $this->pathParameters = $pathParameters;
    }

    /**
     * Get media content for the navigation property messages from users
     * @param array|null $headers Request headers
     * @param array|null $options Request options
     * @return RequestInformation
    */
    public function createGetRequestInformation(?array $headers, ?array $options): RequestInformation {
        $requestInfo = new RequestInformation();
        $requestInfo->urlTemplate = $this->urlTemplate;
        $requestInfo->pathParameters = $this->pathParameters;
        $requestInfo->httpMethod = HttpMethod::GET;
        $requestInfo->setHeadersFromRawObject($headers);
        $requestInfo->addRequestOptions(...$options);
        return $requestInfo;
    }

    /**
     * Update media content for the navigation property messages in users
     * @param StreamInterface $body Binary request body
     * @param array|null $headers Request headers
     * @param array|null $options Request options
     * @return RequestInformation
    */
    public function createPutRequestInformation(StreamInterface $body, ?array $headers, ?array $options): RequestInformation {
        $requestInfo = new RequestInformation();
        $requestInfo->urlTemplate = $this->urlTemplate;
        $requestInfo->pathParameters = $this->pathParameters;
        $requestInfo->httpMethod = HttpMethod::PUT;
        $requestInfo->setHeadersFromRawObject($headers);
        $requestInfo->setStreamContent($body);
        $requestInfo->addRequestOptions(...$options);
        return $requestInfo;
    }

    /**
     * Get media content for the navigation property messages from users
     * @param array|null $headers Request headers
     * @param array|null $options Request options
     * @param ResponseHandler|null $responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return StreamInterface|null
    */
    public function get(?array $headers, ?array $options, ?ResponseHandler $responseHandler): ?StreamInterface {
        $requestInfo = $this->createGetRequestInformation($headers, $options);
    }

    /**
     * Update media content for the navigation property messages in users
     * @param StreamInterface $body Binary request body
     * @param array|null $headers Request headers
     * @param array|null $options Request options
     * @param ResponseHandler|null $responseHandler Response handler to use in place of the default response handling provided by the core service
    */
    public function put(StreamInterface $body, ?array $headers, ?array $options, ?ResponseHandler $responseHandler): void {
        $requestInfo = $this->createPutRequestInformation($body, $headers, $options);
    }

}
