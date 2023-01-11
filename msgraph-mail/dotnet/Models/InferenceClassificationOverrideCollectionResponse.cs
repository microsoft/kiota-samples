using Microsoft.Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
namespace Graphdotnetv4.Models {
    public class InferenceClassificationOverrideCollectionResponse : IAdditionalDataHolder, IParsable {
        /// <summary>Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.</summary>
        public IDictionary<string, object> AdditionalData { get; set; }
        /// <summary>The OdataNextLink property</summary>
#if NETSTANDARD2_1_OR_GREATER || NET6_0_OR_GREATER
        public string? OdataNextLink { get; set; }
#else
        public string OdataNextLink { get; set; }
#endif
        /// <summary>The value property</summary>
#if NETSTANDARD2_1_OR_GREATER || NET6_0_OR_GREATER
        public List<InferenceClassificationOverride>? Value { get; set; }
#else
        public List<InferenceClassificationOverride> Value { get; set; }
#endif
        /// <summary>
        /// Instantiates a new InferenceClassificationOverrideCollectionResponse and sets the default values.
        /// </summary>
        public InferenceClassificationOverrideCollectionResponse() {
            AdditionalData = new Dictionary<string, object>();
        }
        /// <summary>
        /// Creates a new instance of the appropriate class based on discriminator value
        /// </summary>
        /// <param name="parseNode">The parse node to use to read the discriminator value and create the object</param>
        public static InferenceClassificationOverrideCollectionResponse CreateFromDiscriminatorValue(IParseNode parseNode) {
            _ = parseNode ?? throw new ArgumentNullException(nameof(parseNode));
            return new InferenceClassificationOverrideCollectionResponse();
        }
        /// <summary>
        /// The deserialization information for the current model
        /// </summary>
        public IDictionary<string, Action<IParseNode>> GetFieldDeserializers() {
            return new Dictionary<string, Action<IParseNode>> {
                {"@odata.nextLink", n => { OdataNextLink = n.GetStringValue(); } },
                {"value", n => { Value = n.GetCollectionOfObjectValues<InferenceClassificationOverride>(InferenceClassificationOverride.CreateFromDiscriminatorValue)?.ToList(); } },
            };
        }
        /// <summary>
        /// Serializes information the current object
        /// </summary>
        /// <param name="writer">Serialization writer to use to serialize this model</param>
        public void Serialize(ISerializationWriter writer) {
            _ = writer ?? throw new ArgumentNullException(nameof(writer));
            writer.WriteStringValue("@odata.nextLink", OdataNextLink);
            writer.WriteCollectionOfObjectValues<InferenceClassificationOverride>("value", Value);
            writer.WriteAdditionalData(AdditionalData);
        }
    }
}
