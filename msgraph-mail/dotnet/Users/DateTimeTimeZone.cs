using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users {
    public class DateTimeTimeZone : IParsable<DateTimeTimeZone> {
        /// <summary>Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.</summary>
        public IDictionary<string, object> AdditionalData { get; private set; } = new Dictionary<string, object>();
        /// <summary>A single point of time in a combined date and time representation ({date}T{time}; for example, 2017-08-29T04:00:00.0000000).</summary>
        public string DateTime { get; set; }
        /// <summary>The serialization information for the current model</summary>
        public IDictionary<string, Action<DateTimeTimeZone, IParseNode>> DeserializeFields => new Dictionary<string, Action<DateTimeTimeZone, IParseNode>> {
            {
                "dateTime", (o,n) => { o.DateTime = n.GetStringValue(); }
            },
            {
                "timeZone", (o,n) => { o.TimeZone = n.GetStringValue(); }
            },
        };
        /// <summary>Represents a time zone, for example, 'Pacific Standard Time'. See below for more possible values.</summary>
        public string TimeZone { get; set; }
        /// <summary>
        /// Serialiazes information the current object
        /// <param name="writer">Serialization writer to use to serialize this model</param>
        /// </summary>
        public void Serialize(ISerializationWriter writer) {
            writer.WriteStringValue("dateTime", DateTime);
            writer.WriteStringValue("timeZone", TimeZone);
            writer.WriteAdditionalData(AdditionalData);
        }
    }
}
