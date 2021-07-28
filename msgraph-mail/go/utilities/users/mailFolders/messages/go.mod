module messages

replace github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailFolders/messages/attachments => ./attachments
replace github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailFolders/messages/Content => ./Content
replace github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailFolders/messages/extensions => ./extensions
replace github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailFolders/messages/item => ./item
replace github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailFolders/messages/multiValueExtendedProperties => ./multiValueExtendedProperties
replace github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailFolders/messages/singleValueExtendedProperties => ./singleValueExtendedProperties
replace github.com/microsoft/kiota/abstractions/go => ../../../../../../../abstractions/go
replace github.com/microsoft/kiota/abstractions/go/serialization => ../../../../../../../abstractions/go/serialization

go 1.16

require (
    github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailFolders/messages/attachments v0.0.0-00010101000000-000000000000
    github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailFolders/messages/Content v0.0.0-00010101000000-000000000000
    github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailFolders/messages/extensions v0.0.0-00010101000000-000000000000
    github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailFolders/messages/item v0.0.0-00010101000000-000000000000
    github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailFolders/messages/multiValueExtendedProperties v0.0.0-00010101000000-000000000000
    github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailFolders/messages/singleValueExtendedProperties v0.0.0-00010101000000-000000000000
    github.com/microsoft/kiota/abstractions/go v0.0.0-00010101000000-000000000000
    github.com/microsoft/kiota/abstractions/go/serialization v0.0.0-00010101000000-000000000000
)
