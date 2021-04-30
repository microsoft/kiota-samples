using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users.InferenceClassification.Overrides {
    public class OverridesResponse : IParsable<OverridesResponse> {
        /// <summary>Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.</summary>
        public IDictionary<string, object> AdditionalData { get; private set; } = new Dictionary<string, object>();
        /// <summary>The serialization information for the current model</summary>
        public IDictionary<string, Action<OverridesResponse, IParseNode>> DeserializeFields => new Dictionary<string, Action<OverridesResponse, IParseNode>> {
            {
                "@odata.nextLink", (o,n) => { o.NextLink = n.GetStringValue(); }
            },
            {
                "value", (o,n) => { o.Value = n.GetCollectionOfObjectValues<InferenceClassificationOverride>().ToList(); }
            },
        };
        public string NextLink { get; set; }
        public List<InferenceClassificationOverride> Value { get; set; }
        /// <summary>
        /// Serialiazes information the current object
        /// <param name="writer">Serialization writer to use to serialize this model</param>
        /// </summary>
        public void Serialize(ISerializationWriter writer) {
            writer.WriteStringValue("@odata.nextLink", NextLink);
            writer.WriteCollectionOfObjectValues<InferenceClassificationOverride>("value", Value);
            writer.WriteAdditionalData(AdditionalData);
        }
    }
}
