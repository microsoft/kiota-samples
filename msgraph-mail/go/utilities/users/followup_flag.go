package users
import (
    "github.com/microsoft/kiota/abstractions/go/serialization"
)
type FollowupFlag struct {
    additionalData Map<String, Object>;
    completedDateTime DateTimeTimeZone;
    dueDateTime DateTimeTimeZone;
    flagStatus FollowupFlagStatus;
    startDateTime DateTimeTimeZone;
}
