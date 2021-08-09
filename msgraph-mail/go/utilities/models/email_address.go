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
func (m *EmailAddress) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["address"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        o.(*EmailAddress).SetAddress(val)
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        o.(*EmailAddress).SetName(val)
        return nil
    }
    return res
}
func (m *EmailAddress) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := writer.WritePrimitiveValue("address", m.GetAddress())
    if err != nil {
        return err
    }
    err = writer.WritePrimitiveValue("name", m.GetName())
    if err != nil {
        return err
    }
    err = writer.WriteAdditionalData(m.GetAdditionalData())
    if err != nil {
        return err
    }
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
