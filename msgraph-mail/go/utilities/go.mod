module utilities

replace github.com/microsoft/kiota/abstractions/go => ../../../../abstractions/go

replace github.com/microsoft/kiota/abstractions/go/serialization => ../../../../abstractions/go/serialization

replace github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users => ./users
replace github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item => ./users/item

go 1.16

require (
    github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users v0.0.0-00010101000000-000000000000
	github.com/microsoft/kiota/abstractions/go v0.0.0-00010101000000-000000000000
	github.com/microsoft/kiota/abstractions/go/serialization v0.0.0-00010101000000-000000000000
)
