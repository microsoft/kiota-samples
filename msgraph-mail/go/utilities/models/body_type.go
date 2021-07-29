package models
type BodyType int

const (
    TEXT BodyType = iota
    HTML
)

func (i BodyType) String() string {
    return []string{"TEXT", "HTML"}[i]
}
