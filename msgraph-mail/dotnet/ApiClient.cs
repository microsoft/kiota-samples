using Graphdotnetv4.Users;
using Microsoft.Kiota.Abstractions;
using Microsoft.Kiota.Serialization.Json;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Threading.Tasks;
namespace Graphdotnetv4 {
    /// <summary>The main entry point of the SDK, exposes the configuration and the fluent API.</summary>
    public class ApiClient {
        /// <summary>Path segment to use to build the URL for the current request builder</summary>
        private string PathSegment { get; set; }
        /// <summary>The http core service to use to execute the requests.</summary>
        private IRequestAdapter RequestAdapter { get; set; }
        public UsersRequestBuilder Users { get =>
            new UsersRequestBuilder(PathSegment , RequestAdapter, false);
        }
        /// <summary>
        /// Instantiates a new ApiClient and sets the default values.
        /// <param name="requestAdapter">The http core service to use to execute the requests.</param>
        /// </summary>
        public ApiClient(IRequestAdapter requestAdapter) {
            _ = requestAdapter ?? throw new ArgumentNullException(nameof(requestAdapter));
            PathSegment = "https://graph.microsoft.com/v1.0";
            RequestAdapter = requestAdapter;
            ApiClientBuilder.RegisterDefaultSerializer<JsonSerializationWriterFactory>();
            ApiClientBuilder.RegisterDefaultDeserializer<JsonParseNodeFactory>();
        }
    }
}
