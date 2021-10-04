<?php

namespace Microsoft\Graph\Users\Item;

use Microsoft\Graph\Users\InferenceClassification\InferenceClassificationRequestBuilder;
use Microsoft\Graph\Users\MailFolders\MailFoldersRequestBuilder;
use Microsoft\Graph\Users\Messages\MessagesRequestBuilder;
use Microsoft\Kiota\Abstractions\HttpCore;
use Microsoft\Kiota\Abstractions\HttpMethod;
use Microsoft\Kiota\Abstractions\MiddlewareOption;
use Microsoft\Kiota\Abstractions\QueryParametersBase;
use Microsoft\Kiota\Abstractions\RequestInformation;
use Microsoft\Kiota\Abstractions\ResponseHandler;
use Microsoft\Kiota\Abstractions\Serialization\ParseNode;

class UserRequestBuilder 
{
    /** @var string $currentPath Current path for the request */
    private string $currentPath;
    
    private HttpCoreInterface $httpCore;
    
    public function inferenceClassification(): InferenceClassificationRequestBuilder {
        return new InferenceClassificationRequestBuilder($this->currentPath . $this->pathSegment, $this->httpCore);
    }
    
    /** @var bool $isRawUrl Whether the current path is a raw URL */
    private bool $isRawUrl;
    
    public function mailFolders(): MailFoldersRequestBuilder {
        return new MailFoldersRequestBuilder($this->currentPath . $this->pathSegment, $this->httpCore);
    }
    
    public function messages(): MessagesRequestBuilder {
        return new MessagesRequestBuilder($this->currentPath . $this->pathSegment, $this->httpCore);
    }
    
    /** @var string $pathSegment Path segment to use to build the URL for the current request builder */
    private string $pathSegment;
    
    /**
     * Instantiates a new UserRequestBuilder and sets the default values.
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
     * Gets an item from the Microsoft\Graph.users.mailFolders collection
     * @param ?String $id $id Unique identifier of the item
     * @return MailFolderRequestBuilder
    */
    public function mailFoldersById(?String $id): MailFolderRequestBuilder {
    }

    /**
     * Gets an item from the Microsoft\Graph.users.messages collection
     * @param ?String $id $id Unique identifier of the item
     * @return MessageRequestBuilder
    */
    public function messagesById(?String $id): MessageRequestBuilder {
    }

}
