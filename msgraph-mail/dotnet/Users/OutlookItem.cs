using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users {
    public class OutlookItem : Entity, IParsable<OutlookItem> {
        public List<object> Categories { get; set; }
        public string ChangeKey { get; set; }
        public string CreatedDateTime { get; set; }
        public string LastModifiedDateTime { get; set; }
        public new IDictionary<string, Action<OutlookItem, IParseNode>> DeserializeFields => new Dictionary<string, Action<OutlookItem, IParseNode>> {
            {
                "categories", (o,n) => { o.Categories = n.GetCollectionOfPrimitiveValues<Object>().ToList(); }
            },
            {
                "changeKey", (o,n) => { o.ChangeKey = n.GetStringValue(); }
            },
            {
                "createdDateTime", (o,n) => { o.CreatedDateTime = n.GetStringValue(); }
            },
            {
                "lastModifiedDateTime", (o,n) => { o.LastModifiedDateTime = n.GetStringValue(); }
            },
        };
        public new void Serialize(ISerializationWriter writer) {
            base.Serialize(writer);
            writer.WriteCollectionOfPrimitiveValues<Object>("categories", Categories);
            writer.WriteStringValue("changeKey", ChangeKey);
            writer.WriteStringValue("createdDateTime", CreatedDateTime);
            writer.WriteStringValue("lastModifiedDateTime", LastModifiedDateTime);
        }
    }
}
