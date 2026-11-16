package main

import (
	"context"
	"fmt"
	"log"

	absauth "github.com/microsoft/kiota-abstractions-go/authentication"
	r "github.com/microsoft/kiota-http-go"
	u "github.com/microsoft/kiota-samples/petstore/go/utilities"
)

func main() {
	// The Swagger Petstore OpenAPI endpoints do not require authentication
	authProvider := &absauth.AnonymousAuthenticationProvider{}

	adapter, err := r.NewNetHttpRequestAdapter(authProvider)
	if err != nil {
		log.Fatalf("Error creating adapter: %v\n", err)
	}

	client := u.NewApiClient(adapter)
	response, err := client.Pet().ByPetIdInt64(10).Get(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error getting pet: %v\n", err)
	}

	if response != nil && response.GetName() != nil {
		fmt.Printf("Pet : %v\n", *response.GetName())
	}
}

func StrToPtr(s string) *string {
	return &s
}
