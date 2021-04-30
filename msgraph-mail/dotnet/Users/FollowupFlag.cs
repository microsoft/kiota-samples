using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users {
    public class FollowupFlag : IParsable<FollowupFlag> {
        /// <summary>Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.</summary>
        public IDictionary<string, object> AdditionalData { get; private set; } = new Dictionary<string, object>();
        public DateTimeTimeZone CompletedDateTime { get; set; }
        /// <summary>The serialization information for the current model</summary>
        public IDictionary<string, Action<FollowupFlag, IParseNode>> DeserializeFields => new Dictionary<string, Action<FollowupFlag, IParseNode>> {
            {
                "completedDateTime", (o,n) => { o.CompletedDateTime = n.GetObjectValue<DateTimeTimeZone>(); }
            },
            {
                "dueDateTime", (o,n) => { o.DueDateTime = n.GetObjectValue<DateTimeTimeZone>(); }
            },
            {
                "flagStatus", (o,n) => { o.FlagStatus = n.GetEnumValue<FollowupFlagStatus>(); }
            },
            {
                "startDateTime", (o,n) => { o.StartDateTime = n.GetObjectValue<DateTimeTimeZone>(); }
            },
        };
        public DateTimeTimeZone DueDateTime { get; set; }
        public FollowupFlagStatus? FlagStatus { get; set; }
        public DateTimeTimeZone StartDateTime { get; set; }
        /// <summary>
        /// Serialiazes information the current object
        /// <param name="writer">Serialization writer to use to serialize this model</param>
        /// </summary>
        public void Serialize(ISerializationWriter writer) {
            writer.WriteObjectValue<DateTimeTimeZone>("completedDateTime", CompletedDateTime);
            writer.WriteObjectValue<DateTimeTimeZone>("dueDateTime", DueDateTime);
            writer.WriteEnumValue<FollowupFlagStatus>("flagStatus", FlagStatus);
            writer.WriteObjectValue<DateTimeTimeZone>("startDateTime", StartDateTime);
            writer.WriteAdditionalData(AdditionalData);
        }
    }
}
