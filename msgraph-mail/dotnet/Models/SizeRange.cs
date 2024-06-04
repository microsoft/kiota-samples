// <auto-generated/>
using Microsoft.Kiota.Abstractions.Serialization;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System;
namespace Graphdotnetv4.Models
{
    #pragma warning disable CS1591
    public class SizeRange : IAdditionalDataHolder, IParsable
    #pragma warning restore CS1591
    {
        /// <summary>Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.</summary>
        public IDictionary<string, object> AdditionalData { get; set; }
        /// <summary>The maximum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.</summary>
        public int? MaximumSize { get; set; }
        /// <summary>The minimum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.</summary>
        public int? MinimumSize { get; set; }
        /// <summary>
        /// Instantiates a new <see cref="Graphdotnetv4.Models.SizeRange"/> and sets the default values.
        /// </summary>
        public SizeRange()
        {
            AdditionalData = new Dictionary<string, object>();
        }
        /// <summary>
        /// Creates a new instance of the appropriate class based on discriminator value
        /// </summary>
        /// <returns>A <see cref="Graphdotnetv4.Models.SizeRange"/></returns>
        /// <param name="parseNode">The parse node to use to read the discriminator value and create the object</param>
        public static Graphdotnetv4.Models.SizeRange CreateFromDiscriminatorValue(IParseNode parseNode)
        {
            _ = parseNode ?? throw new ArgumentNullException(nameof(parseNode));
            return new Graphdotnetv4.Models.SizeRange();
        }
        /// <summary>
        /// The deserialization information for the current model
        /// </summary>
        /// <returns>A IDictionary&lt;string, Action&lt;IParseNode&gt;&gt;</returns>
        public virtual IDictionary<string, Action<IParseNode>> GetFieldDeserializers()
        {
            return new Dictionary<string, Action<IParseNode>>
            {
                { "maximumSize", n => { MaximumSize = n.GetIntValue(); } },
                { "minimumSize", n => { MinimumSize = n.GetIntValue(); } },
            };
        }
        /// <summary>
        /// Serializes information the current object
        /// </summary>
        /// <param name="writer">Serialization writer to use to serialize this model</param>
        public virtual void Serialize(ISerializationWriter writer)
        {
            _ = writer ?? throw new ArgumentNullException(nameof(writer));
            writer.WriteIntValue("maximumSize", MaximumSize);
            writer.WriteIntValue("minimumSize", MinimumSize);
            writer.WriteAdditionalData(AdditionalData);
        }
    }
}
