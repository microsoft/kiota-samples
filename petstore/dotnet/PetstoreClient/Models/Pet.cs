using Microsoft.Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
namespace KiotaSamples.PetStoreSdk.Models {
    public class Pet : IParsable {
        /// <summary>The category property</summary>
        public KiotaSamples.PetStoreSdk.Models.Category Category { get; set; }
        /// <summary>The id property</summary>
        public long? Id { get; set; }
        /// <summary>The name property</summary>
        public string Name { get; set; }
        /// <summary>The photoUrls property</summary>
        public List<string> PhotoUrls { get; set; }
        /// <summary>pet status in the store</summary>
        public Pet_status? Status { get; set; }
        /// <summary>The tags property</summary>
        public List<Tag> Tags { get; set; }
        /// <summary>
        /// Creates a new instance of the appropriate class based on discriminator value
        /// <param name="parseNode">The parse node to use to read the discriminator value and create the object</param>
        /// </summary>
        public static Pet CreateFromDiscriminatorValue(IParseNode parseNode) {
            _ = parseNode ?? throw new ArgumentNullException(nameof(parseNode));
            return new Pet();
        }
        /// <summary>
        /// The deserialization information for the current model
        /// </summary>
        public IDictionary<string, Action<IParseNode>> GetFieldDeserializers() {
            return new Dictionary<string, Action<IParseNode>> {
                {"category", n => { Category = n.GetObjectValue<KiotaSamples.PetStoreSdk.Models.Category>(KiotaSamples.PetStoreSdk.Models.Category.CreateFromDiscriminatorValue); } },
                {"id", n => { Id = n.GetLongValue(); } },
                {"name", n => { Name = n.GetStringValue(); } },
                {"photoUrls", n => { PhotoUrls = n.GetCollectionOfPrimitiveValues<string>()?.ToList(); } },
                {"status", n => { Status = n.GetEnumValue<Pet_status>(); } },
                {"tags", n => { Tags = n.GetCollectionOfObjectValues<Tag>(Tag.CreateFromDiscriminatorValue)?.ToList(); } },
            };
        }
        /// <summary>
        /// Serializes information the current object
        /// <param name="writer">Serialization writer to use to serialize this model</param>
        /// </summary>
        public void Serialize(ISerializationWriter writer) {
            _ = writer ?? throw new ArgumentNullException(nameof(writer));
            writer.WriteObjectValue<KiotaSamples.PetStoreSdk.Models.Category>("category", Category);
            writer.WriteLongValue("id", Id);
            writer.WriteStringValue("name", Name);
            writer.WriteCollectionOfPrimitiveValues<string>("photoUrls", PhotoUrls);
            writer.WriteEnumValue<Pet_status>("status", Status);
            writer.WriteCollectionOfObjectValues<Tag>("tags", Tags);
        }
    }
}
