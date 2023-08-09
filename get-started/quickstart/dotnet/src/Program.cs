// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

// <ProgramSnippet>
using KiotaPosts.Client;
using KiotaPosts.Client.Models;
using Microsoft.Kiota.Abstractions.Authentication;
using Microsoft.Kiota.Http.HttpClientLibrary;

// API requires no authentication, so use the anonymous
// authentication provider
var authProvider = new AnonymousAuthenticationProvider();
// Create request adapter using the HttpClient-based implementation
var adapter = new HttpClientRequestAdapter(authProvider);
// Create the API client
var client = new PostsClient(adapter);

try
{
    // GET /posts
    var allPosts = await client.Posts.GetAsync();
    Console.WriteLine($"Retrieved {allPosts?.Count} posts.");

    // GET /posts/{id}
    var specificPostId = 5;
    var specificPost = await client.Posts[specificPostId].GetAsync();
    Console.WriteLine($"Retrieved post - ID: {specificPost?.Id}, Title: {specificPost?.Title}, Body: {specificPost?.Body}");

    // POST /posts
    var newPost = new Post
    {
        UserId = 42,
        Title = "Testing Kiota-generated API client",
        Body = "Hello world!"
    };

    var createdPost = await client.Posts.PostAsync(newPost);
    Console.WriteLine($"Created new post with ID: {createdPost?.Id}");

    // PATCH /posts/{id}
    var update = new Post
    {
        // Only update title
        Title = "Updated title"
    };

    var updatedPost = await client.Posts[specificPostId].PatchAsync(update);
    Console.WriteLine($"Updated post - ID: {updatedPost?.Id}, Title: {updatedPost?.Title}, Body: {updatedPost?.Body}");

    // DELETE /posts/{id}
    await client.Posts[specificPostId].DeleteAsync();
}
catch (Exception ex)
{
    Console.WriteLine($"ERROR: {ex.Message}");
    Console.WriteLine(ex.StackTrace);
}
// </ProgramSnippet>
