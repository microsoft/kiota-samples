using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users {
    public class InferenceClassificationType : IParsable<InferenceClassificationType> {
        public IDictionary<string, Action<InferenceClassificationType, IParseNode>> DeserializeFields => new Dictionary<string, Action<InferenceClassificationType, IParseNode>> {
        };
        public void Serialize(ISerializationWriter writer) {
        }
    }
}
