using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users {
    public class Recipient : IParsable<Recipient> {
        public EmailAddress EmailAddress { get; set; }
        public IDictionary<string, Action<Recipient, IParseNode>> DeserializeFields => new Dictionary<string, Action<Recipient, IParseNode>> {
            {
                "emailAddress", (o,n) => { o.EmailAddress = n.GetObjectValue<EmailAddress>(); }
            },
        };
        public void Serialize(ISerializationWriter writer) {
            writer.WriteObjectValue<EmailAddress>("emailAddress", EmailAddress);
        }
    }
}
