package models

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type EmailAddress struct {
    additionalData map[string]interface{};
    address *string;
    name *string;
}
func NewEmailAddress()(*EmailAddress) {
    m := &EmailAddress{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *EmailAddress) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
func (m *EmailAddress) GetAddress()(*string) {
    return m.address
}
func (m *EmailAddress) GetName()(*string) {
    return m.name
}
func (m *EmailAddress) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error), error) {
    return nil, nil
}
func (m *EmailAddress) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    return nil
}
func (m *EmailAddress) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *EmailAddress) SetAddress(value *string)() {
    m.address = value
}
func (m *EmailAddress) SetName(value *string)() {
    m.name = value
}
