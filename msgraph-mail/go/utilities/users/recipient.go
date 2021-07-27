package users
import (
    "github.com/microsoft/kiota/abstractions/go/serialization"
)
type Recipient struct {
    additionalData Map<String, Object>;
    emailAddress EmailAddress;
}
