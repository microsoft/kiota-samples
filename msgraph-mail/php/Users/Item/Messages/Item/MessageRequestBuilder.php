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
     * Gets an item from the Microsoft\Graph.users.item.messages.item.attachments.item collection
     * @param String $id $id Unique identifier of the item
     * @return AttachmentRequestBuilder
    */
    public function attachmentsById(String $id): AttachmentRequestBuilder {
    }

    /**
     * Instantiates a new MessageRequestBuilder and sets the default values.
     * @param string $currentPath $currentPath Current path for the request
     * @param bool|null $isRawUrl $isRawUrl Whether the current path is a raw URL
     * @param IRequestAdapter $$requestAdapter $requestAdapter The http core service to use to execute the requests.
    */
    public function __construct(string $currentPath, ?bool $isRawUrl, IRequestAdapter $$requestAdapter) {
        $this->pathSegment = '';
        $this->requestAdapter = $requestAdapter;
        $this->currentPath = $currentPath;
    }

    /**
     * The messages in a mailbox or folder. Read-only. Nullable.
     * @param array|null $ $h Request headers
     * @param array|null $options $o Request options
     * @return RequestInformation
    */
    public function createDeleteRequestInformation(?array $, ?array $options): RequestInformation {
    }

    /**
     * The messages in a mailbox or folder. Read-only. Nullable.
     * @param GetQueryParameters|null $q $q Request query parameters
     * @param array|null $ $h Request headers
     * @param array|null $options $o Request options
     * @return RequestInformation
    */
    public function createGetRequestInformation(?GetQueryParameters $q, ?array $, ?array $options): RequestInformation {
    }

    /**
     * The messages in a mailbox or folder. Read-only. Nullable.
     * @param Message $body $body 
     * @param array|null $ $h Request headers
     * @param array|null $options $o Request options
     * @return RequestInformation
    */
    public function createPatchRequestInformation(Message $body, ?array $, ?array $options): RequestInformation {
    }

    /**
     * The messages in a mailbox or folder. Read-only. Nullable.
     * @param array|null $ $h Request headers
     * @param array|null $options $o Request options
     * @param IResponseHandler|null $$responseHandler $responseHandler Response handler to use in place of the default response handling provided by the core service
    */
    public function delete(?array $, ?array $options, ?IResponseHandler $$responseHandler): void {
    }

    /**
     * Gets an item from the Microsoft\Graph.users.item.messages.item.extensions.item collection
     * @param String $id $id Unique identifier of the item
     * @return ExtensionRequestBuilder
    */
    public function extensionsById(String $id): ExtensionRequestBuilder {
    }

    /**
     * The messages in a mailbox or folder. Read-only. Nullable.
     * @param GetQueryParameters|null $q $q Request query parameters
     * @param array|null $ $h Request headers
     * @param array|null $options $o Request options
     * @param IResponseHandler|null $$responseHandler $responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return Message|null
    */
    public function get(?GetQueryParameters $q, ?array $, ?array $options, ?IResponseHandler $$responseHandler): ?Message {
    }

    /**
     * Gets an item from the Microsoft\Graph.users.item.messages.item.multiValueExtendedProperties.item collection
     * @param String $id $id Unique identifier of the item
     * @return MultiValueLegacyExtendedPropertyRequestBuilder
    */
    public function multiValueExtendedPropertiesById(String $id): MultiValueLegacyExtendedPropertyRequestBuilder {
    }

    /**
     * The messages in a mailbox or folder. Read-only. Nullable.
     * @param Message $body $body 
     * @param array|null $ $h Request headers
     * @param array|null $options $o Request options
     * @param IResponseHandler|null $$responseHandler $responseHandler Response handler to use in place of the default response handling provided by the core service
    */
    public function patch(Message $body, ?array $, ?array $options, ?IResponseHandler $$responseHandler): void {
    }

    /**
     * Gets an item from the Microsoft\Graph.users.item.messages.item.singleValueExtendedProperties.item collection
     * @param String $id $id Unique identifier of the item
     * @return SingleValueLegacyExtendedPropertyRequestBuilder
    */
    public function singleValueExtendedPropertiesById(String $id): SingleValueLegacyExtendedPropertyRequestBuilder {
    }

}
