package main

import (
	"context"
	"testing"

	absauth "github.com/microsoft/kiota-abstractions-go/authentication"
	r "github.com/microsoft/kiota-http-go"
	u "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities"
	"github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/models"
	"github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users"
)

func TestApiClient_Initialization(t *testing.T) {
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
	expectedBaseUrl := "https://graph.microsoft.com/v1.0"
	if baseUrl != expectedBaseUrl {
		t.Errorf("expected baseUrl %q, got %q", expectedBaseUrl, baseUrl)
	}
}

func TestApiClient_UserMessagesToGetRequestInformation(t *testing.T) {
	authProvider := &absauth.AnonymousAuthenticationProvider{}
	adapter, err := r.NewNetHttpRequestAdapter(authProvider)
	if err != nil {
		t.Fatalf("failed to create request adapter: %v", err)
	}

	client := u.NewApiClient(adapter)
	userEmail := "testuser@contoso.com"
	reqInfo, err := client.Users().ByUserId(userEmail).Messages().ToGetRequestInformation(context.Background(), nil)
	if err != nil {
		t.Fatalf("failed to construct get messages request info: %v", err)
	}
	if reqInfo == nil {
		t.Fatal("expected non-nil RequestInformation")
	}

	uri, err := reqInfo.GetUri()
	if err != nil {
		t.Fatalf("failed to get URI: %v", err)
	}
	expectedUri := "https://graph.microsoft.com/v1.0/users/testuser%40contoso.com/messages"
	if uri.String() != expectedUri {
		t.Errorf("expected URI %q, got %q", expectedUri, uri.String())
	}
}

func TestApiClient_UserMessagesWithQueryParameters(t *testing.T) {
	authProvider := &absauth.AnonymousAuthenticationProvider{}
	adapter, err := r.NewNetHttpRequestAdapter(authProvider)
	if err != nil {
		t.Fatalf("failed to create request adapter: %v", err)
	}

	client := u.NewApiClient(adapter)
	top := int32(10)
	config := &users.ItemMessagesRequestBuilderGetRequestConfiguration{
		QueryParameters: &users.ItemMessagesRequestBuilderGetQueryParameters{
			Top:    &top,
			Select: []string{"subject", "sender", "receivedDateTime"},
		},
	}
	reqInfo, err := client.Users().ByUserId("testuser@contoso.com").Messages().ToGetRequestInformation(context.Background(), config)
	if err != nil {
		t.Fatalf("failed to construct get messages request info with query params: %v", err)
	}

	uri, err := reqInfo.GetUri()
	if err != nil {
		t.Fatalf("failed to get URI: %v", err)
	}
	uriStr := uri.String()
	if uriStr == "" {
		t.Fatal("expected non-empty URI string")
	}
}

func TestApiClient_UserMailFoldersToGetRequestInformation(t *testing.T) {
	authProvider := &absauth.AnonymousAuthenticationProvider{}
	adapter, err := r.NewNetHttpRequestAdapter(authProvider)
	if err != nil {
		t.Fatalf("failed to create request adapter: %v", err)
	}

	client := u.NewApiClient(adapter)
	reqInfo, err := client.Users().ByUserId("testuser@contoso.com").MailFolders().ToGetRequestInformation(context.Background(), nil)
	if err != nil {
		t.Fatalf("failed to construct get mail folders request info: %v", err)
	}

	uri, err := reqInfo.GetUri()
	if err != nil {
		t.Fatalf("failed to get URI: %v", err)
	}
	expectedUri := "https://graph.microsoft.com/v1.0/users/testuser%40contoso.com/mailFolders"
	if uri.String() != expectedUri {
		t.Errorf("expected URI %q, got %q", expectedUri, uri.String())
	}
}

func TestMailModel_MessageProperties(t *testing.T) {
	msg := models.NewMessage()
	if msg == nil {
		t.Fatal("expected non-nil Message")
	}

	subject := "Quarterly Project Status Update"
	isRead := true
	id := "msg-12345"
	msg.SetId(&id)
	msg.SetSubject(&subject)
	msg.SetIsRead(&isRead)

	if msg.GetId() == nil || *msg.GetId() != id {
		t.Errorf("expected id %q, got %v", id, msg.GetId())
	}
	if msg.GetSubject() == nil || *msg.GetSubject() != subject {
		t.Errorf("expected subject %q, got %v", subject, msg.GetSubject())
	}
	if msg.GetIsRead() == nil || !*msg.GetIsRead() {
		t.Errorf("expected isRead true, got %v", msg.GetIsRead())
	}

	fields := msg.GetFieldDeserializers()
	if len(fields) == 0 {
		t.Error("expected non-empty field deserializers for Message")
	}
}

func TestMailModel_MailFolderProperties(t *testing.T) {
	folder := models.NewMailFolder()
	if folder == nil {
		t.Fatal("expected non-nil MailFolder")
	}

	name := "Archive"
	totalCount := int32(42)
	unreadCount := int32(0)
	folder.SetDisplayName(&name)
	folder.SetTotalItemCount(&totalCount)
	folder.SetUnreadItemCount(&unreadCount)

	if folder.GetDisplayName() == nil || *folder.GetDisplayName() != name {
		t.Errorf("expected displayName %q, got %v", name, folder.GetDisplayName())
	}
	if folder.GetTotalItemCount() == nil || *folder.GetTotalItemCount() != totalCount {
		t.Errorf("expected totalItemCount %d, got %v", totalCount, folder.GetTotalItemCount())
	}
	if folder.GetUnreadItemCount() == nil || *folder.GetUnreadItemCount() != unreadCount {
		t.Errorf("expected unreadItemCount %d, got %v", unreadCount, folder.GetUnreadItemCount())
	}
}

func TestStrToPtr(t *testing.T) {
	s := "test-string"
	ptr := StrToPtr(s)
	if ptr == nil || *ptr != s {
		t.Errorf("expected pointer to %q, got %v", s, ptr)
	}
}
