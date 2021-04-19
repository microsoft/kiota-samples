using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users.MailFolders {
    public class MessageRuleActions : IParsable<MessageRuleActions> {
        public List<string> AssignCategories { get; set; }
        public string CopyToFolder { get; set; }
        public bool? Delete { get; set; }
        public List<Recipient> ForwardAsAttachmentTo { get; set; }
        public List<Recipient> ForwardTo { get; set; }
        public bool? MarkAsRead { get; set; }
        public Importance? MarkImportance { get; set; }
        public string MoveToFolder { get; set; }
        public bool? PermanentDelete { get; set; }
        public List<Recipient> RedirectTo { get; set; }
        public bool? StopProcessingRules { get; set; }
        public IDictionary<string, Action<MessageRuleActions, IParseNode>> DeserializeFields => new Dictionary<string, Action<MessageRuleActions, IParseNode>> {
            {
                "assignCategories", (o,n) => { o.AssignCategories = n.GetCollectionOfPrimitiveValues<string>().ToList(); }
            },
            {
                "copyToFolder", (o,n) => { o.CopyToFolder = n.GetStringValue(); }
            },
            {
                "delete", (o,n) => { o.Delete = n.GetBoolValue(); }
            },
            {
                "forwardAsAttachmentTo", (o,n) => { o.ForwardAsAttachmentTo = n.GetCollectionOfObjectValues<Recipient>().ToList(); }
            },
            {
                "forwardTo", (o,n) => { o.ForwardTo = n.GetCollectionOfObjectValues<Recipient>().ToList(); }
            },
            {
                "markAsRead", (o,n) => { o.MarkAsRead = n.GetBoolValue(); }
            },
            {
                "markImportance", (o,n) => { o.MarkImportance = n.GetEnumValue<Importance>(); }
            },
            {
                "moveToFolder", (o,n) => { o.MoveToFolder = n.GetStringValue(); }
            },
            {
                "permanentDelete", (o,n) => { o.PermanentDelete = n.GetBoolValue(); }
            },
            {
                "redirectTo", (o,n) => { o.RedirectTo = n.GetCollectionOfObjectValues<Recipient>().ToList(); }
            },
            {
                "stopProcessingRules", (o,n) => { o.StopProcessingRules = n.GetBoolValue(); }
            },
        };
        public void Serialize(ISerializationWriter writer) {
            writer.WriteCollectionOfPrimitiveValues<string>("assignCategories", AssignCategories);
            writer.WriteStringValue("copyToFolder", CopyToFolder);
            writer.WriteBoolValue("delete", Delete);
            writer.WriteCollectionOfObjectValues<Recipient>("forwardAsAttachmentTo", ForwardAsAttachmentTo);
            writer.WriteCollectionOfObjectValues<Recipient>("forwardTo", ForwardTo);
            writer.WriteBoolValue("markAsRead", MarkAsRead);
            writer.WriteEnumValue<Importance>("markImportance", MarkImportance);
            writer.WriteStringValue("moveToFolder", MoveToFolder);
            writer.WriteBoolValue("permanentDelete", PermanentDelete);
            writer.WriteCollectionOfObjectValues<Recipient>("redirectTo", RedirectTo);
            writer.WriteBoolValue("stopProcessingRules", StopProcessingRules);
        }
    }
}
