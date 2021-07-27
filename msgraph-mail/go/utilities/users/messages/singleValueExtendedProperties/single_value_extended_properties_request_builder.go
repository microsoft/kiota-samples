package singlevalueextendedproperties
import (
    "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users"
    "github.com/microsoft/kiota/abstractions/go/serialization"
    "github.com/microsoft/kiota/abstractions/go"
)
type SingleValueExtendedPropertiesRequestBuilder struct {
    currentPath *string;
    httpCore HttpCore;
    pathSegment *string;
}
