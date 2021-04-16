using Kiota.Abstractions;
using System;
using System.Collections.Generic;
using System.IO;
using System.Threading.Tasks;
using Graphdotnetv4.Users.Item;
namespace Graphdotnetv4.Users {
    public class UsersRequestBuilder {
        public UserRequestBuilder this[string position] { get {
            return new UserRequestBuilder { HttpCore = HttpCore, CurrentPath = CurrentPath + PathSegment  + "/" + position};
        } }
        private string PathSegment { get; } = "/users";
        public string CurrentPath { get; set; }
        public IHttpCore HttpCore { get; set; }
    }
}
