package models

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Entity struct {
    additionalData map[string]interface{};
    id *string;
}
func (m *Entity) getAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
func (m *Entity) getId()(*string) {
    return m.id
}
func (m *Entity) getFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error), error) {
    return nil
}
func (m *Entity) serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    return nil
}
func (m *Entity) setAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *Entity) setId(value *string)() {
    m.id = value
}
