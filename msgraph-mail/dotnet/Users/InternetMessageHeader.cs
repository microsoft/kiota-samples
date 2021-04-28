using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users {
    public class InternetMessageHeader : IParsable<InternetMessageHeader> {
        /// <summary>Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.</summary>
        public IDictionary<string, object> AdditionalData { get; private set; } = new Dictionary<string, object>();
        /// <summary>The serialization information for the current model</summary>
        public IDictionary<string, Action<InternetMessageHeader, IParseNode>> DeserializeFields => new Dictionary<string, Action<InternetMessageHeader, IParseNode>> {
            {
                "name", (o,n) => { o.Name = n.GetStringValue(); }
            },
            {
                "value", (o,n) => { o.Value = n.GetStringValue(); }
            },
        };
        /// <summary>Represents the key in a key-value pair.</summary>
        public string Name { get; set; }
        /// <summary>The value in a key-value pair.</summary>
        public string Value { get; set; }
        /// <summary>
        /// Serialiazes information the current object
        /// <param name="writer">Serialization writer to use to serialize this model</param>
        /// </summary>
        public void Serialize(ISerializationWriter writer) {
            writer.WriteStringValue("name", Name);
            writer.WriteStringValue("value", Value);
            writer.WriteAdditionalData(AdditionalData);
        }
    }
}
