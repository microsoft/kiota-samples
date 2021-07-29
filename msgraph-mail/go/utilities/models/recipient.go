package models

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Recipient struct {
    additionalData map[string]interface{};
    emailAddress EmailAddress;
}
func NewRecipient()(*Recipient) {
    return nil
}
func (m *Recipient) getAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
func (m *Recipient) getEmailAddress()(EmailAddress) {
    return m.emailAddress
}
func (m *Recipient) getFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error), error) {
    return nil
}
func (m *Recipient) serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    return nil
}
func (m *Recipient) setAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *Recipient) setEmailAddress(value EmailAddress)() {
    m.emailAddress = value
}
