// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

// <ProgramSnippet>
package kiotaposts;

import com.microsoft.kiota.authentication.AnonymousAuthenticationProvider;
import com.microsoft.kiota.http.OkHttpRequestAdapter;

import kiotaposts.client.PostsClient;
import kiotaposts.client.models.Post;

public class App {

    public static void main(String[] args) {
        // API requires no authentication, so use the anonymous
        // authentication provider
        final AnonymousAuthenticationProvider authProvider =
            new AnonymousAuthenticationProvider();
        final OkHttpRequestAdapter adapter = new OkHttpRequestAdapter(authProvider);

        final PostsClient client = new PostsClient(adapter);

        // GET /posts
        client.posts().get()
            .thenAccept(allPosts -> {
                System.out.printf("Retrieved %d posts.%n", allPosts.size());
            })
            .exceptionally(err -> {
                System.out.printf("Error: %s%n", err.getMessage());
                return null;
            })
            .join();

        // GET /posts/{id}
        final String specificPostId = "5";
        client.posts(specificPostId).get()
            .thenAccept(specificPost -> {
                System.out.printf("Retrieved post - ID: %d, Title: %s, Body: %s%n",
                    specificPost.getId(), specificPost.getTitle(), specificPost.getBody());
            })
            .exceptionally(err -> {
                System.out.printf("Error: %s%n", err.getMessage());
                return null;
            })
            .join();

        // POST /posts
        final Post newPost = new Post();
        newPost.setUserId(42);
        newPost.setTitle("Testing Kiota-generated API client");
        newPost.setBody("Hello world!");

        client.posts().post(newPost)
            .thenAccept(createdPost -> {
                System.out.printf("Created new post with ID: %d%n", createdPost.getId());
            })
            .exceptionally(err -> {
                System.out.printf("Error: %s%n", err.getMessage());
                return null;
            })
            .join();

        // PATCH /posts/{id}
        final Post update = new Post();
        // Only update title
        update.setTitle("Updated title");

        client.posts(specificPostId).patch(update)
            .thenAccept(updatedPost -> {
                System.out.printf("Updated post - ID: %d, Title: %s, Body: %s%n",
                    updatedPost.getId(), updatedPost.getTitle(), updatedPost.getBody());
            })
            .exceptionally(err -> {
                System.out.printf("Error: %s%n", err.getMessage());
                return null;
            })
            .join();

        // DELETE /posts/{id}
        client.posts(specificPostId).delete()
            .thenAccept(s -> {
                System.out.printf("Deleted post%n");
            })
            .exceptionally(err -> {
                System.out.printf("Error: %s%n", err.getMessage());
                return null;
            })
            .join();
    }
}
// </ProgramSnippet>
