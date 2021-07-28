module users

replace github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/inferenceClassification => ./inferenceClassification
replace github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item => ./item
replace github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailFolders => ./mailFolders
replace github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/messages => ./messages
replace github.com/microsoft/kiota/abstractions/go => ../../../../../abstractions/go
replace github.com/microsoft/kiota/abstractions/go/serialization => ../../../../../abstractions/go/serialization

go 1.16

require (
    github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/inferenceClassification v0.0.0-00010101000000-000000000000
    github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item v0.0.0-00010101000000-000000000000
    github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailFolders v0.0.0-00010101000000-000000000000
    github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/messages v0.0.0-00010101000000-000000000000
    github.com/microsoft/kiota/abstractions/go v0.0.0-00010101000000-000000000000
    github.com/microsoft/kiota/abstractions/go/serialization v0.0.0-00010101000000-000000000000
)
