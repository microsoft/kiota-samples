// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// <ProgramSnippet>
package main

import (
	"context"
	"fmt"
	"log"

	"kiota_posts/client"
	"kiota_posts/client/models"

	auth "github.com/microsoft/kiota-abstractions-go/authentication"
	bundle "github.com/microsoft/kiota-bundle-go"
)

func main() {
	// API requires no authentication, so use the anonymous
	// authentication provider
	authProvider := auth.AnonymousAuthenticationProvider{}

	// Create request adapter using the net/http-based implementation
	adapter, err := bundle.NewDefaultRequestAdapter(&authProvider)
	if err != nil {
		log.Fatalf("Error creating request adapter: %v\n", err)
	}

	// Create the API client
	client := client.NewPostsClient(adapter)

	// GET /posts
	allPosts, err := client.Posts().Get(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error getting posts: %v\n", err)
	}
	fmt.Printf("Retrieved %d posts.\n", len(allPosts))

	// GET /posts/{id}
	specificPostId := int32(5)
	specificPost, err := client.Posts().ByPostIdInteger(specificPostId).Get(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error getting post by ID: %v\n", err)
	}
	fmt.Printf("Retrieved post - ID: %d, Title: %s, Body: %s\n", *specificPost.GetId(), *specificPost.GetTitle(), *specificPost.GetBody())

	// POST /posts
	newPost := models.NewPost()
	userId := int32(42)
	newPost.SetUserId(&userId)
	title := "Testing Kiota-generated API client"
	newPost.SetTitle(&title)
	body := "Hello world!"
	newPost.SetBody(&body)

	createdPost, err := client.Posts().Post(context.Background(), newPost, nil)
	if err != nil {
		log.Fatalf("Error creating post: %v\n", err)
	}
	fmt.Printf("Created new post with ID: %d\n", *createdPost.GetId())

	// PATCH /posts/{id}
	update := models.NewPost()
	newTitle := "Updated title"
	update.SetTitle(&newTitle)

	updatedPost, err := client.Posts().ByPostIdInteger(specificPostId).Patch(context.Background(), update, nil)
	if err != nil {
		log.Fatalf("Error updating post: %v\n", err)
	}
	fmt.Printf("Updated post - ID: %d, Title: %s, Body: %s\n", *updatedPost.GetId(), *updatedPost.GetTitle(), *updatedPost.GetBody())

	// DELETE /posts/{id}
	_, err = client.Posts().ByPostIdInteger(specificPostId).Delete(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error deleting post: %v\n", err)
	}
	fmt.Printf("Deleted post\n")
}

// </ProgramSnippet>
