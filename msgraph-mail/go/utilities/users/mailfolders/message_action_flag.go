package mailfolders
type MessageActionFlag int

const (
    ANY MessageActionFlag = iota
    CALL
    DONOTFORWARD
    FOLLOWUP
    FYI
    FORWARD
    NORESPONSENECESSARY
    READ
    REPLY
    REPLYTOALL
    REVIEW
)

func (i MessageActionFlag) String() string {
    return []string{"ANY", "CALL", "DONOTFORWARD", "FOLLOWUP", "FYI", "FORWARD", "NORESPONSENECESSARY", "READ", "REPLY", "REPLYTOALL", "REVIEW"}[i]
}
