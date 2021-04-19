using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users {
    public class FollowupFlagStatus : IParsable<FollowupFlagStatus> {
        public IDictionary<string, Action<FollowupFlagStatus, IParseNode>> DeserializeFields => new Dictionary<string, Action<FollowupFlagStatus, IParseNode>> {
        };
        public void Serialize(ISerializationWriter writer) {
        }
    }
}
