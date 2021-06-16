using Microsoft.Kiota.Abstractions;
using Microsoft.Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Threading.Tasks;
using Graphdotnetv4.Users;
namespace Graphdotnetv4 {
    /// <summary>The main entry point of the SDK, exposes the configuration and the fluent API.</summary>
    public class ApiClient {
        /// <summary>Current path for the request</summary>
        public string CurrentPath { get; set; }
        /// <summary>Core service to use to execute the requests</summary>
        public IHttpCore HttpCore { get; set; }
        /// <summary>Path segment to use to build the URL for the current request builder</summary>
        private string PathSegment { get; set; }
        /// <summary>Factory to use to get a serializer for payload serialization</summary>
        public ISerializationWriterFactory SerializerFactory { get; set; }
        public UsersRequestBuilder Users { get =>
            new UsersRequestBuilder { HttpCore = HttpCore, SerializerFactory = SerializerFactory, CurrentPath = CurrentPath + PathSegment };
        }
        /// <summary>
        /// Instantiates a new ApiClient and sets the default values.
        /// </summary>
        public ApiClient() {
            PathSegment = "https://graph.microsoft.com/v1.0";
        }
    }
}
