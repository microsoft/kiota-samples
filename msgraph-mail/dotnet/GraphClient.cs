using Kiota.Abstractions;
using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Threading.Tasks;
using Graphdotnetv4.Users;
namespace Graphdotnetv4 {
    public class GraphClient {
        public UsersRequestBuilder Users { get =>
            new UsersRequestBuilder { HttpCore = HttpCore, SerializerFactory = SerializerFactory, CurrentPath = CurrentPath + PathSegment };
        }
        private string PathSegment { get; } = "https://graph.microsoft.com/v1.0";
        public string CurrentPath { get; set; }
        public IHttpCore HttpCore { get; set; }
        public Func<string, ISerializationWriter> SerializerFactory { get; set; }
    }
}
