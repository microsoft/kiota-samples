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
    m := &SingleValueExtendedPropertiesResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *SingleValueExtendedPropertiesResponse) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
func (m *SingleValueExtendedPropertiesResponse) GetNextLink()(*string) {
    return m.nextLink
}
func (m *SingleValueExtendedPropertiesResponse) GetValue()([]ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.SingleValueLegacyExtendedProperty) {
    return m.value
}
func (m *SingleValueExtendedPropertiesResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["@odata.nextLink"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        o.(*SingleValueExtendedPropertiesResponse).SetNextLink(val)
        return nil
    }
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () interface{} { return ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.NewSingleValueLegacyExtendedProperty() })
        if err != nil {
            return err
        }
        res := make([]ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.SingleValueLegacyExtendedProperty, len(val))
        for i, v := range val {
            res[i] = *(v.(*ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.SingleValueLegacyExtendedProperty))
        }
        o.(*SingleValueExtendedPropertiesResponse).SetValue(res)
        return nil
    }
    return res
}
func (m *SingleValueExtendedPropertiesResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := writer.WritePrimitiveValue("@odata.nextLink", m.GetNextLink())
    if err != nil {
        return err
    }
    err = writer.WriteCollectionOfObjectValues("value", i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ConvertToArrayOfParsable(m.GetValue()))
    if err != nil {
        return err
    }
    err = writer.WriteAdditionalData(m.GetAdditionalData())
    if err != nil {
        return err
    }
    return nil
}
func (m *SingleValueExtendedPropertiesResponse) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *SingleValueExtendedPropertiesResponse) SetNextLink(value *string)() {
    m.nextLink = value
}
func (m *SingleValueExtendedPropertiesResponse) SetValue(value []ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.SingleValueLegacyExtendedProperty)() {
    m.value = value
}
