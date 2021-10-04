package graph
import (
    "errors"
)
type InferenceClassificationType int

const (
    FOCUSED_INFERENCECLASSIFICATIONTYPE InferenceClassificationType = iota
    OTHER_INFERENCECLASSIFICATIONTYPE
)

func (i InferenceClassificationType) String() string {
    return []string{"FOCUSED", "OTHER"}[i]
}
func ParseInferenceClassificationType(v string) (interface{}, error) {
    switch v {
        case "FOCUSED":
            return FOCUSED_INFERENCECLASSIFICATIONTYPE, nil
        case "OTHER":
            return OTHER_INFERENCECLASSIFICATIONTYPE, nil
    }
    return 0, errors.New("Unknown InferenceClassificationType value: " + v)
}
