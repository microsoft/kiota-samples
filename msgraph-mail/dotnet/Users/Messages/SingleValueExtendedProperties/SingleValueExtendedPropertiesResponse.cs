using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users.Messages.SingleValueExtendedProperties {
    public class SingleValueExtendedPropertiesResponse : IParsable<SingleValueExtendedPropertiesResponse> {
        public List<SingleValueLegacyExtendedProperty> Value { get; set; }
        public string NextLink { get; set; }
        /// <summary>The serialization information for the current model</summary>
        public IDictionary<string, Action<SingleValueExtendedPropertiesResponse, IParseNode>> DeserializeFields => new Dictionary<string, Action<SingleValueExtendedPropertiesResponse, IParseNode>> {
            {
                "value", (o,n) => { o.Value = n.GetCollectionOfObjectValues<SingleValueLegacyExtendedProperty>().ToList(); }
            },
            {
                "nextLink", (o,n) => { o.NextLink = n.GetStringValue(); }
            },
        };
        /// <summary>
        /// Serialiazes information the current object
        /// <param name="writer">Serialization writer to use to serialize this model</param>
        /// </summary>
        public void Serialize(ISerializationWriter writer) {
            writer.WriteCollectionOfObjectValues<SingleValueLegacyExtendedProperty>("value", Value);
            writer.WriteStringValue("nextLink", NextLink);
            writer.WriteAdditionalData(AdditionalData);
        }
        /// <summary>Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.</summary>
        public IDictionary<string, object> AdditionalData { get; private set; } = new Dictionary<string, object>();
    }
}
