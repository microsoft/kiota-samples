package models

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DateTimeTimeZone struct {
    additionalData map[string]interface{};
    dateTime *string;
    timeZone *string;
}
func NewDateTimeTimeZone()(*DateTimeTimeZone) {
    m := &DateTimeTimeZone{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *DateTimeTimeZone) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
func (m *DateTimeTimeZone) GetDateTime()(*string) {
    return m.dateTime
}
func (m *DateTimeTimeZone) GetTimeZone()(*string) {
    return m.timeZone
}
func (m *DateTimeTimeZone) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error), error) {
    return nil, nil
}
func (m *DateTimeTimeZone) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := writer.WritePrimitiveValue("dateTime", m.GetDateTime())
    if err != nil {
        return err
    }
    err = writer.WritePrimitiveValue("timeZone", m.GetTimeZone())
    if err != nil {
        return err
    }
    err = writer.WriteAdditionalData(m.GetAdditionalData())
    if err != nil {
        return err
    }
    return nil
}
func (m *DateTimeTimeZone) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *DateTimeTimeZone) SetDateTime(value *string)() {
    m.dateTime = value
}
func (m *DateTimeTimeZone) SetTimeZone(value *string)() {
    m.timeZone = value
}
