using Microsoft.Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
namespace Graphdotnetv4.Models.Microsoft.Graph {
    public class Recipient : IParsable {
        /// <summary>Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.</summary>
        public IDictionary<string, object> AdditionalData { get; set; }
        public Graphdotnetv4.Models.Microsoft.Graph.EmailAddress EmailAddress { get; set; }
        /// <summary>
        /// Instantiates a new recipient and sets the default values.
        /// </summary>
        public Recipient() {
            AdditionalData = new Dictionary<string, object>();
        }
        public static Recipient CreateFromDiscriminatorValue(IParseNode parseNode) {
            _ = parseNode ?? throw new ArgumentNullException(nameof(parseNode));
            return new Recipient();
        }
        /// <summary>
        /// The deserialization information for the current model
        /// </summary>
        public IDictionary<string, Action<T, IParseNode>> GetFieldDeserializers<T>() {
            return new Dictionary<string, Action<T, IParseNode>> {
                {"emailAddress", (o,n) => { (o as Recipient).EmailAddress = n.GetObjectValue<Graphdotnetv4.Models.Microsoft.Graph.EmailAddress>(Graphdotnetv4.Models.Microsoft.Graph.EmailAddress.CreateFromDiscriminatorValue); } },
            };
        }
        /// <summary>
        /// Serializes information the current object
        /// <param name="writer">Serialization writer to use to serialize this model</param>
        /// </summary>
        public void Serialize(ISerializationWriter writer) {
            _ = writer ?? throw new ArgumentNullException(nameof(writer));
            writer.WriteObjectValue<Graphdotnetv4.Models.Microsoft.Graph.EmailAddress>("emailAddress", EmailAddress);
            writer.WriteAdditionalData(AdditionalData);
        }
    }
}
