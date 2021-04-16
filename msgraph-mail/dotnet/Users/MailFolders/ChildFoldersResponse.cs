using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users.MailFolders {
    public class ChildFoldersResponse : IParsable<ChildFoldersResponse> {
        public List<MailFolder> Value { get; set; }
        public string NextLink { get; set; }
        public IDictionary<string, Action<ChildFoldersResponse, IParseNode>> DeserializeFields => new Dictionary<string, Action<ChildFoldersResponse, IParseNode>> {
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
