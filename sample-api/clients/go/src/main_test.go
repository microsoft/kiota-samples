// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package main

import (
	"context"
	"testing"
	"time"

	"todoclient/client"
	"todoclient/client/models"
	"todoclient/client/todoitems"

	"github.com/google/uuid"
	absauth "github.com/microsoft/kiota-abstractions-go/authentication"
	http "github.com/microsoft/kiota-http-go"
)

func TestToDoApiClient_Initialization(t *testing.T) {
	authProvider := &absauth.AnonymousAuthenticationProvider{}
	adapter, err := http.NewNetHttpRequestAdapter(authProvider)
	if err != nil {
		t.Fatalf("failed to create request adapter: %v", err)
	}

	apiClient := client.NewToDoApiClient(adapter)
	if apiClient == nil {
		t.Fatal("expected non-nil ToDoApiClient")
	}

	baseUrl := adapter.GetBaseUrl()
	expectedBaseUrl := "https://localhost:7206"
	if baseUrl != expectedBaseUrl {
		t.Errorf("expected baseUrl %q, got %q", expectedBaseUrl, baseUrl)
	}
}

func TestToDoApiClient_ToDoItemsToGetRequestInformation(t *testing.T) {
	authProvider := &absauth.AnonymousAuthenticationProvider{}
	adapter, err := http.NewNetHttpRequestAdapter(authProvider)
	if err != nil {
		t.Fatalf("failed to create request adapter: %v", err)
	}

	apiClient := client.NewToDoApiClient(adapter)
	reqInfo, err := apiClient.ToDoItems().ToGetRequestInformation(context.Background(), nil)
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
	expectedUri := "https://localhost:7206/ToDoItems"
	if uri.String() != expectedUri {
		t.Errorf("expected URI %q, got %q", expectedUri, uri.String())
	}
}

func TestToDoApiClient_ToDoItemsWithQueryParameters(t *testing.T) {
	authProvider := &absauth.AnonymousAuthenticationProvider{}
	adapter, err := http.NewNetHttpRequestAdapter(authProvider)
	if err != nil {
		t.Fatalf("failed to create request adapter: %v", err)
	}

	apiClient := client.NewToDoApiClient(adapter)
	incomplete := false
	config := &todoitems.ToDoItemsRequestBuilderGetRequestConfiguration{
		QueryParameters: &todoitems.ToDoItemsRequestBuilderGetQueryParameters{
			IsComplete: &incomplete,
		},
	}

	reqInfo, err := apiClient.ToDoItems().ToGetRequestInformation(context.Background(), config)
	if err != nil {
		t.Fatalf("failed to construct get request info with filter: %v", err)
	}

	uri, err := reqInfo.GetUri()
	if err != nil {
		t.Fatalf("failed to get URI: %v", err)
	}
	expectedUri := "https://localhost:7206/ToDoItems?isComplete=false"
	if uri.String() != expectedUri {
		t.Errorf("expected URI %q, got %q", expectedUri, uri.String())
	}
}

func TestToDoApiClient_ToDoItemsToPostRequestInformation(t *testing.T) {
	authProvider := &absauth.AnonymousAuthenticationProvider{}
	adapter, err := http.NewNetHttpRequestAdapter(authProvider)
	if err != nil {
		t.Fatalf("failed to create request adapter: %v", err)
	}

	apiClient := client.NewToDoApiClient(adapter)
	newItem := models.NewToDoItem()
	name := "Unit Test Task"
	newItem.SetName(&name)

	reqInfo, err := apiClient.ToDoItems().ToPostRequestInformation(context.Background(), newItem, nil)
	if err != nil {
		t.Fatalf("failed to construct post request info: %v", err)
	}

	uri, err := reqInfo.GetUri()
	if err != nil {
		t.Fatalf("failed to get URI: %v", err)
	}
	expectedUri := "https://localhost:7206/ToDoItems"
	if uri.String() != expectedUri {
		t.Errorf("expected URI %q, got %q", expectedUri, uri.String())
	}
}

func TestToDoApiClient_ToDoItemsByIdGuid(t *testing.T) {
	authProvider := &absauth.AnonymousAuthenticationProvider{}
	adapter, err := http.NewNetHttpRequestAdapter(authProvider)
	if err != nil {
		t.Fatalf("failed to create request adapter: %v", err)
	}

	apiClient := client.NewToDoApiClient(adapter)
	testId := uuid.New()

	// GET by ID
	getReqInfo, err := apiClient.ToDoItems().ByIdGuid(testId).ToGetRequestInformation(context.Background(), nil)
	if err != nil {
		t.Fatalf("failed to construct get by ID request info: %v", err)
	}
	getUri, err := getReqInfo.GetUri()
	if err != nil {
		t.Fatalf("failed to get URI: %v", err)
	}
	expectedUri := "https://localhost:7206/ToDoItems/" + testId.String()
	if getUri.String() != expectedUri {
		t.Errorf("expected URI %q, got %q", expectedUri, getUri.String())
	}

	// PATCH by ID
	updateItem := models.NewToDoItemUpdate()
	patchReqInfo, err := apiClient.ToDoItems().ByIdGuid(testId).ToPatchRequestInformation(context.Background(), updateItem, nil)
	if err != nil {
		t.Fatalf("failed to construct patch by ID request info: %v", err)
	}
	patchUri, err := patchReqInfo.GetUri()
	if err != nil {
		t.Fatalf("failed to get PATCH URI: %v", err)
	}
	if patchUri.String() != expectedUri {
		t.Errorf("expected PATCH URI %q, got %q", expectedUri, patchUri.String())
	}

	// DELETE by ID
	delReqInfo, err := apiClient.ToDoItems().ByIdGuid(testId).ToDeleteRequestInformation(context.Background(), nil)
	if err != nil {
		t.Fatalf("failed to construct delete by ID request info: %v", err)
	}
	delUri, err := delReqInfo.GetUri()
	if err != nil {
		t.Fatalf("failed to get DELETE URI: %v", err)
	}
	if delUri.String() != expectedUri {
		t.Errorf("expected DELETE URI %q, got %q", expectedUri, delUri.String())
	}
}

func TestToDoModel_ToDoItemProperties(t *testing.T) {
	item := models.NewToDoItem()
	if item == nil {
		t.Fatal("expected non-nil ToDoItem")
	}

	testId := uuid.New()
	name := "Sample Task"
	itemType := "work"
	isComplete := false
	dueDate := time.Now()
	priority := int32(1)

	item.SetId(&testId)
	item.SetName(&name)
	item.SetTypeEscaped(&itemType)
	item.SetIsComplete(&isComplete)
	item.SetDueDate(&dueDate)
	item.SetPriority(&priority)

	if item.GetId() == nil || *item.GetId() != testId {
		t.Errorf("expected ID %v, got %v", testId, item.GetId())
	}
	if item.GetName() == nil || *item.GetName() != name {
		t.Errorf("expected Name %q, got %v", name, item.GetName())
	}
	if item.GetTypeEscaped() == nil || *item.GetTypeEscaped() != itemType {
		t.Errorf("expected Type %q, got %v", itemType, item.GetTypeEscaped())
	}
	if item.GetIsComplete() == nil || *item.GetIsComplete() != isComplete {
		t.Errorf("expected IsComplete %v, got %v", isComplete, item.GetIsComplete())
	}
	if item.GetPriority() == nil || *item.GetPriority() != priority {
		t.Errorf("expected Priority %d, got %v", priority, item.GetPriority())
	}

	fields := item.GetFieldDeserializers()
	if len(fields) == 0 {
		t.Error("expected non-empty field deserializers for ToDoItem")
	}
}

func TestToDoModel_ToDoItemUpdateProperties(t *testing.T) {
	update := models.NewToDoItemUpdate()
	if update == nil {
		t.Fatal("expected non-nil ToDoItemUpdate")
	}

	isComplete := true
	update.SetIsComplete(&isComplete)
	if update.GetIsComplete() == nil || !*update.GetIsComplete() {
		t.Errorf("expected IsComplete true, got %v", update.GetIsComplete())
	}

	fields := update.GetFieldDeserializers()
	if len(fields) == 0 {
		t.Error("expected non-empty field deserializers for ToDoItemUpdate")
	}
}

func TestToDoModel_ProblemDetailsProperties(t *testing.T) {
	problem := models.NewProblemDetails()
	if problem == nil {
		t.Fatal("expected non-nil ProblemDetails")
	}

	detail := "Item not found"
	status := int32(404)
	title := "Not Found"

	problem.SetDetail(&detail)
	problem.SetStatus(&status)
	problem.SetTitle(&title)

	if problem.GetDetail() == nil || *problem.GetDetail() != detail {
		t.Errorf("expected Detail %q, got %v", detail, problem.GetDetail())
	}
	if problem.GetStatus() == nil || *problem.GetStatus() != status {
		t.Errorf("expected Status %d, got %v", status, problem.GetStatus())
	}
	if problem.GetTitle() == nil || *problem.GetTitle() != title {
		t.Errorf("expected Title %q, got %v", title, problem.GetTitle())
	}
}
