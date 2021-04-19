using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users.MailFolders {
    public class MultiValueExtendedPropertiesResponse : IParsable<MultiValueExtendedPropertiesResponse> {
        public List<MultiValueLegacyExtendedProperty> Value { get; set; }
        public string NextLink { get; set; }
        public IDictionary<string, Action<MultiValueExtendedPropertiesResponse, IParseNode>> DeserializeFields => new Dictionary<string, Action<MultiValueExtendedPropertiesResponse, IParseNode>> {
            {
                "value", (o,n) => { o.Value = n.GetCollectionOfObjectValues<MultiValueLegacyExtendedProperty>().ToList(); }
            },
            {
                "nextLink", (o,n) => { o.NextLink = n.GetStringValue(); }
            },
        };
        public void Serialize(ISerializationWriter writer) {
            writer.WriteCollectionOfObjectValues<MultiValueLegacyExtendedProperty>("value", Value);
            writer.WriteStringValue("nextLink", NextLink);
        }
    }
}
