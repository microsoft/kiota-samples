using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users {
    public class EmailAddress : IParsable<EmailAddress> {
        /// <summary>Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.</summary>
        public IDictionary<string, object> AdditionalData { get; private set; } = new Dictionary<string, object>();
        /// <summary>The email address of an entity instance.</summary>
        public string Address { get; set; }
        /// <summary>The serialization information for the current model</summary>
        public IDictionary<string, Action<EmailAddress, IParseNode>> DeserializeFields => new Dictionary<string, Action<EmailAddress, IParseNode>> {
            {
                "address", (o,n) => { o.Address = n.GetStringValue(); }
            },
            {
                "name", (o,n) => { o.Name = n.GetStringValue(); }
            },
        };
        /// <summary>The display name of an entity instance.</summary>
        public string Name { get; set; }
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
