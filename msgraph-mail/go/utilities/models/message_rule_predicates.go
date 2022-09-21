package models

import (
    core "github.com/microsoftgraph/msgraph-sdk-go-core"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MessageRulePredicates 
type MessageRulePredicates struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Represents the strings that should appear in the body of an incoming message in order for the condition or exception to apply.
    bodyContains []string
    // Represents the strings that should appear in the body or subject of an incoming message in order for the condition or exception to apply.
    bodyOrSubjectContains []string
    // Represents the categories that an incoming message should be labeled with in order for the condition or exception to apply.
    categories []string
    // Represents the specific sender email addresses of an incoming message in order for the condition or exception to apply.
    fromAddresses []Recipientable
    // Indicates whether an incoming message must have attachments in order for the condition or exception to apply.
    hasAttachments *bool
    // Represents the strings that appear in the headers of an incoming message in order for the condition or exception to apply.
    headerContains []string
    // The importance property
    importance *Importance
    // Indicates whether an incoming message must be an approval request in order for the condition or exception to apply.
    isApprovalRequest *bool
    // Indicates whether an incoming message must be automatically forwarded in order for the condition or exception to apply.
    isAutomaticForward *bool
    // Indicates whether an incoming message must be an auto reply in order for the condition or exception to apply.
    isAutomaticReply *bool
    // Indicates whether an incoming message must be encrypted in order for the condition or exception to apply.
    isEncrypted *bool
    // Indicates whether an incoming message must be a meeting request in order for the condition or exception to apply.
    isMeetingRequest *bool
    // Indicates whether an incoming message must be a meeting response in order for the condition or exception to apply.
    isMeetingResponse *bool
    // Indicates whether an incoming message must be a non-delivery report in order for the condition or exception to apply.
    isNonDeliveryReport *bool
    // Indicates whether an incoming message must be permission controlled (RMS-protected) in order for the condition or exception to apply.
    isPermissionControlled *bool
    // Indicates whether an incoming message must be a read receipt in order for the condition or exception to apply.
    isReadReceipt *bool
    // Indicates whether an incoming message must be S/MIME-signed in order for the condition or exception to apply.
    isSigned *bool
    // Indicates whether an incoming message must be a voice mail in order for the condition or exception to apply.
    isVoicemail *bool
    // The messageActionFlag property
    messageActionFlag *MessageActionFlag
    // Indicates whether the owner of the mailbox must not be a recipient of an incoming message in order for the condition or exception to apply.
    notSentToMe *bool
    // Represents the strings that appear in either the toRecipients or ccRecipients properties of an incoming message in order for the condition or exception to apply.
    recipientContains []string
    // Represents the strings that appear in the from property of an incoming message in order for the condition or exception to apply.
    senderContains []string
    // The sensitivity property
    sensitivity *Sensitivity
    // Indicates whether the owner of the mailbox must be in the ccRecipients property of an incoming message in order for the condition or exception to apply.
    sentCcMe *bool
    // Indicates whether the owner of the mailbox must be the only recipient in an incoming message in order for the condition or exception to apply.
    sentOnlyToMe *bool
    // Represents the email addresses that an incoming message must have been sent to in order for the condition or exception to apply.
    sentToAddresses []Recipientable
    // Indicates whether the owner of the mailbox must be in the toRecipients property of an incoming message in order for the condition or exception to apply.
    sentToMe *bool
    // Indicates whether the owner of the mailbox must be in either a toRecipients or ccRecipients property of an incoming message in order for the condition or exception to apply.
    sentToOrCcMe *bool
    // Represents the strings that appear in the subject of an incoming message in order for the condition or exception to apply.
    subjectContains []string
    // The withinSizeRange property
    withinSizeRange SizeRangeable
}
// NewMessageRulePredicates instantiates a new messageRulePredicates and sets the default values.
func NewMessageRulePredicates()(*MessageRulePredicates) {
    m := &MessageRulePredicates{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMessageRulePredicatesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMessageRulePredicatesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMessageRulePredicates(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MessageRulePredicates) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetBodyContains gets the bodyContains property value. Represents the strings that should appear in the body of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetBodyContains()([]string) {
    return m.bodyContains
}
// GetBodyOrSubjectContains gets the bodyOrSubjectContains property value. Represents the strings that should appear in the body or subject of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetBodyOrSubjectContains()([]string) {
    return m.bodyOrSubjectContains
}
// GetCategories gets the categories property value. Represents the categories that an incoming message should be labeled with in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetCategories()([]string) {
    return m.categories
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MessageRulePredicates) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["bodyContains"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        return core.SetCollectionOfReferencedPrimitiveValue(n.GetCollectionOfPrimitiveValues , string , m.SetBodyContains)
    }
    res["bodyOrSubjectContains"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        return core.SetCollectionOfReferencedPrimitiveValue(n.GetCollectionOfPrimitiveValues , string , m.SetBodyOrSubjectContains)
    }
    res["categories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        return core.SetCollectionOfReferencedPrimitiveValue(n.GetCollectionOfPrimitiveValues , string , m.SetCategories)
    }
    res["fromAddresses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        return core.SetCollectionValue(n.GetCollectionOfObjectValues , CreateRecipientFromDiscriminatorValue , m.SetFromAddresses)
    }
    res["hasAttachments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        return core.SetValue(n.GetBoolValue , m.SetHasAttachments)
    }
    res["headerContains"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        return core.SetCollectionOfReferencedPrimitiveValue(n.GetCollectionOfPrimitiveValues , string , m.SetHeaderContains)
    }
    res["importance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        return core.SetReferencedEnumValue(n.GetEnumValue , ParseImportance , m.SetImportance)
    }
    res["isApprovalRequest"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        return core.SetValue(n.GetBoolValue , m.SetIsApprovalRequest)
    }
    res["isAutomaticForward"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        return core.SetValue(n.GetBoolValue , m.SetIsAutomaticForward)
    }
    res["isAutomaticReply"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        return core.SetValue(n.GetBoolValue , m.SetIsAutomaticReply)
    }
    res["isEncrypted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        return core.SetValue(n.GetBoolValue , m.SetIsEncrypted)
    }
    res["isMeetingRequest"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        return core.SetValue(n.GetBoolValue , m.SetIsMeetingRequest)
    }
    res["isMeetingResponse"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        return core.SetValue(n.GetBoolValue , m.SetIsMeetingResponse)
    }
    res["isNonDeliveryReport"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        return core.SetValue(n.GetBoolValue , m.SetIsNonDeliveryReport)
    }
    res["isPermissionControlled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        return core.SetValue(n.GetBoolValue , m.SetIsPermissionControlled)
    }
    res["isReadReceipt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        return core.SetValue(n.GetBoolValue , m.SetIsReadReceipt)
    }
    res["isSigned"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        return core.SetValue(n.GetBoolValue , m.SetIsSigned)
    }
    res["isVoicemail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        return core.SetValue(n.GetBoolValue , m.SetIsVoicemail)
    }
    res["messageActionFlag"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        return core.SetReferencedEnumValue(n.GetEnumValue , ParseMessageActionFlag , m.SetMessageActionFlag)
    }
    res["notSentToMe"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        return core.SetValue(n.GetBoolValue , m.SetNotSentToMe)
    }
    res["recipientContains"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        return core.SetCollectionOfReferencedPrimitiveValue(n.GetCollectionOfPrimitiveValues , string , m.SetRecipientContains)
    }
    res["senderContains"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        return core.SetCollectionOfReferencedPrimitiveValue(n.GetCollectionOfPrimitiveValues , string , m.SetSenderContains)
    }
    res["sensitivity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        return core.SetReferencedEnumValue(n.GetEnumValue , ParseSensitivity , m.SetSensitivity)
    }
    res["sentCcMe"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        return core.SetValue(n.GetBoolValue , m.SetSentCcMe)
    }
    res["sentOnlyToMe"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        return core.SetValue(n.GetBoolValue , m.SetSentOnlyToMe)
    }
    res["sentToAddresses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        return core.SetCollectionValue(n.GetCollectionOfObjectValues , CreateRecipientFromDiscriminatorValue , m.SetSentToAddresses)
    }
    res["sentToMe"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        return core.SetValue(n.GetBoolValue , m.SetSentToMe)
    }
    res["sentToOrCcMe"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        return core.SetValue(n.GetBoolValue , m.SetSentToOrCcMe)
    }
    res["subjectContains"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        return core.SetCollectionOfReferencedPrimitiveValue(n.GetCollectionOfPrimitiveValues , string , m.SetSubjectContains)
    }
    res["withinSizeRange"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        return core.SetObjectValue(n.GetObjectValue , CreateSizeRangeFromDiscriminatorValue , m.SetWithinSizeRange)
    }
    return res
}
// GetFromAddresses gets the fromAddresses property value. Represents the specific sender email addresses of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetFromAddresses()([]Recipientable) {
    return m.fromAddresses
}
// GetHasAttachments gets the hasAttachments property value. Indicates whether an incoming message must have attachments in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetHasAttachments()(*bool) {
    return m.hasAttachments
}
// GetHeaderContains gets the headerContains property value. Represents the strings that appear in the headers of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetHeaderContains()([]string) {
    return m.headerContains
}
// GetImportance gets the importance property value. The importance property
func (m *MessageRulePredicates) GetImportance()(*Importance) {
    return m.importance
}
// GetIsApprovalRequest gets the isApprovalRequest property value. Indicates whether an incoming message must be an approval request in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetIsApprovalRequest()(*bool) {
    return m.isApprovalRequest
}
// GetIsAutomaticForward gets the isAutomaticForward property value. Indicates whether an incoming message must be automatically forwarded in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetIsAutomaticForward()(*bool) {
    return m.isAutomaticForward
}
// GetIsAutomaticReply gets the isAutomaticReply property value. Indicates whether an incoming message must be an auto reply in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetIsAutomaticReply()(*bool) {
    return m.isAutomaticReply
}
// GetIsEncrypted gets the isEncrypted property value. Indicates whether an incoming message must be encrypted in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetIsEncrypted()(*bool) {
    return m.isEncrypted
}
// GetIsMeetingRequest gets the isMeetingRequest property value. Indicates whether an incoming message must be a meeting request in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetIsMeetingRequest()(*bool) {
    return m.isMeetingRequest
}
// GetIsMeetingResponse gets the isMeetingResponse property value. Indicates whether an incoming message must be a meeting response in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetIsMeetingResponse()(*bool) {
    return m.isMeetingResponse
}
// GetIsNonDeliveryReport gets the isNonDeliveryReport property value. Indicates whether an incoming message must be a non-delivery report in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetIsNonDeliveryReport()(*bool) {
    return m.isNonDeliveryReport
}
// GetIsPermissionControlled gets the isPermissionControlled property value. Indicates whether an incoming message must be permission controlled (RMS-protected) in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetIsPermissionControlled()(*bool) {
    return m.isPermissionControlled
}
// GetIsReadReceipt gets the isReadReceipt property value. Indicates whether an incoming message must be a read receipt in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetIsReadReceipt()(*bool) {
    return m.isReadReceipt
}
// GetIsSigned gets the isSigned property value. Indicates whether an incoming message must be S/MIME-signed in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetIsSigned()(*bool) {
    return m.isSigned
}
// GetIsVoicemail gets the isVoicemail property value. Indicates whether an incoming message must be a voice mail in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetIsVoicemail()(*bool) {
    return m.isVoicemail
}
// GetMessageActionFlag gets the messageActionFlag property value. The messageActionFlag property
func (m *MessageRulePredicates) GetMessageActionFlag()(*MessageActionFlag) {
    return m.messageActionFlag
}
// GetNotSentToMe gets the notSentToMe property value. Indicates whether the owner of the mailbox must not be a recipient of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetNotSentToMe()(*bool) {
    return m.notSentToMe
}
// GetRecipientContains gets the recipientContains property value. Represents the strings that appear in either the toRecipients or ccRecipients properties of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetRecipientContains()([]string) {
    return m.recipientContains
}
// GetSenderContains gets the senderContains property value. Represents the strings that appear in the from property of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetSenderContains()([]string) {
    return m.senderContains
}
// GetSensitivity gets the sensitivity property value. The sensitivity property
func (m *MessageRulePredicates) GetSensitivity()(*Sensitivity) {
    return m.sensitivity
}
// GetSentCcMe gets the sentCcMe property value. Indicates whether the owner of the mailbox must be in the ccRecipients property of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetSentCcMe()(*bool) {
    return m.sentCcMe
}
// GetSentOnlyToMe gets the sentOnlyToMe property value. Indicates whether the owner of the mailbox must be the only recipient in an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetSentOnlyToMe()(*bool) {
    return m.sentOnlyToMe
}
// GetSentToAddresses gets the sentToAddresses property value. Represents the email addresses that an incoming message must have been sent to in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetSentToAddresses()([]Recipientable) {
    return m.sentToAddresses
}
// GetSentToMe gets the sentToMe property value. Indicates whether the owner of the mailbox must be in the toRecipients property of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetSentToMe()(*bool) {
    return m.sentToMe
}
// GetSentToOrCcMe gets the sentToOrCcMe property value. Indicates whether the owner of the mailbox must be in either a toRecipients or ccRecipients property of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetSentToOrCcMe()(*bool) {
    return m.sentToOrCcMe
}
// GetSubjectContains gets the subjectContains property value. Represents the strings that appear in the subject of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetSubjectContains()([]string) {
    return m.subjectContains
}
// GetWithinSizeRange gets the withinSizeRange property value. The withinSizeRange property
func (m *MessageRulePredicates) GetWithinSizeRange()(SizeRangeable) {
    return m.withinSizeRange
}
// Serialize serializes information the current object
func (m *MessageRulePredicates) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetBodyContains() != nil {
        err := writer.WriteCollectionOfStringValues("bodyContains", m.GetBodyContains())
        if err != nil {
            return err
        }
    }
    if m.GetBodyOrSubjectContains() != nil {
        err := writer.WriteCollectionOfStringValues("bodyOrSubjectContains", m.GetBodyOrSubjectContains())
        if err != nil {
            return err
        }
    }
    if m.GetCategories() != nil {
        err := writer.WriteCollectionOfStringValues("categories", m.GetCategories())
        if err != nil {
            return err
        }
    }
    if m.GetFromAddresses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFromAddresses()))
        for i, v := range m.GetFromAddresses() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("fromAddresses", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("hasAttachments", m.GetHasAttachments())
        if err != nil {
            return err
        }
    }
    if m.GetHeaderContains() != nil {
        err := writer.WriteCollectionOfStringValues("headerContains", m.GetHeaderContains())
        if err != nil {
            return err
        }
    }
    if m.GetImportance() != nil {
        cast := (*m.GetImportance()).String()
        err := writer.WriteStringValue("importance", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isApprovalRequest", m.GetIsApprovalRequest())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isAutomaticForward", m.GetIsAutomaticForward())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isAutomaticReply", m.GetIsAutomaticReply())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isEncrypted", m.GetIsEncrypted())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isMeetingRequest", m.GetIsMeetingRequest())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isMeetingResponse", m.GetIsMeetingResponse())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isNonDeliveryReport", m.GetIsNonDeliveryReport())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isPermissionControlled", m.GetIsPermissionControlled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isReadReceipt", m.GetIsReadReceipt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isSigned", m.GetIsSigned())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isVoicemail", m.GetIsVoicemail())
        if err != nil {
            return err
        }
    }
    if m.GetMessageActionFlag() != nil {
        cast := (*m.GetMessageActionFlag()).String()
        err := writer.WriteStringValue("messageActionFlag", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("notSentToMe", m.GetNotSentToMe())
        if err != nil {
            return err
        }
    }
    if m.GetRecipientContains() != nil {
        err := writer.WriteCollectionOfStringValues("recipientContains", m.GetRecipientContains())
        if err != nil {
            return err
        }
    }
    if m.GetSenderContains() != nil {
        err := writer.WriteCollectionOfStringValues("senderContains", m.GetSenderContains())
        if err != nil {
            return err
        }
    }
    if m.GetSensitivity() != nil {
        cast := (*m.GetSensitivity()).String()
        err := writer.WriteStringValue("sensitivity", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("sentCcMe", m.GetSentCcMe())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("sentOnlyToMe", m.GetSentOnlyToMe())
        if err != nil {
            return err
        }
    }
    if m.GetSentToAddresses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSentToAddresses()))
        for i, v := range m.GetSentToAddresses() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("sentToAddresses", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("sentToMe", m.GetSentToMe())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("sentToOrCcMe", m.GetSentToOrCcMe())
        if err != nil {
            return err
        }
    }
    if m.GetSubjectContains() != nil {
        err := writer.WriteCollectionOfStringValues("subjectContains", m.GetSubjectContains())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("withinSizeRange", m.GetWithinSizeRange())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MessageRulePredicates) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetBodyContains sets the bodyContains property value. Represents the strings that should appear in the body of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetBodyContains(value []string)() {
    m.bodyContains = value
}
// SetBodyOrSubjectContains sets the bodyOrSubjectContains property value. Represents the strings that should appear in the body or subject of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetBodyOrSubjectContains(value []string)() {
    m.bodyOrSubjectContains = value
}
// SetCategories sets the categories property value. Represents the categories that an incoming message should be labeled with in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetCategories(value []string)() {
    m.categories = value
}
// SetFromAddresses sets the fromAddresses property value. Represents the specific sender email addresses of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetFromAddresses(value []Recipientable)() {
    m.fromAddresses = value
}
// SetHasAttachments sets the hasAttachments property value. Indicates whether an incoming message must have attachments in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetHasAttachments(value *bool)() {
    m.hasAttachments = value
}
// SetHeaderContains sets the headerContains property value. Represents the strings that appear in the headers of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetHeaderContains(value []string)() {
    m.headerContains = value
}
// SetImportance sets the importance property value. The importance property
func (m *MessageRulePredicates) SetImportance(value *Importance)() {
    m.importance = value
}
// SetIsApprovalRequest sets the isApprovalRequest property value. Indicates whether an incoming message must be an approval request in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetIsApprovalRequest(value *bool)() {
    m.isApprovalRequest = value
}
// SetIsAutomaticForward sets the isAutomaticForward property value. Indicates whether an incoming message must be automatically forwarded in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetIsAutomaticForward(value *bool)() {
    m.isAutomaticForward = value
}
// SetIsAutomaticReply sets the isAutomaticReply property value. Indicates whether an incoming message must be an auto reply in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetIsAutomaticReply(value *bool)() {
    m.isAutomaticReply = value
}
// SetIsEncrypted sets the isEncrypted property value. Indicates whether an incoming message must be encrypted in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetIsEncrypted(value *bool)() {
    m.isEncrypted = value
}
// SetIsMeetingRequest sets the isMeetingRequest property value. Indicates whether an incoming message must be a meeting request in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetIsMeetingRequest(value *bool)() {
    m.isMeetingRequest = value
}
// SetIsMeetingResponse sets the isMeetingResponse property value. Indicates whether an incoming message must be a meeting response in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetIsMeetingResponse(value *bool)() {
    m.isMeetingResponse = value
}
// SetIsNonDeliveryReport sets the isNonDeliveryReport property value. Indicates whether an incoming message must be a non-delivery report in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetIsNonDeliveryReport(value *bool)() {
    m.isNonDeliveryReport = value
}
// SetIsPermissionControlled sets the isPermissionControlled property value. Indicates whether an incoming message must be permission controlled (RMS-protected) in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetIsPermissionControlled(value *bool)() {
    m.isPermissionControlled = value
}
// SetIsReadReceipt sets the isReadReceipt property value. Indicates whether an incoming message must be a read receipt in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetIsReadReceipt(value *bool)() {
    m.isReadReceipt = value
}
// SetIsSigned sets the isSigned property value. Indicates whether an incoming message must be S/MIME-signed in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetIsSigned(value *bool)() {
    m.isSigned = value
}
// SetIsVoicemail sets the isVoicemail property value. Indicates whether an incoming message must be a voice mail in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetIsVoicemail(value *bool)() {
    m.isVoicemail = value
}
// SetMessageActionFlag sets the messageActionFlag property value. The messageActionFlag property
func (m *MessageRulePredicates) SetMessageActionFlag(value *MessageActionFlag)() {
    m.messageActionFlag = value
}
// SetNotSentToMe sets the notSentToMe property value. Indicates whether the owner of the mailbox must not be a recipient of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetNotSentToMe(value *bool)() {
    m.notSentToMe = value
}
// SetRecipientContains sets the recipientContains property value. Represents the strings that appear in either the toRecipients or ccRecipients properties of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetRecipientContains(value []string)() {
    m.recipientContains = value
}
// SetSenderContains sets the senderContains property value. Represents the strings that appear in the from property of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetSenderContains(value []string)() {
    m.senderContains = value
}
// SetSensitivity sets the sensitivity property value. The sensitivity property
func (m *MessageRulePredicates) SetSensitivity(value *Sensitivity)() {
    m.sensitivity = value
}
// SetSentCcMe sets the sentCcMe property value. Indicates whether the owner of the mailbox must be in the ccRecipients property of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetSentCcMe(value *bool)() {
    m.sentCcMe = value
}
// SetSentOnlyToMe sets the sentOnlyToMe property value. Indicates whether the owner of the mailbox must be the only recipient in an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetSentOnlyToMe(value *bool)() {
    m.sentOnlyToMe = value
}
// SetSentToAddresses sets the sentToAddresses property value. Represents the email addresses that an incoming message must have been sent to in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetSentToAddresses(value []Recipientable)() {
    m.sentToAddresses = value
}
// SetSentToMe sets the sentToMe property value. Indicates whether the owner of the mailbox must be in the toRecipients property of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetSentToMe(value *bool)() {
    m.sentToMe = value
}
// SetSentToOrCcMe sets the sentToOrCcMe property value. Indicates whether the owner of the mailbox must be in either a toRecipients or ccRecipients property of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetSentToOrCcMe(value *bool)() {
    m.sentToOrCcMe = value
}
// SetSubjectContains sets the subjectContains property value. Represents the strings that appear in the subject of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetSubjectContains(value []string)() {
    m.subjectContains = value
}
// SetWithinSizeRange sets the withinSizeRange property value. The withinSizeRange property
func (m *MessageRulePredicates) SetWithinSizeRange(value SizeRangeable)() {
    m.withinSizeRange = value
}
