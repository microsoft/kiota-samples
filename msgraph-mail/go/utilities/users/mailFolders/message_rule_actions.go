package mailfolders
import (
    "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users"
    "github.com/microsoft/kiota/abstractions/go/serialization"
)
type MessageRuleActions struct {
    additionalData Map<String, Object>;
    assignCategories []*string;
    copyToFolder *string;
    delete *bool;
    forwardAsAttachmentTo []Recipient;
    forwardTo []Recipient;
    markAsRead *bool;
    markImportance Importance;
    moveToFolder *string;
    permanentDelete *bool;
    redirectTo []Recipient;
    stopProcessingRules *bool;
}
