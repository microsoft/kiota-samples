using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users {
    public class Entity : IParsable<Entity> {
        public string Id { get; set; }
        public IDictionary<string, Action<Entity, IParseNode>> DeserializeFields => new Dictionary<string, Action<Entity, IParseNode>> {
            {
                "id", (o,n) => { o.Id = n.GetStringValue(); }
            },
        };
        public void Serialize(ISerializationWriter writer) {
            writer.WriteStringValue("id", Id);
        }
    }
}
