using System;
using System.Collections.Generic;
using System.Threading.Tasks;
namespace Graphdotnetv4.Users.MailFolders {
    public class MessageRulePredicates {
        public List<object> BodyContains { get; set; }
        public List<object> BodyOrSubjectContains { get; set; }
        public List<object> Categories { get; set; }
        public List<Recipient> FromAddresses { get; set; }
        public bool HasAttachments { get; set; }
        public List<object> HeaderContains { get; set; }
        public Importance Importance { get; set; }
        public bool IsApprovalRequest { get; set; }
        public bool IsAutomaticForward { get; set; }
        public bool IsAutomaticReply { get; set; }
        public bool IsEncrypted { get; set; }
        public bool IsMeetingRequest { get; set; }
        public bool IsMeetingResponse { get; set; }
        public bool IsNonDeliveryReport { get; set; }
        public bool IsPermissionControlled { get; set; }
        public bool IsReadReceipt { get; set; }
        public bool IsSigned { get; set; }
        public bool IsVoicemail { get; set; }
        public MessageActionFlag MessageActionFlag { get; set; }
        public bool NotSentToMe { get; set; }
        public List<object> RecipientContains { get; set; }
        public List<object> SenderContains { get; set; }
        public Sensitivity Sensitivity { get; set; }
        public bool SentCcMe { get; set; }
        public bool SentOnlyToMe { get; set; }
        public List<Recipient> SentToAddresses { get; set; }
        public bool SentToMe { get; set; }
        public bool SentToOrCcMe { get; set; }
        public List<object> SubjectContains { get; set; }
        public SizeRange WithinSizeRange { get; set; }
    }
}
