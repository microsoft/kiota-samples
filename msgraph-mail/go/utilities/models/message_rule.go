package models

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type MessageRule struct {
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
    return nil
}
func (m *MessageRule) getActions()(*MessageRuleActions) {
    return m.actions
}
func (m *MessageRule) getConditions()(*MessageRulePredicates) {
    return m.conditions
}
func (m *MessageRule) getDisplayName()(*string) {
    return m.displayName
}
func (m *MessageRule) getExceptions()(*MessageRulePredicates) {
    return m.exceptions
}
func (m *MessageRule) getHasError()(*bool) {
    return m.hasError
}
func (m *MessageRule) getIsEnabled()(*bool) {
    return m.isEnabled
}
func (m *MessageRule) getIsReadOnly()(*bool) {
    return m.isReadOnly
}
func (m *MessageRule) getSequence()(*int32) {
    return m.sequence
}
func (m *MessageRule) getFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error), error) {
    return nil, nil
}
func (m *MessageRule) serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    return nil
}
func (m *MessageRule) setActions(value *MessageRuleActions)() {
    m.actions = value
}
func (m *MessageRule) setConditions(value *MessageRulePredicates)() {
    m.conditions = value
}
func (m *MessageRule) setDisplayName(value *string)() {
    m.displayName = value
}
func (m *MessageRule) setExceptions(value *MessageRulePredicates)() {
    m.exceptions = value
}
func (m *MessageRule) setHasError(value *bool)() {
    m.hasError = value
}
func (m *MessageRule) setIsEnabled(value *bool)() {
    m.isEnabled = value
}
func (m *MessageRule) setIsReadOnly(value *bool)() {
    m.isReadOnly = value
}
func (m *MessageRule) setSequence(value *int32)() {
    m.sequence = value
}
