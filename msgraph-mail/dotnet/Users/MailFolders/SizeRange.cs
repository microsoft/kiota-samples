using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users.MailFolders {
    public class SizeRange : IParsable<SizeRange> {
        public int? MaximumSize { get; set; }
        public int? MinimumSize { get; set; }
        public IDictionary<string, Action<SizeRange, IParseNode>> DeserializeFields => new Dictionary<string, Action<SizeRange, IParseNode>> {
            {
                "maximumSize", (o,n) => { o.MaximumSize = n.GetIntValue(); }
            },
            {
                "minimumSize", (o,n) => { o.MinimumSize = n.GetIntValue(); }
            },
        };
        public void Serialize(ISerializationWriter writer) {
            writer.WriteIntValue("maximumSize", MaximumSize);
            writer.WriteIntValue("minimumSize", MinimumSize);
        }
    }
}
