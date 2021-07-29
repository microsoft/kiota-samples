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
func NewMessageRuleActions()(*MessageRuleActions) {
    return nil
}
func (m *MessageRuleActions) getAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
func (m *MessageRuleActions) getAssignCategories()([]*string) {
    return m.assignCategories
}
func (m *MessageRuleActions) getCopyToFolder()(*string) {
    return m.copyToFolder
}
func (m *MessageRuleActions) getDelete()(*bool) {
    return m.delete
}
func (m *MessageRuleActions) getForwardAsAttachmentTo()([]Recipient) {
    return m.forwardAsAttachmentTo
}
func (m *MessageRuleActions) getForwardTo()([]Recipient) {
    return m.forwardTo
}
func (m *MessageRuleActions) getMarkAsRead()(*bool) {
    return m.markAsRead
}
func (m *MessageRuleActions) getMarkImportance()(Importance) {
    return m.markImportance
}
func (m *MessageRuleActions) getMoveToFolder()(*string) {
    return m.moveToFolder
}
func (m *MessageRuleActions) getPermanentDelete()(*bool) {
    return m.permanentDelete
}
func (m *MessageRuleActions) getRedirectTo()([]Recipient) {
    return m.redirectTo
}
func (m *MessageRuleActions) getStopProcessingRules()(*bool) {
    return m.stopProcessingRules
}
func (m *MessageRuleActions) getFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error), error) {
    return nil
}
func (m *MessageRuleActions) serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    return nil
}
func (m *MessageRuleActions) setAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *MessageRuleActions) setAssignCategories(value []*string)() {
    m.assignCategories = value
}
func (m *MessageRuleActions) setCopyToFolder(value *string)() {
    m.copyToFolder = value
}
func (m *MessageRuleActions) setDelete(value *bool)() {
    m.delete = value
}
func (m *MessageRuleActions) setForwardAsAttachmentTo(value []Recipient)() {
    m.forwardAsAttachmentTo = value
}
func (m *MessageRuleActions) setForwardTo(value []Recipient)() {
    m.forwardTo = value
}
func (m *MessageRuleActions) setMarkAsRead(value *bool)() {
    m.markAsRead = value
}
func (m *MessageRuleActions) setMarkImportance(value Importance)() {
    m.markImportance = value
}
func (m *MessageRuleActions) setMoveToFolder(value *string)() {
    m.moveToFolder = value
}
func (m *MessageRuleActions) setPermanentDelete(value *bool)() {
    m.permanentDelete = value
}
func (m *MessageRuleActions) setRedirectTo(value []Recipient)() {
    m.redirectTo = value
}
func (m *MessageRuleActions) setStopProcessingRules(value *bool)() {
    m.stopProcessingRules = value
}
