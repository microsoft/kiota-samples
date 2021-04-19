using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users {
    public class Message : OutlookItem, IParsable<Message> {
        public List<Recipient> BccRecipients { get; set; }
        public ItemBody Body { get; set; }
        public string BodyPreview { get; set; }
        public List<Recipient> CcRecipients { get; set; }
        public string ConversationId { get; set; }
        public string ConversationIndex { get; set; }
        public FollowupFlag Flag { get; set; }
        public Recipient From { get; set; }
        public bool HasAttachments { get; set; }
        public Importance Importance { get; set; }
        public InferenceClassificationType InferenceClassification { get; set; }
        public List<InternetMessageHeader> InternetMessageHeaders { get; set; }
        public string InternetMessageId { get; set; }
        public bool IsDeliveryReceiptRequested { get; set; }
        public bool IsDraft { get; set; }
        public bool IsRead { get; set; }
        public bool IsReadReceiptRequested { get; set; }
        public string ParentFolderId { get; set; }
        public string ReceivedDateTime { get; set; }
        public List<Recipient> ReplyTo { get; set; }
        public Recipient Sender { get; set; }
        public string SentDateTime { get; set; }
        public string Subject { get; set; }
        public List<Recipient> ToRecipients { get; set; }
        public ItemBody UniqueBody { get; set; }
        public string WebLink { get; set; }
        public List<Attachment> Attachments { get; set; }
        public List<Extension> Extensions { get; set; }
        public List<MultiValueLegacyExtendedProperty> MultiValueExtendedProperties { get; set; }
        public List<SingleValueLegacyExtendedProperty> SingleValueExtendedProperties { get; set; }
        public new IDictionary<string, Action<Message, IParseNode>> DeserializeFields => new Dictionary<string, Action<Message, IParseNode>> {
            {
                "bccRecipients", (o,n) => { o.BccRecipients = n.GetCollectionOfObjectValues<Recipient>().ToList(); }
            },
            {
                "body", (o,n) => { o.Body = n.GetObjectValue<ItemBody>(); }
            },
            {
                "bodyPreview", (o,n) => { o.BodyPreview = n.GetStringValue(); }
            },
            {
                "ccRecipients", (o,n) => { o.CcRecipients = n.GetCollectionOfObjectValues<Recipient>().ToList(); }
            },
            {
                "conversationId", (o,n) => { o.ConversationId = n.GetStringValue(); }
            },
            {
                "conversationIndex", (o,n) => { o.ConversationIndex = n.GetStringValue(); }
            },
            {
                "flag", (o,n) => { o.Flag = n.GetObjectValue<FollowupFlag>(); }
            },
            {
                "from", (o,n) => { o.From = n.GetObjectValue<Recipient>(); }
            },
            {
                "hasAttachments", (o,n) => { o.HasAttachments = n.GetBoolValue(); }
            },
            {
                "importance", (o,n) => { o.Importance = n.GetObjectValue<Importance>(); }
            },
            {
                "inferenceClassification", (o,n) => { o.InferenceClassification = n.GetObjectValue<InferenceClassificationType>(); }
            },
            {
                "internetMessageHeaders", (o,n) => { o.InternetMessageHeaders = n.GetCollectionOfObjectValues<InternetMessageHeader>().ToList(); }
            },
            {
                "internetMessageId", (o,n) => { o.InternetMessageId = n.GetStringValue(); }
            },
            {
                "isDeliveryReceiptRequested", (o,n) => { o.IsDeliveryReceiptRequested = n.GetBoolValue(); }
            },
            {
                "isDraft", (o,n) => { o.IsDraft = n.GetBoolValue(); }
            },
            {
                "isRead", (o,n) => { o.IsRead = n.GetBoolValue(); }
            },
            {
                "isReadReceiptRequested", (o,n) => { o.IsReadReceiptRequested = n.GetBoolValue(); }
            },
            {
                "parentFolderId", (o,n) => { o.ParentFolderId = n.GetStringValue(); }
            },
            {
                "receivedDateTime", (o,n) => { o.ReceivedDateTime = n.GetStringValue(); }
            },
            {
                "replyTo", (o,n) => { o.ReplyTo = n.GetCollectionOfObjectValues<Recipient>().ToList(); }
            },
            {
                "sender", (o,n) => { o.Sender = n.GetObjectValue<Recipient>(); }
            },
            {
                "sentDateTime", (o,n) => { o.SentDateTime = n.GetStringValue(); }
            },
            {
                "subject", (o,n) => { o.Subject = n.GetStringValue(); }
            },
            {
                "toRecipients", (o,n) => { o.ToRecipients = n.GetCollectionOfObjectValues<Recipient>().ToList(); }
            },
            {
                "uniqueBody", (o,n) => { o.UniqueBody = n.GetObjectValue<ItemBody>(); }
            },
            {
                "webLink", (o,n) => { o.WebLink = n.GetStringValue(); }
            },
            {
                "attachments", (o,n) => { o.Attachments = n.GetCollectionOfObjectValues<Attachment>().ToList(); }
            },
            {
                "extensions", (o,n) => { o.Extensions = n.GetCollectionOfObjectValues<Extension>().ToList(); }
            },
            {
                "multiValueExtendedProperties", (o,n) => { o.MultiValueExtendedProperties = n.GetCollectionOfObjectValues<MultiValueLegacyExtendedProperty>().ToList(); }
            },
            {
                "singleValueExtendedProperties", (o,n) => { o.SingleValueExtendedProperties = n.GetCollectionOfObjectValues<SingleValueLegacyExtendedProperty>().ToList(); }
            },
        };
        public new void Serialize(ISerializationWriter writer) {
            base.Serialize(writer);
            writer.WriteCollectionOfObjectValues<Recipient>("bccRecipients", BccRecipients);
            writer.WriteObjectValue<ItemBody>("body", Body);
            writer.WriteStringValue("bodyPreview", BodyPreview);
            writer.WriteCollectionOfObjectValues<Recipient>("ccRecipients", CcRecipients);
            writer.WriteStringValue("conversationId", ConversationId);
            writer.WriteStringValue("conversationIndex", ConversationIndex);
            writer.WriteObjectValue<FollowupFlag>("flag", Flag);
            writer.WriteObjectValue<Recipient>("from", From);
            writer.WriteBoolValue("hasAttachments", HasAttachments);
            writer.WriteObjectValue<Importance>("importance", Importance);
            writer.WriteObjectValue<InferenceClassificationType>("inferenceClassification", InferenceClassification);
            writer.WriteCollectionOfObjectValues<InternetMessageHeader>("internetMessageHeaders", InternetMessageHeaders);
            writer.WriteStringValue("internetMessageId", InternetMessageId);
            writer.WriteBoolValue("isDeliveryReceiptRequested", IsDeliveryReceiptRequested);
            writer.WriteBoolValue("isDraft", IsDraft);
            writer.WriteBoolValue("isRead", IsRead);
            writer.WriteBoolValue("isReadReceiptRequested", IsReadReceiptRequested);
            writer.WriteStringValue("parentFolderId", ParentFolderId);
            writer.WriteStringValue("receivedDateTime", ReceivedDateTime);
            writer.WriteCollectionOfObjectValues<Recipient>("replyTo", ReplyTo);
            writer.WriteObjectValue<Recipient>("sender", Sender);
            writer.WriteStringValue("sentDateTime", SentDateTime);
            writer.WriteStringValue("subject", Subject);
            writer.WriteCollectionOfObjectValues<Recipient>("toRecipients", ToRecipients);
            writer.WriteObjectValue<ItemBody>("uniqueBody", UniqueBody);
            writer.WriteStringValue("webLink", WebLink);
            writer.WriteCollectionOfObjectValues<Attachment>("attachments", Attachments);
            writer.WriteCollectionOfObjectValues<Extension>("extensions", Extensions);
            writer.WriteCollectionOfObjectValues<MultiValueLegacyExtendedProperty>("multiValueExtendedProperties", MultiValueExtendedProperties);
            writer.WriteCollectionOfObjectValues<SingleValueLegacyExtendedProperty>("singleValueExtendedProperties", SingleValueExtendedProperties);
        }
    }
}
