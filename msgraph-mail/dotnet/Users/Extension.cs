using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users {
    public class Extension : Entity, IParsable<Extension> {
        /// <summary>The serialization information for the current model</summary>
        public new IDictionary<string, Action<Extension, IParseNode>> DeserializeFields => new Dictionary<string, Action<Extension, IParseNode>> {
        };
        /// <summary>
        /// Serialiazes information the current object
        /// <param name="writer">Serialization writer to use to serialize this model</param>
        /// </summary>
        public new void Serialize(ISerializationWriter writer) {
            base.Serialize(writer);
        }
    }
}
