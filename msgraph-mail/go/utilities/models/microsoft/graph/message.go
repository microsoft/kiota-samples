package graph

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
    conversationIndex []byte;
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
func (m *Message) GetAttachments()([]Attachment) {
    return m.attachments
}
func (m *Message) GetBccRecipients()([]Recipient) {
    return m.bccRecipients
}
func (m *Message) GetBody()(*ItemBody) {
    return m.body
}
func (m *Message) GetBodyPreview()(*string) {
    return m.bodyPreview
}
func (m *Message) GetCcRecipients()([]Recipient) {
    return m.ccRecipients
}
func (m *Message) GetConversationId()(*string) {
    return m.conversationId
}
func (m *Message) GetConversationIndex()([]byte) {
    return m.conversationIndex
}
func (m *Message) GetExtensions()([]Extension) {
    return m.extensions
}
func (m *Message) GetFlag()(*FollowupFlag) {
    return m.flag
}
func (m *Message) GetFrom()(*Recipient) {
    return m.from
}
func (m *Message) GetHasAttachments()(*bool) {
    return m.hasAttachments
}
func (m *Message) GetImportance()(*Importance) {
    return m.importance
}
func (m *Message) GetInferenceClassification()(*InferenceClassificationType) {
    return m.inferenceClassification
}
func (m *Message) GetInternetMessageHeaders()([]InternetMessageHeader) {
    return m.internetMessageHeaders
}
func (m *Message) GetInternetMessageId()(*string) {
    return m.internetMessageId
}
func (m *Message) GetIsDeliveryReceiptRequested()(*bool) {
    return m.isDeliveryReceiptRequested
}
func (m *Message) GetIsDraft()(*bool) {
    return m.isDraft
}
func (m *Message) GetIsRead()(*bool) {
    return m.isRead
}
func (m *Message) GetIsReadReceiptRequested()(*bool) {
    return m.isReadReceiptRequested
}
func (m *Message) GetMultiValueExtendedProperties()([]MultiValueLegacyExtendedProperty) {
    return m.multiValueExtendedProperties
}
func (m *Message) GetParentFolderId()(*string) {
    return m.parentFolderId
}
func (m *Message) GetReceivedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.receivedDateTime
}
func (m *Message) GetReplyTo()([]Recipient) {
    return m.replyTo
}
func (m *Message) GetSender()(*Recipient) {
    return m.sender
}
func (m *Message) GetSentDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.sentDateTime
}
func (m *Message) GetSingleValueExtendedProperties()([]SingleValueLegacyExtendedProperty) {
    return m.singleValueExtendedProperties
}
func (m *Message) GetSubject()(*string) {
    return m.subject
}
func (m *Message) GetToRecipients()([]Recipient) {
    return m.toRecipients
}
func (m *Message) GetUniqueBody()(*ItemBody) {
    return m.uniqueBody
}
func (m *Message) GetWebLink()(*string) {
    return m.webLink
}
func (m *Message) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.OutlookItem.GetFieldDeserializers()
    res["attachments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAttachment() })
        if err != nil {
            return err
        }
        res := make([]Attachment, len(val))
        for i, v := range val {
            res[i] = *(v.(*Attachment))
        }
        m.SetAttachments(res)
        return nil
    }
    res["bccRecipients"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRecipient() })
        if err != nil {
            return err
        }
        res := make([]Recipient, len(val))
        for i, v := range val {
            res[i] = *(v.(*Recipient))
        }
        m.SetBccRecipients(res)
        return nil
    }
    res["body"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemBody() })
        if err != nil {
            return err
        }
        m.SetBody(val.(*ItemBody))
        return nil
    }
    res["bodyPreview"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetBodyPreview(val)
        return nil
    }
    res["ccRecipients"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRecipient() })
        if err != nil {
            return err
        }
        res := make([]Recipient, len(val))
        for i, v := range val {
            res[i] = *(v.(*Recipient))
        }
        m.SetCcRecipients(res)
        return nil
    }
    res["conversationId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetConversationId(val)
        return nil
    }
    res["conversationIndex"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetConversationIndex(val)
        return nil
    }
    res["extensions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExtension() })
        if err != nil {
            return err
        }
        res := make([]Extension, len(val))
        for i, v := range val {
            res[i] = *(v.(*Extension))
        }
        m.SetExtensions(res)
        return nil
    }
    res["flag"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewFollowupFlag() })
        if err != nil {
            return err
        }
        m.SetFlag(val.(*FollowupFlag))
        return nil
    }
    res["from"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRecipient() })
        if err != nil {
            return err
        }
        m.SetFrom(val.(*Recipient))
        return nil
    }
    res["hasAttachments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetHasAttachments(val)
        return nil
    }
    res["importance"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseImportance)
        if err != nil {
            return err
        }
        cast := val.(Importance)
        m.SetImportance(&cast)
        return nil
    }
    res["inferenceClassification"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseInferenceClassificationType)
        if err != nil {
            return err
        }
        cast := val.(InferenceClassificationType)
        m.SetInferenceClassification(&cast)
        return nil
    }
    res["internetMessageHeaders"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewInternetMessageHeader() })
        if err != nil {
            return err
        }
        res := make([]InternetMessageHeader, len(val))
        for i, v := range val {
            res[i] = *(v.(*InternetMessageHeader))
        }
        m.SetInternetMessageHeaders(res)
        return nil
    }
    res["internetMessageId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetInternetMessageId(val)
        return nil
    }
    res["isDeliveryReceiptRequested"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsDeliveryReceiptRequested(val)
        return nil
    }
    res["isDraft"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsDraft(val)
        return nil
    }
    res["isRead"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsRead(val)
        return nil
    }
    res["isReadReceiptRequested"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsReadReceiptRequested(val)
        return nil
    }
    res["multiValueExtendedProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMultiValueLegacyExtendedProperty() })
        if err != nil {
            return err
        }
        res := make([]MultiValueLegacyExtendedProperty, len(val))
        for i, v := range val {
            res[i] = *(v.(*MultiValueLegacyExtendedProperty))
        }
        m.SetMultiValueExtendedProperties(res)
        return nil
    }
    res["parentFolderId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetParentFolderId(val)
        return nil
    }
    res["receivedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetReceivedDateTime(val)
        return nil
    }
    res["replyTo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRecipient() })
        if err != nil {
            return err
        }
        res := make([]Recipient, len(val))
        for i, v := range val {
            res[i] = *(v.(*Recipient))
        }
        m.SetReplyTo(res)
        return nil
    }
    res["sender"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRecipient() })
        if err != nil {
            return err
        }
        m.SetSender(val.(*Recipient))
        return nil
    }
    res["sentDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetSentDateTime(val)
        return nil
    }
    res["singleValueExtendedProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSingleValueLegacyExtendedProperty() })
        if err != nil {
            return err
        }
        res := make([]SingleValueLegacyExtendedProperty, len(val))
        for i, v := range val {
            res[i] = *(v.(*SingleValueLegacyExtendedProperty))
        }
        m.SetSingleValueExtendedProperties(res)
        return nil
    }
    res["subject"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSubject(val)
        return nil
    }
    res["toRecipients"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRecipient() })
        if err != nil {
            return err
        }
        res := make([]Recipient, len(val))
        for i, v := range val {
            res[i] = *(v.(*Recipient))
        }
        m.SetToRecipients(res)
        return nil
    }
    res["uniqueBody"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemBody() })
        if err != nil {
            return err
        }
        m.SetUniqueBody(val.(*ItemBody))
        return nil
    }
    res["webLink"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetWebLink(val)
        return nil
    }
    return res
}
func (m *Message) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.OutlookItem.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteCollectionOfObjectValues("attachments", i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ConvertToArrayOfParsable(m.GetAttachments()))
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfObjectValues("bccRecipients", i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ConvertToArrayOfParsable(m.GetBccRecipients()))
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
        err = writer.WritePrimitiveValue("bodyPreview", m.GetBodyPreview())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfObjectValues("ccRecipients", i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ConvertToArrayOfParsable(m.GetCcRecipients()))
        if err != nil {
            return err
        }
    }
    {
        err = writer.WritePrimitiveValue("conversationId", m.GetConversationId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WritePrimitiveValue("conversationIndex", m.GetConversationIndex())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfObjectValues("extensions", i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ConvertToArrayOfParsable(m.GetExtensions()))
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
        err = writer.WritePrimitiveValue("hasAttachments", m.GetHasAttachments())
        if err != nil {
            return err
        }
    }
    {
        if m.GetImportance() != nil {
            err = writer.WritePrimitiveValue("importance", m.GetImportance().String())
            if err != nil {
                return err
            }
        }
    }
    {
        if m.GetInferenceClassification() != nil {
            err = writer.WritePrimitiveValue("inferenceClassification", m.GetInferenceClassification().String())
            if err != nil {
                return err
            }
        }
    }
    {
        err = writer.WriteCollectionOfObjectValues("internetMessageHeaders", i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ConvertToArrayOfParsable(m.GetInternetMessageHeaders()))
        if err != nil {
            return err
        }
    }
    {
        err = writer.WritePrimitiveValue("internetMessageId", m.GetInternetMessageId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WritePrimitiveValue("isDeliveryReceiptRequested", m.GetIsDeliveryReceiptRequested())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WritePrimitiveValue("isDraft", m.GetIsDraft())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WritePrimitiveValue("isRead", m.GetIsRead())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WritePrimitiveValue("isReadReceiptRequested", m.GetIsReadReceiptRequested())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfObjectValues("multiValueExtendedProperties", i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ConvertToArrayOfParsable(m.GetMultiValueExtendedProperties()))
        if err != nil {
            return err
        }
    }
    {
        err = writer.WritePrimitiveValue("parentFolderId", m.GetParentFolderId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WritePrimitiveValue("receivedDateTime", m.GetReceivedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfObjectValues("replyTo", i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ConvertToArrayOfParsable(m.GetReplyTo()))
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
        err = writer.WritePrimitiveValue("sentDateTime", m.GetSentDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfObjectValues("singleValueExtendedProperties", i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ConvertToArrayOfParsable(m.GetSingleValueExtendedProperties()))
        if err != nil {
            return err
        }
    }
    {
        err = writer.WritePrimitiveValue("subject", m.GetSubject())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfObjectValues("toRecipients", i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ConvertToArrayOfParsable(m.GetToRecipients()))
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
        err = writer.WritePrimitiveValue("webLink", m.GetWebLink())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *Message) SetAttachments(value []Attachment)() {
    m.attachments = value
}
func (m *Message) SetBccRecipients(value []Recipient)() {
    m.bccRecipients = value
}
func (m *Message) SetBody(value *ItemBody)() {
    m.body = value
}
func (m *Message) SetBodyPreview(value *string)() {
    m.bodyPreview = value
}
func (m *Message) SetCcRecipients(value []Recipient)() {
    m.ccRecipients = value
}
func (m *Message) SetConversationId(value *string)() {
    m.conversationId = value
}
func (m *Message) SetConversationIndex(value []byte)() {
    m.conversationIndex = value
}
func (m *Message) SetExtensions(value []Extension)() {
    m.extensions = value
}
func (m *Message) SetFlag(value *FollowupFlag)() {
    m.flag = value
}
func (m *Message) SetFrom(value *Recipient)() {
    m.from = value
}
func (m *Message) SetHasAttachments(value *bool)() {
    m.hasAttachments = value
}
func (m *Message) SetImportance(value *Importance)() {
    m.importance = value
}
func (m *Message) SetInferenceClassification(value *InferenceClassificationType)() {
    m.inferenceClassification = value
}
func (m *Message) SetInternetMessageHeaders(value []InternetMessageHeader)() {
    m.internetMessageHeaders = value
}
func (m *Message) SetInternetMessageId(value *string)() {
    m.internetMessageId = value
}
func (m *Message) SetIsDeliveryReceiptRequested(value *bool)() {
    m.isDeliveryReceiptRequested = value
}
func (m *Message) SetIsDraft(value *bool)() {
    m.isDraft = value
}
func (m *Message) SetIsRead(value *bool)() {
    m.isRead = value
}
func (m *Message) SetIsReadReceiptRequested(value *bool)() {
    m.isReadReceiptRequested = value
}
func (m *Message) SetMultiValueExtendedProperties(value []MultiValueLegacyExtendedProperty)() {
    m.multiValueExtendedProperties = value
}
func (m *Message) SetParentFolderId(value *string)() {
    m.parentFolderId = value
}
func (m *Message) SetReceivedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.receivedDateTime = value
}
func (m *Message) SetReplyTo(value []Recipient)() {
    m.replyTo = value
}
func (m *Message) SetSender(value *Recipient)() {
    m.sender = value
}
func (m *Message) SetSentDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.sentDateTime = value
}
func (m *Message) SetSingleValueExtendedProperties(value []SingleValueLegacyExtendedProperty)() {
    m.singleValueExtendedProperties = value
}
func (m *Message) SetSubject(value *string)() {
    m.subject = value
}
func (m *Message) SetToRecipients(value []Recipient)() {
    m.toRecipients = value
}
func (m *Message) SetUniqueBody(value *ItemBody)() {
    m.uniqueBody = value
}
func (m *Message) SetWebLink(value *string)() {
    m.webLink = value
}
