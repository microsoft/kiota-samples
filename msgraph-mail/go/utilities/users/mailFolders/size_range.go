package mailfolders
import (
    "github.com/microsoft/kiota/abstractions/go/serialization"
)
type SizeRange struct {
    additionalData Map<String, Object>;
    maximumSize *int32;
    minimumSize *int32;
}
