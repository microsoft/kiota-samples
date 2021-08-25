package models
import (
    "errors"
)
type BodyType int

const (
    TEXT_BODYTYPE BodyType = iota
    HTML_BODYTYPE
)

func (i BodyType) String() string {
    return []string{"TEXT", "HTML"}[i]
}
func ParseBodyType(v string) (interface{}, error) {
    switch v {
        case "TEXT":
            return TEXT_BODYTYPE, nil
        case "HTML":
            return HTML_BODYTYPE, nil
    }
    return 0, errors.New("Unknown BodyType value: " + v)
}
