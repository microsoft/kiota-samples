package models
type Importance int

const (
    LOW_IMPORTANCE Importance = iota
    NORMAL_IMPORTANCE
    HIGH_IMPORTANCE
)

func (i Importance) String() string {
    return []string{"LOW", "NORMAL", "HIGH"}[i]
}
