using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users {
    public class ItemBody : IParsable<ItemBody> {
        public string Content { get; set; }
        public BodyType? ContentType { get; set; }
        public IDictionary<string, Action<ItemBody, IParseNode>> DeserializeFields => new Dictionary<string, Action<ItemBody, IParseNode>> {
            {
                "content", (o,n) => { o.Content = n.GetStringValue(); }
            },
            {
                "contentType", (o,n) => { o.ContentType = n.GetEnumValue<BodyType>(); }
            },
        };
        public void Serialize(ISerializationWriter writer) {
            writer.WriteStringValue("content", Content);
            writer.WriteEnumValue<BodyType>("contentType", ContentType);
        }
    }
}
