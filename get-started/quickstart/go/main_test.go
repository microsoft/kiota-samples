package main

import (
	"context"
	"testing"

	"kiota_posts/client"
	"kiota_posts/client/posts"

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

func TestPostsClient_ToGetRequestInformation(t *testing.T) {
	postsClient, _ := createTestClient(t)
	reqInfo, err := postsClient.Posts().ToGetRequestInformation(context.Background(), nil)
	if err != nil {
		t.Fatalf("failed to create get request information: %v", err)
	}
	if reqInfo.Method != abstractions.GET {
		t.Errorf("expected GET method, got %v", reqInfo.Method)
	}
	if reqInfo.UrlTemplate != "{+baseurl}/posts{?title*,userId*}" {
		t.Errorf("expected url template '{+baseurl}/posts{?title*,userId*}', got '%s'", reqInfo.UrlTemplate)
	}
	acceptHeaders := reqInfo.Headers.Get("Accept")
	if len(acceptHeaders) == 0 || acceptHeaders[0] != "application/json" {
		t.Errorf("expected Accept header 'application/json', got %v", acceptHeaders)
	}
}

func TestPostsClient_ToGetRequestInformation_WithQueryParameters(t *testing.T) {
	postsClient, _ := createTestClient(t)
	title := "Test Title"
	userId := int32(42)
	config := &posts.PostsRequestBuilderGetRequestConfiguration{
		QueryParameters: &posts.PostsRequestBuilderGetQueryParameters{
			Title:  &title,
			UserId: &userId,
		},
	}
	reqInfo, err := postsClient.Posts().ToGetRequestInformation(context.Background(), config)
	if err != nil {
		t.Fatalf("failed to create get request information with config: %v", err)
	}
	if reqInfo.QueryParameters["title"] != "Test Title" {
		t.Errorf("expected query parameter title 'Test Title', got '%s'", reqInfo.QueryParameters["title"])
	}
	if reqInfo.QueryParameters["userId"] != "42" {
		t.Errorf("expected query parameter userId '42', got '%s'", reqInfo.QueryParameters["userId"])
	}
}
