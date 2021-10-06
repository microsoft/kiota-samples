using Graphdotnetv4.Users.Item;
using Microsoft.Kiota.Abstractions;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Threading.Tasks;
namespace Graphdotnetv4.Users {
    /// <summary>Builds and executes requests for operations under \users</summary>
    public class UsersRequestBuilder {
        /// <summary>The request adapter to use to execute the requests.</summary>
        private IRequestAdapter RequestAdapter { get; set; }
        /// <summary>Url template to use to build the URL for the current request builder</summary>
        private string UrlTemplate { get; set; }
        /// <summary>Url template parameters for the request</summary>
        private Dictionary<string, string> UrlTemplateParameters { get; set; }
        /// <summary>Gets an item from the Graphdotnetv4.users.item collection</summary>
        public UserRequestBuilder this[string position] { get {
            var urlTplParams = new Dictionary<string, string>(UrlTemplateParameters);
            urlTplParams.Add("user_id", position);
            return new UserRequestBuilder(urlTplParams, RequestAdapter);
        } }
        /// <summary>
        /// Instantiates a new UsersRequestBuilder and sets the default values.
        /// <param name="requestAdapter">The request adapter to use to execute the requests.</param>
        /// <param name="urlTemplateParameters">Url template parameters for the request</param>
        /// </summary>
        public UsersRequestBuilder(Dictionary<string, string> urlTemplateParameters, IRequestAdapter requestAdapter) {
            _ = requestAdapter ?? throw new ArgumentNullException(nameof(requestAdapter));
            _ = urlTemplateParameters ?? throw new ArgumentNullException(nameof(urlTemplateParameters));
            UrlTemplate = "https://graph.microsoft.com/v1.0/users";
            RequestAdapter = requestAdapter;
            UrlTemplateParameters = urlTemplateParameters;
        }
    }
}
