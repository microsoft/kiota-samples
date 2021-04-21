using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users.MailFolders {
    public class SizeRange : IParsable<SizeRange> {
        /// <summary>The maximum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.</summary>
        public int? MaximumSize { get; set; }
        /// <summary>The minimum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.</summary>
        public int? MinimumSize { get; set; }
        /// <summary>The serialization information for the current model</summary>
        public IDictionary<string, Action<SizeRange, IParseNode>> DeserializeFields => new Dictionary<string, Action<SizeRange, IParseNode>> {
            {
                "maximumSize", (o,n) => { o.MaximumSize = n.GetIntValue(); }
            },
            {
                "minimumSize", (o,n) => { o.MinimumSize = n.GetIntValue(); }
            },
        };
        /// <summary>
        /// Serialiazes information the current object
        /// <param name="writer">Serialization writer to use to serialize this model</param>
        /// </summary>
        public void Serialize(ISerializationWriter writer) {
            writer.WriteIntValue("maximumSize", MaximumSize);
            writer.WriteIntValue("minimumSize", MinimumSize);
            writer.WriteAdditionalData(AdditionalData);
        }
        /// <summary>Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.</summary>
        public IDictionary<string, object> AdditionalData { get; private set; } = new Dictionary<string, object>();
    }
}
