<?php

namespace KiotaPosts\Client\Posts;

/**
 * Get posts
*/
class PostsRequestBuilderGetQueryParameters 
{
    /**
     * @var string|null $title Filter results by title
    */
    public ?string $title = null;
    
    /**
     * @var int|null $userId Filter results by user ID
    */
    public ?int $userId = null;
    
    /**
     * Instantiates a new PostsRequestBuilderGetQueryParameters and sets the default values.
     * @param string|null $title Filter results by title
     * @param int|null $userId Filter results by user ID
    */
    public function __construct(?string $title = null, ?int $userId = null) {
        $this->title = $title;
        $this->userId = $userId;
    }

}
