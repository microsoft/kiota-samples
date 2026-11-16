package main

import (
	"context"
	"fmt"
	"log"

	azidentity "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	a "github.com/microsoft/kiota-authentication-azure-go"
	r "github.com/microsoft/kiota-http-go"
	u "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities"
)

func main() {
	cred, err := azidentity.NewDeviceCodeCredential(&azidentity.DeviceCodeCredentialOptions{
		TenantID: "09988f3c-947e-4986-a87e-37ad49a3b175",
		ClientID: "f19e2a30-d500-4fa7-8582-bd6099088b37",
		UserPrompt: func(ctx context.Context, message azidentity.DeviceCodeMessage) error {
			fmt.Println(message.Message)
			return nil
		},
	})
	if err != nil {
		log.Fatalf("Error creating credentials: %v\n", err)
	}
	auth, err := a.NewAzureIdentityAuthenticationProviderWithScopes(cred, []string{"Mail.Read", "Mail.Send"})
	if err != nil {
		log.Fatalf("Error authentication provider: %v\n", err)
	}
	adapter, err := r.NewNetHttpRequestAdapter(auth)
	if err != nil {
		log.Fatalf("Error creating adapter: %v\n", err)
	}
	client := u.NewApiClient(adapter)
	urb := client.Users().ByUserId("vincent@biret365.onmicrosoft.com")
	fmt.Printf("urb %v\n", urb)
	response, err := client.Users().ByUserId("vincent@biret365.onmicrosoft.com").Messages().Get(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error getting messages: %v\n", err)
	}
	if response != nil {
		for _, message := range response.GetValue() {
			if message != nil && message.GetSubject() != nil {
				fmt.Printf("Message: %v\n", *message.GetSubject())
			}
		}
	}
}

func StrToPtr(s string) *string {
	return &s
}
