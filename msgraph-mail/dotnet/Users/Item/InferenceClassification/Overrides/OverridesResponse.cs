using Graphdotnetv4.Models.Microsoft.Graph;
using Microsoft.Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
namespace Graphdotnetv4.Users.Item.InferenceClassification.Overrides {
    public class OverridesResponse : IParsable {
        /// <summary>Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.</summary>
        public IDictionary<string, object> AdditionalData { get; set; }
        public string NextLink { get; set; }
        public List<InferenceClassificationOverride> Value { get; set; }
        /// <summary>
        /// Instantiates a new overridesResponse and sets the default values.
        /// </summary>
        public OverridesResponse() {
            AdditionalData = new Dictionary<string, object>();
        }
        public static OverridesResponse CreateFromDiscriminatorValue(IParseNode parseNode) {
            _ = parseNode ?? throw new ArgumentNullException(nameof(parseNode));
            return new OverridesResponse();
        }
        /// <summary>
        /// The deserialization information for the current model
        /// </summary>
        public IDictionary<string, Action<T, IParseNode>> GetFieldDeserializers<T>() {
            return new Dictionary<string, Action<T, IParseNode>> {
                {"@odata.nextLink", (o,n) => { (o as OverridesResponse).NextLink = n.GetStringValue(); } },
                {"value", (o,n) => { (o as OverridesResponse).Value = n.GetCollectionOfObjectValues<InferenceClassificationOverride>(InferenceClassificationOverride.CreateFromDiscriminatorValue).ToList(); } },
            };
        }
        /// <summary>
        /// Serializes information the current object
        /// <param name="writer">Serialization writer to use to serialize this model</param>
        /// </summary>
        public void Serialize(ISerializationWriter writer) {
            _ = writer ?? throw new ArgumentNullException(nameof(writer));
            writer.WriteStringValue("@odata.nextLink", NextLink);
            writer.WriteCollectionOfObjectValues<InferenceClassificationOverride>("value", Value);
            writer.WriteAdditionalData(AdditionalData);
        }
    }
}
