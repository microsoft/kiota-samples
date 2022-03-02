package graph
// InferenceClassificationable 
type InferenceClassificationable interface {
    Entityable
    GetOverrides()([]InferenceClassificationOverrideable)
    SetOverrides(value []InferenceClassificationOverrideable)()
}
