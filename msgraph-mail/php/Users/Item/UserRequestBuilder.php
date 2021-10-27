<?php

namespace Microsoft\Graph\Users\Item;

use Microsoft\Graph\Users\Item\InferenceClassification\InferenceClassificationRequestBuilder;
use Microsoft\Graph\Users\Item\MailFolders\MailFoldersRequestBuilder;
use Microsoft\Graph\Users\Item\Messages\MessagesRequestBuilder;
use Microsoft\Kiota\Abstractions\RequestAdapter;

class UserRequestBuilder 
{
    /** @var string $currentPath Current path for the request */
    private string $currentPath;
    
    public function inferenceClassification(): InferenceClassificationRequestBuilder {
        return new InferenceClassificationRequestBuilder($this->currentPath . $this->pathSegment, $this->requestAdapter);
    }
    
    /** @var bool $isRawUrl Whether the current path is a raw URL */
    private bool $isRawUrl;
    
    public function mailFolders(): MailFoldersRequestBuilder {
        return new MailFoldersRequestBuilder($this->currentPath . $this->pathSegment, $this->requestAdapter);
    }
    
    public function messages(): MessagesRequestBuilder {
        return new MessagesRequestBuilder($this->currentPath . $this->pathSegment, $this->requestAdapter);
    }
    
    /** @var string $pathSegment Path segment to use to build the URL for the current request builder */
    private string $pathSegment;
    
    private RequestAdapter $requestAdapter;
    
    /**
     * Instantiates a new UserRequestBuilder and sets the default values.
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
     * Gets an item from the Microsoft\Graph.users.item.mailFolders.item collection
     * @param String $id $id Unique identifier of the item
     * @return MailFolderRequestBuilder
    */
    public function mailFoldersById(String $id): MailFolderRequestBuilder {
    }

    /**
     * Gets an item from the Microsoft\Graph.users.item.messages.item collection
     * @param String $id $id Unique identifier of the item
     * @return MessageRequestBuilder
    */
    public function messagesById(String $id): MessageRequestBuilder {
    }

}
