using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users.InferenceClassification {
    public class InferenceClassificationOverride : Entity, IParsable<InferenceClassificationOverride> {
        public InferenceClassificationType? ClassifyAs { get; set; }
        public EmailAddress SenderEmailAddress { get; set; }
        /// <summary>The serialization information for the current model</summary>
        public new IDictionary<string, Action<InferenceClassificationOverride, IParseNode>> DeserializeFields => new Dictionary<string, Action<InferenceClassificationOverride, IParseNode>> {
            {
                "classifyAs", (o,n) => { o.ClassifyAs = n.GetEnumValue<InferenceClassificationType>(); }
            },
            {
                "senderEmailAddress", (o,n) => { o.SenderEmailAddress = n.GetObjectValue<EmailAddress>(); }
            },
        };
        /// <summary>
        /// Serialiazes information the current object
        /// <param name="writer">Serialization writer to use to serialize this model</param>
        /// </summary>
        public new void Serialize(ISerializationWriter writer) {
            base.Serialize(writer);
            writer.WriteEnumValue<InferenceClassificationType>("classifyAs", ClassifyAs);
            writer.WriteObjectValue<EmailAddress>("senderEmailAddress", SenderEmailAddress);
        }
    }
}
