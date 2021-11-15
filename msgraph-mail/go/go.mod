module github.com/microsoft/kiota-samples/msgraph-mail/go

replace github.com/microsoft/kiota/abstractions/go => ../../../abstractions/go

replace github.com/microsoft/kiota/serialization/go/json => ../../../serialization/go/json

replace github.com/microsoft/kiota/authentication/go/azure => ../../../authentication/go/azure

replace github.com/microsoft/kiota/http/go/nethttp => ../../../http/go/nethttp

go 1.16

require github.com/microsoft/kiota/abstractions/go v0.0.0-20211112084539-17ac73ffdc7c

require (
	github.com/Azure/azure-sdk-for-go/sdk/azidentity v0.12.0
	github.com/microsoft/kiota/authentication/go/azure v0.0.0-00010101000000-000000000000
	github.com/microsoft/kiota/http/go/nethttp v0.0.0-00010101000000-000000000000
	github.com/microsoft/kiota/serialization/go/json v0.0.0-00010101000000-000000000000
	github.com/pkg/browser v0.0.0-20210911075715-681adbf594b8 // indirect
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	golang.org/x/net v0.0.0-20211104170005-ce137452f963 // indirect
	golang.org/x/sys v0.0.0-20211103235746-7861aae1554b // indirect
)
