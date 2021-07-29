package models

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type MailFolder struct {
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
func (m *MailFolder) getChildFolderCount()(*int32) {
    return nil
}
func (m *MailFolder) getChildFolders()([]MailFolder) {
    return nil
}
func (m *MailFolder) getDisplayName()(*string) {
    return nil
}
func (m *MailFolder) getIsHidden()(*bool) {
    return nil
}
func (m *MailFolder) getMessageRules()([]MessageRule) {
    return nil
}
func (m *MailFolder) getMessages()([]Message) {
    return nil
}
func (m *MailFolder) getMultiValueExtendedProperties()([]MultiValueLegacyExtendedProperty) {
    return nil
}
func (m *MailFolder) getParentFolderId()(*string) {
    return nil
}
func (m *MailFolder) getSingleValueExtendedProperties()([]SingleValueLegacyExtendedProperty) {
    return nil
}
func (m *MailFolder) getTotalItemCount()(*int32) {
    return nil
}
func (m *MailFolder) getUnreadItemCount()(*int32) {
    return nil
}
func (m *MailFolder) getFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error), error) {
    return nil
}
func (m *MailFolder) serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    return nil
}
func (m *MailFolder) setChildFolderCount(value *int32)() {
    return nil
}
func (m *MailFolder) setChildFolders(value []MailFolder)() {
    return nil
}
func (m *MailFolder) setDisplayName(value *string)() {
    return nil
}
func (m *MailFolder) setIsHidden(value *bool)() {
    return nil
}
func (m *MailFolder) setMessageRules(value []MessageRule)() {
    return nil
}
func (m *MailFolder) setMessages(value []Message)() {
    return nil
}
func (m *MailFolder) setMultiValueExtendedProperties(value []MultiValueLegacyExtendedProperty)() {
    return nil
}
func (m *MailFolder) setParentFolderId(value *string)() {
    return nil
}
func (m *MailFolder) setSingleValueExtendedProperties(value []SingleValueLegacyExtendedProperty)() {
    return nil
}
func (m *MailFolder) setTotalItemCount(value *int32)() {
    return nil
}
func (m *MailFolder) setUnreadItemCount(value *int32)() {
    return nil
}
