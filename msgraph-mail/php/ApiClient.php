<?php

namespace Microsoft\Graph;

use Microsoft\Graph\Users\UsersRequestBuilder;
use Microsoft\Kiota\Abstractions\HttpCore;
use Microsoft\Kiota\Abstractions\HttpMethod;
use Microsoft\Kiota\Abstractions\MiddlewareOption;
use Microsoft\Kiota\Abstractions\QueryParametersBase;
use Microsoft\Kiota\Abstractions\RequestInformation;
use Microsoft\Kiota\Abstractions\ResponseHandler;
use Microsoft\Kiota\Abstractions\Serialization\ParseNode;

class ApiClient 
{
    private HttpCoreInterface $httpCore;
    
    /** @var string $pathSegment Path segment to use to build the URL for the current request builder */
    private string $pathSegment;
    
    public function users(): UsersRequestBuilder {
        return new UsersRequestBuilder($this->pathSegment, $this->httpCore);
    }
    
    /**
     * Instantiates a new ApiClient and sets the default values.
     * @param ?HttpCore $httpCore $httpCore The http core service to use to execute the requests.
    */
    public function __construct(?HttpCore $httpCore) {
    }

    /**
     * Gets an item from the Microsoft\Graph.users collection
     * @param ?String $id $id Unique identifier of the item
     * @return UserRequestBuilder
    */
    public function usersById(?String $id): UserRequestBuilder {
    }

}
