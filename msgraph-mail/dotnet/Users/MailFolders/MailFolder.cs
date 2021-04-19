using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users.MailFolders {
    public class MailFolder : Entity, IParsable<MailFolder> {
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
        public new IDictionary<string, Action<MailFolder, IParseNode>> DeserializeFields => new Dictionary<string, Action<MailFolder, IParseNode>> {
            {
                "childFolderCount", (o,n) => { o.ChildFolderCount = n.GetIntValue(); }
            },
            {
                "displayName", (o,n) => { o.DisplayName = n.GetStringValue(); }
            },
            {
                "parentFolderId", (o,n) => { o.ParentFolderId = n.GetStringValue(); }
            },
            {
                "totalItemCount", (o,n) => { o.TotalItemCount = n.GetIntValue(); }
            },
            {
                "unreadItemCount", (o,n) => { o.UnreadItemCount = n.GetIntValue(); }
            },
            {
                "childFolders", (o,n) => { o.ChildFolders = n.GetCollectionOfObjectValues<MailFolder>().ToList(); }
            },
            {
                "messageRules", (o,n) => { o.MessageRules = n.GetCollectionOfObjectValues<MessageRule>().ToList(); }
            },
            {
                "messages", (o,n) => { o.Messages = n.GetCollectionOfObjectValues<Message>().ToList(); }
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
            writer.WriteIntValue("childFolderCount", ChildFolderCount);
            writer.WriteStringValue("displayName", DisplayName);
            writer.WriteStringValue("parentFolderId", ParentFolderId);
            writer.WriteIntValue("totalItemCount", TotalItemCount);
            writer.WriteIntValue("unreadItemCount", UnreadItemCount);
            writer.WriteCollectionOfObjectValues<MailFolder>("childFolders", ChildFolders);
            writer.WriteCollectionOfObjectValues<MessageRule>("messageRules", MessageRules);
            writer.WriteCollectionOfObjectValues<Message>("messages", Messages);
            writer.WriteCollectionOfObjectValues<MultiValueLegacyExtendedProperty>("multiValueExtendedProperties", MultiValueExtendedProperties);
            writer.WriteCollectionOfObjectValues<SingleValueLegacyExtendedProperty>("singleValueExtendedProperties", SingleValueExtendedProperties);
        }
    }
}
