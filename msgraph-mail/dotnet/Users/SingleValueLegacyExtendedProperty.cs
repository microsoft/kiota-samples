using Microsoft.Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users {
    public class SingleValueLegacyExtendedProperty : Entity, IParsable {
        /// <summary>A property value.</summary>
        public string Value { get; set; }
        /// <summary>
        /// The deserialization information for the current model
        /// </summary>
        public new IDictionary<string, Action<T, IParseNode>> GetFieldDeserializers() {
            return new Dictionary<string, Action<T, IParseNode>>(base.GetFieldDeserializers()) {
                {"value", (o,n) => { (o as SingleValueLegacyExtendedProperty).Value = n.GetStringValue(); } },
            };
        }
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
