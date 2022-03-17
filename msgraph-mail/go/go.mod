module github.com/microsoft/kiota-samples/msgraph-mail/go

replace github.com/microsoft/kiota/abstractions/go => ../../../abstractions/go

replace github.com/microsoft/kiota/serialization/go/json => ../../../serialization/go/json

replace github.com/microsoft/kiota/serialization/go/text => ../../../serialization/go/text

replace github.com/microsoft/kiota/authentication/go/azure => ../../../authentication/go/azure

replace github.com/microsoft/kiota/http/go/nethttp => ../../../http/go/nethttp

go 1.16

require github.com/microsoft/kiota/abstractions/go v0.0.0-20220316090511-468a319ac5ba

require (
	github.com/Azure/azure-sdk-for-go/sdk/azidentity v0.13.0
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/microsoft/kiota/authentication/go/azure v0.0.0-20220203130215-5883af1759a7
	github.com/microsoft/kiota/http/go/nethttp v0.0.0-20220203130215-5883af1759a7
	github.com/microsoft/kiota/serialization/go/json v0.0.0-20220203130215-5883af1759a7
	github.com/microsoft/kiota/serialization/go/text v0.0.0-20220203130215-5883af1759a7
	github.com/pkg/browser v0.0.0-20210911075715-681adbf594b8 // indirect
	golang.org/x/crypto v0.0.0-20220131195533-30dcbda58838 // indirect
	golang.org/x/sys v0.0.0-20220128215802-99c3d69c2c27 // indirect
)
