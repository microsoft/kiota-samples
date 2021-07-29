package models

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type MessageRuleActions struct {
    additionalData map[string]interface{};
    assignCategories []*string;
    copyToFolder *string;
    delete *bool;
    forwardAsAttachmentTo []Recipient;
    forwardTo []Recipient;
    markAsRead *bool;
    markImportance Importance;
    moveToFolder *string;
    permanentDelete *bool;
    redirectTo []Recipient;
    stopProcessingRules *bool;
}
func (m *MessageRuleActions) getAdditionalData()(map[string]interface{}) {
    return nil
}
func (m *MessageRuleActions) getAssignCategories()([]*string) {
    return nil
}
func (m *MessageRuleActions) getCopyToFolder()(*string) {
    return nil
}
func (m *MessageRuleActions) getDelete()(*bool) {
    return nil
}
func (m *MessageRuleActions) getForwardAsAttachmentTo()([]Recipient) {
    return nil
}
func (m *MessageRuleActions) getForwardTo()([]Recipient) {
    return nil
}
func (m *MessageRuleActions) getMarkAsRead()(*bool) {
    return nil
}
func (m *MessageRuleActions) getMarkImportance()(Importance) {
    return nil
}
func (m *MessageRuleActions) getMoveToFolder()(*string) {
    return nil
}
func (m *MessageRuleActions) getPermanentDelete()(*bool) {
    return nil
}
func (m *MessageRuleActions) getRedirectTo()([]Recipient) {
    return nil
}
func (m *MessageRuleActions) getStopProcessingRules()(*bool) {
    return nil
}
func (m *MessageRuleActions) getFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error), error) {
    return nil
}
func (m *MessageRuleActions) serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    return nil
}
func (m *MessageRuleActions) setAdditionalData(value map[string]interface{})() {
    return nil
}
func (m *MessageRuleActions) setAssignCategories(value []*string)() {
    return nil
}
func (m *MessageRuleActions) setCopyToFolder(value *string)() {
    return nil
}
func (m *MessageRuleActions) setDelete(value *bool)() {
    return nil
}
func (m *MessageRuleActions) setForwardAsAttachmentTo(value []Recipient)() {
    return nil
}
func (m *MessageRuleActions) setForwardTo(value []Recipient)() {
    return nil
}
func (m *MessageRuleActions) setMarkAsRead(value *bool)() {
    return nil
}
func (m *MessageRuleActions) setMarkImportance(value Importance)() {
    return nil
}
func (m *MessageRuleActions) setMoveToFolder(value *string)() {
    return nil
}
func (m *MessageRuleActions) setPermanentDelete(value *bool)() {
    return nil
}
func (m *MessageRuleActions) setRedirectTo(value []Recipient)() {
    return nil
}
func (m *MessageRuleActions) setStopProcessingRules(value *bool)() {
    return nil
}
