package main

import (
	"testing"

	"kiota_posts/client"

	abstractions "github.com/microsoft/kiota-abstractions-go"
	auth "github.com/microsoft/kiota-abstractions-go/authentication"
	bundle "github.com/microsoft/kiota-bundle-go"
)

func createTestClient(t *testing.T) (*client.PostsClient, abstractions.RequestAdapter) {
	t.Helper()
	authProvider := auth.AnonymousAuthenticationProvider{}
	adapter, err := bundle.NewDefaultRequestAdapter(&authProvider)
	if err != nil {
		t.Fatalf("failed to create default request adapter: %v", err)
	}
	return client.NewPostsClient(adapter), adapter
}

func TestPostsClient_Initialization(t *testing.T) {
	postsClient, adapter := createTestClient(t)
	if postsClient == nil {
		t.Fatal("expected non-nil PostsClient")
	}
	if postsClient.Posts() == nil {
		t.Fatal("expected non-nil PostsRequestBuilder")
	}
	if adapter.GetBaseUrl() != "https://jsonplaceholder.typicode.com" {
		t.Errorf("expected base URL 'https://jsonplaceholder.typicode.com', got '%s'", adapter.GetBaseUrl())
	}
}
