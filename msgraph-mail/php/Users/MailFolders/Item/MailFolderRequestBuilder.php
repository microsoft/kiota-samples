<?php

namespace Microsoft\Graph\Users\MailFolders\Item;

use Microsoft\Graph\Users\MailFolder;
use Microsoft\Graph\Users\MailFolders\ChildFolders\ChildFoldersRequestBuilder;
use Microsoft\Graph\Users\MailFolders\MessageRules\MessageRulesRequestBuilder;
use Microsoft\Graph\Users\MailFolders\Messages\MessagesRequestBuilder;
use Microsoft\Graph\Users\MailFolders\MultiValueExtendedProperties\MultiValueExtendedPropertiesRequestBuilder;
use Microsoft\Graph\Users\MailFolders\SingleValueExtendedProperties\SingleValueExtendedPropertiesRequestBuilder;
use Microsoft\Kiota\Abstractions\HttpCore;
use Microsoft\Kiota\Abstractions\HttpMethod;
use Microsoft\Kiota\Abstractions\MiddlewareOption;
use Microsoft\Kiota\Abstractions\QueryParametersBase;
use Microsoft\Kiota\Abstractions\RequestInformation;
use Microsoft\Kiota\Abstractions\ResponseHandler;
use Microsoft\Kiota\Abstractions\Serialization\ParseNode;

class MailFolderRequestBuilder 
{
    public function childFolders(): ChildFoldersRequestBuilder {
        return new ChildFoldersRequestBuilder($this->currentPath . $this->pathSegment, $this->httpCore);
    }
    
    /** @var string $currentPath Current path for the request */
    private string $currentPath;
    
    private HttpCoreInterface $httpCore;
    
    /** @var bool $isRawUrl Whether the current path is a raw URL */
    private bool $isRawUrl;
    
    public function messageRules(): MessageRulesRequestBuilder {
        return new MessageRulesRequestBuilder($this->currentPath . $this->pathSegment, $this->httpCore);
    }
    
    public function messages(): MessagesRequestBuilder {
        return new MessagesRequestBuilder($this->currentPath . $this->pathSegment, $this->httpCore);
    }
    
    public function multiValueExtendedProperties(): MultiValueExtendedPropertiesRequestBuilder {
        return new MultiValueExtendedPropertiesRequestBuilder($this->currentPath . $this->pathSegment, $this->httpCore);
    }
    
    /** @var string $pathSegment Path segment to use to build the URL for the current request builder */
    private string $pathSegment;
    
    public function singleValueExtendedProperties(): SingleValueExtendedPropertiesRequestBuilder {
        return new SingleValueExtendedPropertiesRequestBuilder($this->currentPath . $this->pathSegment, $this->httpCore);
    }
    
    /**
     * Gets an item from the Microsoft\Graph.users.mailFolders.childFolders collection
     * @param ?String $id $id Unique identifier of the item
     * @return MailFolderRequestBuilder
    */
    public function childFoldersById(?String $id): MailFolderRequestBuilder {
    }

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
     * The user's mail folders. Read-only. Nullable.
     * @param array|null $headers $h Request headers
     * @param array|null $options $o Request options for HTTP middlewares
     * @return RequestInformation
    */
    public function createDeleteRequestInformation(array $headers, array $options): RequestInformation {
    }

    /**
     * The user's mail folders. Read-only. Nullable.
     * @param GetQueryParameters|null $q $q Request query parameters
     * @param array|null $headers $h Request headers
     * @param array|null $options $o Request options for HTTP middlewares
     * @return RequestInformation
    */
    public function createGetRequestInformation(GetQueryParameters $q, array $headers, array $options): RequestInformation {
    }

    /**
     * The user's mail folders. Read-only. Nullable.
     * @param ?MailFolder $body $body 
     * @param array|null $headers $h Request headers
     * @param array|null $options $o Request options for HTTP middlewares
     * @return RequestInformation
    */
    public function createPatchRequestInformation(?MailFolder $body, array $headers, array $options): RequestInformation {
    }

    /**
     * The user's mail folders. Read-only. Nullable.
     * @param array|null $headers $h Request headers
     * @param array|null $options $o Request options for HTTP middlewares
     * @param ResponseHandlerInterface|null $responseHandler $responseHandler Response handler to use in place of the default response handling provided by the core service
    */
    public function delete(array $headers, array $options, ResponseHandlerInterface $responseHandler): void {
    }

    /**
     * The user's mail folders. Read-only. Nullable.
     * @param GetQueryParameters|null $q $q Request query parameters
     * @param array|null $headers $h Request headers
     * @param array|null $options $o Request options for HTTP middlewares
     * @param ResponseHandlerInterface|null $responseHandler $responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return MailFolder|null
    */
    public function get(GetQueryParameters $q, array $headers, array $options, ResponseHandlerInterface $responseHandler): ?MailFolder {
    }

    /**
     * Gets an item from the Microsoft\Graph.users.mailFolders.messageRules collection
     * @param ?String $id $id Unique identifier of the item
     * @return MessageRuleRequestBuilder
    */
    public function messageRulesById(?String $id): MessageRuleRequestBuilder {
    }

    /**
     * Gets an item from the Microsoft\Graph.users.mailFolders.messages collection
     * @param ?String $id $id Unique identifier of the item
     * @return MessageRequestBuilder
    */
    public function messagesById(?String $id): MessageRequestBuilder {
    }

    /**
     * Gets an item from the Microsoft\Graph.users.mailFolders.multiValueExtendedProperties collection
     * @param ?String $id $id Unique identifier of the item
     * @return MultiValueLegacyExtendedPropertyRequestBuilder
    */
    public function multiValueExtendedPropertiesById(?String $id): MultiValueLegacyExtendedPropertyRequestBuilder {
    }

    /**
     * The user's mail folders. Read-only. Nullable.
     * @param ?MailFolder $body $body 
     * @param array|null $headers $h Request headers
     * @param array|null $options $o Request options for HTTP middlewares
     * @param ResponseHandlerInterface|null $responseHandler $responseHandler Response handler to use in place of the default response handling provided by the core service
    */
    public function patch(?MailFolder $body, array $headers, array $options, ResponseHandlerInterface $responseHandler): void {
    }

    /**
     * Gets an item from the Microsoft\Graph.users.mailFolders.singleValueExtendedProperties collection
     * @param ?String $id $id Unique identifier of the item
     * @return SingleValueLegacyExtendedPropertyRequestBuilder
    */
    public function singleValueExtendedPropertiesById(?String $id): SingleValueLegacyExtendedPropertyRequestBuilder {
    }

}
