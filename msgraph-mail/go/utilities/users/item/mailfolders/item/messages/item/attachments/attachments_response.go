package attachments

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/models/microsoft/graph"
)

type AttachmentsResponse struct {
    additionalData map[string]interface{};
    nextLink *string;
    value []i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.Attachment;
}
func NewAttachmentsResponse()(*AttachmentsResponse) {
    m := &AttachmentsResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AttachmentsResponse) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
func (m *AttachmentsResponse) GetNextLink()(*string) {
    return m.nextLink
}
func (m *AttachmentsResponse) GetValue()([]i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.Attachment) {
    return m.value
}
func (m *AttachmentsResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["@odata.nextLink"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNextLink(val)
        return nil
    }
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.NewAttachment() })
        if err != nil {
            return err
        }
        res := make([]i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.Attachment, len(val))
        for i, v := range val {
            res[i] = *(v.(*i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.Attachment))
        }
        m.SetValue(res)
        return nil
    }
    return res
}
func (m *AttachmentsResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WritePrimitiveValue("@odata.nextLink", m.GetNextLink())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfObjectValues("value", i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ConvertToArrayOfParsable(m.GetValue()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *AttachmentsResponse) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AttachmentsResponse) SetNextLink(value *string)() {
    m.nextLink = value
}
func (m *AttachmentsResponse) SetValue(value []i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.Attachment)() {
    m.value = value
}
