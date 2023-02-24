// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"todoclient/client"
	"todoclient/client/models/todoapi/models"
	"todoclient/client/todoitems"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/joho/godotenv"
	azure "github.com/microsoft/kiota-authentication-azure-go"
	http "github.com/microsoft/kiota-http-go"
)

func main() {
	// Load .env files
	// .env.local takes precedence (if present)
	godotenv.Load(".env.local")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	clientId := os.Getenv("CLIENT_ID")
	tenantId := os.Getenv("TENANT_ID")
	scopeString := os.Getenv("API_SCOPE")
	scopes := []string{scopeString}

	// The auth provider will only authorize requests to
	// the allowed hosts, in this case Microsoft Graph
	allowedHosts := []string{"localhost:7206"}

	credential, err := azidentity.NewDeviceCodeCredential(&azidentity.DeviceCodeCredentialOptions{
		ClientID: clientId,
		TenantID: tenantId,
		UserPrompt: func(ctx context.Context, dcm azidentity.DeviceCodeMessage) error {
			fmt.Println(dcm.Message)
			return nil
		},
	})
	if err != nil {
		log.Fatalf("Error creating credential: %v\n", err)
	}

	tokenOptions := policy.TokenRequestOptions{Scopes: scopes}
	token, err := credential.GetToken(context.Background(), tokenOptions)
	if err != nil {
		log.Fatalf("Error getting access token: %v\n", err)
	}

	fmt.Printf("Token: %s\n", token.Token)

	authProvider, err := azure.NewAzureIdentityAuthenticationProviderWithScopesAndValidHosts(
		credential, scopes, allowedHosts)
	if err != nil {
		log.Fatalf("Error creating auth provider: %v\n", err)
	}

	adapter, err := http.NewNetHttpRequestAdapter(authProvider)
	if err != nil {
		log.Fatalf("Error creating request adapter: %v\n", err)
	}

	client := client.NewToDoApiClient(adapter)

	// GET /ToDoItems
	items, err := client.ToDoItems().Get(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error getting ToDo items: %v\n", err)
	}

	fmt.Printf("Found %d items\n", len(items.GetValue()))
	for _, todo := range items.GetValue() {
		fmt.Printf("TODO: %s, Due: %v\n", *todo.GetName(), *todo.GetDueDate())
	}

	// Select
	query := todoitems.ToDoItemsRequestBuilderGetQueryParameters{
		Select: []string{"name", "priority", "isComplete"},
	}

	selectItems, err := client.ToDoItems().Get(context.Background(),
		&todoitems.ToDoItemsRequestBuilderGetRequestConfiguration{
			QueryParameters: &query,
		})
	if err != nil {
		log.Fatalf("Error getting ToDo items with select: %v\n", err)
	}

	for _, todo := range selectItems.GetValue() {
		fmt.Printf("%s - Priority: %d - Complete: %v\n",
			*todo.GetName(), *todo.GetPriority(), *todo.GetIsComplete())
	}

	// Get by ID
	specificId := items.GetValue()[3].GetId()
	specificItem, err := client.ToDoItemsById(specificId.String()).
		Get(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error getting item by ID: %v\n", err)
	}

	fmt.Printf("%s\n", *specificItem.GetName())

	// Create an item
	newItem := models.NewToDoItem()
	name := "Start using Kiota generators in my project"
	newItem.SetName(&name)
	itemType := "work"
	newItem.SetType(&itemType)
	isComplete := false
	newItem.SetIsComplete(&isComplete)
	// One day from now
	dueDate := time.Now().Add(time.Hour * 24)
	newItem.SetDueDate(&dueDate)
	priority := int32(1)
	newItem.SetPriority(&priority)

	createdItem, err := client.ToDoItems().Post(context.Background(), newItem, nil)
	if err != nil {
		log.Fatalf("Error creating new item: %v\n", err)
	}

	fmt.Printf("Item created with ID: %v\n", createdItem.GetId())

	// Update an item
	updateItem := models.NewToDoItem()
	isComplete = true
	updateItem.SetIsComplete(&isComplete)

	err = client.ToDoItemsById(createdItem.GetId().String()).
		Patch(context.Background(), updateItem, nil)
	if err != nil {
		log.Fatalf("Error updating item: %v\n", err)
	}

	fmt.Printf("Updated item\n")

	// Delete an item
	err = client.ToDoItemsById(createdItem.GetId().String()).
		Delete(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error deleting item: %v\n", err)
	}

	fmt.Printf("Deleted item\n")
}
