using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users.InferenceClassification {
    public class InferenceClassification : Entity, IParsable<InferenceClassification> {
        public List<InferenceClassificationOverride> Overrides { get; set; }
        public new IDictionary<string, Action<InferenceClassification, IParseNode>> DeserializeFields => new Dictionary<string, Action<InferenceClassification, IParseNode>> {
            {
                "overrides", (o,n) => { o.Overrides = n.GetCollectionOfObjectValues<InferenceClassificationOverride>().ToList(); }
            },
        };
        public new void Serialize(ISerializationWriter writer) {
            base.Serialize(writer);
            writer.WriteCollectionOfObjectValues<InferenceClassificationOverride>("overrides", Overrides);
        }
    }
}
