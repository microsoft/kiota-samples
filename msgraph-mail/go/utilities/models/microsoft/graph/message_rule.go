package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type MessageRule struct {
    Entity
    actions *MessageRuleActions;
    conditions *MessageRulePredicates;
    displayName *string;
    exceptions *MessageRulePredicates;
    hasError *bool;
    isEnabled *bool;
    isReadOnly *bool;
    sequence *int32;
}
func NewMessageRule()(*MessageRule) {
    m := &MessageRule{
        Entity: *NewEntity(),
    }
    return m
}
func (m *MessageRule) GetActions()(*MessageRuleActions) {
    return m.actions
}
func (m *MessageRule) GetConditions()(*MessageRulePredicates) {
    return m.conditions
}
func (m *MessageRule) GetDisplayName()(*string) {
    return m.displayName
}
func (m *MessageRule) GetExceptions()(*MessageRulePredicates) {
    return m.exceptions
}
func (m *MessageRule) GetHasError()(*bool) {
    return m.hasError
}
func (m *MessageRule) GetIsEnabled()(*bool) {
    return m.isEnabled
}
func (m *MessageRule) GetIsReadOnly()(*bool) {
    return m.isReadOnly
}
func (m *MessageRule) GetSequence()(*int32) {
    return m.sequence
}
func (m *MessageRule) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["actions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () interface{} { return NewMessageRuleActions() })
        if err != nil {
            return err
        }
        o.(*MessageRule).SetActions(val.(*MessageRuleActions))
        return nil
    }
    res["conditions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () interface{} { return NewMessageRulePredicates() })
        if err != nil {
            return err
        }
        o.(*MessageRule).SetConditions(val.(*MessageRulePredicates))
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        o.(*MessageRule).SetDisplayName(val)
        return nil
    }
    res["exceptions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () interface{} { return NewMessageRulePredicates() })
        if err != nil {
            return err
        }
        o.(*MessageRule).SetExceptions(val.(*MessageRulePredicates))
        return nil
    }
    res["hasError"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        o.(*MessageRule).SetHasError(val)
        return nil
    }
    res["isEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        o.(*MessageRule).SetIsEnabled(val)
        return nil
    }
    res["isReadOnly"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        o.(*MessageRule).SetIsReadOnly(val)
        return nil
    }
    res["sequence"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        o.(*MessageRule).SetSequence(val)
        return nil
    }
    return res
}
func (m *MessageRule) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("actions", m.GetActions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("conditions", m.GetConditions())
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
        err = writer.WriteObjectValue("exceptions", m.GetExceptions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WritePrimitiveValue("hasError", m.GetHasError())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WritePrimitiveValue("isEnabled", m.GetIsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WritePrimitiveValue("isReadOnly", m.GetIsReadOnly())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WritePrimitiveValue("sequence", m.GetSequence())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *MessageRule) SetActions(value *MessageRuleActions)() {
    m.actions = value
}
func (m *MessageRule) SetConditions(value *MessageRulePredicates)() {
    m.conditions = value
}
func (m *MessageRule) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *MessageRule) SetExceptions(value *MessageRulePredicates)() {
    m.exceptions = value
}
func (m *MessageRule) SetHasError(value *bool)() {
    m.hasError = value
}
func (m *MessageRule) SetIsEnabled(value *bool)() {
    m.isEnabled = value
}
func (m *MessageRule) SetIsReadOnly(value *bool)() {
    m.isReadOnly = value
}
func (m *MessageRule) SetSequence(value *int32)() {
    m.sequence = value
}
