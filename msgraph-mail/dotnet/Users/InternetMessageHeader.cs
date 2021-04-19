using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users {
    public class InternetMessageHeader : IParsable<InternetMessageHeader> {
        /// <summary>Represents the key in a key-value pair.</summary>
        public string Name { get; set; }
        /// <summary>The value in a key-value pair.</summary>
        public string Value { get; set; }
        /// <summary>The serialization information for the current model</summary>
        public IDictionary<string, Action<InternetMessageHeader, IParseNode>> DeserializeFields => new Dictionary<string, Action<InternetMessageHeader, IParseNode>> {
            {
                "name", (o,n) => { o.Name = n.GetStringValue(); }
            },
            {
                "value", (o,n) => { o.Value = n.GetStringValue(); }
            },
        };
        /// <summary>
        /// Serialiazes information the current object
        /// <param name="writer">Serialization writer to use to serialize this model</param>
        /// </summary>
        public void Serialize(ISerializationWriter writer) {
            writer.WriteStringValue("name", Name);
            writer.WriteStringValue("value", Value);
        }
    }
}
