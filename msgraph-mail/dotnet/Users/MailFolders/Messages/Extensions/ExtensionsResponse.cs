using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users.MailFolders.Messages.Extensions {
    public class ExtensionsResponse : IParsable<ExtensionsResponse> {
        public List<Extension> Value { get; set; }
        public string NextLink { get; set; }
        public IDictionary<string, Action<ExtensionsResponse, IParseNode>> DeserializeFields => new Dictionary<string, Action<ExtensionsResponse, IParseNode>> {
            {
                "value", (o,n) => { o.Value = n.GetCollectionOfObjectValues<Extension>().ToList(); }
            },
            {
                "nextLink", (o,n) => { o.NextLink = n.GetStringValue(); }
            },
        };
        public void Serialize(ISerializationWriter writer) {
            writer.WriteCollectionOfObjectValues<Extension>("value", Value);
            writer.WriteStringValue("nextLink", NextLink);
        }
    }
}
