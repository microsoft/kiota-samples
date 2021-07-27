package utilities
import (
    "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users"
    "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item"
    "github.com/microsoft/kiota/abstractions/go/serialization"
    "github.com/microsoft/kiota/abstractions/go"
)
type ApiClient struct {
    httpCore HttpCore;
    pathSegment *string;
    Users UsersRequestBuilder;
}
