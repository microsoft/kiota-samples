package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Message struct {
    OutlookItem
    attachments []Attachment;
    bccRecipients []Recipient;
    body *ItemBody;
    bodyPreview *string;
    ccRecipients []Recipient;
    conversationId *string;
    conversationIndex *string;
    extensions []Extension;
    flag *FollowupFlag;
    from *Recipient;
    hasAttachments *bool;
    importance *Importance;
    inferenceClassification *InferenceClassificationType;
    internetMessageHeaders []InternetMessageHeader;
    internetMessageId *string;
    isDeliveryReceiptRequested *bool;
    isDraft *bool;
    isRead *bool;
    isReadReceiptRequested *bool;
    multiValueExtendedProperties []MultiValueLegacyExtendedProperty;
    parentFolderId *string;
    receivedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    replyTo []Recipient;
    sender *Recipient;
    sentDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    singleValueExtendedProperties []SingleValueLegacyExtendedProperty;
    subject *string;
    toRecipients []Recipient;
    uniqueBody *ItemBody;
    webLink *string;
}
func NewMessage()(*Message) {
    m := &Message{
        OutlookItem: *NewOutlookItem(),
    }
    return m
}
func (m *Message) getAttachments()([]Attachment) {
    return m.attachments
}
func (m *Message) getBccRecipients()([]Recipient) {
    return m.bccRecipients
}
func (m *Message) getBody()(*ItemBody) {
    return m.body
}
func (m *Message) getBodyPreview()(*string) {
    return m.bodyPreview
}
func (m *Message) getCcRecipients()([]Recipient) {
    return m.ccRecipients
}
func (m *Message) getConversationId()(*string) {
    return m.conversationId
}
func (m *Message) getConversationIndex()(*string) {
    return m.conversationIndex
}
func (m *Message) getExtensions()([]Extension) {
    return m.extensions
}
func (m *Message) getFlag()(*FollowupFlag) {
    return m.flag
}
func (m *Message) getFrom()(*Recipient) {
    return m.from
}
func (m *Message) getHasAttachments()(*bool) {
    return m.hasAttachments
}
func (m *Message) getImportance()(*Importance) {
    return m.importance
}
func (m *Message) getInferenceClassification()(*InferenceClassificationType) {
    return m.inferenceClassification
}
func (m *Message) getInternetMessageHeaders()([]InternetMessageHeader) {
    return m.internetMessageHeaders
}
func (m *Message) getInternetMessageId()(*string) {
    return m.internetMessageId
}
func (m *Message) getIsDeliveryReceiptRequested()(*bool) {
    return m.isDeliveryReceiptRequested
}
func (m *Message) getIsDraft()(*bool) {
    return m.isDraft
}
func (m *Message) getIsRead()(*bool) {
    return m.isRead
}
func (m *Message) getIsReadReceiptRequested()(*bool) {
    return m.isReadReceiptRequested
}
func (m *Message) getMultiValueExtendedProperties()([]MultiValueLegacyExtendedProperty) {
    return m.multiValueExtendedProperties
}
func (m *Message) getParentFolderId()(*string) {
    return m.parentFolderId
}
func (m *Message) getReceivedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.receivedDateTime
}
func (m *Message) getReplyTo()([]Recipient) {
    return m.replyTo
}
func (m *Message) getSender()(*Recipient) {
    return m.sender
}
func (m *Message) getSentDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.sentDateTime
}
func (m *Message) getSingleValueExtendedProperties()([]SingleValueLegacyExtendedProperty) {
    return m.singleValueExtendedProperties
}
func (m *Message) getSubject()(*string) {
    return m.subject
}
func (m *Message) getToRecipients()([]Recipient) {
    return m.toRecipients
}
func (m *Message) getUniqueBody()(*ItemBody) {
    return m.uniqueBody
}
func (m *Message) getWebLink()(*string) {
    return m.webLink
}
func (m *Message) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error), error) {
    return nil, nil
}
func (m *Message) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    return nil
}
func (m *Message) setAttachments(value []Attachment)() {
    m.attachments = value
}
func (m *Message) setBccRecipients(value []Recipient)() {
    m.bccRecipients = value
}
func (m *Message) setBody(value *ItemBody)() {
    m.body = value
}
func (m *Message) setBodyPreview(value *string)() {
    m.bodyPreview = value
}
func (m *Message) setCcRecipients(value []Recipient)() {
    m.ccRecipients = value
}
func (m *Message) setConversationId(value *string)() {
    m.conversationId = value
}
func (m *Message) setConversationIndex(value *string)() {
    m.conversationIndex = value
}
func (m *Message) setExtensions(value []Extension)() {
    m.extensions = value
}
func (m *Message) setFlag(value *FollowupFlag)() {
    m.flag = value
}
func (m *Message) setFrom(value *Recipient)() {
    m.from = value
}
func (m *Message) setHasAttachments(value *bool)() {
    m.hasAttachments = value
}
func (m *Message) setImportance(value *Importance)() {
    m.importance = value
}
func (m *Message) setInferenceClassification(value *InferenceClassificationType)() {
    m.inferenceClassification = value
}
func (m *Message) setInternetMessageHeaders(value []InternetMessageHeader)() {
    m.internetMessageHeaders = value
}
func (m *Message) setInternetMessageId(value *string)() {
    m.internetMessageId = value
}
func (m *Message) setIsDeliveryReceiptRequested(value *bool)() {
    m.isDeliveryReceiptRequested = value
}
func (m *Message) setIsDraft(value *bool)() {
    m.isDraft = value
}
func (m *Message) setIsRead(value *bool)() {
    m.isRead = value
}
func (m *Message) setIsReadReceiptRequested(value *bool)() {
    m.isReadReceiptRequested = value
}
func (m *Message) setMultiValueExtendedProperties(value []MultiValueLegacyExtendedProperty)() {
    m.multiValueExtendedProperties = value
}
func (m *Message) setParentFolderId(value *string)() {
    m.parentFolderId = value
}
func (m *Message) setReceivedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.receivedDateTime = value
}
func (m *Message) setReplyTo(value []Recipient)() {
    m.replyTo = value
}
func (m *Message) setSender(value *Recipient)() {
    m.sender = value
}
func (m *Message) setSentDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.sentDateTime = value
}
func (m *Message) setSingleValueExtendedProperties(value []SingleValueLegacyExtendedProperty)() {
    m.singleValueExtendedProperties = value
}
func (m *Message) setSubject(value *string)() {
    m.subject = value
}
func (m *Message) setToRecipients(value []Recipient)() {
    m.toRecipients = value
}
func (m *Message) setUniqueBody(value *ItemBody)() {
    m.uniqueBody = value
}
func (m *Message) setWebLink(value *string)() {
    m.webLink = value
}
