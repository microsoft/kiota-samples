using Kiota.Abstractions;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Graphdotnetv4.Users;
namespace Graphdotnetv4 {
    public class GraphClient {
        public UsersRequestBuilder Users { get =>
            new UsersRequestBuilder { HttpCore = HttpCore, CurrentPath = CurrentPath + PathSegment };
        }
        private string PathSegment { get; } = "https://graph.microsoft.com/v1.0";
        public string CurrentPath { get; set; }
        public IHttpCore HttpCore { get; set; }
    }
}
