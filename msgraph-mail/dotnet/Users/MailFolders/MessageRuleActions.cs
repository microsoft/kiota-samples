using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users.MailFolders {
    public class MessageRuleActions : IParsable<MessageRuleActions> {
        /// <summary>A list of categories to be assigned to a message.</summary>
        public List<string> AssignCategories { get; set; }
        /// <summary>The ID of a folder that a message is to be copied to.</summary>
        public string CopyToFolder { get; set; }
        /// <summary>Indicates whether a message should be moved to the Deleted Items folder.</summary>
        public bool? Delete { get; set; }
        /// <summary>The email addresses of the recipients to which a message should be forwarded as an attachment.</summary>
        public List<Recipient> ForwardAsAttachmentTo { get; set; }
        /// <summary>The email addresses of the recipients to which a message should be forwarded.</summary>
        public List<Recipient> ForwardTo { get; set; }
        /// <summary>Indicates whether a message should be marked as read.</summary>
        public bool? MarkAsRead { get; set; }
        public Importance? MarkImportance { get; set; }
        /// <summary>The ID of the folder that a message will be moved to.</summary>
        public string MoveToFolder { get; set; }
        /// <summary>Indicates whether a message should be permanently deleted and not saved to the Deleted Items folder.</summary>
        public bool? PermanentDelete { get; set; }
        /// <summary>The email address to which a message should be redirected.</summary>
        public List<Recipient> RedirectTo { get; set; }
        /// <summary>Indicates whether subsequent rules should be evaluated.</summary>
        public bool? StopProcessingRules { get; set; }
        /// <summary>The serialization information for the current model</summary>
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
        /// <summary>
        /// Serialiazes information the current object
        /// <param name="writer">Serialization writer to use to serialize this model</param>
        /// </summary>
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
        /// <summary>Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.</summary>
        public IDictionary<string, object> AdditionalData { get; private set; } = new Dictionary<string, object>();
    }
}
