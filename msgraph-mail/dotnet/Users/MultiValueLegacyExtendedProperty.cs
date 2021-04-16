using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users {
    public class MultiValueLegacyExtendedProperty : Entity, IParsable<MultiValueLegacyExtendedProperty> {
        /// <summary>A collection of property values.</summary>
        public List<string> Value { get; set; }
        /// <summary>The serialization information for the current model</summary>
        public new IDictionary<string, Action<MultiValueLegacyExtendedProperty, IParseNode>> DeserializeFields => new Dictionary<string, Action<MultiValueLegacyExtendedProperty, IParseNode>> {
            {
                "value", (o,n) => { o.Value = n.GetCollectionOfPrimitiveValues<string>().ToList(); }
            },
        };
        /// <summary>
        /// Serialiazes information the current object
        /// <param name="writer">Serialization writer to use to serialize this model</param>
        /// </summary>
        public new void Serialize(ISerializationWriter writer) {
            base.Serialize(writer);
            writer.WriteCollectionOfPrimitiveValues<string>("value", Value);
        }
    }
}
