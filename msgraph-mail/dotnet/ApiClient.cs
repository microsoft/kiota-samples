using Graphdotnetv4.Users;
using Microsoft.Kiota.Abstractions;
using Microsoft.Kiota.Abstractions.Serialization;
using Microsoft.Kiota.Serialization.Json;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Threading.Tasks;
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
        /// Instantiates a new Api client and sets the default values.
        /// <param name="httpCore">The http core service to use to execute the requests.</param>
        /// <param name="serializationWriterFactory">Factory to use to get a serializer for payload serialization.</param>
        /// </summary>
        public ApiClient(IHttpCore httpCore, ISerializationWriterFactory serializationWriterFactory = default) {
            _ = httpCore ?? throw new ArgumentNullException(nameof(httpCore));
            PathSegment = "https://graph.microsoft.com/v1.0";
            HttpCore = httpCore;
            ApiClientBuilder.RegisterDefaultSerializers<JsonSerializationWriterFactory>();
            ApiClientBuilder.RegisterDefaultDeSerializers<JsonParseNodeFactory>();
            if(serializationWriterFactory == default && !SerializationWriterFactoryRegistry.DefaultInstance.ContentTypeAssociatedFactories.Any()) throw new InvalidOperationException("The Serialization Writer factory has not been initialized for this client.");
            if(serializationWriterFactory == default && !ParseNodeFactoryRegistry.DefaultInstance.ContentTypeAssociatedFactories.Any()) throw new InvalidOperationException("The Parse Node factory has not been initialized for this client.");
            SerializerFactory = serializationWriterFactory ?? SerializationWriterFactoryRegistry.DefaultInstance;
        }
    }
}
