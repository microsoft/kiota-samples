using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users {
    public class MessagesResponse : IParsable<MessagesResponse> {
        public List<Message> Value { get; set; }
        public string NextLink { get; set; }
        public IDictionary<string, Action<MessagesResponse, IParseNode>> DeserializeFields => new Dictionary<string, Action<MessagesResponse, IParseNode>> {
            {
                "value", (o,n) => { o.Value = n.GetCollectionOfObjectValues<Message>().ToList(); }
            },
            {
                "nextLink", (o,n) => { o.NextLink = n.GetStringValue(); }
            },
        };
        public void Serialize(ISerializationWriter writer) {
            writer.WriteCollectionOfObjectValues<Message>("value", Value);
            writer.WriteStringValue("nextLink", NextLink);
        }
    }
}
