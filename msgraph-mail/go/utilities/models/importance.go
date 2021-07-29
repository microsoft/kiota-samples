package models
type Importance int

const (
    LOW Importance = iota
    NORMAL
    HIGH
)

func (i Importance) String() string {
    return []string{"LOW", "NORMAL", "HIGH"}[i]
}
