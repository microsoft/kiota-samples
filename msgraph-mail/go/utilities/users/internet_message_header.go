package users
import (
    "github.com/microsoft/kiota/abstractions/go/serialization"
)
type InternetMessageHeader struct {
    additionalData Map<String, Object>;
    name *string;
    value *string;
}
