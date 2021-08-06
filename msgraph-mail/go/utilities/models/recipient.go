package models

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Recipient struct {
    additionalData map[string]interface{};
    emailAddress *EmailAddress;
}
func NewRecipient()(*Recipient) {
    m := &Recipient{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *Recipient) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
func (m *Recipient) GetEmailAddress()(*EmailAddress) {
    return m.emailAddress
}
func (m *Recipient) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error), error) {
    return nil, nil
}
func (m *Recipient) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    return nil
}
func (m *Recipient) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *Recipient) SetEmailAddress(value *EmailAddress)() {
    m.emailAddress = value
}
