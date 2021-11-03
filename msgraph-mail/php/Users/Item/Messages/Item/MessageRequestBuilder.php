<?php

namespace Microsoft\Graph\Users\Item\Messages\Item;

use Microsoft\Graph\Models\Microsoft\Graph\Message;
use Microsoft\Graph\Users\Item\Messages\Item\Attachments\AttachmentsRequestBuilder;
use Microsoft\Graph\Users\Item\Messages\Item\Extensions\ExtensionsRequestBuilder;
use Microsoft\Graph\Users\Item\Messages\Item\MultiValueExtendedProperties\MultiValueExtendedPropertiesRequestBuilder;
use Microsoft\Graph\Users\Item\Messages\Item\SingleValueExtendedProperties\SingleValueExtendedPropertiesRequestBuilder;
use Microsoft\Graph\Users\Item\Messages\Item\Value\ContentRequestBuilder;
use Microsoft\Kiota\Abstractions\HttpMethod;
use Microsoft\Kiota\Abstractions\MiddlewareOption;
use Microsoft\Kiota\Abstractions\RequestAdapter;
use Microsoft\Kiota\Abstractions\RequestInformation;
use Microsoft\Kiota\Abstractions\ResponseHandler;
use Microsoft\Kiota\Abstractions\Serialization\Parsable;

class MessageRequestBuilder 
{
    public function attachments(): AttachmentsRequestBuilder {
        return new AttachmentsRequestBuilder($this->pathParameters, $this->requestAdapter);
    }
    
    public function content(): ContentRequestBuilder {
        return new ContentRequestBuilder($this->pathParameters, $this->requestAdapter);
    }
    
    public function extensions(): ExtensionsRequestBuilder {
        return new ExtensionsRequestBuilder($this->pathParameters, $this->requestAdapter);
    }
    
    public function multiValueExtendedProperties(): MultiValueExtendedPropertiesRequestBuilder {
        return new MultiValueExtendedPropertiesRequestBuilder($this->pathParameters, $this->requestAdapter);
    }
    
    /** @var array $pathParameters Path parameters for the request */
    private array $pathParameters;
    
    /** @var IRequestAdapter $requestAdapter The request adapter to use to execute the requests. */
    private RequestAdapter $requestAdapter;
    
    public function singleValueExtendedProperties(): SingleValueExtendedPropertiesRequestBuilder {
        return new SingleValueExtendedPropertiesRequestBuilder($this->pathParameters, $this->requestAdapter);
    }
    
    /** @var string $urlTemplate Url template to use to build the URL for the current request builder */
    private string $urlTemplate;
    
    /**
     * Gets an item from the Microsoft\Graph.users.item.messages.item.attachments.item collection
     * @param String $id Unique identifier of the item
     * @return AttachmentRequestBuilder
    */
    public function attachmentsById(String $id): AttachmentRequestBuilder {
    }

    /**
     * Instantiates a new MessageRequestBuilder and sets the default values.
     * @param array $pathParameters Path parameters for the request
     * @param RequestAdapter $requestAdapter The request adapter to use to execute the requests.
    */
    public function __construct(array $pathParameters, RequestAdapter $requestAdapter) {
        $this->urlTemplate = 'https://graph.microsoft.com/v1.0/users/{user_id}/messages/{message_id}{?select,expand}';
        $this->requestAdapter = $requestAdapter;
        $this->pathParameters = $pathParameters;
    }

    /**
     * The messages in a mailbox or folder. Read-only. Nullable.
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
     * The messages in a mailbox or folder. Read-only. Nullable.
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
     * The messages in a mailbox or folder. Read-only. Nullable.
     * @param Message $body 
     * @param array|null $headers Request headers
     * @param array|null $options Request options
     * @return RequestInformation
    */
    public function createPatchRequestInformation(Message $body, ?array $headers, ?array $options): RequestInformation {
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
     * The messages in a mailbox or folder. Read-only. Nullable.
     * @param array|null $headers Request headers
     * @param array|null $options Request options
     * @param ResponseHandler|null $responseHandler Response handler to use in place of the default response handling provided by the core service
    */
    public function delete(?array $headers, ?array $options, ?ResponseHandler $responseHandler): void {
    }

    /**
     * Gets an item from the Microsoft\Graph.users.item.messages.item.extensions.item collection
     * @param String $id Unique identifier of the item
     * @return ExtensionRequestBuilder
    */
    public function extensionsById(String $id): ExtensionRequestBuilder {
    }

    /**
     * The messages in a mailbox or folder. Read-only. Nullable.
     * @param GetQueryParameters|null $queryParameters Request query parameters
     * @param array|null $headers Request headers
     * @param array|null $options Request options
     * @param ResponseHandler|null $responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return Message|null
    */
    public function get(?GetQueryParameters $queryParameters, ?array $headers, ?array $options, ?ResponseHandler $responseHandler): ?Message {
    }

    /**
     * Gets an item from the Microsoft\Graph.users.item.messages.item.multiValueExtendedProperties.item collection
     * @param String $id Unique identifier of the item
     * @return MultiValueLegacyExtendedPropertyRequestBuilder
    */
    public function multiValueExtendedPropertiesById(String $id): MultiValueLegacyExtendedPropertyRequestBuilder {
    }

    /**
     * The messages in a mailbox or folder. Read-only. Nullable.
     * @param Message $body 
     * @param array|null $headers Request headers
     * @param array|null $options Request options
     * @param ResponseHandler|null $responseHandler Response handler to use in place of the default response handling provided by the core service
    */
    public function patch(Message $body, ?array $headers, ?array $options, ?ResponseHandler $responseHandler): void {
    }

    /**
     * Gets an item from the Microsoft\Graph.users.item.messages.item.singleValueExtendedProperties.item collection
     * @param String $id Unique identifier of the item
     * @return SingleValueLegacyExtendedPropertyRequestBuilder
    */
    public function singleValueExtendedPropertiesById(String $id): SingleValueLegacyExtendedPropertyRequestBuilder {
    }

}
