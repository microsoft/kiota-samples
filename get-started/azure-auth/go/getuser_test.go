package main

import (
	"context"
	"strings"
	"testing"

	"getuser/client"
	"getuser/client/me"
	"getuser/client/models"

	abstractions "github.com/microsoft/kiota-abstractions-go"
	auth "github.com/microsoft/kiota-abstractions-go/authentication"
	bundle "github.com/microsoft/kiota-bundle-go"
	serjson "github.com/microsoft/kiota-serialization-json-go"
)

func TestGraphApiClient_Initialization(t *testing.T) {
	authProvider := auth.AnonymousAuthenticationProvider{}
	adapter, err := bundle.NewDefaultRequestAdapter(&authProvider)
	if err != nil {
		t.Fatalf("failed to create default request adapter: %v", err)
	}

	apiClient := client.NewGraphApiClient(adapter)
	if apiClient == nil {
		t.Fatal("expected non-nil GraphApiClient")
	}
	if apiClient.Me() == nil {
		t.Fatal("expected non-nil MeRequestBuilder")
	}
	if adapter.GetBaseUrl() != "https://graph.microsoft.com/v1.0" {
		t.Errorf("expected base URL 'https://graph.microsoft.com/v1.0', got '%s'", adapter.GetBaseUrl())
	}
}

func TestGraphApiClient_MeToGetRequestInformation(t *testing.T) {
	authProvider := auth.AnonymousAuthenticationProvider{}
	adapter, err := bundle.NewDefaultRequestAdapter(&authProvider)
	if err != nil {
		t.Fatalf("failed to create default request adapter: %v", err)
	}

	apiClient := client.NewGraphApiClient(adapter)
	reqInfo, err := apiClient.Me().ToGetRequestInformation(context.Background(), nil)
	if err != nil {
		t.Fatalf("failed to create request information: %v", err)
	}
	if reqInfo.Method != abstractions.GET {
		t.Errorf("expected GET method, got %v", reqInfo.Method)
	}
	if reqInfo.UrlTemplate != "{+baseurl}/me" {
		t.Errorf("expected url template '{+baseurl}/me', got '%s'", reqInfo.UrlTemplate)
	}
	acceptHeaders := reqInfo.Headers.Get("Accept")
	if len(acceptHeaders) == 0 || acceptHeaders[0] != "application/json" {
		t.Errorf("expected Accept header 'application/json', got %v", acceptHeaders)
	}
}

func TestUserModel_LifecycleAndSerialization(t *testing.T) {
	parsable, err := models.CreateUserFromDiscriminatorValue(nil)
	if err != nil {
		t.Fatalf("unexpected error from discriminator: %v", err)
	}
	if parsable == nil {
		t.Fatal("expected non-nil parsable from discriminator")
	}

	user := models.NewUser()
	if user == nil {
		t.Fatal("expected non-nil User from NewUser")
	}

	displayName := "Jane Doe"
	id := "user-12345"
	user.SetDisplayName(&displayName)
	user.SetId(&id)

	if user.GetDisplayName() == nil || *user.GetDisplayName() != "Jane Doe" {
		t.Errorf("expected displayName 'Jane Doe', got %v", user.GetDisplayName())
	}
	if user.GetId() == nil || *user.GetId() != "user-12345" {
		t.Errorf("expected ID 'user-12345', got %v", user.GetId())
	}
	if user.GetAdditionalData() == nil {
		t.Error("expected non-nil AdditionalData")
	}

	deserializers := user.GetFieldDeserializers()
	for _, field := range []string{"displayName", "id"} {
		if _, ok := deserializers[field]; !ok {
			t.Errorf("expected field deserializer for '%s'", field)
		}
	}

	writer := serjson.NewJsonSerializationWriter()
	err = user.Serialize(writer)
	if err != nil {
		t.Fatalf("failed to serialize user: %v", err)
	}

	content, err := writer.GetSerializedContent()
	if err != nil {
		t.Fatalf("failed to get serialized content: %v", err)
	}
	jsonStr := string(content)
	if !strings.Contains(jsonStr, `"displayName":"Jane Doe"`) {
		t.Errorf("expected serialized content to contain displayName, got: %s", jsonStr)
	}
	if !strings.Contains(jsonStr, `"id":"user-12345"`) {
		t.Errorf("expected serialized content to contain id, got: %s", jsonStr)
	}
}

func TestGraphApiClient_MeWithUrl(t *testing.T) {
	authProvider := auth.AnonymousAuthenticationProvider{}
	adapter, err := bundle.NewDefaultRequestAdapter(&authProvider)
	if err != nil {
		t.Fatalf("failed to create default request adapter: %v", err)
	}

	apiClient := client.NewGraphApiClient(adapter)
	customMe := apiClient.Me().WithUrl("https://graph.microsoft.com/v1.0/me/custom")
	if customMe == nil {
		t.Fatal("expected non-nil MeRequestBuilder from WithUrl")
	}

	reqInfo, err := customMe.ToGetRequestInformation(context.Background(), nil)
	if err != nil {
		t.Fatalf("failed to create request information from customMe: %v", err)
	}
	uri, err := reqInfo.GetUri()
	if err != nil {
		t.Fatalf("failed to get URI from reqInfo: %v", err)
	}
	if uri.String() != "https://graph.microsoft.com/v1.0/me/custom" {
		t.Errorf("expected URI 'https://graph.microsoft.com/v1.0/me/custom', got '%s'", uri.String())
	}
}

func TestGraphApiClient_MeToGetRequestInformation_WithOptionsAndHeaders(t *testing.T) {
	authProvider := auth.AnonymousAuthenticationProvider{}
	adapter, err := bundle.NewDefaultRequestAdapter(&authProvider)
	if err != nil {
		t.Fatalf("failed to create default request adapter: %v", err)
	}

	apiClient := client.NewGraphApiClient(adapter)
	headers := abstractions.NewRequestHeaders()
	headers.Add("ConsistencyLevel", "eventual")
	config := &me.MeRequestBuilderGetRequestConfiguration{
		Headers: headers,
	}

	reqInfo, err := apiClient.Me().ToGetRequestInformation(context.Background(), config)
	if err != nil {
		t.Fatalf("failed to create request information with config: %v", err)
	}
	consistencyValues := reqInfo.Headers.Get("ConsistencyLevel")
	if len(consistencyValues) == 0 || consistencyValues[0] != "eventual" {
		t.Errorf("expected ConsistencyLevel header 'eventual', got %v", consistencyValues)
	}
}
