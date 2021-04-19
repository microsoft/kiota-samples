using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users.MailFolders {
    public class SingleValueExtendedPropertiesResponse : IParsable<SingleValueExtendedPropertiesResponse> {
        public List<SingleValueLegacyExtendedProperty> Value { get; set; }
        public string NextLink { get; set; }
        public IDictionary<string, Action<SingleValueExtendedPropertiesResponse, IParseNode>> DeserializeFields => new Dictionary<string, Action<SingleValueExtendedPropertiesResponse, IParseNode>> {
            {
                "value", (o,n) => { o.Value = n.GetCollectionOfObjectValues<SingleValueLegacyExtendedProperty>().ToList(); }
            },
            {
                "nextLink", (o,n) => { o.NextLink = n.GetStringValue(); }
            },
        };
        public void Serialize(ISerializationWriter writer) {
            writer.WriteCollectionOfObjectValues<SingleValueLegacyExtendedProperty>("value", Value);
            writer.WriteStringValue("nextLink", NextLink);
        }
    }
}
