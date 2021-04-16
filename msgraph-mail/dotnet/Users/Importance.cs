using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users {
    public class Importance : IParsable<Importance> {
        public IDictionary<string, Action<Importance, IParseNode>> DeserializeFields => new Dictionary<string, Action<Importance, IParseNode>> {
        };
        public void Serialize(ISerializationWriter writer) {
        }
    }
}
