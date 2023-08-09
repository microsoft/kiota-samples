<!-- Copyright (c) Microsoft Corporation.
     Licensed under the MIT License.-->

<!-- <ProgramSnippet> -->
<?php

use KiotaPosts\Client\PostsApiClient;
use KiotaPosts\Client\Models\Post;
use Microsoft\Kiota\Abstractions\ApiException;
use Microsoft\Kiota\Abstractions\Authentication\AnonymousAuthenticationProvider;
use Microsoft\Kiota\Http\GuzzleRequestAdapter;

require __DIR__.'/vendor/autoload.php';

try {
    $authProvider = new AnonymousAuthenticationProvider();
    $requestAdapter = new GuzzleRequestAdapter($authProvider);
    $client = new PostsApiClient($requestAdapter);

    // GET /posts
    $allPosts = $client->posts()->get()->wait();
    $postCount = sizeof($allPosts);
    echo "Retrieved {$postCount} posts.\n";

    // GET /posts/{id}
    $specificPostId = 5;
    $specificPost = $client->posts()->byPostId($specificPostId)->get()->wait();
    echo "Retrieved post - ID: {$specificPost->getId()}, Title: {$specificPost->getTitle()}, Body: {$specificPost->getBody()}\n";

    // POST /posts
    $newPost = new Post();
    $newPost->setUserId(42);
    $newPost->setTitle("Testing Kiota-generated API client");
    $newPost->setBody("Hello world!");

    $createdPost = $client->posts()->post($newPost)->wait();
    echo "Created new post with ID: {$createdPost->getId()}\n";

    // PATCH /posts/{id}
    $update = new Post();
    // Only update title
    $update->setTitle("Updated title");

    $updatedPost = $client->posts()->byPostId($specificPostId)->patch($update)->wait();
    echo "Updated post - ID: {$updatedPost->getId()}, Title: {$updatedPost->getTitle()}, Body: {$updatedPost->getBody()}\n";

    // DELETE /posts/{id}
    $client->posts()->byPostId($specificPostId)->delete()->wait();
}catch (ApiException $ex) {
    echo $ex->getMessage();
}
?>
<!-- </ProgramSnippet> -->
