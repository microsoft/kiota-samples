package models

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type InferenceClassificationOverride struct {
    Entity
    classifyAs *InferenceClassificationType;
    senderEmailAddress *EmailAddress;
}
func NewInferenceClassificationOverride()(*InferenceClassificationOverride) {
    m := &InferenceClassificationOverride{
        Entity: *NewEntity(),
    }
    return m
}
func (m *InferenceClassificationOverride) GetClassifyAs()(*InferenceClassificationType) {
    return m.classifyAs
}
func (m *InferenceClassificationOverride) GetSenderEmailAddress()(*EmailAddress) {
    return m.senderEmailAddress
}
func (m *InferenceClassificationOverride) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error), error) {
    return nil, nil
}
func (m *InferenceClassificationOverride) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    return nil
}
func (m *InferenceClassificationOverride) SetClassifyAs(value *InferenceClassificationType)() {
    m.classifyAs = value
}
func (m *InferenceClassificationOverride) SetSenderEmailAddress(value *EmailAddress)() {
    m.senderEmailAddress = value
}
