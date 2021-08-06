package multivalueextendedproperties

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/models"
)

type MultiValueExtendedPropertiesResponse struct {
    additionalData map[string]interface{};
    nextLink *string;
    value []ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MultiValueLegacyExtendedProperty;
}
func NewMultiValueExtendedPropertiesResponse()(*MultiValueExtendedPropertiesResponse) {
    m := &MultiValueExtendedPropertiesResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *MultiValueExtendedPropertiesResponse) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
func (m *MultiValueExtendedPropertiesResponse) GetNextLink()(*string) {
    return m.nextLink
}
func (m *MultiValueExtendedPropertiesResponse) GetValue()([]ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MultiValueLegacyExtendedProperty) {
    return m.value
}
func (m *MultiValueExtendedPropertiesResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error), error) {
    return nil, nil
}
func (m *MultiValueExtendedPropertiesResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    return nil
}
func (m *MultiValueExtendedPropertiesResponse) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *MultiValueExtendedPropertiesResponse) SetNextLink(value *string)() {
    m.nextLink = value
}
func (m *MultiValueExtendedPropertiesResponse) SetValue(value []ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MultiValueLegacyExtendedProperty)() {
    m.value = value
}
