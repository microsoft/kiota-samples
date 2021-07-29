package models

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type InferenceClassification struct {
    overrides []InferenceClassificationOverride;
}
func (m *InferenceClassification) getOverrides()([]InferenceClassificationOverride) {
    return nil
}
func (m *InferenceClassification) getFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error), error) {
    return nil
}
func (m *InferenceClassification) serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    return nil
}
func (m *InferenceClassification) setOverrides(value []InferenceClassificationOverride)() {
    return nil
}
