package main

import (
	"context"
	"testing"

	"getuser/client"

	abstractions "github.com/microsoft/kiota-abstractions-go"
	auth "github.com/microsoft/kiota-abstractions-go/authentication"
	bundle "github.com/microsoft/kiota-bundle-go"
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
