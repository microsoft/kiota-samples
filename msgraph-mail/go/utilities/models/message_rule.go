package models

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type MessageRule struct {
    actions MessageRuleActions;
    conditions MessageRulePredicates;
    displayName *string;
    exceptions MessageRulePredicates;
    hasError *bool;
    isEnabled *bool;
    isReadOnly *bool;
    sequence *int32;
}
func (m *MessageRule) getActions()(MessageRuleActions) {
    return nil
}
func (m *MessageRule) getConditions()(MessageRulePredicates) {
    return nil
}
func (m *MessageRule) getDisplayName()(*string) {
    return nil
}
func (m *MessageRule) getExceptions()(MessageRulePredicates) {
    return nil
}
func (m *MessageRule) getHasError()(*bool) {
    return nil
}
func (m *MessageRule) getIsEnabled()(*bool) {
    return nil
}
func (m *MessageRule) getIsReadOnly()(*bool) {
    return nil
}
func (m *MessageRule) getSequence()(*int32) {
    return nil
}
func (m *MessageRule) getFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error), error) {
    return nil
}
func (m *MessageRule) serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    return nil
}
func (m *MessageRule) setActions(value MessageRuleActions)() {
    return nil
}
func (m *MessageRule) setConditions(value MessageRulePredicates)() {
    return nil
}
func (m *MessageRule) setDisplayName(value *string)() {
    return nil
}
func (m *MessageRule) setExceptions(value MessageRulePredicates)() {
    return nil
}
func (m *MessageRule) setHasError(value *bool)() {
    return nil
}
func (m *MessageRule) setIsEnabled(value *bool)() {
    return nil
}
func (m *MessageRule) setIsReadOnly(value *bool)() {
    return nil
}
func (m *MessageRule) setSequence(value *int32)() {
    return nil
}
