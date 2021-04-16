using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users {
    public class FollowupFlag : IParsable<FollowupFlag> {
        public DateTimeTimeZone CompletedDateTime { get; set; }
        public DateTimeTimeZone DueDateTime { get; set; }
        public FollowupFlagStatus FlagStatus { get; set; }
        public DateTimeTimeZone StartDateTime { get; set; }
        public IDictionary<string, Action<FollowupFlag, IParseNode>> DeserializeFields => new Dictionary<string, Action<FollowupFlag, IParseNode>> {
            {
                "completedDateTime", (o,n) => { o.CompletedDateTime = n.GetObjectValue<DateTimeTimeZone>(); }
            },
            {
                "dueDateTime", (o,n) => { o.DueDateTime = n.GetObjectValue<DateTimeTimeZone>(); }
            },
            {
                "flagStatus", (o,n) => { o.FlagStatus = n.GetObjectValue<FollowupFlagStatus>(); }
            },
            {
                "startDateTime", (o,n) => { o.StartDateTime = n.GetObjectValue<DateTimeTimeZone>(); }
            },
        };
        public void Serialize(ISerializationWriter writer) {
            writer.WriteObjectValue<DateTimeTimeZone>("completedDateTime", CompletedDateTime);
            writer.WriteObjectValue<DateTimeTimeZone>("dueDateTime", DueDateTime);
            writer.WriteObjectValue<FollowupFlagStatus>("flagStatus", FlagStatus);
            writer.WriteObjectValue<DateTimeTimeZone>("startDateTime", StartDateTime);
        }
    }
}
