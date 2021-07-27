package users
import (
    "github.com/microsoft/kiota/abstractions/go/serialization"
    "github.com/microsoft/kiota/abstractions/go"
)
type UsersRequestBuilder struct {
    currentPath *string;
    httpCore HttpCore;
    pathSegment *string;
}
