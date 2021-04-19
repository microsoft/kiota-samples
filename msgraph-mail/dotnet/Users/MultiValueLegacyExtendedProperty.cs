using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users {
    public class MultiValueLegacyExtendedProperty : Entity, IParsable<MultiValueLegacyExtendedProperty> {
        public List<string> Value { get; set; }
        public new IDictionary<string, Action<MultiValueLegacyExtendedProperty, IParseNode>> DeserializeFields => new Dictionary<string, Action<MultiValueLegacyExtendedProperty, IParseNode>> {
            {
                "value", (o,n) => { o.Value = n.GetCollectionOfPrimitiveValues<string>().ToList(); }
            },
        };
        public new void Serialize(ISerializationWriter writer) {
            base.Serialize(writer);
            writer.WriteCollectionOfPrimitiveValues<string>("value", Value);
        }
    }
}
