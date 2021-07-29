package models
type FollowupFlagStatus int

const (
    NOTFLAGGED FollowupFlagStatus = iota
    COMPLETE
    FLAGGED
)

func (i FollowupFlagStatus) String() string {
    return []string{"NOTFLAGGED", "COMPLETE", "FLAGGED"}[i]
}
