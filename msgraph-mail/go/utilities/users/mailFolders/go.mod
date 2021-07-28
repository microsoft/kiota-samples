module mailfolders

replace github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailFolders/childFolders => ./childFolders
replace github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailFolders/item => ./item
replace github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailFolders/messageRules => ./messageRules
replace github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailFolders/messages => ./messages
replace github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailFolders/multiValueExtendedProperties => ./multiValueExtendedProperties
replace github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailFolders/singleValueExtendedProperties => ./singleValueExtendedProperties
replace github.com/microsoft/kiota/abstractions/go => ../../../../../../abstractions/go
replace github.com/microsoft/kiota/abstractions/go/serialization => ../../../../../../abstractions/go/serialization

go 1.16

require (
    github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailFolders/childFolders v0.0.0-00010101000000-000000000000
    github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailFolders/item v0.0.0-00010101000000-000000000000
    github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailFolders/messageRules v0.0.0-00010101000000-000000000000
    github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailFolders/messages v0.0.0-00010101000000-000000000000
    github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailFolders/multiValueExtendedProperties v0.0.0-00010101000000-000000000000
    github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailFolders/singleValueExtendedProperties v0.0.0-00010101000000-000000000000
    github.com/microsoft/kiota/abstractions/go v0.0.0-00010101000000-000000000000
    github.com/microsoft/kiota/abstractions/go/serialization v0.0.0-00010101000000-000000000000
)
