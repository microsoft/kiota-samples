package main

import (
	"context"
	"testing"

	"kiota_posts/client"
	"kiota_posts/client/models"
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

func TestPostsClient_ToPostRequestInformation(t *testing.T) {
	postsClient, _ := createTestClient(t)
	newPost := models.NewPost()
	userId := int32(42)
	title := "New Post"
	body := "Hello world"
	newPost.SetUserId(&userId)
	newPost.SetTitle(&title)
	newPost.SetBody(&body)

	reqInfo, err := postsClient.Posts().ToPostRequestInformation(context.Background(), newPost, nil)
	if err != nil {
		t.Fatalf("failed to create post request information: %v", err)
	}
	if reqInfo.Method != abstractions.POST {
		t.Errorf("expected POST method, got %v", reqInfo.Method)
	}
	acceptHeaders := reqInfo.Headers.Get("Accept")
	if len(acceptHeaders) == 0 || acceptHeaders[0] != "application/json" {
		t.Errorf("expected Accept header 'application/json', got %v", acceptHeaders)
	}
	if len(reqInfo.Content) == 0 {
		t.Error("expected non-empty post request body content")
	}
}

func TestPostsClient_ByPostIdInteger(t *testing.T) {
	postsClient, _ := createTestClient(t)
	itemBuilder := postsClient.Posts().ByPostIdInteger(5)
	if itemBuilder == nil {
		t.Fatal("expected non-nil PostItemRequestBuilder")
	}

	// GET /posts/{id}
	getReq, err := itemBuilder.ToGetRequestInformation(context.Background(), nil)
	if err != nil {
		t.Fatalf("failed to create get request information: %v", err)
	}
	if getReq.Method != abstractions.GET {
		t.Errorf("expected GET method, got %v", getReq.Method)
	}
	hasId := false
	for _, val := range getReq.PathParameters {
		if val == "5" {
			hasId = true
			break
		}
	}
	if !hasId {
		t.Errorf("expected path parameters to contain '5', got %v", getReq.PathParameters)
	}

	// PATCH /posts/{id}
	update := models.NewPost()
	newTitle := "Updated title"
	update.SetTitle(&newTitle)
	patchReq, err := itemBuilder.ToPatchRequestInformation(context.Background(), update, nil)
	if err != nil {
		t.Fatalf("failed to create patch request information: %v", err)
	}
	if patchReq.Method != abstractions.PATCH {
		t.Errorf("expected PATCH method, got %v", patchReq.Method)
	}
	if len(patchReq.Content) == 0 {
		t.Error("expected non-empty patch request body content")
	}

	// DELETE /posts/{id}
	delReq, err := itemBuilder.ToDeleteRequestInformation(context.Background(), nil)
	if err != nil {
		t.Fatalf("failed to create delete request information: %v", err)
	}
	if delReq.Method != abstractions.DELETE {
		t.Errorf("expected DELETE method, got %v", delReq.Method)
	}
}

func TestPostsClient_ByPostIdString(t *testing.T) {
	postsClient, _ := createTestClient(t)
	itemBuilder := postsClient.Posts().ByPostId("5")
	if itemBuilder == nil {
		t.Fatal("expected non-nil PostItemRequestBuilder")
	}

	getReq, err := itemBuilder.ToGetRequestInformation(context.Background(), nil)
	if err != nil {
		t.Fatalf("failed to create get request information: %v", err)
	}
	if getReq.Method != abstractions.GET {
		t.Errorf("expected GET method, got %v", getReq.Method)
	}
	hasId := false
	for _, val := range getReq.PathParameters {
		if val == "5" {
			hasId = true
			break
		}
	}
	if !hasId {
		t.Errorf("expected path parameters to contain '5', got %v", getReq.PathParameters)
	}
}
