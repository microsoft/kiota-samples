package models
type Sensitivity int

const (
    NORMAL Sensitivity = iota
    PERSONAL
    PRIVATE
    CONFIDENTIAL
)

func (i Sensitivity) String() string {
    return []string{"NORMAL", "PERSONAL", "PRIVATE", "CONFIDENTIAL"}[i]
}
