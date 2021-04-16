using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users {
    public class Extension : Entity, IParsable<Extension> {
        public new IDictionary<string, Action<Extension, IParseNode>> DeserializeFields => new Dictionary<string, Action<Extension, IParseNode>> {
        };
        public new void Serialize(ISerializationWriter writer) {
            base.Serialize(writer);
        }
    }
}
