package models

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type InferenceClassificationOverride struct {
    classifyAs InferenceClassificationType;
    senderEmailAddress EmailAddress;
}
func (m *InferenceClassificationOverride) getClassifyAs()(InferenceClassificationType) {
    return m.classifyAs
}
func (m *InferenceClassificationOverride) getSenderEmailAddress()(EmailAddress) {
    return m.senderEmailAddress
}
func (m *InferenceClassificationOverride) getFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error), error) {
    return nil
}
func (m *InferenceClassificationOverride) serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    return nil
}
func (m *InferenceClassificationOverride) setClassifyAs(value InferenceClassificationType)() {
    m.classifyAs = value
}
func (m *InferenceClassificationOverride) setSenderEmailAddress(value EmailAddress)() {
    m.senderEmailAddress = value
}
