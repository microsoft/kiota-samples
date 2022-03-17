using Microsoft.Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
namespace Graphdotnetv4.Models.Odata.Error {
    public class Main : IAdditionalDataHolder, IParsable {
        /// <summary>Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.</summary>
        public IDictionary<string, object> AdditionalData { get; set; }
        public string Code { get; set; }
        public List<Detail> Details { get; set; }
        /// <summary>The structure of this object is service-specific</summary>
        public Graphdotnetv4.Models.Odata.Error.Innererror Innererror { get; set; }
        public string Message { get; set; }
        public string Target { get; set; }
        /// <summary>
        /// Instantiates a new main and sets the default values.
        /// </summary>
        public Main() {
            AdditionalData = new Dictionary<string, object>();
        }
        /// <summary>
        /// Creates a new instance of the appropriate class based on discriminator value
        /// <param name="parseNode">The parse node to use to read the discriminator value and create the object</param>
        /// </summary>
        public static Main CreateFromDiscriminatorValue(IParseNode parseNode) {
            _ = parseNode ?? throw new ArgumentNullException(nameof(parseNode));
            return new Main();
        }
        /// <summary>
        /// The deserialization information for the current model
        /// </summary>
        public IDictionary<string, Action<T, IParseNode>> GetFieldDeserializers<T>() {
            return new Dictionary<string, Action<T, IParseNode>> {
                {"code", (o,n) => { (o as Main).Code = n.GetStringValue(); } },
                {"details", (o,n) => { (o as Main).Details = n.GetCollectionOfObjectValues<Detail>(Detail.CreateFromDiscriminatorValue).ToList(); } },
                {"innererror", (o,n) => { (o as Main).Innererror = n.GetObjectValue<Graphdotnetv4.Models.Odata.Error.Innererror>(Graphdotnetv4.Models.Odata.Error.Innererror.CreateFromDiscriminatorValue); } },
                {"message", (o,n) => { (o as Main).Message = n.GetStringValue(); } },
                {"target", (o,n) => { (o as Main).Target = n.GetStringValue(); } },
            };
        }
        /// <summary>
        /// Serializes information the current object
        /// <param name="writer">Serialization writer to use to serialize this model</param>
        /// </summary>
        public void Serialize(ISerializationWriter writer) {
            _ = writer ?? throw new ArgumentNullException(nameof(writer));
            writer.WriteStringValue("code", Code);
            writer.WriteCollectionOfObjectValues<Detail>("details", Details);
            writer.WriteObjectValue<Graphdotnetv4.Models.Odata.Error.Innererror>("innererror", Innererror);
            writer.WriteStringValue("message", Message);
            writer.WriteStringValue("target", Target);
            writer.WriteAdditionalData(AdditionalData);
        }
    }
}
