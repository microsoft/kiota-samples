package graph

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
func (m *MessageRulePredicates) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["bodyContains"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        o.(*MessageRulePredicates).SetBodyContains(res)
        return nil
    }
    res["bodyOrSubjectContains"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        o.(*MessageRulePredicates).SetBodyOrSubjectContains(res)
        return nil
    }
    res["categories"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        o.(*MessageRulePredicates).SetCategories(res)
        return nil
    }
    res["fromAddresses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () interface{} { return NewRecipient() })
        if err != nil {
            return err
        }
        res := make([]Recipient, len(val))
        for i, v := range val {
            res[i] = *(v.(*Recipient))
        }
        o.(*MessageRulePredicates).SetFromAddresses(res)
        return nil
    }
    res["hasAttachments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        o.(*MessageRulePredicates).SetHasAttachments(val)
        return nil
    }
    res["headerContains"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        o.(*MessageRulePredicates).SetHeaderContains(res)
        return nil
    }
    res["importance"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseImportance)
        if err != nil {
            return err
        }
        o.(*MessageRulePredicates).SetImportance(val.(*Importance))
        return nil
    }
    res["isApprovalRequest"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        o.(*MessageRulePredicates).SetIsApprovalRequest(val)
        return nil
    }
    res["isAutomaticForward"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        o.(*MessageRulePredicates).SetIsAutomaticForward(val)
        return nil
    }
    res["isAutomaticReply"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        o.(*MessageRulePredicates).SetIsAutomaticReply(val)
        return nil
    }
    res["isEncrypted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        o.(*MessageRulePredicates).SetIsEncrypted(val)
        return nil
    }
    res["isMeetingRequest"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        o.(*MessageRulePredicates).SetIsMeetingRequest(val)
        return nil
    }
    res["isMeetingResponse"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        o.(*MessageRulePredicates).SetIsMeetingResponse(val)
        return nil
    }
    res["isNonDeliveryReport"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        o.(*MessageRulePredicates).SetIsNonDeliveryReport(val)
        return nil
    }
    res["isPermissionControlled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        o.(*MessageRulePredicates).SetIsPermissionControlled(val)
        return nil
    }
    res["isReadReceipt"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        o.(*MessageRulePredicates).SetIsReadReceipt(val)
        return nil
    }
    res["isSigned"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        o.(*MessageRulePredicates).SetIsSigned(val)
        return nil
    }
    res["isVoicemail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        o.(*MessageRulePredicates).SetIsVoicemail(val)
        return nil
    }
    res["messageActionFlag"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseMessageActionFlag)
        if err != nil {
            return err
        }
        o.(*MessageRulePredicates).SetMessageActionFlag(val.(*MessageActionFlag))
        return nil
    }
    res["notSentToMe"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        o.(*MessageRulePredicates).SetNotSentToMe(val)
        return nil
    }
    res["recipientContains"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        o.(*MessageRulePredicates).SetRecipientContains(res)
        return nil
    }
    res["senderContains"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        o.(*MessageRulePredicates).SetSenderContains(res)
        return nil
    }
    res["sensitivity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSensitivity)
        if err != nil {
            return err
        }
        o.(*MessageRulePredicates).SetSensitivity(val.(*Sensitivity))
        return nil
    }
    res["sentCcMe"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        o.(*MessageRulePredicates).SetSentCcMe(val)
        return nil
    }
    res["sentOnlyToMe"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        o.(*MessageRulePredicates).SetSentOnlyToMe(val)
        return nil
    }
    res["sentToAddresses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () interface{} { return NewRecipient() })
        if err != nil {
            return err
        }
        res := make([]Recipient, len(val))
        for i, v := range val {
            res[i] = *(v.(*Recipient))
        }
        o.(*MessageRulePredicates).SetSentToAddresses(res)
        return nil
    }
    res["sentToMe"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        o.(*MessageRulePredicates).SetSentToMe(val)
        return nil
    }
    res["sentToOrCcMe"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        o.(*MessageRulePredicates).SetSentToOrCcMe(val)
        return nil
    }
    res["subjectContains"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        o.(*MessageRulePredicates).SetSubjectContains(res)
        return nil
    }
    res["withinSizeRange"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () interface{} { return NewSizeRange() })
        if err != nil {
            return err
        }
        o.(*MessageRulePredicates).SetWithinSizeRange(val.(*SizeRange))
        return nil
    }
    return res
}
func (m *MessageRulePredicates) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfPrimitiveValues("bodyContains", i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ConvertToArrayOfPrimitives(m.GetBodyContains()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfPrimitiveValues("bodyOrSubjectContains", i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ConvertToArrayOfPrimitives(m.GetBodyOrSubjectContains()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfPrimitiveValues("categories", i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ConvertToArrayOfPrimitives(m.GetCategories()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfObjectValues("fromAddresses", i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ConvertToArrayOfParsable(m.GetFromAddresses()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WritePrimitiveValue("hasAttachments", m.GetHasAttachments())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfPrimitiveValues("headerContains", i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ConvertToArrayOfPrimitives(m.GetHeaderContains()))
        if err != nil {
            return err
        }
    }
    {
        if m.GetImportance() != nil {
            err := writer.WritePrimitiveValue("importance", m.GetImportance().String())
            if err != nil {
                return err
            }
        }
    }
    {
        err := writer.WritePrimitiveValue("isApprovalRequest", m.GetIsApprovalRequest())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WritePrimitiveValue("isAutomaticForward", m.GetIsAutomaticForward())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WritePrimitiveValue("isAutomaticReply", m.GetIsAutomaticReply())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WritePrimitiveValue("isEncrypted", m.GetIsEncrypted())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WritePrimitiveValue("isMeetingRequest", m.GetIsMeetingRequest())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WritePrimitiveValue("isMeetingResponse", m.GetIsMeetingResponse())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WritePrimitiveValue("isNonDeliveryReport", m.GetIsNonDeliveryReport())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WritePrimitiveValue("isPermissionControlled", m.GetIsPermissionControlled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WritePrimitiveValue("isReadReceipt", m.GetIsReadReceipt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WritePrimitiveValue("isSigned", m.GetIsSigned())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WritePrimitiveValue("isVoicemail", m.GetIsVoicemail())
        if err != nil {
            return err
        }
    }
    {
        if m.GetMessageActionFlag() != nil {
            err := writer.WritePrimitiveValue("messageActionFlag", m.GetMessageActionFlag().String())
            if err != nil {
                return err
            }
        }
    }
    {
        err := writer.WritePrimitiveValue("notSentToMe", m.GetNotSentToMe())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfPrimitiveValues("recipientContains", i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ConvertToArrayOfPrimitives(m.GetRecipientContains()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfPrimitiveValues("senderContains", i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ConvertToArrayOfPrimitives(m.GetSenderContains()))
        if err != nil {
            return err
        }
    }
    {
        if m.GetSensitivity() != nil {
            err := writer.WritePrimitiveValue("sensitivity", m.GetSensitivity().String())
            if err != nil {
                return err
            }
        }
    }
    {
        err := writer.WritePrimitiveValue("sentCcMe", m.GetSentCcMe())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WritePrimitiveValue("sentOnlyToMe", m.GetSentOnlyToMe())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfObjectValues("sentToAddresses", i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ConvertToArrayOfParsable(m.GetSentToAddresses()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WritePrimitiveValue("sentToMe", m.GetSentToMe())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WritePrimitiveValue("sentToOrCcMe", m.GetSentToOrCcMe())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfPrimitiveValues("subjectContains", i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ConvertToArrayOfPrimitives(m.GetSubjectContains()))
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
