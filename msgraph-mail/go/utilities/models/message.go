package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

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
    // A collection of message headers defined by RFC5322. The set includes message headers indicating the network path taken by a message from the sender to the recipient. It can also contain custom message headers that hold app data for the message.  Returned only on applying a $select query option. Read-only.
    internetMessageHeaders []InternetMessageHeaderable
    // The message ID in the format specified by RFC2822.
    internetMessageId *string
    // Indicates whether a read receipt is requested for the message.
    isDeliveryReceiptRequested *bool
    // Indicates whether the message is a draft. A message is a draft if it hasn't been sent yet.
    isDraft *bool
    // Indicates whether the message has been read.
    isRead *bool
    // Indicates whether a read receipt is requested for the message.
    isReadReceiptRequested *bool
    // The collection of multi-value extended properties defined for the message. Nullable.
    multiValueExtendedProperties []MultiValueLegacyExtendedPropertyable
    // The unique identifier for the message's parent mailFolder.
    parentFolderId *string
    // The date and time the message was received.  The date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    receivedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The email addresses to use when replying.
    replyTo []Recipientable
    // The sender property
    sender Recipientable
    // The date and time the message was sent.  The date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    sentDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The collection of single-value extended properties defined for the message. Nullable.
    singleValueExtendedProperties []SingleValueLegacyExtendedPropertyable
    // The subject of the message.
    subject *string
    // The To: recipients for the message.
    toRecipients []Recipientable
    // The uniqueBody property
    uniqueBody ItemBodyable
    // The URL to open the message in Outlook on the web.You can append an ispopout argument to the end of the URL to change how the message is displayed. If ispopout is not present or if it is set to 1, then the message is shown in a popout window. If ispopout is set to 0, the browser shows the message in the Outlook on the web review pane.The message opens in the browser if you are signed in to your mailbox via Outlook on the web. You are prompted to sign in if you are not already signed in with the browser.This URL cannot be accessed from within an iFrame.
    webLink *string
}
// NewMessage instantiates a new Message and sets the default values.
func NewMessage()(*Message) {
    m := &Message{
        OutlookItem: *NewOutlookItem(),
    }
    return m
}
// CreateMessageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMessageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMessage(), nil
}
// GetAttachments gets the attachments property value. The fileAttachment and itemAttachment attachments for the message.
// returns a []Attachmentable when successful
func (m *Message) GetAttachments()([]Attachmentable) {
    return m.attachments
}
// GetBccRecipients gets the bccRecipients property value. The Bcc: recipients for the message.
// returns a []Recipientable when successful
func (m *Message) GetBccRecipients()([]Recipientable) {
    return m.bccRecipients
}
// GetBody gets the body property value. The body property
// returns a ItemBodyable when successful
func (m *Message) GetBody()(ItemBodyable) {
    return m.body
}
// GetBodyPreview gets the bodyPreview property value. The first 255 characters of the message body. It is in text format.
// returns a *string when successful
func (m *Message) GetBodyPreview()(*string) {
    return m.bodyPreview
}
// GetCcRecipients gets the ccRecipients property value. The Cc: recipients for the message.
// returns a []Recipientable when successful
func (m *Message) GetCcRecipients()([]Recipientable) {
    return m.ccRecipients
}
// GetConversationId gets the conversationId property value. The ID of the conversation the email belongs to.
// returns a *string when successful
func (m *Message) GetConversationId()(*string) {
    return m.conversationId
}
// GetConversationIndex gets the conversationIndex property value. Indicates the position of the message within the conversation.
// returns a []byte when successful
func (m *Message) GetConversationIndex()([]byte) {
    return m.conversationIndex
}
// GetExtensions gets the extensions property value. The collection of open extensions defined for the message. Nullable.
// returns a []Extensionable when successful
func (m *Message) GetExtensions()([]Extensionable) {
    return m.extensions
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Message) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.OutlookItem.GetFieldDeserializers()
    res["attachments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAttachmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Attachmentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Attachmentable)
                }
            }
            m.SetAttachments(res)
        }
        return nil
    }
    res["bccRecipients"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRecipientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Recipientable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Recipientable)
                }
            }
            m.SetBccRecipients(res)
        }
        return nil
    }
    res["body"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemBodyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBody(val.(ItemBodyable))
        }
        return nil
    }
    res["bodyPreview"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBodyPreview(val)
        }
        return nil
    }
    res["ccRecipients"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRecipientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Recipientable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Recipientable)
                }
            }
            m.SetCcRecipients(res)
        }
        return nil
    }
    res["conversationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConversationId(val)
        }
        return nil
    }
    res["conversationIndex"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConversationIndex(val)
        }
        return nil
    }
    res["extensions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExtensionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Extensionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Extensionable)
                }
            }
            m.SetExtensions(res)
        }
        return nil
    }
    res["flag"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFollowupFlagFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFlag(val.(FollowupFlagable))
        }
        return nil
    }
    res["from"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRecipientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFrom(val.(Recipientable))
        }
        return nil
    }
    res["hasAttachments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasAttachments(val)
        }
        return nil
    }
    res["importance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseImportance)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImportance(val.(*Importance))
        }
        return nil
    }
    res["inferenceClassification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseInferenceClassificationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInferenceClassification(val.(*InferenceClassificationType))
        }
        return nil
    }
    res["internetMessageHeaders"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateInternetMessageHeaderFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]InternetMessageHeaderable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(InternetMessageHeaderable)
                }
            }
            m.SetInternetMessageHeaders(res)
        }
        return nil
    }
    res["internetMessageId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInternetMessageId(val)
        }
        return nil
    }
    res["isDeliveryReceiptRequested"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDeliveryReceiptRequested(val)
        }
        return nil
    }
    res["isDraft"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDraft(val)
        }
        return nil
    }
    res["isRead"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRead(val)
        }
        return nil
    }
    res["isReadReceiptRequested"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsReadReceiptRequested(val)
        }
        return nil
    }
    res["multiValueExtendedProperties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMultiValueLegacyExtendedPropertyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MultiValueLegacyExtendedPropertyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MultiValueLegacyExtendedPropertyable)
                }
            }
            m.SetMultiValueExtendedProperties(res)
        }
        return nil
    }
    res["parentFolderId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentFolderId(val)
        }
        return nil
    }
    res["receivedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReceivedDateTime(val)
        }
        return nil
    }
    res["replyTo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRecipientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Recipientable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Recipientable)
                }
            }
            m.SetReplyTo(res)
        }
        return nil
    }
    res["sender"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRecipientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSender(val.(Recipientable))
        }
        return nil
    }
    res["sentDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSentDateTime(val)
        }
        return nil
    }
    res["singleValueExtendedProperties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSingleValueLegacyExtendedPropertyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SingleValueLegacyExtendedPropertyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SingleValueLegacyExtendedPropertyable)
                }
            }
            m.SetSingleValueExtendedProperties(res)
        }
        return nil
    }
    res["subject"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubject(val)
        }
        return nil
    }
    res["toRecipients"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRecipientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Recipientable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Recipientable)
                }
            }
            m.SetToRecipients(res)
        }
        return nil
    }
    res["uniqueBody"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemBodyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUniqueBody(val.(ItemBodyable))
        }
        return nil
    }
    res["webLink"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebLink(val)
        }
        return nil
    }
    return res
}
// GetFlag gets the flag property value. The flag property
// returns a FollowupFlagable when successful
func (m *Message) GetFlag()(FollowupFlagable) {
    return m.flag
}
// GetFrom gets the from property value. The from property
// returns a Recipientable when successful
func (m *Message) GetFrom()(Recipientable) {
    return m.from
}
// GetHasAttachments gets the hasAttachments property value. Indicates whether the message has attachments. This property doesn't include inline attachments, so if a message contains only inline attachments, this property is false. To verify the existence of inline attachments, parse the body property to look for a src attribute, such as <IMG src='cid:image001.jpg@01D26CD8.6C05F070'>.
// returns a *bool when successful
func (m *Message) GetHasAttachments()(*bool) {
    return m.hasAttachments
}
// GetImportance gets the importance property value. The importance property
// returns a *Importance when successful
func (m *Message) GetImportance()(*Importance) {
    return m.importance
}
// GetInferenceClassification gets the inferenceClassification property value. The inferenceClassification property
// returns a *InferenceClassificationType when successful
func (m *Message) GetInferenceClassification()(*InferenceClassificationType) {
    return m.inferenceClassification
}
// GetInternetMessageHeaders gets the internetMessageHeaders property value. A collection of message headers defined by RFC5322. The set includes message headers indicating the network path taken by a message from the sender to the recipient. It can also contain custom message headers that hold app data for the message.  Returned only on applying a $select query option. Read-only.
// returns a []InternetMessageHeaderable when successful
func (m *Message) GetInternetMessageHeaders()([]InternetMessageHeaderable) {
    return m.internetMessageHeaders
}
// GetInternetMessageId gets the internetMessageId property value. The message ID in the format specified by RFC2822.
// returns a *string when successful
func (m *Message) GetInternetMessageId()(*string) {
    return m.internetMessageId
}
// GetIsDeliveryReceiptRequested gets the isDeliveryReceiptRequested property value. Indicates whether a read receipt is requested for the message.
// returns a *bool when successful
func (m *Message) GetIsDeliveryReceiptRequested()(*bool) {
    return m.isDeliveryReceiptRequested
}
// GetIsDraft gets the isDraft property value. Indicates whether the message is a draft. A message is a draft if it hasn't been sent yet.
// returns a *bool when successful
func (m *Message) GetIsDraft()(*bool) {
    return m.isDraft
}
// GetIsRead gets the isRead property value. Indicates whether the message has been read.
// returns a *bool when successful
func (m *Message) GetIsRead()(*bool) {
    return m.isRead
}
// GetIsReadReceiptRequested gets the isReadReceiptRequested property value. Indicates whether a read receipt is requested for the message.
// returns a *bool when successful
func (m *Message) GetIsReadReceiptRequested()(*bool) {
    return m.isReadReceiptRequested
}
// GetMultiValueExtendedProperties gets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the message. Nullable.
// returns a []MultiValueLegacyExtendedPropertyable when successful
func (m *Message) GetMultiValueExtendedProperties()([]MultiValueLegacyExtendedPropertyable) {
    return m.multiValueExtendedProperties
}
// GetParentFolderId gets the parentFolderId property value. The unique identifier for the message's parent mailFolder.
// returns a *string when successful
func (m *Message) GetParentFolderId()(*string) {
    return m.parentFolderId
}
// GetReceivedDateTime gets the receivedDateTime property value. The date and time the message was received.  The date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *Message) GetReceivedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.receivedDateTime
}
// GetReplyTo gets the replyTo property value. The email addresses to use when replying.
// returns a []Recipientable when successful
func (m *Message) GetReplyTo()([]Recipientable) {
    return m.replyTo
}
// GetSender gets the sender property value. The sender property
// returns a Recipientable when successful
func (m *Message) GetSender()(Recipientable) {
    return m.sender
}
// GetSentDateTime gets the sentDateTime property value. The date and time the message was sent.  The date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *Message) GetSentDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.sentDateTime
}
// GetSingleValueExtendedProperties gets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the message. Nullable.
// returns a []SingleValueLegacyExtendedPropertyable when successful
func (m *Message) GetSingleValueExtendedProperties()([]SingleValueLegacyExtendedPropertyable) {
    return m.singleValueExtendedProperties
}
// GetSubject gets the subject property value. The subject of the message.
// returns a *string when successful
func (m *Message) GetSubject()(*string) {
    return m.subject
}
// GetToRecipients gets the toRecipients property value. The To: recipients for the message.
// returns a []Recipientable when successful
func (m *Message) GetToRecipients()([]Recipientable) {
    return m.toRecipients
}
// GetUniqueBody gets the uniqueBody property value. The uniqueBody property
// returns a ItemBodyable when successful
func (m *Message) GetUniqueBody()(ItemBodyable) {
    return m.uniqueBody
}
// GetWebLink gets the webLink property value. The URL to open the message in Outlook on the web.You can append an ispopout argument to the end of the URL to change how the message is displayed. If ispopout is not present or if it is set to 1, then the message is shown in a popout window. If ispopout is set to 0, the browser shows the message in the Outlook on the web review pane.The message opens in the browser if you are signed in to your mailbox via Outlook on the web. You are prompted to sign in if you are not already signed in with the browser.This URL cannot be accessed from within an iFrame.
// returns a *string when successful
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAttachments()))
        for i, v := range m.GetAttachments() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("attachments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetBccRecipients() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetBccRecipients()))
        for i, v := range m.GetBccRecipients() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCcRecipients()))
        for i, v := range m.GetCcRecipients() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExtensions()))
        for i, v := range m.GetExtensions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetInternetMessageHeaders()))
        for i, v := range m.GetInternetMessageHeaders() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMultiValueExtendedProperties()))
        for i, v := range m.GetMultiValueExtendedProperties() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetReplyTo()))
        for i, v := range m.GetReplyTo() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSingleValueExtendedProperties()))
        for i, v := range m.GetSingleValueExtendedProperties() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetToRecipients()))
        for i, v := range m.GetToRecipients() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
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
// SetInternetMessageHeaders sets the internetMessageHeaders property value. A collection of message headers defined by RFC5322. The set includes message headers indicating the network path taken by a message from the sender to the recipient. It can also contain custom message headers that hold app data for the message.  Returned only on applying a $select query option. Read-only.
func (m *Message) SetInternetMessageHeaders(value []InternetMessageHeaderable)() {
    m.internetMessageHeaders = value
}
// SetInternetMessageId sets the internetMessageId property value. The message ID in the format specified by RFC2822.
func (m *Message) SetInternetMessageId(value *string)() {
    m.internetMessageId = value
}
// SetIsDeliveryReceiptRequested sets the isDeliveryReceiptRequested property value. Indicates whether a read receipt is requested for the message.
func (m *Message) SetIsDeliveryReceiptRequested(value *bool)() {
    m.isDeliveryReceiptRequested = value
}
// SetIsDraft sets the isDraft property value. Indicates whether the message is a draft. A message is a draft if it hasn't been sent yet.
func (m *Message) SetIsDraft(value *bool)() {
    m.isDraft = value
}
// SetIsRead sets the isRead property value. Indicates whether the message has been read.
func (m *Message) SetIsRead(value *bool)() {
    m.isRead = value
}
// SetIsReadReceiptRequested sets the isReadReceiptRequested property value. Indicates whether a read receipt is requested for the message.
func (m *Message) SetIsReadReceiptRequested(value *bool)() {
    m.isReadReceiptRequested = value
}
// SetMultiValueExtendedProperties sets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the message. Nullable.
func (m *Message) SetMultiValueExtendedProperties(value []MultiValueLegacyExtendedPropertyable)() {
    m.multiValueExtendedProperties = value
}
// SetParentFolderId sets the parentFolderId property value. The unique identifier for the message's parent mailFolder.
func (m *Message) SetParentFolderId(value *string)() {
    m.parentFolderId = value
}
// SetReceivedDateTime sets the receivedDateTime property value. The date and time the message was received.  The date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *Message) SetReceivedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.receivedDateTime = value
}
// SetReplyTo sets the replyTo property value. The email addresses to use when replying.
func (m *Message) SetReplyTo(value []Recipientable)() {
    m.replyTo = value
}
// SetSender sets the sender property value. The sender property
func (m *Message) SetSender(value Recipientable)() {
    m.sender = value
}
// SetSentDateTime sets the sentDateTime property value. The date and time the message was sent.  The date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *Message) SetSentDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.sentDateTime = value
}
// SetSingleValueExtendedProperties sets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the message. Nullable.
func (m *Message) SetSingleValueExtendedProperties(value []SingleValueLegacyExtendedPropertyable)() {
    m.singleValueExtendedProperties = value
}
// SetSubject sets the subject property value. The subject of the message.
func (m *Message) SetSubject(value *string)() {
    m.subject = value
}
// SetToRecipients sets the toRecipients property value. The To: recipients for the message.
func (m *Message) SetToRecipients(value []Recipientable)() {
    m.toRecipients = value
}
// SetUniqueBody sets the uniqueBody property value. The uniqueBody property
func (m *Message) SetUniqueBody(value ItemBodyable)() {
    m.uniqueBody = value
}
// SetWebLink sets the webLink property value. The URL to open the message in Outlook on the web.You can append an ispopout argument to the end of the URL to change how the message is displayed. If ispopout is not present or if it is set to 1, then the message is shown in a popout window. If ispopout is set to 0, the browser shows the message in the Outlook on the web review pane.The message opens in the browser if you are signed in to your mailbox via Outlook on the web. You are prompted to sign in if you are not already signed in with the browser.This URL cannot be accessed from within an iFrame.
func (m *Message) SetWebLink(value *string)() {
    m.webLink = value
}
type Messageable interface {
    OutlookItemable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAttachments()([]Attachmentable)
    GetBccRecipients()([]Recipientable)
    GetBody()(ItemBodyable)
    GetBodyPreview()(*string)
    GetCcRecipients()([]Recipientable)
    GetConversationId()(*string)
    GetConversationIndex()([]byte)
    GetExtensions()([]Extensionable)
    GetFlag()(FollowupFlagable)
    GetFrom()(Recipientable)
    GetHasAttachments()(*bool)
    GetImportance()(*Importance)
    GetInferenceClassification()(*InferenceClassificationType)
    GetInternetMessageHeaders()([]InternetMessageHeaderable)
    GetInternetMessageId()(*string)
    GetIsDeliveryReceiptRequested()(*bool)
    GetIsDraft()(*bool)
    GetIsRead()(*bool)
    GetIsReadReceiptRequested()(*bool)
    GetMultiValueExtendedProperties()([]MultiValueLegacyExtendedPropertyable)
    GetParentFolderId()(*string)
    GetReceivedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetReplyTo()([]Recipientable)
    GetSender()(Recipientable)
    GetSentDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSingleValueExtendedProperties()([]SingleValueLegacyExtendedPropertyable)
    GetSubject()(*string)
    GetToRecipients()([]Recipientable)
    GetUniqueBody()(ItemBodyable)
    GetWebLink()(*string)
    SetAttachments(value []Attachmentable)()
    SetBccRecipients(value []Recipientable)()
    SetBody(value ItemBodyable)()
    SetBodyPreview(value *string)()
    SetCcRecipients(value []Recipientable)()
    SetConversationId(value *string)()
    SetConversationIndex(value []byte)()
    SetExtensions(value []Extensionable)()
    SetFlag(value FollowupFlagable)()
    SetFrom(value Recipientable)()
    SetHasAttachments(value *bool)()
    SetImportance(value *Importance)()
    SetInferenceClassification(value *InferenceClassificationType)()
    SetInternetMessageHeaders(value []InternetMessageHeaderable)()
    SetInternetMessageId(value *string)()
    SetIsDeliveryReceiptRequested(value *bool)()
    SetIsDraft(value *bool)()
    SetIsRead(value *bool)()
    SetIsReadReceiptRequested(value *bool)()
    SetMultiValueExtendedProperties(value []MultiValueLegacyExtendedPropertyable)()
    SetParentFolderId(value *string)()
    SetReceivedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetReplyTo(value []Recipientable)()
    SetSender(value Recipientable)()
    SetSentDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSingleValueExtendedProperties(value []SingleValueLegacyExtendedPropertyable)()
    SetSubject(value *string)()
    SetToRecipients(value []Recipientable)()
    SetUniqueBody(value ItemBodyable)()
    SetWebLink(value *string)()
}
