package item
import (
    "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailFolders/multiValueExtendedProperties"
    "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailFolders/singleValueExtendedProperties"
    "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailFolders/messages"
    "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailFolders/messageRules"
    "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailFolders/childFolders"
    "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailFolders/messageRules/item"
    "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailFolders/childFolders/item"
    "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailFolders/multiValueExtendedProperties/item"
    "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users"
    "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailFolders/singleValueExtendedProperties/item"
    "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailFolders/messages/item"
    "github.com/microsoft/kiota/abstractions/go/serialization"
    "github.com/microsoft/kiota/abstractions/go"
)
type MailFolderRequestBuilder struct {
    ChildFolders ChildFoldersRequestBuilder;
    currentPath *string;
    httpCore HttpCore;
    MessageRules MessageRulesRequestBuilder;
    Messages MessagesRequestBuilder;
    MultiValueExtendedProperties MultiValueExtendedPropertiesRequestBuilder;
    pathSegment *string;
    SingleValueExtendedProperties SingleValueExtendedPropertiesRequestBuilder;
}
