using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users.MailFolders.Messages {
    public class AttachmentsResponse : IParsable<AttachmentsResponse> {
        public List<Attachment> Value { get; set; }
        public string NextLink { get; set; }
        public IDictionary<string, Action<AttachmentsResponse, IParseNode>> DeserializeFields => new Dictionary<string, Action<AttachmentsResponse, IParseNode>> {
            {
                "value", (o,n) => { o.Value = n.GetCollectionOfObjectValues<Attachment>().ToList(); }
            },
            {
                "nextLink", (o,n) => { o.NextLink = n.GetStringValue(); }
            },
        };
        public void Serialize(ISerializationWriter writer) {
            writer.WriteCollectionOfObjectValues<Attachment>("value", Value);
            writer.WriteStringValue("nextLink", NextLink);
        }
    }
}
