using System;
using System.Collections.Generic;
using System.Threading.Tasks;
namespace Graphdotnetv4.Users {
    public class Message : OutlookItem {
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
    }
}
