using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users.InferenceClassification {
    public class OverridesResponse : IParsable<OverridesResponse> {
        public List<InferenceClassificationOverride> Value { get; set; }
        public string NextLink { get; set; }
        public IDictionary<string, Action<OverridesResponse, IParseNode>> DeserializeFields => new Dictionary<string, Action<OverridesResponse, IParseNode>> {
            {
                "value", (o,n) => { o.Value = n.GetCollectionOfObjectValues<InferenceClassificationOverride>().ToList(); }
            },
            {
                "nextLink", (o,n) => { o.NextLink = n.GetStringValue(); }
            },
        };
        public void Serialize(ISerializationWriter writer) {
            writer.WriteCollectionOfObjectValues<InferenceClassificationOverride>("value", Value);
            writer.WriteStringValue("nextLink", NextLink);
        }
    }
}
