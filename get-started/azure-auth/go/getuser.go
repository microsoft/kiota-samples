// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// <ProgramSnippet>
package main

import (
	"context"
	"fmt"

	"getuser/client"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	azure "github.com/microsoft/kiota-authentication-azure-go"
	bundle "github.com/microsoft/kiota-bundle-go"
)

func main() {
	clientId := "YOUR_CLIENT_ID"

	// The auth provider will only authorize requests to
	// the allowed hosts, in this case Microsoft Graph
	allowedHosts := []string{"graph.microsoft.com"}
	graphScopes := []string{"User.Read"}

	credential, err := azidentity.NewDeviceCodeCredential(&azidentity.DeviceCodeCredentialOptions{
		ClientID: clientId,
		UserPrompt: func(ctx context.Context, dcm azidentity.DeviceCodeMessage) error {
			fmt.Println(dcm.Message)
			return nil
		},
	})

	if err != nil {
		fmt.Printf("Error creating credential: %v\n", err)
	}

	authProvider, err := azure.NewAzureIdentityAuthenticationProviderWithScopesAndValidHosts(
		credential, graphScopes, allowedHosts)

	if err != nil {
		fmt.Printf("Error creating auth provider: %v\n", err)
	}

	adapter, err := bundle.NewDefaultRequestAdapter(authProvider)

	if err != nil {
		fmt.Printf("Error creating request adapter: %v\n", err)
	}

	client := client.NewGraphApiClient(adapter)

	me, err := client.Me().Get(context.Background(), nil)

	if err != nil {
		fmt.Printf("Error getting user: %v\n", err)
	}

	fmt.Printf("Hello %s, your ID is %s\n", *me.GetDisplayName(), *me.GetId())
}
// </ProgramSnippet>
