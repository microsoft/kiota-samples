package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    utils "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Message 
type Message struct {
    OutlookItem
    // The fileAttachment and itemAttachment attachments for the message.
    attachments []Attachmentable
    // The Bcc: recipients for the message.
    bccRecipients []Recipientable
    // The body property
    body ItemBodyable
    // The first 255 characters of the message body. It is in text format.
    bodyPreview *string
    // The Cc: recipients for the message.
    ccRecipients []Recipientable
    // The ID of the conversation the email belongs to.
    conversationId *string
    // Indicates the position of the message within the conversation.
    conversationIndex []byte
    // The collection of open extensions defined for the message. Nullable.
    extensions []Extensionable
    // The flag property
    flag FollowupFlagable
    // The from property
    from Recipientable
    // Indicates whether the message has attachments. This property doesn't include inline attachments, so if a message contains only inline attachments, this property is false. To verify the existence of inline attachments, parse the body property to look for a src attribute, such as <IMG src='cid:image001.jpg@01D26CD8.6C05F070'>.
    hasAttachments *bool
    // The importance property
    importance *Importance
    // The inferenceClassification property
    inferenceClassification *InferenceClassificationType
    // The internetMessageHeaders property
    internetMessageHeaders []InternetMessageHeaderable
    // The internetMessageId property
    internetMessageId *string
    // The isDeliveryReceiptRequested property
    isDeliveryReceiptRequested *bool
    // The isDraft property
    isDraft *bool
    // The isRead property
    isRead *bool
    // The isReadReceiptRequested property
    isReadReceiptRequested *bool
    // The collection of multi-value extended properties defined for the message. Nullable.
    multiValueExtendedProperties []MultiValueLegacyExtendedPropertyable
    // The parentFolderId property
    parentFolderId *string
    // The receivedDateTime property
    receivedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The replyTo property
    replyTo []Recipientable
    // The sender property
    sender Recipientable
    // The sentDateTime property
    sentDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The collection of single-value extended properties defined for the message. Nullable.
    singleValueExtendedProperties []SingleValueLegacyExtendedPropertyable
    // The subject property
    subject *string
    // The toRecipients property
    toRecipients []Recipientable
    // The uniqueBody property
    uniqueBody ItemBodyable
    // The webLink property
    webLink *string
}
// NewMessage instantiates a new message and sets the default values.
func NewMessage()(*Message) {
    m := &Message{
        OutlookItem: *NewOutlookItem(),
    }
    return m
}
// CreateMessageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMessageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMessage(), nil
}
// GetAttachments gets the attachments property value. The fileAttachment and itemAttachment attachments for the message.
func (m *Message) GetAttachments()([]Attachmentable) {
    return m.attachments
}
// GetBccRecipients gets the bccRecipients property value. The Bcc: recipients for the message.
func (m *Message) GetBccRecipients()([]Recipientable) {
    return m.bccRecipients
}
// GetBody gets the body property value. The body property
func (m *Message) GetBody()(ItemBodyable) {
    return m.body
}
// GetBodyPreview gets the bodyPreview property value. The first 255 characters of the message body. It is in text format.
func (m *Message) GetBodyPreview()(*string) {
    return m.bodyPreview
}
// GetCcRecipients gets the ccRecipients property value. The Cc: recipients for the message.
func (m *Message) GetCcRecipients()([]Recipientable) {
    return m.ccRecipients
}
// GetConversationId gets the conversationId property value. The ID of the conversation the email belongs to.
func (m *Message) GetConversationId()(*string) {
    return m.conversationId
}
// GetConversationIndex gets the conversationIndex property value. Indicates the position of the message within the conversation.
func (m *Message) GetConversationIndex()([]byte) {
    return m.conversationIndex
}
// GetExtensions gets the extensions property value. The collection of open extensions defined for the message. Nullable.
func (m *Message) GetExtensions()([]Extensionable) {
    return m.extensions
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Message) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.OutlookItem.GetFieldDeserializers()
    res["attachments"] = utils.SetCollectionOfObjectValues(CreateAttachmentFromDiscriminatorValue , m.SetAttachments)
    res["bccRecipients"] = utils.SetCollectionOfObjectValues(CreateRecipientFromDiscriminatorValue , m.SetBccRecipients)
    res["body"] = utils.SetObjectValue(CreateItemBodyFromDiscriminatorValue , m.SetBody)
    res["bodyPreview"] = utils.SetStringValue(m.SetBodyPreview)
    res["ccRecipients"] = utils.SetCollectionOfObjectValues(CreateRecipientFromDiscriminatorValue , m.SetCcRecipients)
    res["conversationId"] = utils.SetStringValue(m.SetConversationId)
    res["conversationIndex"] = utils.SetByteArrayValue(m.SetConversationIndex)
    res["extensions"] = utils.SetCollectionOfObjectValues(CreateExtensionFromDiscriminatorValue , m.SetExtensions)
    res["flag"] = utils.SetObjectValue(CreateFollowupFlagFromDiscriminatorValue , m.SetFlag)
    res["from"] = utils.SetObjectValue(CreateRecipientFromDiscriminatorValue , m.SetFrom)
    res["hasAttachments"] = utils.SetBoolValue(m.SetHasAttachments)
    res["importance"] = utils.SetEnumValue(ParseImportance , m.SetImportance)
    res["inferenceClassification"] = utils.SetEnumValue(ParseInferenceClassificationType , m.SetInferenceClassification)
    res["internetMessageHeaders"] = utils.SetCollectionOfObjectValues(CreateInternetMessageHeaderFromDiscriminatorValue , m.SetInternetMessageHeaders)
    res["internetMessageId"] = utils.SetStringValue(m.SetInternetMessageId)
    res["isDeliveryReceiptRequested"] = utils.SetBoolValue(m.SetIsDeliveryReceiptRequested)
    res["isDraft"] = utils.SetBoolValue(m.SetIsDraft)
    res["isRead"] = utils.SetBoolValue(m.SetIsRead)
    res["isReadReceiptRequested"] = utils.SetBoolValue(m.SetIsReadReceiptRequested)
    res["multiValueExtendedProperties"] = utils.SetCollectionOfObjectValues(CreateMultiValueLegacyExtendedPropertyFromDiscriminatorValue , m.SetMultiValueExtendedProperties)
    res["parentFolderId"] = utils.SetStringValue(m.SetParentFolderId)
    res["receivedDateTime"] = utils.SetTimeValue(m.SetReceivedDateTime)
    res["replyTo"] = utils.SetCollectionOfObjectValues(CreateRecipientFromDiscriminatorValue , m.SetReplyTo)
    res["sender"] = utils.SetObjectValue(CreateRecipientFromDiscriminatorValue , m.SetSender)
    res["sentDateTime"] = utils.SetTimeValue(m.SetSentDateTime)
    res["singleValueExtendedProperties"] = utils.SetCollectionOfObjectValues(CreateSingleValueLegacyExtendedPropertyFromDiscriminatorValue , m.SetSingleValueExtendedProperties)
    res["subject"] = utils.SetStringValue(m.SetSubject)
    res["toRecipients"] = utils.SetCollectionOfObjectValues(CreateRecipientFromDiscriminatorValue , m.SetToRecipients)
    res["uniqueBody"] = utils.SetObjectValue(CreateItemBodyFromDiscriminatorValue , m.SetUniqueBody)
    res["webLink"] = utils.SetStringValue(m.SetWebLink)
    return res
}
// GetFlag gets the flag property value. The flag property
func (m *Message) GetFlag()(FollowupFlagable) {
    return m.flag
}
// GetFrom gets the from property value. The from property
func (m *Message) GetFrom()(Recipientable) {
    return m.from
}
// GetHasAttachments gets the hasAttachments property value. Indicates whether the message has attachments. This property doesn't include inline attachments, so if a message contains only inline attachments, this property is false. To verify the existence of inline attachments, parse the body property to look for a src attribute, such as <IMG src='cid:image001.jpg@01D26CD8.6C05F070'>.
func (m *Message) GetHasAttachments()(*bool) {
    return m.hasAttachments
}
// GetImportance gets the importance property value. The importance property
func (m *Message) GetImportance()(*Importance) {
    return m.importance
}
// GetInferenceClassification gets the inferenceClassification property value. The inferenceClassification property
func (m *Message) GetInferenceClassification()(*InferenceClassificationType) {
    return m.inferenceClassification
}
// GetInternetMessageHeaders gets the internetMessageHeaders property value. The internetMessageHeaders property
func (m *Message) GetInternetMessageHeaders()([]InternetMessageHeaderable) {
    return m.internetMessageHeaders
}
// GetInternetMessageId gets the internetMessageId property value. The internetMessageId property
func (m *Message) GetInternetMessageId()(*string) {
    return m.internetMessageId
}
// GetIsDeliveryReceiptRequested gets the isDeliveryReceiptRequested property value. The isDeliveryReceiptRequested property
func (m *Message) GetIsDeliveryReceiptRequested()(*bool) {
    return m.isDeliveryReceiptRequested
}
// GetIsDraft gets the isDraft property value. The isDraft property
func (m *Message) GetIsDraft()(*bool) {
    return m.isDraft
}
// GetIsRead gets the isRead property value. The isRead property
func (m *Message) GetIsRead()(*bool) {
    return m.isRead
}
// GetIsReadReceiptRequested gets the isReadReceiptRequested property value. The isReadReceiptRequested property
func (m *Message) GetIsReadReceiptRequested()(*bool) {
    return m.isReadReceiptRequested
}
// GetMultiValueExtendedProperties gets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the message. Nullable.
func (m *Message) GetMultiValueExtendedProperties()([]MultiValueLegacyExtendedPropertyable) {
    return m.multiValueExtendedProperties
}
// GetParentFolderId gets the parentFolderId property value. The parentFolderId property
func (m *Message) GetParentFolderId()(*string) {
    return m.parentFolderId
}
// GetReceivedDateTime gets the receivedDateTime property value. The receivedDateTime property
func (m *Message) GetReceivedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.receivedDateTime
}
// GetReplyTo gets the replyTo property value. The replyTo property
func (m *Message) GetReplyTo()([]Recipientable) {
    return m.replyTo
}
// GetSender gets the sender property value. The sender property
func (m *Message) GetSender()(Recipientable) {
    return m.sender
}
// GetSentDateTime gets the sentDateTime property value. The sentDateTime property
func (m *Message) GetSentDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.sentDateTime
}
// GetSingleValueExtendedProperties gets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the message. Nullable.
func (m *Message) GetSingleValueExtendedProperties()([]SingleValueLegacyExtendedPropertyable) {
    return m.singleValueExtendedProperties
}
// GetSubject gets the subject property value. The subject property
func (m *Message) GetSubject()(*string) {
    return m.subject
}
// GetToRecipients gets the toRecipients property value. The toRecipients property
func (m *Message) GetToRecipients()([]Recipientable) {
    return m.toRecipients
}
// GetUniqueBody gets the uniqueBody property value. The uniqueBody property
func (m *Message) GetUniqueBody()(ItemBodyable) {
    return m.uniqueBody
}
// GetWebLink gets the webLink property value. The webLink property
func (m *Message) GetWebLink()(*string) {
    return m.webLink
}
// Serialize serializes information the current object
func (m *Message) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.OutlookItem.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAttachments() != nil {
        cast := utils.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAttachments())
        err = writer.WriteCollectionOfObjectValues("attachments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetBccRecipients() != nil {
        cast := utils.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetBccRecipients())
        err = writer.WriteCollectionOfObjectValues("bccRecipients", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("body", m.GetBody())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("bodyPreview", m.GetBodyPreview())
        if err != nil {
            return err
        }
    }
    if m.GetCcRecipients() != nil {
        cast := utils.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCcRecipients())
        err = writer.WriteCollectionOfObjectValues("ccRecipients", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("conversationId", m.GetConversationId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("conversationIndex", m.GetConversationIndex())
        if err != nil {
            return err
        }
    }
    if m.GetExtensions() != nil {
        cast := utils.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetExtensions())
        err = writer.WriteCollectionOfObjectValues("extensions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("flag", m.GetFlag())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("from", m.GetFrom())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hasAttachments", m.GetHasAttachments())
        if err != nil {
            return err
        }
    }
    if m.GetImportance() != nil {
        cast := (*m.GetImportance()).String()
        err = writer.WriteStringValue("importance", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetInferenceClassification() != nil {
        cast := (*m.GetInferenceClassification()).String()
        err = writer.WriteStringValue("inferenceClassification", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetInternetMessageHeaders() != nil {
        cast := utils.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetInternetMessageHeaders())
        err = writer.WriteCollectionOfObjectValues("internetMessageHeaders", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("internetMessageId", m.GetInternetMessageId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDeliveryReceiptRequested", m.GetIsDeliveryReceiptRequested())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDraft", m.GetIsDraft())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isRead", m.GetIsRead())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isReadReceiptRequested", m.GetIsReadReceiptRequested())
        if err != nil {
            return err
        }
    }
    if m.GetMultiValueExtendedProperties() != nil {
        cast := utils.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetMultiValueExtendedProperties())
        err = writer.WriteCollectionOfObjectValues("multiValueExtendedProperties", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("parentFolderId", m.GetParentFolderId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("receivedDateTime", m.GetReceivedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetReplyTo() != nil {
        cast := utils.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetReplyTo())
        err = writer.WriteCollectionOfObjectValues("replyTo", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sender", m.GetSender())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("sentDateTime", m.GetSentDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetSingleValueExtendedProperties() != nil {
        cast := utils.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSingleValueExtendedProperties())
        err = writer.WriteCollectionOfObjectValues("singleValueExtendedProperties", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subject", m.GetSubject())
        if err != nil {
            return err
        }
    }
    if m.GetToRecipients() != nil {
        cast := utils.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetToRecipients())
        err = writer.WriteCollectionOfObjectValues("toRecipients", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("uniqueBody", m.GetUniqueBody())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("webLink", m.GetWebLink())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAttachments sets the attachments property value. The fileAttachment and itemAttachment attachments for the message.
func (m *Message) SetAttachments(value []Attachmentable)() {
    m.attachments = value
}
// SetBccRecipients sets the bccRecipients property value. The Bcc: recipients for the message.
func (m *Message) SetBccRecipients(value []Recipientable)() {
    m.bccRecipients = value
}
// SetBody sets the body property value. The body property
func (m *Message) SetBody(value ItemBodyable)() {
    m.body = value
}
// SetBodyPreview sets the bodyPreview property value. The first 255 characters of the message body. It is in text format.
func (m *Message) SetBodyPreview(value *string)() {
    m.bodyPreview = value
}
// SetCcRecipients sets the ccRecipients property value. The Cc: recipients for the message.
func (m *Message) SetCcRecipients(value []Recipientable)() {
    m.ccRecipients = value
}
// SetConversationId sets the conversationId property value. The ID of the conversation the email belongs to.
func (m *Message) SetConversationId(value *string)() {
    m.conversationId = value
}
// SetConversationIndex sets the conversationIndex property value. Indicates the position of the message within the conversation.
func (m *Message) SetConversationIndex(value []byte)() {
    m.conversationIndex = value
}
// SetExtensions sets the extensions property value. The collection of open extensions defined for the message. Nullable.
func (m *Message) SetExtensions(value []Extensionable)() {
    m.extensions = value
}
// SetFlag sets the flag property value. The flag property
func (m *Message) SetFlag(value FollowupFlagable)() {
    m.flag = value
}
// SetFrom sets the from property value. The from property
func (m *Message) SetFrom(value Recipientable)() {
    m.from = value
}
// SetHasAttachments sets the hasAttachments property value. Indicates whether the message has attachments. This property doesn't include inline attachments, so if a message contains only inline attachments, this property is false. To verify the existence of inline attachments, parse the body property to look for a src attribute, such as <IMG src='cid:image001.jpg@01D26CD8.6C05F070'>.
func (m *Message) SetHasAttachments(value *bool)() {
    m.hasAttachments = value
}
// SetImportance sets the importance property value. The importance property
func (m *Message) SetImportance(value *Importance)() {
    m.importance = value
}
// SetInferenceClassification sets the inferenceClassification property value. The inferenceClassification property
func (m *Message) SetInferenceClassification(value *InferenceClassificationType)() {
    m.inferenceClassification = value
}
// SetInternetMessageHeaders sets the internetMessageHeaders property value. The internetMessageHeaders property
func (m *Message) SetInternetMessageHeaders(value []InternetMessageHeaderable)() {
    m.internetMessageHeaders = value
}
// SetInternetMessageId sets the internetMessageId property value. The internetMessageId property
func (m *Message) SetInternetMessageId(value *string)() {
    m.internetMessageId = value
}
// SetIsDeliveryReceiptRequested sets the isDeliveryReceiptRequested property value. The isDeliveryReceiptRequested property
func (m *Message) SetIsDeliveryReceiptRequested(value *bool)() {
    m.isDeliveryReceiptRequested = value
}
// SetIsDraft sets the isDraft property value. The isDraft property
func (m *Message) SetIsDraft(value *bool)() {
    m.isDraft = value
}
// SetIsRead sets the isRead property value. The isRead property
func (m *Message) SetIsRead(value *bool)() {
    m.isRead = value
}
// SetIsReadReceiptRequested sets the isReadReceiptRequested property value. The isReadReceiptRequested property
func (m *Message) SetIsReadReceiptRequested(value *bool)() {
    m.isReadReceiptRequested = value
}
// SetMultiValueExtendedProperties sets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the message. Nullable.
func (m *Message) SetMultiValueExtendedProperties(value []MultiValueLegacyExtendedPropertyable)() {
    m.multiValueExtendedProperties = value
}
// SetParentFolderId sets the parentFolderId property value. The parentFolderId property
func (m *Message) SetParentFolderId(value *string)() {
    m.parentFolderId = value
}
// SetReceivedDateTime sets the receivedDateTime property value. The receivedDateTime property
func (m *Message) SetReceivedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.receivedDateTime = value
}
// SetReplyTo sets the replyTo property value. The replyTo property
func (m *Message) SetReplyTo(value []Recipientable)() {
    m.replyTo = value
}
// SetSender sets the sender property value. The sender property
func (m *Message) SetSender(value Recipientable)() {
    m.sender = value
}
// SetSentDateTime sets the sentDateTime property value. The sentDateTime property
func (m *Message) SetSentDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.sentDateTime = value
}
// SetSingleValueExtendedProperties sets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the message. Nullable.
func (m *Message) SetSingleValueExtendedProperties(value []SingleValueLegacyExtendedPropertyable)() {
    m.singleValueExtendedProperties = value
}
// SetSubject sets the subject property value. The subject property
func (m *Message) SetSubject(value *string)() {
    m.subject = value
}
// SetToRecipients sets the toRecipients property value. The toRecipients property
func (m *Message) SetToRecipients(value []Recipientable)() {
    m.toRecipients = value
}
// SetUniqueBody sets the uniqueBody property value. The uniqueBody property
func (m *Message) SetUniqueBody(value ItemBodyable)() {
    m.uniqueBody = value
}
// SetWebLink sets the webLink property value. The webLink property
func (m *Message) SetWebLink(value *string)() {
    m.webLink = value
}
