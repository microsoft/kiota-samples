<?php

namespace Microsoft\Graph;

use Microsoft\Graph\Users\Item\UserRequestBuilder;
use Microsoft\Graph\Users\UsersRequestBuilder;
use Microsoft\Kiota\Abstractions\RequestAdapter;

class ApiClient 
{
    /** @var array|null $pathParameters Path parameters for the request */
    private array $pathParameters;
    
    /** @var RequestAdapter|null $requestAdapter The request adapter to use to execute the requests. */
    private RequestAdapter $requestAdapter;
    
    /** @var string|null $urlTemplate Url template to use to build the URL for the current request builder */
    private ?string $urlTemplate;
    
    public function users(): UsersRequestBuilder {
        return new UsersRequestBuilder($this->pathParameters, $this->requestAdapter);
    }
    
    /**
     * Instantiates a new ApiClient and sets the default values.
     * @param RequestAdapter $requestAdapter The request adapter to use to execute the requests.
    */
    public function __construct(RequestAdapter $requestAdapter) {
        $this->pathParameters = [];
        $this->urlTemplate = '{+baseurl}';
        $this->requestAdapter = $requestAdapter;
        $this->requestAdapter->setBaseUrl('https://graph.microsoft.com/v1.0');
    }

    /**
     * Gets an item from the Microsoft\Graph.users.item collection
     * @param string $id Unique identifier of the item
     * @return UserRequestBuilder
    */
    public function usersById(string $id): UserRequestBuilder {
        $urlTplParams = $this->pathParameters;
        $urlTplParams['user_id'] = $id;
        return new UserRequestBuilder($urlTplParams, $this->requestAdapter);
    }

}
