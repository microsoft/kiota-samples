package users
import (
    "github.com/microsoft/kiota/abstractions/go/serialization"
    "time"
)
type OutlookItem struct {
    categories []*string;
    changeKey *string;
    createdDateTime Time;
    lastModifiedDateTime Time;
}
