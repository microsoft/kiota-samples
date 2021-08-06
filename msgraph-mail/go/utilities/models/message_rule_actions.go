package models

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
func (m *MessageRuleActions) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error), error) {
    return nil, nil
}
func (m *MessageRuleActions) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
