package users
import (
    "github.com/microsoft/kiota/abstractions/go/serialization"
)
type EmailAddress struct {
    additionalData Map<String, Object>;
    address *string;
    name *string;
}
