package models

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type MultiValueLegacyExtendedProperty struct {
    Entity
    value []string;
}
func NewMultiValueLegacyExtendedProperty()(*MultiValueLegacyExtendedProperty) {
    m := &MultiValueLegacyExtendedProperty{
        Entity: *NewEntity(),
    }
    return m
}
func (m *MultiValueLegacyExtendedProperty) getValue()([]string) {
    return m.value
}
func (m *MultiValueLegacyExtendedProperty) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error), error) {
    return nil, nil
}
func (m *MultiValueLegacyExtendedProperty) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    return nil
}
func (m *MultiValueLegacyExtendedProperty) setValue(value []string)() {
    m.value = value
}
