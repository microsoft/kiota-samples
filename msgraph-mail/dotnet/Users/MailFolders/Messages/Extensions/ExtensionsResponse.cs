using Microsoft.Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users.MailFolders.Messages.Extensions {
    public class ExtensionsResponse : IParsable {
        /// <summary>Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.</summary>
        public IDictionary<string, object> AdditionalData { get; private set; } = new Dictionary<string, object>();
        public string NextLink { get; set; }
        public List<Extension> Value { get; set; }
        /// <summary>
        /// The deserialization information for the current model
        /// </summary>
        public IDictionary<string, Action<T, IParseNode>> GetFieldDeserializers() {
            return new Dictionary<string, Action<T, IParseNode>> {
                {"@odata.nextLink", (o,n) => { (o as ExtensionsResponse).NextLink = n.GetStringValue(); } },
                {"value", (o,n) => { (o as ExtensionsResponse).Value = n.GetCollectionOfObjectValues<Extension>().ToList(); } },
            };
        }
        /// <summary>
        /// Serialiazes information the current object
        /// <param name="writer">Serialization writer to use to serialize this model</param>
        /// </summary>
        public void Serialize(ISerializationWriter writer) {
            writer.WriteStringValue("@odata.nextLink", NextLink);
            writer.WriteCollectionOfObjectValues<Extension>("value", Value);
            writer.WriteAdditionalData(AdditionalData);
        }
    }
}
