using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users.MailFolders.Messages.Attachments {
    public class AttachmentsResponse : IParsable<AttachmentsResponse> {
        public List<Attachment> Value { get; set; }
        public string NextLink { get; set; }
        /// <summary>The serialization information for the current model</summary>
        public IDictionary<string, Action<AttachmentsResponse, IParseNode>> DeserializeFields => new Dictionary<string, Action<AttachmentsResponse, IParseNode>> {
            {
                "value", (o,n) => { o.Value = n.GetCollectionOfObjectValues<Attachment>().ToList(); }
            },
            {
                "nextLink", (o,n) => { o.NextLink = n.GetStringValue(); }
            },
        };
        /// <summary>
        /// Serialiazes information the current object
        /// <param name="writer">Serialization writer to use to serialize this model</param>
        /// </summary>
        public void Serialize(ISerializationWriter writer) {
            writer.WriteCollectionOfObjectValues<Attachment>("value", Value);
            writer.WriteStringValue("nextLink", NextLink);
        }
    }
}
