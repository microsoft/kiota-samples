using System;
using System.Threading.Tasks;
namespace Graphdotnetv4.Users.MailFolders {
    public class MessageRule : Entity {
        public MessageRuleActions Actions { get; set; }
        public MessageRulePredicates Conditions { get; set; }
        public string DisplayName { get; set; }
        public MessageRulePredicates Exceptions { get; set; }
        public bool HasError { get; set; }
        public bool IsEnabled { get; set; }
        public bool IsReadOnly { get; set; }
        public int Sequence { get; set; }
    }
}
