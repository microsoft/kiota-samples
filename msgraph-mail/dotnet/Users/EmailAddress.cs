using Microsoft.Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users {
    public class EmailAddress : IParsable {
        /// <summary>Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.</summary>
        public IDictionary<string, object> AdditionalData { get; private set; } = new Dictionary<string, object>();
        /// <summary>The email address of an entity instance.</summary>
        public string Address { get; set; }
        /// <summary>The display name of an entity instance.</summary>
        public string Name { get; set; }
        /// <summary>
        /// The deserialization information for the current model
        /// </summary>
        public IDictionary<string, Action<T, IParseNode>> GetFieldDeserializers() {
            return new Dictionary<string, Action<T, IParseNode>> {
                {"address", (o,n) => { (o as EmailAddress).Address = n.GetStringValue(); } },
                {"name", (o,n) => { (o as EmailAddress).Name = n.GetStringValue(); } },
            };
        }
        /// <summary>
        /// Serialiazes information the current object
        /// <param name="writer">Serialization writer to use to serialize this model</param>
        /// </summary>
        public void Serialize(ISerializationWriter writer) {
            writer.WriteStringValue("address", Address);
            writer.WriteStringValue("name", Name);
            writer.WriteAdditionalData(AdditionalData);
        }
    }
}
