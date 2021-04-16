using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users.MailFolders {
    public class MessageRule : Entity, IParsable<MessageRule> {
        public MessageRuleActions Actions { get; set; }
        public MessageRulePredicates Conditions { get; set; }
        public string DisplayName { get; set; }
        public MessageRulePredicates Exceptions { get; set; }
        public bool? HasError { get; set; }
        public bool? IsEnabled { get; set; }
        public bool? IsReadOnly { get; set; }
        public int? Sequence { get; set; }
        public new IDictionary<string, Action<MessageRule, IParseNode>> DeserializeFields => new Dictionary<string, Action<MessageRule, IParseNode>> {
            {
                "actions", (o,n) => { o.Actions = n.GetObjectValue<MessageRuleActions>(); }
            },
            {
                "conditions", (o,n) => { o.Conditions = n.GetObjectValue<MessageRulePredicates>(); }
            },
            {
                "displayName", (o,n) => { o.DisplayName = n.GetStringValue(); }
            },
            {
                "exceptions", (o,n) => { o.Exceptions = n.GetObjectValue<MessageRulePredicates>(); }
            },
            {
                "hasError", (o,n) => { o.HasError = n.GetBoolValue(); }
            },
            {
                "isEnabled", (o,n) => { o.IsEnabled = n.GetBoolValue(); }
            },
            {
                "isReadOnly", (o,n) => { o.IsReadOnly = n.GetBoolValue(); }
            },
            {
                "sequence", (o,n) => { o.Sequence = n.GetIntValue(); }
            },
        };
        public new void Serialize(ISerializationWriter writer) {
            base.Serialize(writer);
            writer.WriteObjectValue<MessageRuleActions>("actions", Actions);
            writer.WriteObjectValue<MessageRulePredicates>("conditions", Conditions);
            writer.WriteStringValue("displayName", DisplayName);
            writer.WriteObjectValue<MessageRulePredicates>("exceptions", Exceptions);
            writer.WriteBoolValue("hasError", HasError);
            writer.WriteBoolValue("isEnabled", IsEnabled);
            writer.WriteBoolValue("isReadOnly", IsReadOnly);
            writer.WriteIntValue("sequence", Sequence);
        }
    }
}
