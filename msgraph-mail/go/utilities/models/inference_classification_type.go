package models
type InferenceClassificationType int

const (
    FOCUSED_INFERENCECLASSIFICATIONTYPE InferenceClassificationType = iota
    OTHER_INFERENCECLASSIFICATIONTYPE
)

func (i InferenceClassificationType) String() string {
    return []string{"FOCUSED", "OTHER"}[i]
}
