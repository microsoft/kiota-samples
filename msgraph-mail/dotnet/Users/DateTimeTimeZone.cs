using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users {
    public class DateTimeTimeZone : IParsable<DateTimeTimeZone> {
        /// <summary>A single point of time in a combined date and time representation ({date}T{time}). For example, '2019-04-16T09:00:00'.</summary>
        public string DateTime { get; set; }
        /// <summary>Represents a time zone, for example, 'Pacific Standard Time'. See below for possible values.</summary>
        public string TimeZone { get; set; }
        /// <summary>The serialization information for the current model</summary>
        public IDictionary<string, Action<DateTimeTimeZone, IParseNode>> DeserializeFields => new Dictionary<string, Action<DateTimeTimeZone, IParseNode>> {
            {
                "dateTime", (o,n) => { o.DateTime = n.GetStringValue(); }
            },
            {
                "timeZone", (o,n) => { o.TimeZone = n.GetStringValue(); }
            },
        };
        /// <summary>
        /// Serialiazes information the current object
        /// <param name="writer">Serialization writer to use to serialize this model</param>
        /// </summary>
        public void Serialize(ISerializationWriter writer) {
            writer.WriteStringValue("dateTime", DateTime);
            writer.WriteStringValue("timeZone", TimeZone);
        }
        /// <summary>Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.</summary>
        public IDictionary<string, object> AdditionalData { get; private set; } = new Dictionary<string, object>();
    }
}
