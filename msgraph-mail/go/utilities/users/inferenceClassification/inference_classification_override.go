package inferenceclassification
import (
    "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users"
    "github.com/microsoft/kiota/abstractions/go/serialization"
)
type InferenceClassificationOverride struct {
    classifyAs InferenceClassificationType;
    senderEmailAddress EmailAddress;
}
