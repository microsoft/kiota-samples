package models
type FollowupFlag struct {
    additionalData map[string]interface{};
    completedDateTime DateTimeTimeZone;
    dueDateTime DateTimeTimeZone;
    flagStatus FollowupFlagStatus;
    startDateTime DateTimeTimeZone;
}
