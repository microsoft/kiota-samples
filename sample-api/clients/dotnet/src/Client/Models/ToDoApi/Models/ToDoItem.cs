using Microsoft.Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
namespace ToDoClient.ApiClient.Models.ToDoApi.Models {
    public class ToDoItem : IAdditionalDataHolder, IParsable {
        /// <summary>Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.</summary>
        public IDictionary<string, object> AdditionalData { get; set; }
        /// <summary>The dueDate property</summary>
        public DateTimeOffset? DueDate { get; set; }
        /// <summary>The id property</summary>
        public string Id { get; set; }
        /// <summary>The isComplete property</summary>
        public bool? IsComplete { get; set; }
        /// <summary>The name property</summary>
        public string Name { get; set; }
        /// <summary>The priority property</summary>
        public int? Priority { get; set; }
        /// <summary>The type property</summary>
        public string Type { get; set; }
        /// <summary>
        /// Instantiates a new ToDoItem and sets the default values.
        /// </summary>
        public ToDoItem() {
            AdditionalData = new Dictionary<string, object>();
        }
        /// <summary>
        /// Creates a new instance of the appropriate class based on discriminator value
        /// <param name="parseNode">The parse node to use to read the discriminator value and create the object</param>
        /// </summary>
        public static ToDoItem CreateFromDiscriminatorValue(IParseNode parseNode) {
            _ = parseNode ?? throw new ArgumentNullException(nameof(parseNode));
            return new ToDoItem();
        }
        /// <summary>
        /// The deserialization information for the current model
        /// </summary>
        public IDictionary<string, Action<IParseNode>> GetFieldDeserializers() {
            return new Dictionary<string, Action<IParseNode>> {
                {"dueDate", n => { DueDate = n.GetDateTimeOffsetValue(); } },
                {"id", n => { Id = n.GetStringValue(); } },
                {"isComplete", n => { IsComplete = n.GetBoolValue(); } },
                {"name", n => { Name = n.GetStringValue(); } },
                {"priority", n => { Priority = n.GetIntValue(); } },
                {"type", n => { Type = n.GetStringValue(); } },
            };
        }
        /// <summary>
        /// Serializes information the current object
        /// <param name="writer">Serialization writer to use to serialize this model</param>
        /// </summary>
        public void Serialize(ISerializationWriter writer) {
            _ = writer ?? throw new ArgumentNullException(nameof(writer));
            writer.WriteDateTimeOffsetValue("dueDate", DueDate);
            writer.WriteStringValue("id", Id);
            writer.WriteBoolValue("isComplete", IsComplete);
            writer.WriteStringValue("name", Name);
            writer.WriteIntValue("priority", Priority);
            writer.WriteStringValue("type", Type);
            writer.WriteAdditionalData(AdditionalData);
        }
    }
}
