using Microsoft.Kiota.Abstractions;
using Microsoft.Kiota.Serialization.Json;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Threading.Tasks;
using ToDoClient.ApiClient.ToDoItems;
namespace ToDoClient.ApiClient {
    /// <summary>The main entry point of the SDK, exposes the configuration and the fluent API.</summary>
    public class ToDoApiClient {
        /// <summary>Path parameters for the request</summary>
        private Dictionary<string, object> PathParameters { get; set; }
        /// <summary>The request adapter to use to execute the requests.</summary>
        private IRequestAdapter RequestAdapter { get; set; }
        public ToDoItemsRequestBuilder ToDoItems { get =>
            new ToDoItemsRequestBuilder(PathParameters, RequestAdapter);
        }
        /// <summary>Url template to use to build the URL for the current request builder</summary>
        private string UrlTemplate { get; set; }
        /// <summary>
        /// Instantiates a new ToDoApiClient and sets the default values.
        /// <param name="requestAdapter">The request adapter to use to execute the requests.</param>
        /// </summary>
        public ToDoApiClient(IRequestAdapter requestAdapter) {
            _ = requestAdapter ?? throw new ArgumentNullException(nameof(requestAdapter));
            PathParameters = new Dictionary<string, object>();
            UrlTemplate = "{+baseurl}";
            RequestAdapter = requestAdapter;
            ApiClientBuilder.RegisterDefaultSerializer<JsonSerializationWriterFactory>();
            ApiClientBuilder.RegisterDefaultDeserializer<JsonParseNodeFactory>();
            RequestAdapter.BaseUrl = "https://localhost:7206";
        }
    }
}
