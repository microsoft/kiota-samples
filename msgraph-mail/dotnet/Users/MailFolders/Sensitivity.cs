using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users.MailFolders {
    public class Sensitivity : IParsable<Sensitivity> {
        public IDictionary<string, Action<Sensitivity, IParseNode>> DeserializeFields => new Dictionary<string, Action<Sensitivity, IParseNode>> {
        };
        public void Serialize(ISerializationWriter writer) {
        }
    }
}
