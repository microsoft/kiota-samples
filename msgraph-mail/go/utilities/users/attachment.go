package users
import (
    "github.com/microsoft/kiota/abstractions/go/serialization"
    "time"
)
type Attachment struct {
    contentType *string;
    isInline *bool;
    lastModifiedDateTime Time;
    name *string;
    size *int32;
}
