using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users {
    public class Entity : IParsable<Entity> {
        /// <summary>Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.</summary>
        public IDictionary<string, object> AdditionalData { get; private set; } = new Dictionary<string, object>();
        /// <summary>The serialization information for the current model</summary>
        public IDictionary<string, Action<Entity, IParseNode>> DeserializeFields => new Dictionary<string, Action<Entity, IParseNode>> {
            {
                "id", (o,n) => { o.Id = n.GetStringValue(); }
            },
        };
        /// <summary>Read-only.</summary>
        public string Id { get; set; }
        /// <summary>
        /// Serialiazes information the current object
        /// <param name="writer">Serialization writer to use to serialize this model</param>
        /// </summary>
        public void Serialize(ISerializationWriter writer) {
            writer.WriteStringValue("id", Id);
            writer.WriteAdditionalData(AdditionalData);
        }
    }
}
