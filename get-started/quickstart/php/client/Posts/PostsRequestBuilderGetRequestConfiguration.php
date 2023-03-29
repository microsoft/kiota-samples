<?php

namespace KiotaPosts\Client\Posts;

use Microsoft\Kiota\Abstractions\RequestOption;

/**
 * Configuration for the request such as headers, query parameters, and middleware options.
*/
class PostsRequestBuilderGetRequestConfiguration 
{
    /**
     * @var array<string, array<string>|string>|null $headers Request headers
    */
    public ?array $headers = null;
    
    /**
     * @var array<RequestOption>|null $options Request options
    */
    public ?array $options = null;
    
    /**
     * @var PostsRequestBuilderGetQueryParameters|null $queryParameters Request query parameters
    */
    public ?PostsRequestBuilderGetQueryParameters $queryParameters = null;
    
    /**
     * Instantiates a new postsRequestBuilderGetQueryParameters.
     * @param string|null $title Filter results by title
     * @param int|null $userId Filter results by user ID
     * @return PostsRequestBuilderGetQueryParameters
    */
    public static function addQueryParameters(?string $title = null, ?int $userId = null): PostsRequestBuilderGetQueryParameters {
        return new PostsRequestBuilderGetQueryParameters($title, $userId);
    }

    /**
     * Instantiates a new postsRequestBuilderGetRequestConfiguration and sets the default values.
     * @param array<string, array<string>|string>|null $headers Request headers
     * @param array<RequestOption>|null $options Request options
     * @param PostsRequestBuilderGetQueryParameters|null $queryParameters Request query parameters
    */
    public function __construct(?array $headers = null, ?array $options = null, ?PostsRequestBuilderGetQueryParameters $queryParameters = null) {
        $this->headers = $headers;
        $this->options = $options;
        $this->queryParameters = $queryParameters;
    }

}
