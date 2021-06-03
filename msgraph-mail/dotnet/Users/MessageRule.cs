using Microsoft.Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
using Graphdotnetv4.Users.MailFolders;
namespace Graphdotnetv4.Users {
    public class MessageRule : Entity, IParsable<MessageRule> {
        public MessageRuleActions Actions { get; set; }
        public MessageRulePredicates Conditions { get; set; }
        /// <summary>The serialization information for the current model</summary>
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
        /// <summary>The display name of the rule.</summary>
        public string DisplayName { get; set; }
        public MessageRulePredicates Exceptions { get; set; }
        /// <summary>Indicates whether the rule is in an error condition. Read-only.</summary>
        public bool? HasError { get; set; }
        /// <summary>Indicates whether the rule is enabled to be applied to messages.</summary>
        public bool? IsEnabled { get; set; }
        /// <summary>Indicates if the rule is read-only and cannot be modified or deleted by the rules REST API.</summary>
        public bool? IsReadOnly { get; set; }
        /// <summary>Indicates the order in which the rule is executed, among other rules.</summary>
        public int? Sequence { get; set; }
        /// <summary>
        /// Serialiazes information the current object
        /// <param name="writer">Serialization writer to use to serialize this model</param>
        /// </summary>
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
