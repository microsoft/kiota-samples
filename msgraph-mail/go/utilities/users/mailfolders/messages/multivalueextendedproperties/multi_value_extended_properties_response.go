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
func (m *MultiValueExtendedPropertiesResponse) getAdditionalData()(map[string]interface{}) {
    return nil
}
func (m *MultiValueExtendedPropertiesResponse) getNextLink()(*string) {
    return nil
}
func (m *MultiValueExtendedPropertiesResponse) getValue()([]ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MultiValueLegacyExtendedProperty) {
    return nil
}
func (m *MultiValueExtendedPropertiesResponse) getFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error), error) {
    return nil
}
func (m *MultiValueExtendedPropertiesResponse) serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    return nil
}
func (m *MultiValueExtendedPropertiesResponse) setAdditionalData(value map[string]interface{})() {
    return nil
}
func (m *MultiValueExtendedPropertiesResponse) setNextLink(value *string)() {
    return nil
}
func (m *MultiValueExtendedPropertiesResponse) setValue(value []ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MultiValueLegacyExtendedProperty)() {
    return nil
}
