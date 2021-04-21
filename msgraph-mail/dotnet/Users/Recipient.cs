using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users {
    public class Recipient : IParsable<Recipient> {
        public EmailAddress EmailAddress { get; set; }
        /// <summary>The serialization information for the current model</summary>
        public IDictionary<string, Action<Recipient, IParseNode>> DeserializeFields => new Dictionary<string, Action<Recipient, IParseNode>> {
            {
                "emailAddress", (o,n) => { o.EmailAddress = n.GetObjectValue<EmailAddress>(); }
            },
        };
        /// <summary>
        /// Serialiazes information the current object
        /// <param name="writer">Serialization writer to use to serialize this model</param>
        /// </summary>
        public void Serialize(ISerializationWriter writer) {
            writer.WriteObjectValue<EmailAddress>("emailAddress", EmailAddress);
        }
        /// <summary>Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.</summary>
        public IDictionary<string, object> AdditionalData { get; private set; } = new Dictionary<string, object>();
    }
}
