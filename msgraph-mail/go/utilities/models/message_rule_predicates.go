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
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *MessageRulePredicates) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
func (m *MessageRulePredicates) GetBodyContains()([]string) {
    return m.bodyContains
}
func (m *MessageRulePredicates) GetBodyOrSubjectContains()([]string) {
    return m.bodyOrSubjectContains
}
func (m *MessageRulePredicates) GetCategories()([]string) {
    return m.categories
}
func (m *MessageRulePredicates) GetFromAddresses()([]Recipient) {
    return m.fromAddresses
}
func (m *MessageRulePredicates) GetHasAttachments()(*bool) {
    return m.hasAttachments
}
func (m *MessageRulePredicates) GetHeaderContains()([]string) {
    return m.headerContains
}
func (m *MessageRulePredicates) GetImportance()(*Importance) {
    return m.importance
}
func (m *MessageRulePredicates) GetIsApprovalRequest()(*bool) {
    return m.isApprovalRequest
}
func (m *MessageRulePredicates) GetIsAutomaticForward()(*bool) {
    return m.isAutomaticForward
}
func (m *MessageRulePredicates) GetIsAutomaticReply()(*bool) {
    return m.isAutomaticReply
}
func (m *MessageRulePredicates) GetIsEncrypted()(*bool) {
    return m.isEncrypted
}
func (m *MessageRulePredicates) GetIsMeetingRequest()(*bool) {
    return m.isMeetingRequest
}
func (m *MessageRulePredicates) GetIsMeetingResponse()(*bool) {
    return m.isMeetingResponse
}
func (m *MessageRulePredicates) GetIsNonDeliveryReport()(*bool) {
    return m.isNonDeliveryReport
}
func (m *MessageRulePredicates) GetIsPermissionControlled()(*bool) {
    return m.isPermissionControlled
}
func (m *MessageRulePredicates) GetIsReadReceipt()(*bool) {
    return m.isReadReceipt
}
func (m *MessageRulePredicates) GetIsSigned()(*bool) {
    return m.isSigned
}
func (m *MessageRulePredicates) GetIsVoicemail()(*bool) {
    return m.isVoicemail
}
func (m *MessageRulePredicates) GetMessageActionFlag()(*MessageActionFlag) {
    return m.messageActionFlag
}
func (m *MessageRulePredicates) GetNotSentToMe()(*bool) {
    return m.notSentToMe
}
func (m *MessageRulePredicates) GetRecipientContains()([]string) {
    return m.recipientContains
}
func (m *MessageRulePredicates) GetSenderContains()([]string) {
    return m.senderContains
}
func (m *MessageRulePredicates) GetSensitivity()(*Sensitivity) {
    return m.sensitivity
}
func (m *MessageRulePredicates) GetSentCcMe()(*bool) {
    return m.sentCcMe
}
func (m *MessageRulePredicates) GetSentOnlyToMe()(*bool) {
    return m.sentOnlyToMe
}
func (m *MessageRulePredicates) GetSentToAddresses()([]Recipient) {
    return m.sentToAddresses
}
func (m *MessageRulePredicates) GetSentToMe()(*bool) {
    return m.sentToMe
}
func (m *MessageRulePredicates) GetSentToOrCcMe()(*bool) {
    return m.sentToOrCcMe
}
func (m *MessageRulePredicates) GetSubjectContains()([]string) {
    return m.subjectContains
}
func (m *MessageRulePredicates) GetWithinSizeRange()(*SizeRange) {
    return m.withinSizeRange
}
func (m *MessageRulePredicates) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error), error) {
    return nil, nil
}
func (m *MessageRulePredicates) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    return nil
}
func (m *MessageRulePredicates) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *MessageRulePredicates) SetBodyContains(value []string)() {
    m.bodyContains = value
}
func (m *MessageRulePredicates) SetBodyOrSubjectContains(value []string)() {
    m.bodyOrSubjectContains = value
}
func (m *MessageRulePredicates) SetCategories(value []string)() {
    m.categories = value
}
func (m *MessageRulePredicates) SetFromAddresses(value []Recipient)() {
    m.fromAddresses = value
}
func (m *MessageRulePredicates) SetHasAttachments(value *bool)() {
    m.hasAttachments = value
}
func (m *MessageRulePredicates) SetHeaderContains(value []string)() {
    m.headerContains = value
}
func (m *MessageRulePredicates) SetImportance(value *Importance)() {
    m.importance = value
}
func (m *MessageRulePredicates) SetIsApprovalRequest(value *bool)() {
    m.isApprovalRequest = value
}
func (m *MessageRulePredicates) SetIsAutomaticForward(value *bool)() {
    m.isAutomaticForward = value
}
func (m *MessageRulePredicates) SetIsAutomaticReply(value *bool)() {
    m.isAutomaticReply = value
}
func (m *MessageRulePredicates) SetIsEncrypted(value *bool)() {
    m.isEncrypted = value
}
func (m *MessageRulePredicates) SetIsMeetingRequest(value *bool)() {
    m.isMeetingRequest = value
}
func (m *MessageRulePredicates) SetIsMeetingResponse(value *bool)() {
    m.isMeetingResponse = value
}
func (m *MessageRulePredicates) SetIsNonDeliveryReport(value *bool)() {
    m.isNonDeliveryReport = value
}
func (m *MessageRulePredicates) SetIsPermissionControlled(value *bool)() {
    m.isPermissionControlled = value
}
func (m *MessageRulePredicates) SetIsReadReceipt(value *bool)() {
    m.isReadReceipt = value
}
func (m *MessageRulePredicates) SetIsSigned(value *bool)() {
    m.isSigned = value
}
func (m *MessageRulePredicates) SetIsVoicemail(value *bool)() {
    m.isVoicemail = value
}
func (m *MessageRulePredicates) SetMessageActionFlag(value *MessageActionFlag)() {
    m.messageActionFlag = value
}
func (m *MessageRulePredicates) SetNotSentToMe(value *bool)() {
    m.notSentToMe = value
}
func (m *MessageRulePredicates) SetRecipientContains(value []string)() {
    m.recipientContains = value
}
func (m *MessageRulePredicates) SetSenderContains(value []string)() {
    m.senderContains = value
}
func (m *MessageRulePredicates) SetSensitivity(value *Sensitivity)() {
    m.sensitivity = value
}
func (m *MessageRulePredicates) SetSentCcMe(value *bool)() {
    m.sentCcMe = value
}
func (m *MessageRulePredicates) SetSentOnlyToMe(value *bool)() {
    m.sentOnlyToMe = value
}
func (m *MessageRulePredicates) SetSentToAddresses(value []Recipient)() {
    m.sentToAddresses = value
}
func (m *MessageRulePredicates) SetSentToMe(value *bool)() {
    m.sentToMe = value
}
func (m *MessageRulePredicates) SetSentToOrCcMe(value *bool)() {
    m.sentToOrCcMe = value
}
func (m *MessageRulePredicates) SetSubjectContains(value []string)() {
    m.subjectContains = value
}
func (m *MessageRulePredicates) SetWithinSizeRange(value *SizeRange)() {
    m.withinSizeRange = value
}
