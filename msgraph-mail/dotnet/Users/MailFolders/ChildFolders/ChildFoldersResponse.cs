using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.Linq;
namespace Graphdotnetv4.Users.MailFolders.ChildFolders {
    public class ChildFoldersResponse : IParsable<ChildFoldersResponse> {
        /// <summary>Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.</summary>
        public IDictionary<string, object> AdditionalData { get; private set; } = new Dictionary<string, object>();
        /// <summary>The serialization information for the current model</summary>
        public IDictionary<string, Action<ChildFoldersResponse, IParseNode>> DeserializeFields => new Dictionary<string, Action<ChildFoldersResponse, IParseNode>> {
            {
                "nextLink", (o,n) => { o.NextLink = n.GetStringValue(); }
            },
            {
                "value", (o,n) => { o.Value = n.GetCollectionOfObjectValues<MailFolder>().ToList(); }
            },
        };
        public string NextLink { get; set; }
        public List<MailFolder> Value { get; set; }
        /// <summary>
        /// Serialiazes information the current object
        /// <param name="writer">Serialization writer to use to serialize this model</param>
        /// </summary>
        public void Serialize(ISerializationWriter writer) {
            writer.WriteStringValue("nextLink", NextLink);
            writer.WriteCollectionOfObjectValues<MailFolder>("value", Value);
            writer.WriteAdditionalData(AdditionalData);
        }
    }
}
