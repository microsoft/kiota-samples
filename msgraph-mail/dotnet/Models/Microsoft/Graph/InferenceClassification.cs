using Microsoft.Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
namespace Graphdotnetv4.Models.Microsoft.Graph {
    public class InferenceClassification : Entity, IParsable {
        /// <summary>A set of overrides for a user to always classify messages from specific senders in certain ways: focused, or other. Read-only. Nullable.</summary>
        public List<InferenceClassificationOverride> Overrides { get; set; }
        public static new Graphdotnetv4.Models.Microsoft.Graph.InferenceClassification CreateFromDiscriminatorValue(IParseNode parseNode) {
            _ = parseNode ?? throw new ArgumentNullException(nameof(parseNode));
            return new InferenceClassification();
        }
        /// <summary>
        /// The deserialization information for the current model
        /// </summary>
        public new IDictionary<string, Action<T, IParseNode>> GetFieldDeserializers<T>() {
            return new Dictionary<string, Action<T, IParseNode>>(base.GetFieldDeserializers<T>()) {
                {"overrides", (o,n) => { (o as InferenceClassification).Overrides = n.GetCollectionOfObjectValues<InferenceClassificationOverride>(InferenceClassificationOverride.CreateFromDiscriminatorValue).ToList(); } },
            };
        }
        /// <summary>
        /// Serializes information the current object
        /// <param name="writer">Serialization writer to use to serialize this model</param>
        /// </summary>
        public new void Serialize(ISerializationWriter writer) {
            _ = writer ?? throw new ArgumentNullException(nameof(writer));
            base.Serialize(writer);
            writer.WriteCollectionOfObjectValues<InferenceClassificationOverride>("overrides", Overrides);
        }
    }
}
