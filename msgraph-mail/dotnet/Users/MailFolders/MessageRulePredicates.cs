using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users.MailFolders {
    public class MessageRulePredicates : IParsable<MessageRulePredicates> {
        public List<string> BodyContains { get; set; }
        public List<string> BodyOrSubjectContains { get; set; }
        public List<string> Categories { get; set; }
        public List<Recipient> FromAddresses { get; set; }
        public bool? HasAttachments { get; set; }
        public List<string> HeaderContains { get; set; }
        public Importance? Importance { get; set; }
        public bool? IsApprovalRequest { get; set; }
        public bool? IsAutomaticForward { get; set; }
        public bool? IsAutomaticReply { get; set; }
        public bool? IsEncrypted { get; set; }
        public bool? IsMeetingRequest { get; set; }
        public bool? IsMeetingResponse { get; set; }
        public bool? IsNonDeliveryReport { get; set; }
        public bool? IsPermissionControlled { get; set; }
        public bool? IsReadReceipt { get; set; }
        public bool? IsSigned { get; set; }
        public bool? IsVoicemail { get; set; }
        public MessageActionFlag? MessageActionFlag { get; set; }
        public bool? NotSentToMe { get; set; }
        public List<string> RecipientContains { get; set; }
        public List<string> SenderContains { get; set; }
        public Sensitivity? Sensitivity { get; set; }
        public bool? SentCcMe { get; set; }
        public bool? SentOnlyToMe { get; set; }
        public List<Recipient> SentToAddresses { get; set; }
        public bool? SentToMe { get; set; }
        public bool? SentToOrCcMe { get; set; }
        public List<string> SubjectContains { get; set; }
        public SizeRange WithinSizeRange { get; set; }
        public IDictionary<string, Action<MessageRulePredicates, IParseNode>> DeserializeFields => new Dictionary<string, Action<MessageRulePredicates, IParseNode>> {
            {
                "bodyContains", (o,n) => { o.BodyContains = n.GetCollectionOfPrimitiveValues<string>().ToList(); }
            },
            {
                "bodyOrSubjectContains", (o,n) => { o.BodyOrSubjectContains = n.GetCollectionOfPrimitiveValues<string>().ToList(); }
            },
            {
                "categories", (o,n) => { o.Categories = n.GetCollectionOfPrimitiveValues<string>().ToList(); }
            },
            {
                "fromAddresses", (o,n) => { o.FromAddresses = n.GetCollectionOfObjectValues<Recipient>().ToList(); }
            },
            {
                "hasAttachments", (o,n) => { o.HasAttachments = n.GetBoolValue(); }
            },
            {
                "headerContains", (o,n) => { o.HeaderContains = n.GetCollectionOfPrimitiveValues<string>().ToList(); }
            },
            {
                "importance", (o,n) => { o.Importance = n.GetEnumValue<Importance>(); }
            },
            {
                "isApprovalRequest", (o,n) => { o.IsApprovalRequest = n.GetBoolValue(); }
            },
            {
                "isAutomaticForward", (o,n) => { o.IsAutomaticForward = n.GetBoolValue(); }
            },
            {
                "isAutomaticReply", (o,n) => { o.IsAutomaticReply = n.GetBoolValue(); }
            },
            {
                "isEncrypted", (o,n) => { o.IsEncrypted = n.GetBoolValue(); }
            },
            {
                "isMeetingRequest", (o,n) => { o.IsMeetingRequest = n.GetBoolValue(); }
            },
            {
                "isMeetingResponse", (o,n) => { o.IsMeetingResponse = n.GetBoolValue(); }
            },
            {
                "isNonDeliveryReport", (o,n) => { o.IsNonDeliveryReport = n.GetBoolValue(); }
            },
            {
                "isPermissionControlled", (o,n) => { o.IsPermissionControlled = n.GetBoolValue(); }
            },
            {
                "isReadReceipt", (o,n) => { o.IsReadReceipt = n.GetBoolValue(); }
            },
            {
                "isSigned", (o,n) => { o.IsSigned = n.GetBoolValue(); }
            },
            {
                "isVoicemail", (o,n) => { o.IsVoicemail = n.GetBoolValue(); }
            },
            {
                "messageActionFlag", (o,n) => { o.MessageActionFlag = n.GetEnumValue<MessageActionFlag>(); }
            },
            {
                "notSentToMe", (o,n) => { o.NotSentToMe = n.GetBoolValue(); }
            },
            {
                "recipientContains", (o,n) => { o.RecipientContains = n.GetCollectionOfPrimitiveValues<string>().ToList(); }
            },
            {
                "senderContains", (o,n) => { o.SenderContains = n.GetCollectionOfPrimitiveValues<string>().ToList(); }
            },
            {
                "sensitivity", (o,n) => { o.Sensitivity = n.GetEnumValue<Sensitivity>(); }
            },
            {
                "sentCcMe", (o,n) => { o.SentCcMe = n.GetBoolValue(); }
            },
            {
                "sentOnlyToMe", (o,n) => { o.SentOnlyToMe = n.GetBoolValue(); }
            },
            {
                "sentToAddresses", (o,n) => { o.SentToAddresses = n.GetCollectionOfObjectValues<Recipient>().ToList(); }
            },
            {
                "sentToMe", (o,n) => { o.SentToMe = n.GetBoolValue(); }
            },
            {
                "sentToOrCcMe", (o,n) => { o.SentToOrCcMe = n.GetBoolValue(); }
            },
            {
                "subjectContains", (o,n) => { o.SubjectContains = n.GetCollectionOfPrimitiveValues<string>().ToList(); }
            },
            {
                "withinSizeRange", (o,n) => { o.WithinSizeRange = n.GetObjectValue<SizeRange>(); }
            },
        };
        public void Serialize(ISerializationWriter writer) {
            writer.WriteCollectionOfPrimitiveValues<string>("bodyContains", BodyContains);
            writer.WriteCollectionOfPrimitiveValues<string>("bodyOrSubjectContains", BodyOrSubjectContains);
            writer.WriteCollectionOfPrimitiveValues<string>("categories", Categories);
            writer.WriteCollectionOfObjectValues<Recipient>("fromAddresses", FromAddresses);
            writer.WriteBoolValue("hasAttachments", HasAttachments);
            writer.WriteCollectionOfPrimitiveValues<string>("headerContains", HeaderContains);
            writer.WriteEnumValue<Importance>("importance", Importance);
            writer.WriteBoolValue("isApprovalRequest", IsApprovalRequest);
            writer.WriteBoolValue("isAutomaticForward", IsAutomaticForward);
            writer.WriteBoolValue("isAutomaticReply", IsAutomaticReply);
            writer.WriteBoolValue("isEncrypted", IsEncrypted);
            writer.WriteBoolValue("isMeetingRequest", IsMeetingRequest);
            writer.WriteBoolValue("isMeetingResponse", IsMeetingResponse);
            writer.WriteBoolValue("isNonDeliveryReport", IsNonDeliveryReport);
            writer.WriteBoolValue("isPermissionControlled", IsPermissionControlled);
            writer.WriteBoolValue("isReadReceipt", IsReadReceipt);
            writer.WriteBoolValue("isSigned", IsSigned);
            writer.WriteBoolValue("isVoicemail", IsVoicemail);
            writer.WriteEnumValue<MessageActionFlag>("messageActionFlag", MessageActionFlag);
            writer.WriteBoolValue("notSentToMe", NotSentToMe);
            writer.WriteCollectionOfPrimitiveValues<string>("recipientContains", RecipientContains);
            writer.WriteCollectionOfPrimitiveValues<string>("senderContains", SenderContains);
            writer.WriteEnumValue<Sensitivity>("sensitivity", Sensitivity);
            writer.WriteBoolValue("sentCcMe", SentCcMe);
            writer.WriteBoolValue("sentOnlyToMe", SentOnlyToMe);
            writer.WriteCollectionOfObjectValues<Recipient>("sentToAddresses", SentToAddresses);
            writer.WriteBoolValue("sentToMe", SentToMe);
            writer.WriteBoolValue("sentToOrCcMe", SentToOrCcMe);
            writer.WriteCollectionOfPrimitiveValues<string>("subjectContains", SubjectContains);
            writer.WriteObjectValue<SizeRange>("withinSizeRange", WithinSizeRange);
        }
    }
}
