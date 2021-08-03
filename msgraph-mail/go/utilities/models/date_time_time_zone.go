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
    return nil
}
func (m *DateTimeTimeZone) getAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
func (m *DateTimeTimeZone) getDateTime()(*string) {
    return m.dateTime
}
func (m *DateTimeTimeZone) getTimeZone()(*string) {
    return m.timeZone
}
func (m *DateTimeTimeZone) getFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error), error) {
    return nil, nil
}
func (m *DateTimeTimeZone) serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    return nil
}
func (m *DateTimeTimeZone) setAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *DateTimeTimeZone) setDateTime(value *string)() {
    m.dateTime = value
}
func (m *DateTimeTimeZone) setTimeZone(value *string)() {
    m.timeZone = value
}
