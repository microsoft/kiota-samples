using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users.InferenceClassification {
    public class InferenceClassificationOverride : Entity, IParsable<InferenceClassificationOverride> {
        public InferenceClassificationType ClassifyAs { get; set; }
        public EmailAddress SenderEmailAddress { get; set; }
        public new IDictionary<string, Action<InferenceClassificationOverride, IParseNode>> DeserializeFields => new Dictionary<string, Action<InferenceClassificationOverride, IParseNode>> {
            {
                "classifyAs", (o,n) => { o.ClassifyAs = n.GetObjectValue<InferenceClassificationType>(); }
            },
            {
                "senderEmailAddress", (o,n) => { o.SenderEmailAddress = n.GetObjectValue<EmailAddress>(); }
            },
        };
        public new void Serialize(ISerializationWriter writer) {
            base.Serialize(writer);
            writer.WriteObjectValue<InferenceClassificationType>("classifyAs", ClassifyAs);
            writer.WriteObjectValue<EmailAddress>("senderEmailAddress", SenderEmailAddress);
        }
    }
}
