package main

import (
	"context"
	"strings"
	"testing"

	absauth "github.com/microsoft/kiota-abstractions-go/authentication"
	r "github.com/microsoft/kiota-http-go"
	serjson "github.com/microsoft/kiota-serialization-json-go"
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

func TestPetAndOrder_SerializationAndDiscriminators(t *testing.T) {
	petParsable, err := models.CreatePetFromDiscriminatorValue(nil)
	if err != nil {
		t.Fatalf("unexpected error from pet discriminator: %v", err)
	}
	if petParsable == nil {
		t.Fatal("expected non-nil parsable from pet discriminator")
	}

	orderParsable, err := models.CreateOrderFromDiscriminatorValue(nil)
	if err != nil {
		t.Fatalf("unexpected error from order discriminator: %v", err)
	}
	if orderParsable == nil {
		t.Fatal("expected non-nil parsable from order discriminator")
	}

	petObj := models.NewPet()
	id := int64(77)
	name := "Buddy"
	petObj.SetId(&id)
	petObj.SetName(&name)

	writer := serjson.NewJsonSerializationWriter()
	err = petObj.Serialize(writer)
	if err != nil {
		t.Fatalf("failed to serialize pet: %v", err)
	}
	content, err := writer.GetSerializedContent()
	if err != nil {
		t.Fatalf("failed to get serialized pet content: %v", err)
	}
	strContent := string(content)
	if !strings.Contains(strContent, "Buddy") || !strings.Contains(strContent, "77") {
		t.Errorf("serialized pet missing expected fields: %s", strContent)
	}

	orderObj := models.NewOrder()
	orderId := int64(88)
	qty := int32(3)
	complete := true
	orderObj.SetId(&orderId)
	orderObj.SetQuantity(&qty)
	orderObj.SetComplete(&complete)

	orderWriter := serjson.NewJsonSerializationWriter()
	err = orderObj.Serialize(orderWriter)
	if err != nil {
		t.Fatalf("failed to serialize order: %v", err)
	}
	orderContent, err := orderWriter.GetSerializedContent()
	if err != nil {
		t.Fatalf("failed to get serialized order content: %v", err)
	}
	orderStr := string(orderContent)
	if !strings.Contains(orderStr, "88") || !strings.Contains(orderStr, "3") || !strings.Contains(orderStr, "true") {
		t.Errorf("serialized order missing expected fields: %s", orderStr)
	}
}

func TestPetstoreApiClient_WithUrl(t *testing.T) {
	authProvider := &absauth.AnonymousAuthenticationProvider{}
	adapter, err := r.NewNetHttpRequestAdapter(authProvider)
	if err != nil {
		t.Fatalf("failed to create request adapter: %v", err)
	}
	client := u.NewApiClient(adapter)

	customPetUrl := "https://petstore.swagger.io/v2/custom-pet"
	customPetBuilder := client.Pet().WithUrl(customPetUrl)
	reqInfo, err := customPetBuilder.ToPostRequestInformation(context.Background(), models.NewPet(), nil)
	if err != nil {
		t.Fatalf("failed to get request information for WithUrl pet: %v", err)
	}
	uri, err := reqInfo.GetUri()
	if err != nil {
		t.Fatalf("failed to get URI: %v", err)
	}
	if uri.String() != customPetUrl {
		t.Errorf("expected URI %q, got %q", customPetUrl, uri.String())
	}

	customItemUrl := "https://petstore.swagger.io/v2/custom-pet/99"
	customItemBuilder := client.Pet().ByPetIdInt64(99).WithUrl(customItemUrl)
	itemReqInfo, err := customItemBuilder.ToGetRequestInformation(context.Background(), nil)
	if err != nil {
		t.Fatalf("failed to get request information for WithUrl pet item: %v", err)
	}
	itemUri, err := itemReqInfo.GetUri()
	if err != nil {
		t.Fatalf("failed to get URI: %v", err)
	}
	if itemUri.String() != customItemUrl {
		t.Errorf("expected URI %q, got %q", customItemUrl, itemUri.String())
	}
}
