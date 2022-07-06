using Graphdotnetv4.Models;
using Microsoft.Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
namespace Graphdotnetv4.Models {
    public class Entity : IAdditionalDataHolder, IParsable {
        /// <summary>Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.</summary>
        public IDictionary<string, object> AdditionalData { get; set; }
        /// <summary>The id property</summary>
        public string Id { get; set; }
        /// <summary>The type property</summary>
        public string Type { get; set; }
        /// <summary>
        /// Instantiates a new entity and sets the default values.
        /// </summary>
        public Entity() {
            AdditionalData = new Dictionary<string, object>();
            Type = "#microsoft.graph.entity";
        }
        /// <summary>
        /// Creates a new instance of the appropriate class based on discriminator value
        /// <param name="parseNode">The parse node to use to read the discriminator value and create the object</param>
        /// </summary>
        public static Entity CreateFromDiscriminatorValue(IParseNode parseNode) {
            _ = parseNode ?? throw new ArgumentNullException(nameof(parseNode));
            var mappingValueNode = parseNode.GetChildNode("@odata.type");
            var mappingValue = mappingValueNode?.GetStringValue();
            return mappingValue switch {
                "#microsoft.graph.attachment" => new Attachment(),
                "#microsoft.graph.extension" => new Extension(),
                "#microsoft.graph.message" => new Message(),
                "#microsoft.graph.multiValueLegacyExtendedProperty" => new MultiValueLegacyExtendedProperty(),
                "#microsoft.graph.outlookItem" => new OutlookItem(),
                "#microsoft.graph.singleValueLegacyExtendedProperty" => new SingleValueLegacyExtendedProperty(),
                _ => new Entity(),
            };
        }
        /// <summary>
        /// The deserialization information for the current model
        /// </summary>
        public IDictionary<string, Action<IParseNode>> GetFieldDeserializers() {
            return new Dictionary<string, Action<IParseNode>> {
                {"id", n => { Id = n.GetStringValue(); } },
                {"@odata.type", n => { Type = n.GetStringValue(); } },
            };
        }
        /// <summary>
        /// Serializes information the current object
        /// <param name="writer">Serialization writer to use to serialize this model</param>
        /// </summary>
        public void Serialize(ISerializationWriter writer) {
            _ = writer ?? throw new ArgumentNullException(nameof(writer));
            writer.WriteStringValue("id", Id);
            writer.WriteStringValue("@odata.type", Type);
            writer.WriteAdditionalData(AdditionalData);
        }
    }
}
