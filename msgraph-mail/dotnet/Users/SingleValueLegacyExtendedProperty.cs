using Microsoft.Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users {
    public class SingleValueLegacyExtendedProperty : Entity, IParsable<SingleValueLegacyExtendedProperty> {
        /// <summary>The serialization information for the current model</summary>
        public new IDictionary<string, Action<SingleValueLegacyExtendedProperty, IParseNode>> DeserializeFields => new Dictionary<string, Action<SingleValueLegacyExtendedProperty, IParseNode>> {
            {
                "value", (o,n) => { o.Value = n.GetStringValue(); }
            },
        };
        /// <summary>A property value.</summary>
        public string Value { get; set; }
        /// <summary>
        /// Serialiazes information the current object
        /// <param name="writer">Serialization writer to use to serialize this model</param>
        /// </summary>
        public new void Serialize(ISerializationWriter writer) {
            base.Serialize(writer);
            writer.WriteStringValue("value", Value);
        }
    }
}
