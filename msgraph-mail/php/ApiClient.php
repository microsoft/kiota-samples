<?php

namespace Microsoft\Graph;

use Microsoft\Graph\Users\UsersRequestBuilder;
use Microsoft\Kiota\Abstractions\RequestAdapter;

class ApiClient 
{
    /** @var string $pathSegment Path segment to use to build the URL for the current request builder */
    private string $pathSegment;
    
    private RequestAdapter $requestAdapter;
    
    public function users(): UsersRequestBuilder {
        return new UsersRequestBuilder($this->pathSegment, true , $this->requestAdapter);
    }

    /**
     * Instantiates a new ApiClient and sets the default values.
     * @param RequestAdapter $requestAdapter The http core service to use to execute the requests.
     * @throws \Exception
     */
    public function __construct(RequestAdapter $requestAdapter) {
        if (is_null($requestAdapter)) {
            throw new \Exception('$requestAdapter cannot be null');
        }
    }

    /**
     * Gets an item from the Microsoft\Graph.users.item collection
     * @param String $id Unique identifier of the item
     * @return UserRequestBuilder
     * @throws \Exception
     */
    public function usersById(String $id): UserRequestBuilder {
        if (is_null($id)) {
            throw new \Exception('$id cannot be null');
        }
    }

}
