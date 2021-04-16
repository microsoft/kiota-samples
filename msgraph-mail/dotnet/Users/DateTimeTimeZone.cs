using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users {
    public class DateTimeTimeZone : IParsable<DateTimeTimeZone> {
        public string DateTime { get; set; }
        public string TimeZone { get; set; }
        public IDictionary<string, Action<DateTimeTimeZone, IParseNode>> DeserializeFields => new Dictionary<string, Action<DateTimeTimeZone, IParseNode>> {
            {
                "dateTime", (o,n) => { o.DateTime = n.GetStringValue(); }
            },
            {
                "timeZone", (o,n) => { o.TimeZone = n.GetStringValue(); }
            },
        };
        public void Serialize(ISerializationWriter writer) {
            writer.WriteStringValue("dateTime", DateTime);
            writer.WriteStringValue("timeZone", TimeZone);
        }
    }
}
