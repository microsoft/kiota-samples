package models

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type InternetMessageHeader struct {
    additionalData map[string]interface{};
    name *string;
    value *string;
}
func NewInternetMessageHeader()(*InternetMessageHeader) {
    m := &InternetMessageHeader{
    }
    m.setAdditionalData(make(map[string]interface{}));
    return m
}
func (m *InternetMessageHeader) getAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
func (m *InternetMessageHeader) getName()(*string) {
    return m.name
}
func (m *InternetMessageHeader) getValue()(*string) {
    return m.value
}
func (m *InternetMessageHeader) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error), error) {
    return nil, nil
}
func (m *InternetMessageHeader) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    return nil
}
func (m *InternetMessageHeader) setAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *InternetMessageHeader) setName(value *string)() {
    m.name = value
}
func (m *InternetMessageHeader) setValue(value *string)() {
    m.value = value
}
