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
    return nil
}
func (m *EmailAddress) getAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
func (m *EmailAddress) getAddress()(*string) {
    return m.address
}
func (m *EmailAddress) getName()(*string) {
    return m.name
}
func (m *EmailAddress) getFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error), error) {
    return nil
}
func (m *EmailAddress) serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    return nil
}
func (m *EmailAddress) setAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *EmailAddress) setAddress(value *string)() {
    m.address = value
}
func (m *EmailAddress) setName(value *string)() {
    m.name = value
}
