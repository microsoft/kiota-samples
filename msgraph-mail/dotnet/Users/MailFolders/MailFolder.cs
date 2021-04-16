using System;
using System.Collections.Generic;
using System.Threading.Tasks;
namespace Graphdotnetv4.Users.MailFolders {
    public class MailFolder : Entity {
        public int ChildFolderCount { get; set; }
        public string DisplayName { get; set; }
        public string ParentFolderId { get; set; }
        public int TotalItemCount { get; set; }
        public int UnreadItemCount { get; set; }
        public List<MailFolder> ChildFolders { get; set; }
        public List<MessageRule> MessageRules { get; set; }
        public List<Message> Messages { get; set; }
        public List<MultiValueLegacyExtendedProperty> MultiValueExtendedProperties { get; set; }
        public List<SingleValueLegacyExtendedProperty> SingleValueExtendedProperties { get; set; }
    }
}
