package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type MessageRule struct {
    Entity
    // The actions property
    actions MessageRuleActionsable
    // The conditions property
    conditions MessageRulePredicatesable
    // The display name of the rule.
    displayName *string
    // The exceptions property
    exceptions MessageRulePredicatesable
    // Indicates whether the rule is in an error condition. Read-only.
    hasError *bool
    // Indicates whether the rule is enabled to be applied to messages.
    isEnabled *bool
    // Indicates if the rule is read-only and cannot be modified or deleted by the rules REST API.
    isReadOnly *bool
    // Indicates the order in which the rule is executed, among other rules.
    sequence *int32
}
// NewMessageRule instantiates a new MessageRule and sets the default values.
func NewMessageRule()(*MessageRule) {
    m := &MessageRule{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMessageRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMessageRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMessageRule(), nil
}
// GetActions gets the actions property value. The actions property
// returns a MessageRuleActionsable when successful
func (m *MessageRule) GetActions()(MessageRuleActionsable) {
    return m.actions
}
// GetConditions gets the conditions property value. The conditions property
// returns a MessageRulePredicatesable when successful
func (m *MessageRule) GetConditions()(MessageRulePredicatesable) {
    return m.conditions
}
// GetDisplayName gets the displayName property value. The display name of the rule.
// returns a *string when successful
func (m *MessageRule) GetDisplayName()(*string) {
    return m.displayName
}
// GetExceptions gets the exceptions property value. The exceptions property
// returns a MessageRulePredicatesable when successful
func (m *MessageRule) GetExceptions()(MessageRulePredicatesable) {
    return m.exceptions
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MessageRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["actions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMessageRuleActionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActions(val.(MessageRuleActionsable))
        }
        return nil
    }
    res["conditions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMessageRulePredicatesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConditions(val.(MessageRulePredicatesable))
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["exceptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMessageRulePredicatesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExceptions(val.(MessageRulePredicatesable))
        }
        return nil
    }
    res["hasError"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasError(val)
        }
        return nil
    }
    res["isEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabled(val)
        }
        return nil
    }
    res["isReadOnly"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsReadOnly(val)
        }
        return nil
    }
    res["sequence"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSequence(val)
        }
        return nil
    }
    return res
}
// GetHasError gets the hasError property value. Indicates whether the rule is in an error condition. Read-only.
// returns a *bool when successful
func (m *MessageRule) GetHasError()(*bool) {
    return m.hasError
}
// GetIsEnabled gets the isEnabled property value. Indicates whether the rule is enabled to be applied to messages.
// returns a *bool when successful
func (m *MessageRule) GetIsEnabled()(*bool) {
    return m.isEnabled
}
// GetIsReadOnly gets the isReadOnly property value. Indicates if the rule is read-only and cannot be modified or deleted by the rules REST API.
// returns a *bool when successful
func (m *MessageRule) GetIsReadOnly()(*bool) {
    return m.isReadOnly
}
// GetSequence gets the sequence property value. Indicates the order in which the rule is executed, among other rules.
// returns a *int32 when successful
func (m *MessageRule) GetSequence()(*int32) {
    return m.sequence
}
// Serialize serializes information the current object
func (m *MessageRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
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
        err = writer.WriteBoolValue("hasError", m.GetHasError())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isReadOnly", m.GetIsReadOnly())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("sequence", m.GetSequence())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActions sets the actions property value. The actions property
func (m *MessageRule) SetActions(value MessageRuleActionsable)() {
    m.actions = value
}
// SetConditions sets the conditions property value. The conditions property
func (m *MessageRule) SetConditions(value MessageRulePredicatesable)() {
    m.conditions = value
}
// SetDisplayName sets the displayName property value. The display name of the rule.
func (m *MessageRule) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetExceptions sets the exceptions property value. The exceptions property
func (m *MessageRule) SetExceptions(value MessageRulePredicatesable)() {
    m.exceptions = value
}
// SetHasError sets the hasError property value. Indicates whether the rule is in an error condition. Read-only.
func (m *MessageRule) SetHasError(value *bool)() {
    m.hasError = value
}
// SetIsEnabled sets the isEnabled property value. Indicates whether the rule is enabled to be applied to messages.
func (m *MessageRule) SetIsEnabled(value *bool)() {
    m.isEnabled = value
}
// SetIsReadOnly sets the isReadOnly property value. Indicates if the rule is read-only and cannot be modified or deleted by the rules REST API.
func (m *MessageRule) SetIsReadOnly(value *bool)() {
    m.isReadOnly = value
}
// SetSequence sets the sequence property value. Indicates the order in which the rule is executed, among other rules.
func (m *MessageRule) SetSequence(value *int32)() {
    m.sequence = value
}
type MessageRuleable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActions()(MessageRuleActionsable)
    GetConditions()(MessageRulePredicatesable)
    GetDisplayName()(*string)
    GetExceptions()(MessageRulePredicatesable)
    GetHasError()(*bool)
    GetIsEnabled()(*bool)
    GetIsReadOnly()(*bool)
    GetSequence()(*int32)
    SetActions(value MessageRuleActionsable)()
    SetConditions(value MessageRulePredicatesable)()
    SetDisplayName(value *string)()
    SetExceptions(value MessageRulePredicatesable)()
    SetHasError(value *bool)()
    SetIsEnabled(value *bool)()
    SetIsReadOnly(value *bool)()
    SetSequence(value *int32)()
}
