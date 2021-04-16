using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users {
    public class InternetMessageHeader : IParsable<InternetMessageHeader> {
        public string Name { get; set; }
        public string Value { get; set; }
        public IDictionary<string, Action<InternetMessageHeader, IParseNode>> DeserializeFields => new Dictionary<string, Action<InternetMessageHeader, IParseNode>> {
            {
                "name", (o,n) => { o.Name = n.GetStringValue(); }
            },
            {
                "value", (o,n) => { o.Value = n.GetStringValue(); }
            },
        };
        public void Serialize(ISerializationWriter writer) {
            writer.WriteStringValue("name", Name);
            writer.WriteStringValue("value", Value);
        }
    }
}
