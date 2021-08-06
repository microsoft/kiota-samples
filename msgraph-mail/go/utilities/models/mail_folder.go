package models

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
func (m *MailFolder) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error), error) {
    return nil, nil
}
func (m *MailFolder) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
