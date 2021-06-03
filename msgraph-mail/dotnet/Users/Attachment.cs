using Microsoft.Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users {
    public class Attachment : Entity, IParsable<Attachment> {
        /// <summary>The MIME type.</summary>
        public string ContentType { get; set; }
        /// <summary>The serialization information for the current model</summary>
        public new IDictionary<string, Action<Attachment, IParseNode>> DeserializeFields => new Dictionary<string, Action<Attachment, IParseNode>> {
            {
                "contentType", (o,n) => { o.ContentType = n.GetStringValue(); }
            },
            {
                "isInline", (o,n) => { o.IsInline = n.GetBoolValue(); }
            },
            {
                "lastModifiedDateTime", (o,n) => { o.LastModifiedDateTime = n.GetDateTimeOffsetValue(); }
            },
            {
                "name", (o,n) => { o.Name = n.GetStringValue(); }
            },
            {
                "size", (o,n) => { o.Size = n.GetIntValue(); }
            },
        };
        /// <summary>true if the attachment is an inline attachment; otherwise, false.</summary>
        public bool? IsInline { get; set; }
        /// <summary>The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z</summary>
        public DateTimeOffset? LastModifiedDateTime { get; set; }
        /// <summary>The display name of the attachment. This does not need to be the actual file name.</summary>
        public string Name { get; set; }
        /// <summary>The length of the attachment in bytes.</summary>
        public int? Size { get; set; }
        /// <summary>
        /// Serialiazes information the current object
        /// <param name="writer">Serialization writer to use to serialize this model</param>
        /// </summary>
        public new void Serialize(ISerializationWriter writer) {
            base.Serialize(writer);
            writer.WriteStringValue("contentType", ContentType);
            writer.WriteBoolValue("isInline", IsInline);
            writer.WriteDateTimeOffsetValue("lastModifiedDateTime", LastModifiedDateTime);
            writer.WriteStringValue("name", Name);
            writer.WriteIntValue("size", Size);
        }
    }
}
