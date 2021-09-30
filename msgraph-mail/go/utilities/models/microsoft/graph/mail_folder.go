package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type MailFolder struct {
    Entity
    childFolderCount *int32;
    childFolders []MailFolder;
    displayName *string;
    isHidden *bool;
    messageRules []MessageRule;
    messages []Message;
    multiValueExtendedProperties []MultiValueLegacyExtendedProperty;
    parentFolderId *string;
    singleValueExtendedProperties []SingleValueLegacyExtendedProperty;
    totalItemCount *int32;
    unreadItemCount *int32;
}
func NewMailFolder()(*MailFolder) {
    m := &MailFolder{
        Entity: *NewEntity(),
    }
    return m
}
func (m *MailFolder) GetChildFolderCount()(*int32) {
    return m.childFolderCount
}
func (m *MailFolder) GetChildFolders()([]MailFolder) {
    return m.childFolders
}
func (m *MailFolder) GetDisplayName()(*string) {
    return m.displayName
}
func (m *MailFolder) GetIsHidden()(*bool) {
    return m.isHidden
}
func (m *MailFolder) GetMessageRules()([]MessageRule) {
    return m.messageRules
}
func (m *MailFolder) GetMessages()([]Message) {
    return m.messages
}
func (m *MailFolder) GetMultiValueExtendedProperties()([]MultiValueLegacyExtendedProperty) {
    return m.multiValueExtendedProperties
}
func (m *MailFolder) GetParentFolderId()(*string) {
    return m.parentFolderId
}
func (m *MailFolder) GetSingleValueExtendedProperties()([]SingleValueLegacyExtendedProperty) {
    return m.singleValueExtendedProperties
}
func (m *MailFolder) GetTotalItemCount()(*int32) {
    return m.totalItemCount
}
func (m *MailFolder) GetUnreadItemCount()(*int32) {
    return m.unreadItemCount
}
func (m *MailFolder) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["childFolderCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        o.(*MailFolder).SetChildFolderCount(val)
        return nil
    }
    res["childFolders"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () interface{} { return NewMailFolder() })
        if err != nil {
            return err
        }
        res := make([]MailFolder, len(val))
        for i, v := range val {
            res[i] = *(v.(*MailFolder))
        }
        o.(*MailFolder).SetChildFolders(res)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        o.(*MailFolder).SetDisplayName(val)
        return nil
    }
    res["isHidden"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        o.(*MailFolder).SetIsHidden(val)
        return nil
    }
    res["messageRules"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () interface{} { return NewMessageRule() })
        if err != nil {
            return err
        }
        res := make([]MessageRule, len(val))
        for i, v := range val {
            res[i] = *(v.(*MessageRule))
        }
        o.(*MailFolder).SetMessageRules(res)
        return nil
    }
    res["messages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () interface{} { return NewMessage() })
        if err != nil {
            return err
        }
        res := make([]Message, len(val))
        for i, v := range val {
            res[i] = *(v.(*Message))
        }
        o.(*MailFolder).SetMessages(res)
        return nil
    }
    res["multiValueExtendedProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () interface{} { return NewMultiValueLegacyExtendedProperty() })
        if err != nil {
            return err
        }
        res := make([]MultiValueLegacyExtendedProperty, len(val))
        for i, v := range val {
            res[i] = *(v.(*MultiValueLegacyExtendedProperty))
        }
        o.(*MailFolder).SetMultiValueExtendedProperties(res)
        return nil
    }
    res["parentFolderId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        o.(*MailFolder).SetParentFolderId(val)
        return nil
    }
    res["singleValueExtendedProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () interface{} { return NewSingleValueLegacyExtendedProperty() })
        if err != nil {
            return err
        }
        res := make([]SingleValueLegacyExtendedProperty, len(val))
        for i, v := range val {
            res[i] = *(v.(*SingleValueLegacyExtendedProperty))
        }
        o.(*MailFolder).SetSingleValueExtendedProperties(res)
        return nil
    }
    res["totalItemCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        o.(*MailFolder).SetTotalItemCount(val)
        return nil
    }
    res["unreadItemCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        o.(*MailFolder).SetUnreadItemCount(val)
        return nil
    }
    return res
}
func (m *MailFolder) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WritePrimitiveValue("childFolderCount", m.GetChildFolderCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfObjectValues("childFolders", i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ConvertToArrayOfParsable(m.GetChildFolders()))
        if err != nil {
            return err
        }
    }
    {
        err = writer.WritePrimitiveValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WritePrimitiveValue("isHidden", m.GetIsHidden())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfObjectValues("messageRules", i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ConvertToArrayOfParsable(m.GetMessageRules()))
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfObjectValues("messages", i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ConvertToArrayOfParsable(m.GetMessages()))
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfObjectValues("multiValueExtendedProperties", i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ConvertToArrayOfParsable(m.GetMultiValueExtendedProperties()))
        if err != nil {
            return err
        }
    }
    {
        err = writer.WritePrimitiveValue("parentFolderId", m.GetParentFolderId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfObjectValues("singleValueExtendedProperties", i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ConvertToArrayOfParsable(m.GetSingleValueExtendedProperties()))
        if err != nil {
            return err
        }
    }
    {
        err = writer.WritePrimitiveValue("totalItemCount", m.GetTotalItemCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WritePrimitiveValue("unreadItemCount", m.GetUnreadItemCount())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *MailFolder) SetChildFolderCount(value *int32)() {
    m.childFolderCount = value
}
func (m *MailFolder) SetChildFolders(value []MailFolder)() {
    m.childFolders = value
}
func (m *MailFolder) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *MailFolder) SetIsHidden(value *bool)() {
    m.isHidden = value
}
func (m *MailFolder) SetMessageRules(value []MessageRule)() {
    m.messageRules = value
}
func (m *MailFolder) SetMessages(value []Message)() {
    m.messages = value
}
func (m *MailFolder) SetMultiValueExtendedProperties(value []MultiValueLegacyExtendedProperty)() {
    m.multiValueExtendedProperties = value
}
func (m *MailFolder) SetParentFolderId(value *string)() {
    m.parentFolderId = value
}
func (m *MailFolder) SetSingleValueExtendedProperties(value []SingleValueLegacyExtendedProperty)() {
    m.singleValueExtendedProperties = value
}
func (m *MailFolder) SetTotalItemCount(value *int32)() {
    m.totalItemCount = value
}
func (m *MailFolder) SetUnreadItemCount(value *int32)() {
    m.unreadItemCount = value
}
