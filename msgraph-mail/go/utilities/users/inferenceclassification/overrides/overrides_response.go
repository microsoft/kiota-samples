package overrides

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/models"
)

type OverridesResponse struct {
    additionalData map[string]interface{};
    nextLink *string;
    value []ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.InferenceClassificationOverride;
}
func NewOverridesResponse()(*OverridesResponse) {
    return nil
}
func (m *OverridesResponse) getAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
func (m *OverridesResponse) getNextLink()(*string) {
    return m.nextLink
}
func (m *OverridesResponse) getValue()([]ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.InferenceClassificationOverride) {
    return m.value
}
func (m *OverridesResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error), error) {
    return nil, nil
}
func (m *OverridesResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    return nil
}
func (m *OverridesResponse) setAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *OverridesResponse) setNextLink(value *string)() {
    m.nextLink = value
}
func (m *OverridesResponse) setValue(value []ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.InferenceClassificationOverride)() {
    m.value = value
}
