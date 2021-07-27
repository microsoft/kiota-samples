package content
import (
    "github.com/microsoft/kiota/abstractions/go/serialization"
    "github.com/microsoft/kiota/abstractions/go"
)
type ContentRequestBuilder struct {
    currentPath *string;
    httpCore HttpCore;
    pathSegment *string;
}
