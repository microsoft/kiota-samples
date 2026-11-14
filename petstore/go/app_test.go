package main

import (
	"context"
	"testing"

	absauth "github.com/microsoft/kiota-abstractions-go/authentication"
	r "github.com/microsoft/kiota-http-go"
	u "github.com/microsoft/kiota-samples/petstore/go/utilities"
	"github.com/microsoft/kiota-samples/petstore/go/utilities/models"
	"github.com/microsoft/kiota-samples/petstore/go/utilities/pet"
)

func TestPetstoreApiClient_Initialization(t *testing.T) {
	authProvider := &absauth.AnonymousAuthenticationProvider{}
	adapter, err := r.NewNetHttpRequestAdapter(authProvider)
	if err != nil {
		t.Fatalf("failed to create request adapter: %v", err)
	}

	client := u.NewApiClient(adapter)
	if client == nil {
		t.Fatal("expected non-nil ApiClient")
	}

	baseUrl := adapter.GetBaseUrl()
	expectedBaseUrl := "https://petstore.swagger.io/v2"
	if baseUrl != expectedBaseUrl {
		t.Errorf("expected baseUrl %q, got %q", expectedBaseUrl, baseUrl)
	}
}

func TestPetstoreApiClient_PetToGetRequestInformation(t *testing.T) {
	authProvider := &absauth.AnonymousAuthenticationProvider{}
	adapter, err := r.NewNetHttpRequestAdapter(authProvider)
	if err != nil {
		t.Fatalf("failed to create request adapter: %v", err)
	}

	client := u.NewApiClient(adapter)
	reqInfo, err := client.Pet().ByPetIdInt64(42).ToGetRequestInformation(context.Background(), nil)
	if err != nil {
		t.Fatalf("failed to construct get request info: %v", err)
	}
	if reqInfo == nil {
		t.Fatal("expected non-nil RequestInformation")
	}

	uri, err := reqInfo.GetUri()
	if err != nil {
		t.Fatalf("failed to get URI: %v", err)
	}
	expectedUri := "https://petstore.swagger.io/v2/pet/42"
	if uri.String() != expectedUri {
		t.Errorf("expected URI %q, got %q", expectedUri, uri.String())
	}
}

func TestPetstoreApiClient_FindByStatusToGetRequestInformation(t *testing.T) {
	authProvider := &absauth.AnonymousAuthenticationProvider{}
	adapter, err := r.NewNetHttpRequestAdapter(authProvider)
	if err != nil {
		t.Fatalf("failed to create request adapter: %v", err)
	}

	client := u.NewApiClient(adapter)
	config := &pet.FindByStatusRequestBuilderGetRequestConfiguration{
		QueryParameters: &pet.FindByStatusRequestBuilderGetQueryParameters{
			Status: []string{"available"},
		},
	}
	reqInfo, err := client.Pet().FindByStatus().ToGetRequestInformation(context.Background(), config)
	if err != nil {
		t.Fatalf("failed to construct findByStatus request info: %v", err)
	}
	if reqInfo == nil {
		t.Fatal("expected non-nil RequestInformation")
	}

	uri, err := reqInfo.GetUri()
	if err != nil {
		t.Fatalf("failed to get URI: %v", err)
	}
	expectedUri := "https://petstore.swagger.io/v2/pet/findByStatus?status=available"
	if uri.String() != expectedUri {
		t.Errorf("expected URI %q, got %q", expectedUri, uri.String())
	}
}

func TestPetstoreApiClient_StoreInventoryToGetRequestInformation(t *testing.T) {
	authProvider := &absauth.AnonymousAuthenticationProvider{}
	adapter, err := r.NewNetHttpRequestAdapter(authProvider)
	if err != nil {
		t.Fatalf("failed to create request adapter: %v", err)
	}

	client := u.NewApiClient(adapter)
	reqInfo, err := client.Store().Inventory().ToGetRequestInformation(context.Background(), nil)
	if err != nil {
		t.Fatalf("failed to construct inventory request info: %v", err)
	}
	if reqInfo == nil {
		t.Fatal("expected non-nil RequestInformation")
	}

	uri, err := reqInfo.GetUri()
	if err != nil {
		t.Fatalf("failed to get URI: %v", err)
	}
	expectedUri := "https://petstore.swagger.io/v2/store/inventory"
	if uri.String() != expectedUri {
		t.Errorf("expected URI %q, got %q", expectedUri, uri.String())
	}
}

func TestPetstoreModel_PetProperties(t *testing.T) {
	pet := models.NewPet()
	if pet == nil {
		t.Fatal("expected non-nil Pet")
	}

	id := int64(100)
	name := "Fido"
	pet.SetId(&id)
	pet.SetName(&name)

	if pet.GetId() == nil || *pet.GetId() != 100 {
		t.Errorf("expected id 100, got %v", pet.GetId())
	}
	if pet.GetName() == nil || *pet.GetName() != "Fido" {
		t.Errorf("expected name Fido, got %v", pet.GetName())
	}

	fields := pet.GetFieldDeserializers()
	if len(fields) == 0 {
		t.Error("expected non-empty field deserializers for Pet")
	}
}

func TestPetstoreModel_OrderProperties(t *testing.T) {
	order := models.NewOrder()
	if order == nil {
		t.Fatal("expected non-nil Order")
	}

	id := int64(5)
	qty := int32(2)
	complete := true
	order.SetId(&id)
	order.SetQuantity(&qty)
	order.SetComplete(&complete)

	if order.GetId() == nil || *order.GetId() != 5 {
		t.Errorf("expected order id 5, got %v", order.GetId())
	}
	if order.GetQuantity() == nil || *order.GetQuantity() != 2 {
		t.Errorf("expected quantity 2, got %v", order.GetQuantity())
	}
	if order.GetComplete() == nil || !*order.GetComplete() {
		t.Errorf("expected complete true, got %v", order.GetComplete())
	}
}

func TestStrToPtr(t *testing.T) {
	val := "hello"
	ptr := StrToPtr(val)
	if ptr == nil || *ptr != "hello" {
		t.Errorf("expected pointer to %q, got %v", val, ptr)
	}
}
