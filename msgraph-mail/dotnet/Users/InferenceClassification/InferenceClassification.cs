using Microsoft.Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users.InferenceClassification {
    public class InferenceClassification : Entity, IParsable {
        /// <summary>A set of overrides for a user to always classify messages from specific senders in certain ways: focused, or other. Read-only. Nullable.</summary>
        public List<InferenceClassificationOverride> Overrides { get; set; }
        /// <summary>
        /// The deserialization information for the current model
        /// </summary>
        public new IDictionary<string, Action<T, IParseNode>> GetFieldDeserializers() {
            return new Dictionary<string, Action<T, IParseNode>>(base.GetFieldDeserializers()) {
                {"overrides", (o,n) => { (o as InferenceClassification).Overrides = n.GetCollectionOfObjectValues<InferenceClassificationOverride>().ToList(); } },
            };
        }
        /// <summary>
        /// Serialiazes information the current object
        /// <param name="writer">Serialization writer to use to serialize this model</param>
        /// </summary>
        public new void Serialize(ISerializationWriter writer) {
            base.Serialize(writer);
            writer.WriteCollectionOfObjectValues<InferenceClassificationOverride>("overrides", Overrides);
        }
    }
}
