using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users {
    public class BodyType : IParsable<BodyType> {
        public IDictionary<string, Action<BodyType, IParseNode>> DeserializeFields => new Dictionary<string, Action<BodyType, IParseNode>> {
        };
        public void Serialize(ISerializationWriter writer) {
        }
    }
}
