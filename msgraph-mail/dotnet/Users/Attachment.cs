using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users {
    public class Attachment : Entity, IParsable<Attachment> {
        public string ContentType { get; set; }
        public bool IsInline { get; set; }
        public string LastModifiedDateTime { get; set; }
        public string Name { get; set; }
        public int Size { get; set; }
        public new IDictionary<string, Action<Attachment, IParseNode>> DeserializeFields => new Dictionary<string, Action<Attachment, IParseNode>> {
            {
                "contentType", (o,n) => { o.ContentType = n.GetStringValue(); }
            },
            {
                "isInline", (o,n) => { o.IsInline = n.GetBoolValue(); }
            },
            {
                "lastModifiedDateTime", (o,n) => { o.LastModifiedDateTime = n.GetStringValue(); }
            },
            {
                "name", (o,n) => { o.Name = n.GetStringValue(); }
            },
            {
                "size", (o,n) => { o.Size = n.GetIntValue(); }
            },
        };
        public new void Serialize(ISerializationWriter writer) {
            base.Serialize(writer);
            writer.WriteStringValue("contentType", ContentType);
            writer.WriteBoolValue("isInline", IsInline);
            writer.WriteStringValue("lastModifiedDateTime", LastModifiedDateTime);
            writer.WriteStringValue("name", Name);
            writer.WriteIntValue("size", Size);
        }
    }
}
