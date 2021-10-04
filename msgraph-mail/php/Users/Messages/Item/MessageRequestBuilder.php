<?php

namespace Microsoft\Graph\Users\Messages\Item;

use Microsoft\Graph\Users\Message;
use Microsoft\Graph\Users\Messages\Attachments\AttachmentsRequestBuilder;
use Microsoft\Graph\Users\Messages\Content\ContentRequestBuilder;
use Microsoft\Graph\Users\Messages\Extensions\ExtensionsRequestBuilder;
use Microsoft\Graph\Users\Messages\MultiValueExtendedProperties\MultiValueExtendedPropertiesRequestBuilder;
use Microsoft\Graph\Users\Messages\SingleValueExtendedProperties\SingleValueExtendedPropertiesRequestBuilder;
use Microsoft\Kiota\Abstractions\HttpCore;
use Microsoft\Kiota\Abstractions\HttpMethod;
use Microsoft\Kiota\Abstractions\MiddlewareOption;
use Microsoft\Kiota\Abstractions\QueryParametersBase;
use Microsoft\Kiota\Abstractions\RequestInformation;
use Microsoft\Kiota\Abstractions\ResponseHandler;
use Microsoft\Kiota\Abstractions\Serialization\ParseNode;

class MessageRequestBuilder 
{
    public function attachments(): AttachmentsRequestBuilder {
        return new AttachmentsRequestBuilder($this->currentPath . $this->pathSegment, $this->httpCore);
    }
    
    public function content(): ContentRequestBuilder {
        return new ContentRequestBuilder($this->currentPath . $this->pathSegment, $this->httpCore);
    }
    
    /** @var string $currentPath Current path for the request */
    private string $currentPath;
    
    public function extensions(): ExtensionsRequestBuilder {
        return new ExtensionsRequestBuilder($this->currentPath . $this->pathSegment, $this->httpCore);
    }
    
    private HttpCoreInterface $httpCore;
    
    /** @var bool $isRawUrl Whether the current path is a raw URL */
    private bool $isRawUrl;
    
    public function multiValueExtendedProperties(): MultiValueExtendedPropertiesRequestBuilder {
        return new MultiValueExtendedPropertiesRequestBuilder($this->currentPath . $this->pathSegment, $this->httpCore);
    }
    
    /** @var string $pathSegment Path segment to use to build the URL for the current request builder */
    private string $pathSegment;
    
    public function singleValueExtendedProperties(): SingleValueExtendedPropertiesRequestBuilder {
        return new SingleValueExtendedPropertiesRequestBuilder($this->currentPath . $this->pathSegment, $this->httpCore);
    }
    
    /**
     * Gets an item from the Microsoft\Graph.users.messages.attachments collection
     * @param ?String $id $id Unique identifier of the item
     * @return AttachmentRequestBuilder
    */
    public function attachmentsById(?String $id): AttachmentRequestBuilder {
    }

    /**
     * Instantiates a new MessageRequestBuilder and sets the default values.
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
     * The messages in a mailbox or folder. Read-only. Nullable.
     * @param array|null $headers $h Request headers
     * @param array|null $options $o Request options for HTTP middlewares
     * @return RequestInformation
    */
    public function createDeleteRequestInformation(array $headers, array $options): RequestInformation {
    }

    /**
     * The messages in a mailbox or folder. Read-only. Nullable.
     * @param GetQueryParameters|null $q $q Request query parameters
     * @param array|null $headers $h Request headers
     * @param array|null $options $o Request options for HTTP middlewares
     * @return RequestInformation
    */
    public function createGetRequestInformation(GetQueryParameters $q, array $headers, array $options): RequestInformation {
    }

    /**
     * The messages in a mailbox or folder. Read-only. Nullable.
     * @param ?Message $body $body 
     * @param array|null $headers $h Request headers
     * @param array|null $options $o Request options for HTTP middlewares
     * @return RequestInformation
    */
    public function createPatchRequestInformation(?Message $body, array $headers, array $options): RequestInformation {
    }

    /**
     * The messages in a mailbox or folder. Read-only. Nullable.
     * @param array|null $headers $h Request headers
     * @param array|null $options $o Request options for HTTP middlewares
     * @param ResponseHandlerInterface|null $responseHandler $responseHandler Response handler to use in place of the default response handling provided by the core service
    */
    public function delete(array $headers, array $options, ResponseHandlerInterface $responseHandler): void {
    }

    /**
     * Gets an item from the Microsoft\Graph.users.messages.extensions collection
     * @param ?String $id $id Unique identifier of the item
     * @return ExtensionRequestBuilder
    */
    public function extensionsById(?String $id): ExtensionRequestBuilder {
    }

    /**
     * The messages in a mailbox or folder. Read-only. Nullable.
     * @param GetQueryParameters|null $q $q Request query parameters
     * @param array|null $headers $h Request headers
     * @param array|null $options $o Request options for HTTP middlewares
     * @param ResponseHandlerInterface|null $responseHandler $responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return Message|null
    */
    public function get(GetQueryParameters $q, array $headers, array $options, ResponseHandlerInterface $responseHandler): ?Message {
    }

    /**
     * Gets an item from the Microsoft\Graph.users.messages.multiValueExtendedProperties collection
     * @param ?String $id $id Unique identifier of the item
     * @return MultiValueLegacyExtendedPropertyRequestBuilder
    */
    public function multiValueExtendedPropertiesById(?String $id): MultiValueLegacyExtendedPropertyRequestBuilder {
    }

    /**
     * The messages in a mailbox or folder. Read-only. Nullable.
     * @param ?Message $body $body 
     * @param array|null $headers $h Request headers
     * @param array|null $options $o Request options for HTTP middlewares
     * @param ResponseHandlerInterface|null $responseHandler $responseHandler Response handler to use in place of the default response handling provided by the core service
    */
    public function patch(?Message $body, array $headers, array $options, ResponseHandlerInterface $responseHandler): void {
    }

    /**
     * Gets an item from the Microsoft\Graph.users.messages.singleValueExtendedProperties collection
     * @param ?String $id $id Unique identifier of the item
     * @return SingleValueLegacyExtendedPropertyRequestBuilder
    */
    public function singleValueExtendedPropertiesById(?String $id): SingleValueLegacyExtendedPropertyRequestBuilder {
    }

}
