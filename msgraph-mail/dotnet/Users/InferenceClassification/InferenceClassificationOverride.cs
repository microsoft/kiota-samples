using Microsoft.Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users.InferenceClassification {
    public class InferenceClassificationOverride : Entity, IParsable {
        public InferenceClassificationType? ClassifyAs { get; set; }
        public EmailAddress SenderEmailAddress { get; set; }
        /// <summary>
        /// The deserialization information for the current model
        /// </summary>
        public new IDictionary<string, Action<T, IParseNode>> GetFieldDeserializers<T>() {
            return new Dictionary<string, Action<T, IParseNode>>(base.GetFieldDeserializers<T>()) {
                {"classifyAs", (o,n) => { (o as InferenceClassificationOverride).ClassifyAs = n.GetEnumValue<InferenceClassificationType>(); } },
                {"senderEmailAddress", (o,n) => { (o as InferenceClassificationOverride).SenderEmailAddress = n.GetObjectValue<EmailAddress>(); } },
            };
        }
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
