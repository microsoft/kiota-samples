using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
using Graphdotnetv4.Users.MailFolders;
namespace Graphdotnetv4.Users {
    public class MailFoldersResponse : IParsable<MailFoldersResponse> {
        public List<MailFolder> Value { get; set; }
        public string NextLink { get; set; }
        public IDictionary<string, Action<MailFoldersResponse, IParseNode>> DeserializeFields => new Dictionary<string, Action<MailFoldersResponse, IParseNode>> {
            {
                "value", (o,n) => { o.Value = n.GetCollectionOfObjectValues<MailFolder>().ToList(); }
            },
            {
                "nextLink", (o,n) => { o.NextLink = n.GetStringValue(); }
            },
        };
        public void Serialize(ISerializationWriter writer) {
            writer.WriteCollectionOfObjectValues<MailFolder>("value", Value);
            writer.WriteStringValue("nextLink", NextLink);
        }
    }
}
