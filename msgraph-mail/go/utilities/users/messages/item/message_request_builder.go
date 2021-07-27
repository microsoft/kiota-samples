package item
import (
    "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/messages/Content"
    "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/messages/multiValueExtendedProperties"
    "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/messages/singleValueExtendedProperties"
    "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/messages/attachments"
    "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/messages/extensions"
    "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/messages/attachments/item"
    "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users"
    "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/messages/multiValueExtendedProperties/item"
    "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/messages/singleValueExtendedProperties/item"
    "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/messages/extensions/item"
    "github.com/microsoft/kiota/abstractions/go/serialization"
    "github.com/microsoft/kiota/abstractions/go"
)
type MessageRequestBuilder struct {
    Attachments AttachmentsRequestBuilder;
    Content ContentRequestBuilder;
    currentPath *string;
    Extensions ExtensionsRequestBuilder;
    httpCore HttpCore;
    MultiValueExtendedProperties MultiValueExtendedPropertiesRequestBuilder;
    pathSegment *string;
    SingleValueExtendedProperties SingleValueExtendedPropertiesRequestBuilder;
}
