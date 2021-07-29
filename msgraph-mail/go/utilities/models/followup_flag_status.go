package models
type FollowupFlagStatus int

const (
    NOTFLAGGED_FOLLOWUPFLAGSTATUS FollowupFlagStatus = iota
    COMPLETE_FOLLOWUPFLAGSTATUS
    FLAGGED_FOLLOWUPFLAGSTATUS
)

func (i FollowupFlagStatus) String() string {
    return []string{"NOTFLAGGED", "COMPLETE", "FLAGGED"}[i]
}
