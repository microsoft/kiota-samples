using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users {
    public class MultiValueLegacyExtendedProperty : Entity, IParsable<MultiValueLegacyExtendedProperty> {
        /// <summary>The serialization information for the current model</summary>
        public new IDictionary<string, Action<MultiValueLegacyExtendedProperty, IParseNode>> DeserializeFields => new Dictionary<string, Action<MultiValueLegacyExtendedProperty, IParseNode>> {
            {
                "value", (o,n) => { o.Value = n.GetCollectionOfPrimitiveValues<Array>().ToList(); }
            },
        };
        /// <summary>A collection of property values.</summary>
        public List<Array> Value { get; set; }
        /// <summary>
        /// Serialiazes information the current object
        /// <param name="writer">Serialization writer to use to serialize this model</param>
        /// </summary>
        public new void Serialize(ISerializationWriter writer) {
            base.Serialize(writer);
            writer.WriteCollectionOfPrimitiveValues<Array>("value", Value);
        }
    }
}
