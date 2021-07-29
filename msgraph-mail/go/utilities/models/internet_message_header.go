package models

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type InternetMessageHeader struct {
    additionalData map[string]interface{};
    name *string;
    value *string;
}
func (m *InternetMessageHeader) getAdditionalData()(map[string]interface{}) {
    return nil
}
func (m *InternetMessageHeader) getName()(*string) {
    return nil
}
func (m *InternetMessageHeader) getValue()(*string) {
    return nil
}
func (m *InternetMessageHeader) getFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error), error) {
    return nil
}
func (m *InternetMessageHeader) serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    return nil
}
func (m *InternetMessageHeader) setAdditionalData(value map[string]interface{})() {
    return nil
}
func (m *InternetMessageHeader) setName(value *string)() {
    return nil
}
func (m *InternetMessageHeader) setValue(value *string)() {
    return nil
}
