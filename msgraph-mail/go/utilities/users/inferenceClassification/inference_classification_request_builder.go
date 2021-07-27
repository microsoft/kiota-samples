package inferenceclassification
import (
    "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/inferenceClassification/overrides"
    "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/inferenceClassification/overrides/item"
    "github.com/microsoft/kiota/abstractions/go/serialization"
    "github.com/microsoft/kiota/abstractions/go"
)
type InferenceClassificationRequestBuilder struct {
    currentPath *string;
    httpCore HttpCore;
    Overrides OverridesRequestBuilder;
    pathSegment *string;
}
