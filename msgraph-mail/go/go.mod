module github.com/microsoft/kiota-samples/msgraph-mail/go

replace github.com/microsoft/kiota/abstractions/go => ../../../abstractions/go

replace github.com/microsoft/kiota/serialization/go/json => ../../../serialization/go/json

replace github.com/microsoft/kiota/authentication/go/azure => ../../../authentication/go/azure

replace github.com/microsoft/kiota/http/go/nethttp => ../../../http/go/nethttp

go 1.16

require github.com/microsoft/kiota/abstractions/go v0.0.0-20211013091133-b793efa27646

require (
	github.com/Azure/azure-sdk-for-go/sdk/azidentity v0.11.0 // indirect
	github.com/microsoft/kiota/authentication/go/azure v0.0.0-00010101000000-000000000000 // indirect
	github.com/microsoft/kiota/http/go/nethttp v0.0.0-00010101000000-000000000000 // indirect
	github.com/microsoft/kiota/serialization/go/json v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20211014222326-fd004c51d1d6 // indirect
)
