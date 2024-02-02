package main

import (
	"context"
	"fmt"

	azidentity "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	a "github.com/microsoft/kiota-authentication-azure-go"
	r "github.com/microsoft/kiota-http-go"
	u "github.com/microsoft/kiota-samples/petstore/go/utilities"
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
		fmt.Printf("Error creating credentials: %v\n", err)
	}
	auth, err := a.NewAzureIdentityAuthenticationProviderWithScopes(cred, []string{"Mail.Read", "Mail.Send"})
	if err != nil {
		fmt.Printf("Error authentication provider: %v\n", err)
		return
	}
	adapter, err := r.NewNetHttpRequestAdapter(auth)
	if err != nil {
		fmt.Printf("Error creating adapter: %v\n", err)
		return
	}
	client := u.NewApiClient(adapter)
	response, err := client.Pet().ByPetIdInt64(10).Get(context.Background(), nil)
	if err != nil {
		fmt.Printf("Error getting pet: %v\n", err)
		return
	}
	fmt.Printf("Pet : %v\n", *(response.GetName()))
}

func StrToPtr(s string) *string {
	return &s
}
