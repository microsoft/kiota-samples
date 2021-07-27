package overrides
import (
    "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/inferenceClassification"
    "github.com/microsoft/kiota/abstractions/go/serialization"
)
type OverridesResponse struct {
    additionalData Map<String, Object>;
    nextLink *string;
    value []InferenceClassificationOverride;
}
