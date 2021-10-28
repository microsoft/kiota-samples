<?php

namespace Microsoft\Graph\Users\Item\MailFolders\Item\Messages\Item;

use Microsoft\Graph\Models\Microsoft\Graph\Message;
use Microsoft\Graph\Users\Item\MailFolders\Item\Messages\Item\Attachments\AttachmentsRequestBuilder;
use Microsoft\Graph\Users\Item\MailFolders\Item\Messages\Item\Extensions\ExtensionsRequestBuilder;
use Microsoft\Graph\Users\Item\MailFolders\Item\Messages\Item\MultiValueExtendedProperties\MultiValueExtendedPropertiesRequestBuilder;
use Microsoft\Graph\Users\Item\MailFolders\Item\Messages\Item\SingleValueExtendedProperties\SingleValueExtendedPropertiesRequestBuilder;
use Microsoft\Graph\Users\Item\MailFolders\Item\Messages\Item\Value\ContentRequestBuilder;
use Microsoft\Kiota\Abstractions\HttpMethod;
use Microsoft\Kiota\Abstractions\MiddlewareOption;
use Microsoft\Kiota\Abstractions\RequestAdapter;
use Microsoft\Kiota\Abstractions\RequestInformation;
use Microsoft\Kiota\Abstractions\ResponseHandler;
use Microsoft\Kiota\Abstractions\Serialization\Parsable;

class MessageRequestBuilder 
{
    public function attachments(): AttachmentsRequestBuilder {
        return new AttachmentsRequestBuilder($this->currentPath . $this->pathSegment, $this->requestAdapter);
    }
    
    public function content(): ContentRequestBuilder {
        return new ContentRequestBuilder($this->currentPath . $this->pathSegment, $this->requestAdapter);
    }
    
    /** @var string $currentPath Current path for the request */
    private string $currentPath;
    
    public function extensions(): ExtensionsRequestBuilder {
        return new ExtensionsRequestBuilder($this->currentPath . $this->pathSegment, $this->requestAdapter);
    }
    
    /** @var bool $isRawUrl Whether the current path is a raw URL */
    private bool $isRawUrl;
    
    public function multiValueExtendedProperties(): MultiValueExtendedPropertiesRequestBuilder {
        return new MultiValueExtendedPropertiesRequestBuilder($this->currentPath . $this->pathSegment, $this->requestAdapter);
    }
    
    /** @var string $pathSegment Path segment to use to build the URL for the current request builder */
    private string $pathSegment;
    
    private RequestAdapter $requestAdapter;
    
    public function singleValueExtendedProperties(): SingleValueExtendedPropertiesRequestBuilder {
        return new SingleValueExtendedPropertiesRequestBuilder($this->currentPath . $this->pathSegment, $this->requestAdapter);
    }
    
    /**
     * Gets an item from the Microsoft\Graph.users.item.mailFolders.item.messages.item.attachments.item collection
     * @param String $id Unique identifier of the item
     * @return AttachmentRequestBuilder
    */
    public function attachmentsById(String $id): AttachmentRequestBuilder {
        if (is_null($id)) {
            throw new \Exception('$id cannot be null');
        }
    }

    /**
     * Instantiates a new MessageRequestBuilder and sets the default values.
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
        $this->pathSegment = '';
        $this->requestAdapter = $requestAdapter;
        $this->currentPath = $currentPath;
    }

    /**
     * The collection of messages in the mailFolder.
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
     * The collection of messages in the mailFolder.
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
     * The collection of messages in the mailFolder.
     * @param Message $body 
     * @param array|null $headers Request headers
     * @param array|null $options Request options
     * @return RequestInformation
    */
    public function createPatchRequestInformation(Message $body, ?array $headers, ?array $options): RequestInformation {
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
     * The collection of messages in the mailFolder.
     * @param array|null $headers Request headers
     * @param array|null $options Request options
     * @param ResponseHandler|null $responseHandler Response handler to use in place of the default response handling provided by the core service
    */
    public function delete(?array $headers, ?array $options, ?ResponseHandler $responseHandler): void {
    }

    /**
     * Gets an item from the Microsoft\Graph.users.item.mailFolders.item.messages.item.extensions.item collection
     * @param String $id Unique identifier of the item
     * @return ExtensionRequestBuilder
    */
    public function extensionsById(String $id): ExtensionRequestBuilder {
        if (is_null($id)) {
            throw new \Exception('$id cannot be null');
        }
    }

    /**
     * The collection of messages in the mailFolder.
     * @param GetQueryParameters|null $queryParameters Request query parameters
     * @param array|null $headers Request headers
     * @param array|null $options Request options
     * @param ResponseHandler|null $responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return Message|null
    */
    public function get(?GetQueryParameters $queryParameters, ?array $headers, ?array $options, ?ResponseHandler $responseHandler): ?Message {
    }

    /**
     * Gets an item from the Microsoft\Graph.users.item.mailFolders.item.messages.item.multiValueExtendedProperties.item collection
     * @param String $id Unique identifier of the item
     * @return MultiValueLegacyExtendedPropertyRequestBuilder
    */
    public function multiValueExtendedPropertiesById(String $id): MultiValueLegacyExtendedPropertyRequestBuilder {
        if (is_null($id)) {
            throw new \Exception('$id cannot be null');
        }
    }

    /**
     * The collection of messages in the mailFolder.
     * @param Message $body 
     * @param array|null $headers Request headers
     * @param array|null $options Request options
     * @param ResponseHandler|null $responseHandler Response handler to use in place of the default response handling provided by the core service
    */
    public function patch(Message $body, ?array $headers, ?array $options, ?ResponseHandler $responseHandler): void {
        if (is_null($body)) {
            throw new \Exception('$body cannot be null');
        }
    }

    /**
     * Gets an item from the Microsoft\Graph.users.item.mailFolders.item.messages.item.singleValueExtendedProperties.item collection
     * @param String $id Unique identifier of the item
     * @return SingleValueLegacyExtendedPropertyRequestBuilder
    */
    public function singleValueExtendedPropertiesById(String $id): SingleValueLegacyExtendedPropertyRequestBuilder {
        if (is_null($id)) {
            throw new \Exception('$id cannot be null');
        }
    }

}
