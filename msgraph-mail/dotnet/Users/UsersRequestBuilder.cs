using Kiota.Abstractions;
using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Threading.Tasks;
using Graphdotnetv4.Users.Item;
namespace Graphdotnetv4.Users {
    /// <summary>Builds and executes requests for operations under \users</summary>
    public class UsersRequestBuilder {
        /// <summary>Gets an item from the users collection</summary>
        public UserRequestBuilder this[string position] { get {
            return new UserRequestBuilder { HttpCore = HttpCore, SerializerFactory = SerializerFactory, CurrentPath = CurrentPath + PathSegment  + "/" + position};
        } }
        /// <summary>Path segment to use to build the URL for the current request builder</summary>
        private string PathSegment { get; } = "/users";
        /// <summary>Current path for the request</summary>
        public string CurrentPath { get; set; }
        /// <summary>Core service to use to execute the requests</summary>
        public IHttpCore HttpCore { get; set; }
        /// <summary>Factory to use to get a serializer for payload serialization</summary>
        public Func<string, ISerializationWriter> SerializerFactory { get; set; }
    }
}
