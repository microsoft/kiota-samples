package graph
import (
    "errors"
)
type Importance int

const (
    LOW_IMPORTANCE Importance = iota
    NORMAL_IMPORTANCE
    HIGH_IMPORTANCE
)

func (i Importance) String() string {
    return []string{"LOW", "NORMAL", "HIGH"}[i]
}
func ParseImportance(v string) (interface{}, error) {
    switch v {
        case "LOW":
            return LOW_IMPORTANCE, nil
        case "NORMAL":
            return NORMAL_IMPORTANCE, nil
        case "HIGH":
            return HIGH_IMPORTANCE, nil
    }
    return 0, errors.New("Unknown Importance value: " + v)
}
