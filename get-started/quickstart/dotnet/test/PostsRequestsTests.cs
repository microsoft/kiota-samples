// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

// <UnitTestsSnippet>
using System.Text.Json;
using KiotaPosts.Client;
using KiotaPosts.Client.Models;
using Microsoft.Kiota.Abstractions;
using Microsoft.Kiota.Abstractions.Serialization;
using Microsoft.Kiota.Serialization.Json;
using NSubstitute;

namespace KiotaPostsTests;

public class PostsRequestsTests
{
    private static readonly List<Post> postsMock = new()
    {
        new Post
        {
            Id = 1,
            UserId = 1,
            Title = "Post Title #1",
            Body = "Posts Body #1",
        },
        new Post
        {
            Id = 2,
            UserId = 2,
            Title = "Post Title #2",
            Body = "Post Body #2",
        },
    };

    // Basic test for GET /posts
    [Fact]
    public async Task All_Posts()
    {
        // Arrange
        var adapter = Substitute.For<IRequestAdapter>();
        var postsClient = new PostsClient(adapter);

        // Return the mocked list of posts
        adapter.SendCollectionAsync(
            Arg.Any<RequestInformation>(),
            Arg.Any<ParsableFactory<Post>>(),
            Arg.Any<Dictionary<string, ParsableFactory<IParsable>>>(),
            Arg.Any<CancellationToken>())
            .ReturnsForAnyArgs(postsMock);

        // Act
        var posts = await postsClient.Posts.GetAsync();

        // Assert
        Assert.NotNull(posts);
        Assert.Equal(postsMock.Count, posts.Count);
    }

    // Basic test for GET /posts/{id}
    [Fact]
    public async Task Post_By_Id()
    {
        // Arrange
        var adapter = Substitute.For<IRequestAdapter>();
        var postsClient = new PostsClient(adapter);

        // Return the mocked post #1 object if the path
        // contains the ID of post #1
        adapter.SendAsync(
            Arg.Is<RequestInformation>(req => req.PathParameters.Values.Contains(1)),
            Arg.Any<ParsableFactory<Post>>(),
            Arg.Any<Dictionary<string, ParsableFactory<IParsable>>>(),
            Arg.Any<CancellationToken>())
            .Returns(postsMock[0]);

        // Return the mocked post #2 object if the path
        // contains the ID of post #2
        adapter.SendAsync(
            Arg.Is<RequestInformation>(req => req.PathParameters.Values.Contains(2)),
            Arg.Any<ParsableFactory<Post>>(),
            Arg.Any<Dictionary<string, ParsableFactory<IParsable>>>(),
            Arg.Any<CancellationToken>())
            .Returns(postsMock[1]);

        // Act
        var post1 = await postsClient.Posts[1].GetAsync();
        var post2 = await postsClient.Posts[2].GetAsync();

        // Assert
        Assert.NotNull(post1);
        Assert.Equal(postsMock[0].Title, post1.Title);
        Assert.NotNull(post2);
        Assert.Equal(postsMock[1].Title, post2.Title);
    }

    // Basic test for POST /posts
    [Fact]
    public async Task Post_Create_New()
    {
        // Arrange
        var adapter = Substitute.For<IRequestAdapter>();
        var postsClient = new PostsClient(adapter);

        // Add JsonSerializationWriter to serialize Post objects
        adapter.SerializationWriterFactory.GetSerializationWriter("application/json")
            .Returns(sw => new JsonSerializationWriter());

        // When the request is sent through the adapter
        // save it so we can validate the content of the request
        RequestInformation? postRequest = null;
        await adapter.SendAsync(
            Arg.Do<RequestInformation>(req => postRequest = req),
            Arg.Any<ParsableFactory<Post>>(),
            Arg.Any<Dictionary<string, ParsableFactory<IParsable>>?>(),
            Arg.Any<CancellationToken>());

        var newPost = new Post
        {
            UserId = 3,
            Title = "Created new post",
            Body = "For testing purposes",
        };

        // Act
        await postsClient.Posts.PostAsync(newPost);

        // Assert
        Assert.NotNull(postRequest);
        // Deserialize the request body
        var postFromBody = JsonSerializer.Deserialize<Post>(
            postRequest.Content,
            new JsonSerializerOptions
            {
                PropertyNamingPolicy = JsonNamingPolicy.CamelCase,
            });

        Assert.NotNull(postFromBody);
        Assert.Equal(newPost.UserId, postFromBody.UserId);
        Assert.Equal(newPost.Title, postFromBody.Title);
        Assert.Equal(newPost.Body, postFromBody.Body);
    }
}
// </UnitTestsSnippet>
