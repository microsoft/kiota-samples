using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users {
    public class EmailAddress : IParsable<EmailAddress> {
        public string Address { get; set; }
        public string Name { get; set; }
        public IDictionary<string, Action<EmailAddress, IParseNode>> DeserializeFields => new Dictionary<string, Action<EmailAddress, IParseNode>> {
            {
                "address", (o,n) => { o.Address = n.GetStringValue(); }
            },
            {
                "name", (o,n) => { o.Name = n.GetStringValue(); }
            },
        };
        public void Serialize(ISerializationWriter writer) {
            writer.WriteStringValue("address", Address);
            writer.WriteStringValue("name", Name);
        }
    }
}
