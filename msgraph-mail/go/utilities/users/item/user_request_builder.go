package item
import (
    "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/messages"
    "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailFolders"
    "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/inferenceClassification"
    "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailFolders/item"
    "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/messages/item"
    "github.com/microsoft/kiota/abstractions/go/serialization"
    "github.com/microsoft/kiota/abstractions/go"
)
type UserRequestBuilder struct {
    currentPath *string;
    httpCore HttpCore;
    InferenceClassification InferenceClassificationRequestBuilder;
    MailFolders MailFoldersRequestBuilder;
    Messages MessagesRequestBuilder;
    pathSegment *string;
}
