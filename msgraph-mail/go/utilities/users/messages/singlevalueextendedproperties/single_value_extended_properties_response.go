package singlevalueextendedproperties

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/models"
)

type SingleValueExtendedPropertiesResponse struct {
    additionalData map[string]interface{};
    nextLink *string;
    value []ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.SingleValueLegacyExtendedProperty;
}
func NewSingleValueExtendedPropertiesResponse()(*SingleValueExtendedPropertiesResponse) {
    return nil
}
func (m *SingleValueExtendedPropertiesResponse) getAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
func (m *SingleValueExtendedPropertiesResponse) getNextLink()(*string) {
    return m.nextLink
}
func (m *SingleValueExtendedPropertiesResponse) getValue()([]ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.SingleValueLegacyExtendedProperty) {
    return m.value
}
func (m *SingleValueExtendedPropertiesResponse) getFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error), error) {
    return nil
}
func (m *SingleValueExtendedPropertiesResponse) serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    return nil
}
func (m *SingleValueExtendedPropertiesResponse) setAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *SingleValueExtendedPropertiesResponse) setNextLink(value *string)() {
    m.nextLink = value
}
func (m *SingleValueExtendedPropertiesResponse) setValue(value []ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.SingleValueLegacyExtendedProperty)() {
    m.value = value
}
