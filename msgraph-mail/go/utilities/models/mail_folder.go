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
    return nil
}
func (m *MailFolder) getChildFolderCount()(*int32) {
    return m.childFolderCount
}
func (m *MailFolder) getChildFolders()([]MailFolder) {
    return m.childFolders
}
func (m *MailFolder) getDisplayName()(*string) {
    return m.displayName
}
func (m *MailFolder) getIsHidden()(*bool) {
    return m.isHidden
}
func (m *MailFolder) getMessageRules()([]MessageRule) {
    return m.messageRules
}
func (m *MailFolder) getMessages()([]Message) {
    return m.messages
}
func (m *MailFolder) getMultiValueExtendedProperties()([]MultiValueLegacyExtendedProperty) {
    return m.multiValueExtendedProperties
}
func (m *MailFolder) getParentFolderId()(*string) {
    return m.parentFolderId
}
func (m *MailFolder) getSingleValueExtendedProperties()([]SingleValueLegacyExtendedProperty) {
    return m.singleValueExtendedProperties
}
func (m *MailFolder) getTotalItemCount()(*int32) {
    return m.totalItemCount
}
func (m *MailFolder) getUnreadItemCount()(*int32) {
    return m.unreadItemCount
}
func (m *MailFolder) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error), error) {
    return nil, nil
}
func (m *MailFolder) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    return nil
}
func (m *MailFolder) setChildFolderCount(value *int32)() {
    m.childFolderCount = value
}
func (m *MailFolder) setChildFolders(value []MailFolder)() {
    m.childFolders = value
}
func (m *MailFolder) setDisplayName(value *string)() {
    m.displayName = value
}
func (m *MailFolder) setIsHidden(value *bool)() {
    m.isHidden = value
}
func (m *MailFolder) setMessageRules(value []MessageRule)() {
    m.messageRules = value
}
func (m *MailFolder) setMessages(value []Message)() {
    m.messages = value
}
func (m *MailFolder) setMultiValueExtendedProperties(value []MultiValueLegacyExtendedProperty)() {
    m.multiValueExtendedProperties = value
}
func (m *MailFolder) setParentFolderId(value *string)() {
    m.parentFolderId = value
}
func (m *MailFolder) setSingleValueExtendedProperties(value []SingleValueLegacyExtendedProperty)() {
    m.singleValueExtendedProperties = value
}
func (m *MailFolder) setTotalItemCount(value *int32)() {
    m.totalItemCount = value
}
func (m *MailFolder) setUnreadItemCount(value *int32)() {
    m.unreadItemCount = value
}
