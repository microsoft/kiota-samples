package models

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SizeRange struct {
    additionalData map[string]interface{};
    maximumSize *int32;
    minimumSize *int32;
}
func NewSizeRange()(*SizeRange) {
    return nil
}
func (m *SizeRange) getAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
func (m *SizeRange) getMaximumSize()(*int32) {
    return m.maximumSize
}
func (m *SizeRange) getMinimumSize()(*int32) {
    return m.minimumSize
}
func (m *SizeRange) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error), error) {
    return nil, nil
}
func (m *SizeRange) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    return nil
}
func (m *SizeRange) setAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *SizeRange) setMaximumSize(value *int32)() {
    m.maximumSize = value
}
func (m *SizeRange) setMinimumSize(value *int32)() {
    m.minimumSize = value
}
