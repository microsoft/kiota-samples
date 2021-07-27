package mailfolders
import (
    "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users"
    "github.com/microsoft/kiota/abstractions/go/serialization"
)
type MailFoldersResponse struct {
    additionalData Map<String, Object>;
    nextLink *string;
    value []MailFolder;
}
