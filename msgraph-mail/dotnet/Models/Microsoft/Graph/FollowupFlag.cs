using Microsoft.Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
namespace Graphdotnetv4.Models.Microsoft.Graph {
    public class FollowupFlag : IParsable {
        /// <summary>Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.</summary>
        public IDictionary<string, object> AdditionalData { get; set; }
        public DateTimeTimeZone CompletedDateTime { get; set; }
        public DateTimeTimeZone DueDateTime { get; set; }
        public FollowupFlagStatus? FlagStatus { get; set; }
        public DateTimeTimeZone StartDateTime { get; set; }
        /// <summary>
        /// Instantiates a new followupFlag and sets the default values.
        /// </summary>
        public FollowupFlag() {
            AdditionalData = new Dictionary<string, object>();
        }
        public static FollowupFlag CreateFromDiscriminatorValue(IParseNode parseNode) {
            _ = parseNode ?? throw new ArgumentNullException(nameof(parseNode));
            return new FollowupFlag();
        }
        /// <summary>
        /// The deserialization information for the current model
        /// </summary>
        public IDictionary<string, Action<T, IParseNode>> GetFieldDeserializers<T>() {
            return new Dictionary<string, Action<T, IParseNode>> {
                {"completedDateTime", (o,n) => { (o as FollowupFlag).CompletedDateTime = n.GetObjectValue<DateTimeTimeZone>(DateTimeTimeZone.CreateFromDiscriminatorValue); } },
                {"dueDateTime", (o,n) => { (o as FollowupFlag).DueDateTime = n.GetObjectValue<DateTimeTimeZone>(DateTimeTimeZone.CreateFromDiscriminatorValue); } },
                {"flagStatus", (o,n) => { (o as FollowupFlag).FlagStatus = n.GetEnumValue<FollowupFlagStatus>(); } },
                {"startDateTime", (o,n) => { (o as FollowupFlag).StartDateTime = n.GetObjectValue<DateTimeTimeZone>(DateTimeTimeZone.CreateFromDiscriminatorValue); } },
            };
        }
        /// <summary>
        /// Serializes information the current object
        /// <param name="writer">Serialization writer to use to serialize this model</param>
        /// </summary>
        public void Serialize(ISerializationWriter writer) {
            _ = writer ?? throw new ArgumentNullException(nameof(writer));
            writer.WriteObjectValue<DateTimeTimeZone>("completedDateTime", CompletedDateTime);
            writer.WriteObjectValue<DateTimeTimeZone>("dueDateTime", DueDateTime);
            writer.WriteEnumValue<FollowupFlagStatus>("flagStatus", FlagStatus);
            writer.WriteObjectValue<DateTimeTimeZone>("startDateTime", StartDateTime);
            writer.WriteAdditionalData(AdditionalData);
        }
    }
}
