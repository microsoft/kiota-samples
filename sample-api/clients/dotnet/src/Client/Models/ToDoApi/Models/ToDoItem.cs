using Microsoft.Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
namespace ToDoClient.ApiClient.Models.ToDoApi.Models {
    public class ToDoItem : IParsable {
        /// <summary>Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.</summary>
        public IDictionary<string, object> AdditionalData { get; set; }
        public DateTimeOffset? DueDate { get; set; }
        public string Id { get; set; }
        public bool? IsComplete { get; set; }
        public string Name { get; set; }
        public int? Priority { get; set; }
        public string Type { get; set; }
        /// <summary>
        /// Instantiates a new ToDoItem and sets the default values.
        /// </summary>
        public ToDoItem() {
            AdditionalData = new Dictionary<string, object>();
        }
        /// <summary>
        /// The deserialization information for the current model
        /// </summary>
        public IDictionary<string, Action<T, IParseNode>> GetFieldDeserializers<T>() {
            return new Dictionary<string, Action<T, IParseNode>> {
                {"dueDate", (o,n) => { (o as ToDoItem).DueDate = n.GetDateTimeOffsetValue(); } },
                {"id", (o,n) => { (o as ToDoItem).Id = n.GetStringValue(); } },
                {"isComplete", (o,n) => { (o as ToDoItem).IsComplete = n.GetBoolValue(); } },
                {"name", (o,n) => { (o as ToDoItem).Name = n.GetStringValue(); } },
                {"priority", (o,n) => { (o as ToDoItem).Priority = n.GetIntValue(); } },
                {"type", (o,n) => { (o as ToDoItem).Type = n.GetStringValue(); } },
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
