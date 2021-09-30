package mailfolders

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/models/microsoft/graph"
)

type MailFoldersResponse struct {
    additionalData map[string]interface{};
    nextLink *string;
    value []i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.MailFolder;
}
func NewMailFoldersResponse()(*MailFoldersResponse) {
    m := &MailFoldersResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *MailFoldersResponse) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
func (m *MailFoldersResponse) GetNextLink()(*string) {
    return m.nextLink
}
func (m *MailFoldersResponse) GetValue()([]i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.MailFolder) {
    return m.value
}
func (m *MailFoldersResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["@odata.nextLink"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        o.(*MailFoldersResponse).SetNextLink(val)
        return nil
    }
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () interface{} { return i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.NewMailFolder() })
        if err != nil {
            return err
        }
        res := make([]i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.MailFolder, len(val))
        for i, v := range val {
            res[i] = *(v.(*i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.MailFolder))
        }
        o.(*MailFoldersResponse).SetValue(res)
        return nil
    }
    return res
}
func (m *MailFoldersResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
func (m *MailFoldersResponse) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *MailFoldersResponse) SetNextLink(value *string)() {
    m.nextLink = value
}
func (m *MailFoldersResponse) SetValue(value []i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.MailFolder)() {
    m.value = value
}
