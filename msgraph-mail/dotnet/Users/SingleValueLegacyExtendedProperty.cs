using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users {
    public class SingleValueLegacyExtendedProperty : Entity, IParsable<SingleValueLegacyExtendedProperty> {
        public string Value { get; set; }
        public new IDictionary<string, Action<SingleValueLegacyExtendedProperty, IParseNode>> DeserializeFields => new Dictionary<string, Action<SingleValueLegacyExtendedProperty, IParseNode>> {
            {
                "value", (o,n) => { o.Value = n.GetStringValue(); }
            },
        };
        public new void Serialize(ISerializationWriter writer) {
            base.Serialize(writer);
            writer.WriteStringValue("value", Value);
        }
    }
}
