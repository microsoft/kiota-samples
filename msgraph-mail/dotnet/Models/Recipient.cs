using Microsoft.Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
namespace Graphdotnetv4.Models {
    public class Recipient : IAdditionalDataHolder, IParsable {
        /// <summary>Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.</summary>
        public IDictionary<string, object> AdditionalData { get; set; }
        /// <summary>The emailAddress property</summary>
#if NETSTANDARD2_1_OR_GREATER || NET6_0_OR_GREATER
        public Graphdotnetv4.Models.EmailAddress? EmailAddress { get; set; }
#else
        public Graphdotnetv4.Models.EmailAddress EmailAddress { get; set; }
#endif
        /// <summary>
        /// Instantiates a new recipient and sets the default values.
        /// </summary>
        public Recipient() {
            AdditionalData = new Dictionary<string, object>();
        }
        /// <summary>
        /// Creates a new instance of the appropriate class based on discriminator value
        /// </summary>
        /// <param name="parseNode">The parse node to use to read the discriminator value and create the object</param>
        public static Recipient CreateFromDiscriminatorValue(IParseNode parseNode) {
            _ = parseNode ?? throw new ArgumentNullException(nameof(parseNode));
            return new Recipient();
        }
        /// <summary>
        /// The deserialization information for the current model
        /// </summary>
        public IDictionary<string, Action<IParseNode>> GetFieldDeserializers() {
            return new Dictionary<string, Action<IParseNode>> {
                {"emailAddress", n => { EmailAddress = n.GetObjectValue<Graphdotnetv4.Models.EmailAddress>(Graphdotnetv4.Models.EmailAddress.CreateFromDiscriminatorValue); } },
            };
        }
        /// <summary>
        /// Serializes information the current object
        /// </summary>
        /// <param name="writer">Serialization writer to use to serialize this model</param>
        public void Serialize(ISerializationWriter writer) {
            _ = writer ?? throw new ArgumentNullException(nameof(writer));
#if NETSTANDARD2_1_OR_GREATER || NET6_0_OR_GREATER
            writer.WriteObjectValue<Graphdotnetv4.Models.EmailAddress?>("emailAddress", EmailAddress);
#else
            writer.WriteObjectValue<Graphdotnetv4.Models.EmailAddress>("emailAddress", EmailAddress);
#endif
            writer.WriteAdditionalData(AdditionalData);
        }
    }
}
