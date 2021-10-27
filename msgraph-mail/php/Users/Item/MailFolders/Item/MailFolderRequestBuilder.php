<?php

namespace Microsoft\Graph\Users\Item\MailFolders\Item;

use Microsoft\Graph\Models\Microsoft\Graph\MailFolder;
use Microsoft\Graph\Users\Item\MailFolders\Item\ChildFolders\ChildFoldersRequestBuilder;
use Microsoft\Graph\Users\Item\MailFolders\Item\MessageRules\MessageRulesRequestBuilder;
use Microsoft\Graph\Users\Item\MailFolders\Item\Messages\MessagesRequestBuilder;
use Microsoft\Graph\Users\Item\MailFolders\Item\MultiValueExtendedProperties\MultiValueExtendedPropertiesRequestBuilder;
use Microsoft\Graph\Users\Item\MailFolders\Item\SingleValueExtendedProperties\SingleValueExtendedPropertiesRequestBuilder;
use Microsoft\Kiota\Abstractions\HttpMethod;
use Microsoft\Kiota\Abstractions\MiddlewareOption;
use Microsoft\Kiota\Abstractions\RequestAdapter;
use Microsoft\Kiota\Abstractions\RequestInformation;
use Microsoft\Kiota\Abstractions\ResponseHandler;
use Microsoft\Kiota\Abstractions\Serialization\Parsable;

class MailFolderRequestBuilder 
{
    public function childFolders(): ChildFoldersRequestBuilder {
        return new ChildFoldersRequestBuilder($this->currentPath . $this->pathSegment, $this->requestAdapter);
    }
    
    /** @var string $currentPath Current path for the request */
    private string $currentPath;
    
    /** @var bool $isRawUrl Whether the current path is a raw URL */
    private bool $isRawUrl;
    
    public function messageRules(): MessageRulesRequestBuilder {
        return new MessageRulesRequestBuilder($this->currentPath . $this->pathSegment, $this->requestAdapter);
    }
    
    public function messages(): MessagesRequestBuilder {
        return new MessagesRequestBuilder($this->currentPath . $this->pathSegment, $this->requestAdapter);
    }
    
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
     * Gets an item from the Microsoft\Graph.users.item.mailFolders.item.childFolders.item collection
     * @param String $id $id Unique identifier of the item
     * @return MailFolderRequestBuilder
    */
    public function childFoldersById(String $id): MailFolderRequestBuilder {
    }

    /**
     * Instantiates a new MailFolderRequestBuilder and sets the default values.
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
     * The user's mail folders. Read-only. Nullable.
     * @param array|null $ $h Request headers
     * @param array|null $options $o Request options
     * @return RequestInformation
    */
    public function createDeleteRequestInformation(?array $, ?array $options): RequestInformation {
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
    public function createPatchRequestInformation(MailFolder $body, ?array $, ?array $options): RequestInformation {
    }

    /**
     * The user's mail folders. Read-only. Nullable.
     * @param array|null $ $h Request headers
     * @param array|null $options $o Request options
     * @param IResponseHandler|null $$responseHandler $responseHandler Response handler to use in place of the default response handling provided by the core service
    */
    public function delete(?array $, ?array $options, ?IResponseHandler $$responseHandler): void {
    }

    /**
     * The user's mail folders. Read-only. Nullable.
     * @param GetQueryParameters|null $q $q Request query parameters
     * @param array|null $ $h Request headers
     * @param array|null $options $o Request options
     * @param IResponseHandler|null $$responseHandler $responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return MailFolder|null
    */
    public function get(?GetQueryParameters $q, ?array $, ?array $options, ?IResponseHandler $$responseHandler): ?MailFolder {
    }

    /**
     * Gets an item from the Microsoft\Graph.users.item.mailFolders.item.messageRules.item collection
     * @param String $id $id Unique identifier of the item
     * @return MessageRuleRequestBuilder
    */
    public function messageRulesById(String $id): MessageRuleRequestBuilder {
    }

    /**
     * Gets an item from the Microsoft\Graph.users.item.mailFolders.item.messages.item collection
     * @param String $id $id Unique identifier of the item
     * @return MessageRequestBuilder
    */
    public function messagesById(String $id): MessageRequestBuilder {
    }

    /**
     * Gets an item from the Microsoft\Graph.users.item.mailFolders.item.multiValueExtendedProperties.item collection
     * @param String $id $id Unique identifier of the item
     * @return MultiValueLegacyExtendedPropertyRequestBuilder
    */
    public function multiValueExtendedPropertiesById(String $id): MultiValueLegacyExtendedPropertyRequestBuilder {
    }

    /**
     * The user's mail folders. Read-only. Nullable.
     * @param MailFolder $body $body 
     * @param array|null $ $h Request headers
     * @param array|null $options $o Request options
     * @param IResponseHandler|null $$responseHandler $responseHandler Response handler to use in place of the default response handling provided by the core service
    */
    public function patch(MailFolder $body, ?array $, ?array $options, ?IResponseHandler $$responseHandler): void {
    }

    /**
     * Gets an item from the Microsoft\Graph.users.item.mailFolders.item.singleValueExtendedProperties.item collection
     * @param String $id $id Unique identifier of the item
     * @return SingleValueLegacyExtendedPropertyRequestBuilder
    */
    public function singleValueExtendedPropertiesById(String $id): SingleValueLegacyExtendedPropertyRequestBuilder {
    }

}
