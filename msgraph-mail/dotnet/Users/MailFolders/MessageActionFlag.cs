using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users.MailFolders {
    public class MessageActionFlag : IParsable<MessageActionFlag> {
        public IDictionary<string, Action<MessageActionFlag, IParseNode>> DeserializeFields => new Dictionary<string, Action<MessageActionFlag, IParseNode>> {
        };
        public void Serialize(ISerializationWriter writer) {
        }
    }
}
