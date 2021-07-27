package users
import (
    "github.com/microsoft/kiota/abstractions/go/serialization"
)
type ItemBody struct {
    additionalData Map<String, Object>;
    content *string;
    contentType BodyType;
}
