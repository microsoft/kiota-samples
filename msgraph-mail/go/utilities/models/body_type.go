package models
type BodyType int

const (
    TEXT_BODYTYPE BodyType = iota
    HTML_BODYTYPE
)

func (i BodyType) String() string {
    return []string{"TEXT", "HTML"}[i]
}
