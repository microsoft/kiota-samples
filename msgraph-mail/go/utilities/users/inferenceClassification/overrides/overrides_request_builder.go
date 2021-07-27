package overrides
import (
    "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/inferenceClassification"
    "github.com/microsoft/kiota/abstractions/go/serialization"
    "github.com/microsoft/kiota/abstractions/go"
)
type OverridesRequestBuilder struct {
    currentPath *string;
    httpCore HttpCore;
    pathSegment *string;
}
