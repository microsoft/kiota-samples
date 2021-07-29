package models
type InferenceClassificationType int

const (
    FOCUSED InferenceClassificationType = iota
    OTHER
)

func (i InferenceClassificationType) String() string {
    return []string{"FOCUSED", "OTHER"}[i]
}
