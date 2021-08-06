package models

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type MessageRulePredicates struct {
    additionalData map[string]interface{};
    bodyContains []string;
    bodyOrSubjectContains []string;
    categories []string;
    fromAddresses []Recipient;
    hasAttachments *bool;
    headerContains []string;
    importance *Importance;
    isApprovalRequest *bool;
    isAutomaticForward *bool;
    isAutomaticReply *bool;
    isEncrypted *bool;
    isMeetingRequest *bool;
    isMeetingResponse *bool;
    isNonDeliveryReport *bool;
    isPermissionControlled *bool;
    isReadReceipt *bool;
    isSigned *bool;
    isVoicemail *bool;
    messageActionFlag *MessageActionFlag;
    notSentToMe *bool;
    recipientContains []string;
    senderContains []string;
    sensitivity *Sensitivity;
    sentCcMe *bool;
    sentOnlyToMe *bool;
    sentToAddresses []Recipient;
    sentToMe *bool;
    sentToOrCcMe *bool;
    subjectContains []string;
    withinSizeRange *SizeRange;
}
func NewMessageRulePredicates()(*MessageRulePredicates) {
    m := &MessageRulePredicates{
    }
    m.setAdditionalData(make(map[string]interface{}));
    return m
}
func (m *MessageRulePredicates) getAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
func (m *MessageRulePredicates) getBodyContains()([]string) {
    return m.bodyContains
}
func (m *MessageRulePredicates) getBodyOrSubjectContains()([]string) {
    return m.bodyOrSubjectContains
}
func (m *MessageRulePredicates) getCategories()([]string) {
    return m.categories
}
func (m *MessageRulePredicates) getFromAddresses()([]Recipient) {
    return m.fromAddresses
}
func (m *MessageRulePredicates) getHasAttachments()(*bool) {
    return m.hasAttachments
}
func (m *MessageRulePredicates) getHeaderContains()([]string) {
    return m.headerContains
}
func (m *MessageRulePredicates) getImportance()(*Importance) {
    return m.importance
}
func (m *MessageRulePredicates) getIsApprovalRequest()(*bool) {
    return m.isApprovalRequest
}
func (m *MessageRulePredicates) getIsAutomaticForward()(*bool) {
    return m.isAutomaticForward
}
func (m *MessageRulePredicates) getIsAutomaticReply()(*bool) {
    return m.isAutomaticReply
}
func (m *MessageRulePredicates) getIsEncrypted()(*bool) {
    return m.isEncrypted
}
func (m *MessageRulePredicates) getIsMeetingRequest()(*bool) {
    return m.isMeetingRequest
}
func (m *MessageRulePredicates) getIsMeetingResponse()(*bool) {
    return m.isMeetingResponse
}
func (m *MessageRulePredicates) getIsNonDeliveryReport()(*bool) {
    return m.isNonDeliveryReport
}
func (m *MessageRulePredicates) getIsPermissionControlled()(*bool) {
    return m.isPermissionControlled
}
func (m *MessageRulePredicates) getIsReadReceipt()(*bool) {
    return m.isReadReceipt
}
func (m *MessageRulePredicates) getIsSigned()(*bool) {
    return m.isSigned
}
func (m *MessageRulePredicates) getIsVoicemail()(*bool) {
    return m.isVoicemail
}
func (m *MessageRulePredicates) getMessageActionFlag()(*MessageActionFlag) {
    return m.messageActionFlag
}
func (m *MessageRulePredicates) getNotSentToMe()(*bool) {
    return m.notSentToMe
}
func (m *MessageRulePredicates) getRecipientContains()([]string) {
    return m.recipientContains
}
func (m *MessageRulePredicates) getSenderContains()([]string) {
    return m.senderContains
}
func (m *MessageRulePredicates) getSensitivity()(*Sensitivity) {
    return m.sensitivity
}
func (m *MessageRulePredicates) getSentCcMe()(*bool) {
    return m.sentCcMe
}
func (m *MessageRulePredicates) getSentOnlyToMe()(*bool) {
    return m.sentOnlyToMe
}
func (m *MessageRulePredicates) getSentToAddresses()([]Recipient) {
    return m.sentToAddresses
}
func (m *MessageRulePredicates) getSentToMe()(*bool) {
    return m.sentToMe
}
func (m *MessageRulePredicates) getSentToOrCcMe()(*bool) {
    return m.sentToOrCcMe
}
func (m *MessageRulePredicates) getSubjectContains()([]string) {
    return m.subjectContains
}
func (m *MessageRulePredicates) getWithinSizeRange()(*SizeRange) {
    return m.withinSizeRange
}
func (m *MessageRulePredicates) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error), error) {
    return nil, nil
}
func (m *MessageRulePredicates) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    return nil
}
func (m *MessageRulePredicates) setAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *MessageRulePredicates) setBodyContains(value []string)() {
    m.bodyContains = value
}
func (m *MessageRulePredicates) setBodyOrSubjectContains(value []string)() {
    m.bodyOrSubjectContains = value
}
func (m *MessageRulePredicates) setCategories(value []string)() {
    m.categories = value
}
func (m *MessageRulePredicates) setFromAddresses(value []Recipient)() {
    m.fromAddresses = value
}
func (m *MessageRulePredicates) setHasAttachments(value *bool)() {
    m.hasAttachments = value
}
func (m *MessageRulePredicates) setHeaderContains(value []string)() {
    m.headerContains = value
}
func (m *MessageRulePredicates) setImportance(value *Importance)() {
    m.importance = value
}
func (m *MessageRulePredicates) setIsApprovalRequest(value *bool)() {
    m.isApprovalRequest = value
}
func (m *MessageRulePredicates) setIsAutomaticForward(value *bool)() {
    m.isAutomaticForward = value
}
func (m *MessageRulePredicates) setIsAutomaticReply(value *bool)() {
    m.isAutomaticReply = value
}
func (m *MessageRulePredicates) setIsEncrypted(value *bool)() {
    m.isEncrypted = value
}
func (m *MessageRulePredicates) setIsMeetingRequest(value *bool)() {
    m.isMeetingRequest = value
}
func (m *MessageRulePredicates) setIsMeetingResponse(value *bool)() {
    m.isMeetingResponse = value
}
func (m *MessageRulePredicates) setIsNonDeliveryReport(value *bool)() {
    m.isNonDeliveryReport = value
}
func (m *MessageRulePredicates) setIsPermissionControlled(value *bool)() {
    m.isPermissionControlled = value
}
func (m *MessageRulePredicates) setIsReadReceipt(value *bool)() {
    m.isReadReceipt = value
}
func (m *MessageRulePredicates) setIsSigned(value *bool)() {
    m.isSigned = value
}
func (m *MessageRulePredicates) setIsVoicemail(value *bool)() {
    m.isVoicemail = value
}
func (m *MessageRulePredicates) setMessageActionFlag(value *MessageActionFlag)() {
    m.messageActionFlag = value
}
func (m *MessageRulePredicates) setNotSentToMe(value *bool)() {
    m.notSentToMe = value
}
func (m *MessageRulePredicates) setRecipientContains(value []string)() {
    m.recipientContains = value
}
func (m *MessageRulePredicates) setSenderContains(value []string)() {
    m.senderContains = value
}
func (m *MessageRulePredicates) setSensitivity(value *Sensitivity)() {
    m.sensitivity = value
}
func (m *MessageRulePredicates) setSentCcMe(value *bool)() {
    m.sentCcMe = value
}
func (m *MessageRulePredicates) setSentOnlyToMe(value *bool)() {
    m.sentOnlyToMe = value
}
func (m *MessageRulePredicates) setSentToAddresses(value []Recipient)() {
    m.sentToAddresses = value
}
func (m *MessageRulePredicates) setSentToMe(value *bool)() {
    m.sentToMe = value
}
func (m *MessageRulePredicates) setSentToOrCcMe(value *bool)() {
    m.sentToOrCcMe = value
}
func (m *MessageRulePredicates) setSubjectContains(value []string)() {
    m.subjectContains = value
}
func (m *MessageRulePredicates) setWithinSizeRange(value *SizeRange)() {
    m.withinSizeRange = value
}
