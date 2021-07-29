package models

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SizeRange struct {
    additionalData map[string]interface{};
    maximumSize *int32;
    minimumSize *int32;
}
func (m *SizeRange) getAdditionalData()(map[string]interface{}) {
    return nil
}
func (m *SizeRange) getMaximumSize()(*int32) {
    return nil
}
func (m *SizeRange) getMinimumSize()(*int32) {
    return nil
}
func (m *SizeRange) getFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error), error) {
    return nil
}
func (m *SizeRange) serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    return nil
}
func (m *SizeRange) setAdditionalData(value map[string]interface{})() {
    return nil
}
func (m *SizeRange) setMaximumSize(value *int32)() {
    return nil
}
func (m *SizeRange) setMinimumSize(value *int32)() {
    return nil
}
