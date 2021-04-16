using System;
using System.Collections.Generic;
using System.Threading.Tasks;
namespace Graphdotnetv4.Users.MailFolders {
    public class MessageRuleActions {
        public List<object> AssignCategories { get; set; }
        public string CopyToFolder { get; set; }
        public bool Delete { get; set; }
        public List<Recipient> ForwardAsAttachmentTo { get; set; }
        public List<Recipient> ForwardTo { get; set; }
        public bool MarkAsRead { get; set; }
        public Importance MarkImportance { get; set; }
        public string MoveToFolder { get; set; }
        public bool PermanentDelete { get; set; }
        public List<Recipient> RedirectTo { get; set; }
        public bool StopProcessingRules { get; set; }
    }
}
