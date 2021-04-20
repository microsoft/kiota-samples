using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users {
    public class OutlookItem : Entity, IParsable<OutlookItem> {
        /// <summary>The categories associated with the item</summary>
        public List<string> Categories { get; set; }
        /// <summary>Identifies the version of the item. Every time the item is changed, changeKey changes as well. This allows Exchange to apply changes to the correct version of the object. Read-only.</summary>
        public string ChangeKey { get; set; }
        /// <summary>The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z</summary>
        public DateTimeOffset? CreatedDateTime { get; set; }
        /// <summary>The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z</summary>
        public DateTimeOffset? LastModifiedDateTime { get; set; }
        /// <summary>The serialization information for the current model</summary>
        public new IDictionary<string, Action<OutlookItem, IParseNode>> DeserializeFields => new Dictionary<string, Action<OutlookItem, IParseNode>> {
            {
                "categories", (o,n) => { o.Categories = n.GetCollectionOfPrimitiveValues<string>().ToList(); }
            },
            {
                "changeKey", (o,n) => { o.ChangeKey = n.GetStringValue(); }
            },
            {
                "createdDateTime", (o,n) => { o.CreatedDateTime = n.GetDateTimeOffsetValue(); }
            },
            {
                "lastModifiedDateTime", (o,n) => { o.LastModifiedDateTime = n.GetDateTimeOffsetValue(); }
            },
        };
        /// <summary>
        /// Serialiazes information the current object
        /// <param name="writer">Serialization writer to use to serialize this model</param>
        /// </summary>
        public new void Serialize(ISerializationWriter writer) {
            base.Serialize(writer);
            writer.WriteCollectionOfPrimitiveValues<string>("categories", Categories);
            writer.WriteStringValue("changeKey", ChangeKey);
            writer.WriteDateTimeOffsetValue("createdDateTime", CreatedDateTime);
            writer.WriteDateTimeOffsetValue("lastModifiedDateTime", LastModifiedDateTime);
        }
    }
}
