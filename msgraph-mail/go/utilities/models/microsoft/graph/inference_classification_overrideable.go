package graph
// InferenceClassificationOverrideable 
type InferenceClassificationOverrideable interface {
    Entityable
    GetClassifyAs()(*InferenceClassificationType)
    GetSenderEmailAddress()(EmailAddressable)
    SetClassifyAs(value *InferenceClassificationType)()
    SetSenderEmailAddress(value EmailAddressable)()
}
