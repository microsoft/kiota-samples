using Microsoft.Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users {
    public class ItemBody : IParsable {
        /// <summary>Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.</summary>
        public IDictionary<string, object> AdditionalData { get; private set; } = new Dictionary<string, object>();
        /// <summary>The content of the item.</summary>
        public string Content { get; set; }
        public BodyType? ContentType { get; set; }
        /// <summary>
        /// The deserialization information for the current model
        /// </summary>
        public IDictionary<string, Action<T, IParseNode>> GetFieldDeserializers() {
            return new Dictionary<string, Action<T, IParseNode>> {
                {"content", (o,n) => { (o as ItemBody).Content = n.GetStringValue(); } },
                {"contentType", (o,n) => { (o as ItemBody).ContentType = n.GetEnumValue<BodyType>(); } },
            };
        }
        /// <summary>
        /// Serialiazes information the current object
        /// <param name="writer">Serialization writer to use to serialize this model</param>
        /// </summary>
        public void Serialize(ISerializationWriter writer) {
            writer.WriteStringValue("content", Content);
            writer.WriteEnumValue<BodyType>("contentType", ContentType);
            writer.WriteAdditionalData(AdditionalData);
        }
    }
}
