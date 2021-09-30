package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type MessageRuleActions struct {
    additionalData map[string]interface{};
    assignCategories []string;
    copyToFolder *string;
    delete *bool;
    forwardAsAttachmentTo []Recipient;
    forwardTo []Recipient;
    markAsRead *bool;
    markImportance *Importance;
    moveToFolder *string;
    permanentDelete *bool;
    redirectTo []Recipient;
    stopProcessingRules *bool;
}
func NewMessageRuleActions()(*MessageRuleActions) {
    m := &MessageRuleActions{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *MessageRuleActions) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
func (m *MessageRuleActions) GetAssignCategories()([]string) {
    return m.assignCategories
}
func (m *MessageRuleActions) GetCopyToFolder()(*string) {
    return m.copyToFolder
}
func (m *MessageRuleActions) GetDelete()(*bool) {
    return m.delete
}
func (m *MessageRuleActions) GetForwardAsAttachmentTo()([]Recipient) {
    return m.forwardAsAttachmentTo
}
func (m *MessageRuleActions) GetForwardTo()([]Recipient) {
    return m.forwardTo
}
func (m *MessageRuleActions) GetMarkAsRead()(*bool) {
    return m.markAsRead
}
func (m *MessageRuleActions) GetMarkImportance()(*Importance) {
    return m.markImportance
}
func (m *MessageRuleActions) GetMoveToFolder()(*string) {
    return m.moveToFolder
}
func (m *MessageRuleActions) GetPermanentDelete()(*bool) {
    return m.permanentDelete
}
func (m *MessageRuleActions) GetRedirectTo()([]Recipient) {
    return m.redirectTo
}
func (m *MessageRuleActions) GetStopProcessingRules()(*bool) {
    return m.stopProcessingRules
}
func (m *MessageRuleActions) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["assignCategories"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        o.(*MessageRuleActions).SetAssignCategories(res)
        return nil
    }
    res["copyToFolder"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        o.(*MessageRuleActions).SetCopyToFolder(val)
        return nil
    }
    res["delete"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        o.(*MessageRuleActions).SetDelete(val)
        return nil
    }
    res["forwardAsAttachmentTo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () interface{} { return NewRecipient() })
        if err != nil {
            return err
        }
        res := make([]Recipient, len(val))
        for i, v := range val {
            res[i] = *(v.(*Recipient))
        }
        o.(*MessageRuleActions).SetForwardAsAttachmentTo(res)
        return nil
    }
    res["forwardTo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () interface{} { return NewRecipient() })
        if err != nil {
            return err
        }
        res := make([]Recipient, len(val))
        for i, v := range val {
            res[i] = *(v.(*Recipient))
        }
        o.(*MessageRuleActions).SetForwardTo(res)
        return nil
    }
    res["markAsRead"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        o.(*MessageRuleActions).SetMarkAsRead(val)
        return nil
    }
    res["markImportance"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseImportance)
        if err != nil {
            return err
        }
        o.(*MessageRuleActions).SetMarkImportance(val.(*Importance))
        return nil
    }
    res["moveToFolder"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        o.(*MessageRuleActions).SetMoveToFolder(val)
        return nil
    }
    res["permanentDelete"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        o.(*MessageRuleActions).SetPermanentDelete(val)
        return nil
    }
    res["redirectTo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () interface{} { return NewRecipient() })
        if err != nil {
            return err
        }
        res := make([]Recipient, len(val))
        for i, v := range val {
            res[i] = *(v.(*Recipient))
        }
        o.(*MessageRuleActions).SetRedirectTo(res)
        return nil
    }
    res["stopProcessingRules"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        o.(*MessageRuleActions).SetStopProcessingRules(val)
        return nil
    }
    return res
}
func (m *MessageRuleActions) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfPrimitiveValues("assignCategories", i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ConvertToArrayOfPrimitives(m.GetAssignCategories()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WritePrimitiveValue("copyToFolder", m.GetCopyToFolder())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WritePrimitiveValue("delete", m.GetDelete())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfObjectValues("forwardAsAttachmentTo", i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ConvertToArrayOfParsable(m.GetForwardAsAttachmentTo()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfObjectValues("forwardTo", i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ConvertToArrayOfParsable(m.GetForwardTo()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WritePrimitiveValue("markAsRead", m.GetMarkAsRead())
        if err != nil {
            return err
        }
    }
    {
        if m.GetMarkImportance() != nil {
            err := writer.WritePrimitiveValue("markImportance", m.GetMarkImportance().String())
            if err != nil {
                return err
            }
        }
    }
    {
        err := writer.WritePrimitiveValue("moveToFolder", m.GetMoveToFolder())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WritePrimitiveValue("permanentDelete", m.GetPermanentDelete())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfObjectValues("redirectTo", i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ConvertToArrayOfParsable(m.GetRedirectTo()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WritePrimitiveValue("stopProcessingRules", m.GetStopProcessingRules())
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
func (m *MessageRuleActions) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *MessageRuleActions) SetAssignCategories(value []string)() {
    m.assignCategories = value
}
func (m *MessageRuleActions) SetCopyToFolder(value *string)() {
    m.copyToFolder = value
}
func (m *MessageRuleActions) SetDelete(value *bool)() {
    m.delete = value
}
func (m *MessageRuleActions) SetForwardAsAttachmentTo(value []Recipient)() {
    m.forwardAsAttachmentTo = value
}
func (m *MessageRuleActions) SetForwardTo(value []Recipient)() {
    m.forwardTo = value
}
func (m *MessageRuleActions) SetMarkAsRead(value *bool)() {
    m.markAsRead = value
}
func (m *MessageRuleActions) SetMarkImportance(value *Importance)() {
    m.markImportance = value
}
func (m *MessageRuleActions) SetMoveToFolder(value *string)() {
    m.moveToFolder = value
}
func (m *MessageRuleActions) SetPermanentDelete(value *bool)() {
    m.permanentDelete = value
}
func (m *MessageRuleActions) SetRedirectTo(value []Recipient)() {
    m.redirectTo = value
}
func (m *MessageRuleActions) SetStopProcessingRules(value *bool)() {
    m.stopProcessingRules = value
}
