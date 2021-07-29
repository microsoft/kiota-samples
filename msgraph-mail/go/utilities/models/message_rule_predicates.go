package models

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type MessageRulePredicates struct {
    additionalData map[string]interface{};
    bodyContains []*string;
    bodyOrSubjectContains []*string;
    categories []*string;
    fromAddresses []Recipient;
    hasAttachments *bool;
    headerContains []*string;
    importance Importance;
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
    messageActionFlag MessageActionFlag;
    notSentToMe *bool;
    recipientContains []*string;
    senderContains []*string;
    sensitivity Sensitivity;
    sentCcMe *bool;
    sentOnlyToMe *bool;
    sentToAddresses []Recipient;
    sentToMe *bool;
    sentToOrCcMe *bool;
    subjectContains []*string;
    withinSizeRange SizeRange;
}
func (m *MessageRulePredicates) getAdditionalData()(map[string]interface{}) {
    return nil
}
func (m *MessageRulePredicates) getBodyContains()([]*string) {
    return nil
}
func (m *MessageRulePredicates) getBodyOrSubjectContains()([]*string) {
    return nil
}
func (m *MessageRulePredicates) getCategories()([]*string) {
    return nil
}
func (m *MessageRulePredicates) getFromAddresses()([]Recipient) {
    return nil
}
func (m *MessageRulePredicates) getHasAttachments()(*bool) {
    return nil
}
func (m *MessageRulePredicates) getHeaderContains()([]*string) {
    return nil
}
func (m *MessageRulePredicates) getImportance()(Importance) {
    return nil
}
func (m *MessageRulePredicates) getIsApprovalRequest()(*bool) {
    return nil
}
func (m *MessageRulePredicates) getIsAutomaticForward()(*bool) {
    return nil
}
func (m *MessageRulePredicates) getIsAutomaticReply()(*bool) {
    return nil
}
func (m *MessageRulePredicates) getIsEncrypted()(*bool) {
    return nil
}
func (m *MessageRulePredicates) getIsMeetingRequest()(*bool) {
    return nil
}
func (m *MessageRulePredicates) getIsMeetingResponse()(*bool) {
    return nil
}
func (m *MessageRulePredicates) getIsNonDeliveryReport()(*bool) {
    return nil
}
func (m *MessageRulePredicates) getIsPermissionControlled()(*bool) {
    return nil
}
func (m *MessageRulePredicates) getIsReadReceipt()(*bool) {
    return nil
}
func (m *MessageRulePredicates) getIsSigned()(*bool) {
    return nil
}
func (m *MessageRulePredicates) getIsVoicemail()(*bool) {
    return nil
}
func (m *MessageRulePredicates) getMessageActionFlag()(MessageActionFlag) {
    return nil
}
func (m *MessageRulePredicates) getNotSentToMe()(*bool) {
    return nil
}
func (m *MessageRulePredicates) getRecipientContains()([]*string) {
    return nil
}
func (m *MessageRulePredicates) getSenderContains()([]*string) {
    return nil
}
func (m *MessageRulePredicates) getSensitivity()(Sensitivity) {
    return nil
}
func (m *MessageRulePredicates) getSentCcMe()(*bool) {
    return nil
}
func (m *MessageRulePredicates) getSentOnlyToMe()(*bool) {
    return nil
}
func (m *MessageRulePredicates) getSentToAddresses()([]Recipient) {
    return nil
}
func (m *MessageRulePredicates) getSentToMe()(*bool) {
    return nil
}
func (m *MessageRulePredicates) getSentToOrCcMe()(*bool) {
    return nil
}
func (m *MessageRulePredicates) getSubjectContains()([]*string) {
    return nil
}
func (m *MessageRulePredicates) getWithinSizeRange()(SizeRange) {
    return nil
}
func (m *MessageRulePredicates) getFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error), error) {
    return nil
}
func (m *MessageRulePredicates) serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    return nil
}
func (m *MessageRulePredicates) setAdditionalData(value map[string]interface{})() {
    return nil
}
func (m *MessageRulePredicates) setBodyContains(value []*string)() {
    return nil
}
func (m *MessageRulePredicates) setBodyOrSubjectContains(value []*string)() {
    return nil
}
func (m *MessageRulePredicates) setCategories(value []*string)() {
    return nil
}
func (m *MessageRulePredicates) setFromAddresses(value []Recipient)() {
    return nil
}
func (m *MessageRulePredicates) setHasAttachments(value *bool)() {
    return nil
}
func (m *MessageRulePredicates) setHeaderContains(value []*string)() {
    return nil
}
func (m *MessageRulePredicates) setImportance(value Importance)() {
    return nil
}
func (m *MessageRulePredicates) setIsApprovalRequest(value *bool)() {
    return nil
}
func (m *MessageRulePredicates) setIsAutomaticForward(value *bool)() {
    return nil
}
func (m *MessageRulePredicates) setIsAutomaticReply(value *bool)() {
    return nil
}
func (m *MessageRulePredicates) setIsEncrypted(value *bool)() {
    return nil
}
func (m *MessageRulePredicates) setIsMeetingRequest(value *bool)() {
    return nil
}
func (m *MessageRulePredicates) setIsMeetingResponse(value *bool)() {
    return nil
}
func (m *MessageRulePredicates) setIsNonDeliveryReport(value *bool)() {
    return nil
}
func (m *MessageRulePredicates) setIsPermissionControlled(value *bool)() {
    return nil
}
func (m *MessageRulePredicates) setIsReadReceipt(value *bool)() {
    return nil
}
func (m *MessageRulePredicates) setIsSigned(value *bool)() {
    return nil
}
func (m *MessageRulePredicates) setIsVoicemail(value *bool)() {
    return nil
}
func (m *MessageRulePredicates) setMessageActionFlag(value MessageActionFlag)() {
    return nil
}
func (m *MessageRulePredicates) setNotSentToMe(value *bool)() {
    return nil
}
func (m *MessageRulePredicates) setRecipientContains(value []*string)() {
    return nil
}
func (m *MessageRulePredicates) setSenderContains(value []*string)() {
    return nil
}
func (m *MessageRulePredicates) setSensitivity(value Sensitivity)() {
    return nil
}
func (m *MessageRulePredicates) setSentCcMe(value *bool)() {
    return nil
}
func (m *MessageRulePredicates) setSentOnlyToMe(value *bool)() {
    return nil
}
func (m *MessageRulePredicates) setSentToAddresses(value []Recipient)() {
    return nil
}
func (m *MessageRulePredicates) setSentToMe(value *bool)() {
    return nil
}
func (m *MessageRulePredicates) setSentToOrCcMe(value *bool)() {
    return nil
}
func (m *MessageRulePredicates) setSubjectContains(value []*string)() {
    return nil
}
func (m *MessageRulePredicates) setWithinSizeRange(value SizeRange)() {
    return nil
}
