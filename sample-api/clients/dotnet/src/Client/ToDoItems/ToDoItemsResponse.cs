using Microsoft.Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using ToDoClient.ApiClient.Models.ToDoApi.Models;
namespace ToDoClient.ApiClient.ToDoItems {
    public class ToDoItemsResponse : IParsable {
        /// <summary>Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.</summary>
        public IDictionary<string, object> AdditionalData { get; set; }
        public List<ToDoItem> Value { get; set; }
        /// <summary>
        /// Instantiates a new ToDoItemsResponse and sets the default values.
        /// </summary>
        public ToDoItemsResponse() {
            AdditionalData = new Dictionary<string, object>();
        }
        /// <summary>
        /// The deserialization information for the current model
        /// </summary>
        public IDictionary<string, Action<T, IParseNode>> GetFieldDeserializers<T>() {
            return new Dictionary<string, Action<T, IParseNode>> {
                {"value", (o,n) => { (o as ToDoItemsResponse).Value = n.GetCollectionOfObjectValues<ToDoItem>().ToList(); } },
            };
        }
        /// <summary>
        /// Serializes information the current object
        /// <param name="writer">Serialization writer to use to serialize this model</param>
        /// </summary>
        public void Serialize(ISerializationWriter writer) {
            _ = writer ?? throw new ArgumentNullException(nameof(writer));
            writer.WriteCollectionOfObjectValues<ToDoItem>("value", Value);
            writer.WriteAdditionalData(AdditionalData);
        }
    }
}
