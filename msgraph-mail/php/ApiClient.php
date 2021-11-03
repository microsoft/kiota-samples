<?php

namespace Microsoft\Graph;

use Microsoft\Graph\Users\UsersRequestBuilder;
use Microsoft\Kiota\Abstractions\RequestAdapter;

class ApiClient 
{
    /** @var array $pathParameters Path parameters for the request */
    private array $pathParameters;
    
    /** @var IRequestAdapter $requestAdapter The request adapter to use to execute the requests. */
    private RequestAdapter $requestAdapter;
    
    /** @var string $urlTemplate Url template to use to build the URL for the current request builder */
    private string $urlTemplate;
    
    public function users(): UsersRequestBuilder {
        return new UsersRequestBuilder($this->pathParameters, $this->requestAdapter);
    }
    
    /**
     * Instantiates a new ApiClient and sets the default values.
     * @param RequestAdapter $requestAdapter The request adapter to use to execute the requests.
    */
    public function __construct(RequestAdapter $requestAdapter) {
    }

    /**
     * Gets an item from the Microsoft\Graph.users.item collection
     * @param String $id Unique identifier of the item
     * @return UserRequestBuilder
    */
    public function usersById(String $id): UserRequestBuilder {
    }

}
