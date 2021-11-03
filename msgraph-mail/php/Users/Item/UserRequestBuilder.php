<?php

namespace Microsoft\Graph\Users\Item;

use Microsoft\Graph\Users\Item\InferenceClassification\InferenceClassificationRequestBuilder;
use Microsoft\Graph\Users\Item\MailFolders\MailFoldersRequestBuilder;
use Microsoft\Graph\Users\Item\Messages\MessagesRequestBuilder;
use Microsoft\Kiota\Abstractions\RequestAdapter;

class UserRequestBuilder 
{
    public function inferenceClassification(): InferenceClassificationRequestBuilder {
        return new InferenceClassificationRequestBuilder($this->currentPath . $this->pathSegment, $this->requestAdapter);
    }
    
    public function mailFolders(): MailFoldersRequestBuilder {
        return new MailFoldersRequestBuilder($this->currentPath . $this->pathSegment, $this->requestAdapter);
    }
    
    public function messages(): MessagesRequestBuilder {
        return new MessagesRequestBuilder($this->currentPath . $this->pathSegment, $this->requestAdapter);
    }
    
    /** @var array $pathParameters Path parameters for the request */
    private array $pathParameters;
    
    /** @var IRequestAdapter $requestAdapter The request adapter to use to execute the requests. */
    private RequestAdapter $requestAdapter;
    
    /** @var string $urlTemplate Url template to use to build the URL for the current request builder */
    private string $urlTemplate;
    
    /**
     * Instantiates a new UserRequestBuilder and sets the default values.
     * @param array $pathParameters Path parameters for the request
     * @param RequestAdapter $requestAdapter The request adapter to use to execute the requests.
    */
    public function __construct(array $pathParameters, RequestAdapter $requestAdapter) {
        if (is_null($pathParameters)) {
            throw new \Exception('$pathParameters cannot be null');
        }
        if (is_null($requestAdapter)) {
            throw new \Exception('$requestAdapter cannot be null');
        }
        $this->urlTemplate = 'https://graph.microsoft.com/v1.0/users/{user_id}';
        $this->requestAdapter = $requestAdapter;
        $this->pathParameters = $pathParameters;
    }

    /**
     * Gets an item from the Microsoft\Graph.users.item.mailFolders.item collection
     * @param String $id Unique identifier of the item
     * @return MailFolderRequestBuilder
    */
    public function mailFoldersById(String $id): MailFolderRequestBuilder {
        if (is_null($id)) {
            throw new \Exception('$id cannot be null');
        }
    }

    /**
     * Gets an item from the Microsoft\Graph.users.item.messages.item collection
     * @param String $id Unique identifier of the item
     * @return MessageRequestBuilder
    */
    public function messagesById(String $id): MessageRequestBuilder {
        if (is_null($id)) {
            throw new \Exception('$id cannot be null');
        }
    }

}
