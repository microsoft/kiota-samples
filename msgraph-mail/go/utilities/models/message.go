package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Message struct {
    attachments []Attachment;
    bccRecipients []Recipient;
    body ItemBody;
    bodyPreview *string;
    ccRecipients []Recipient;
    conversationId *string;
    conversationIndex *string;
    extensions []Extension;
    flag FollowupFlag;
    from Recipient;
    hasAttachments *bool;
    importance Importance;
    inferenceClassification InferenceClassificationType;
    internetMessageHeaders []InternetMessageHeader;
    internetMessageId *string;
    isDeliveryReceiptRequested *bool;
    isDraft *bool;
    isRead *bool;
    isReadReceiptRequested *bool;
    multiValueExtendedProperties []MultiValueLegacyExtendedProperty;
    parentFolderId *string;
    receivedDateTime i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    replyTo []Recipient;
    sender Recipient;
    sentDateTime i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    singleValueExtendedProperties []SingleValueLegacyExtendedProperty;
    subject *string;
    toRecipients []Recipient;
    uniqueBody ItemBody;
    webLink *string;
}
func (m *Message) getAttachments()([]Attachment) {
    return nil
}
func (m *Message) getBccRecipients()([]Recipient) {
    return nil
}
func (m *Message) getBody()(ItemBody) {
    return nil
}
func (m *Message) getBodyPreview()(*string) {
    return nil
}
func (m *Message) getCcRecipients()([]Recipient) {
    return nil
}
func (m *Message) getConversationId()(*string) {
    return nil
}
func (m *Message) getConversationIndex()(*string) {
    return nil
}
func (m *Message) getExtensions()([]Extension) {
    return nil
}
func (m *Message) getFlag()(FollowupFlag) {
    return nil
}
func (m *Message) getFrom()(Recipient) {
    return nil
}
func (m *Message) getHasAttachments()(*bool) {
    return nil
}
func (m *Message) getImportance()(Importance) {
    return nil
}
func (m *Message) getInferenceClassification()(InferenceClassificationType) {
    return nil
}
func (m *Message) getInternetMessageHeaders()([]InternetMessageHeader) {
    return nil
}
func (m *Message) getInternetMessageId()(*string) {
    return nil
}
func (m *Message) getIsDeliveryReceiptRequested()(*bool) {
    return nil
}
func (m *Message) getIsDraft()(*bool) {
    return nil
}
func (m *Message) getIsRead()(*bool) {
    return nil
}
func (m *Message) getIsReadReceiptRequested()(*bool) {
    return nil
}
func (m *Message) getMultiValueExtendedProperties()([]MultiValueLegacyExtendedProperty) {
    return nil
}
func (m *Message) getParentFolderId()(*string) {
    return nil
}
func (m *Message) getReceivedDateTime()(i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return nil
}
func (m *Message) getReplyTo()([]Recipient) {
    return nil
}
func (m *Message) getSender()(Recipient) {
    return nil
}
func (m *Message) getSentDateTime()(i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return nil
}
func (m *Message) getSingleValueExtendedProperties()([]SingleValueLegacyExtendedProperty) {
    return nil
}
func (m *Message) getSubject()(*string) {
    return nil
}
func (m *Message) getToRecipients()([]Recipient) {
    return nil
}
func (m *Message) getUniqueBody()(ItemBody) {
    return nil
}
func (m *Message) getWebLink()(*string) {
    return nil
}
func (m *Message) getFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error), error) {
    return nil
}
func (m *Message) serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    return nil
}
func (m *Message) setAttachments(value []Attachment)() {
    return nil
}
func (m *Message) setBccRecipients(value []Recipient)() {
    return nil
}
func (m *Message) setBody(value ItemBody)() {
    return nil
}
func (m *Message) setBodyPreview(value *string)() {
    return nil
}
func (m *Message) setCcRecipients(value []Recipient)() {
    return nil
}
func (m *Message) setConversationId(value *string)() {
    return nil
}
func (m *Message) setConversationIndex(value *string)() {
    return nil
}
func (m *Message) setExtensions(value []Extension)() {
    return nil
}
func (m *Message) setFlag(value FollowupFlag)() {
    return nil
}
func (m *Message) setFrom(value Recipient)() {
    return nil
}
func (m *Message) setHasAttachments(value *bool)() {
    return nil
}
func (m *Message) setImportance(value Importance)() {
    return nil
}
func (m *Message) setInferenceClassification(value InferenceClassificationType)() {
    return nil
}
func (m *Message) setInternetMessageHeaders(value []InternetMessageHeader)() {
    return nil
}
func (m *Message) setInternetMessageId(value *string)() {
    return nil
}
func (m *Message) setIsDeliveryReceiptRequested(value *bool)() {
    return nil
}
func (m *Message) setIsDraft(value *bool)() {
    return nil
}
func (m *Message) setIsRead(value *bool)() {
    return nil
}
func (m *Message) setIsReadReceiptRequested(value *bool)() {
    return nil
}
func (m *Message) setMultiValueExtendedProperties(value []MultiValueLegacyExtendedProperty)() {
    return nil
}
func (m *Message) setParentFolderId(value *string)() {
    return nil
}
func (m *Message) setReceivedDateTime(value i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    return nil
}
func (m *Message) setReplyTo(value []Recipient)() {
    return nil
}
func (m *Message) setSender(value Recipient)() {
    return nil
}
func (m *Message) setSentDateTime(value i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    return nil
}
func (m *Message) setSingleValueExtendedProperties(value []SingleValueLegacyExtendedProperty)() {
    return nil
}
func (m *Message) setSubject(value *string)() {
    return nil
}
func (m *Message) setToRecipients(value []Recipient)() {
    return nil
}
func (m *Message) setUniqueBody(value ItemBody)() {
    return nil
}
func (m *Message) setWebLink(value *string)() {
    return nil
}
