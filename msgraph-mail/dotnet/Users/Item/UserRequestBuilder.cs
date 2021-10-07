using Graphdotnetv4.Users.Item.InferenceClassification;
using Graphdotnetv4.Users.Item.MailFolders;
using Graphdotnetv4.Users.Item.Messages;
using Microsoft.Kiota.Abstractions;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Threading.Tasks;
namespace Graphdotnetv4.Users.Item {
    /// <summary>Builds and executes requests for operations under \users\{user-id}</summary>
    public class UserRequestBuilder {
        public InferenceClassificationRequestBuilder InferenceClassification { get =>
            new InferenceClassificationRequestBuilder(UrlTemplateParameters, RequestAdapter);
        }
        public MailFoldersRequestBuilder MailFolders { get =>
            new MailFoldersRequestBuilder(UrlTemplateParameters, RequestAdapter);
        }
        public MessagesRequestBuilder Messages { get =>
            new MessagesRequestBuilder(UrlTemplateParameters, RequestAdapter);
        }
        /// <summary>The request adapter to use to execute the requests.</summary>
        private IRequestAdapter RequestAdapter { get; set; }
        /// <summary>Url template to use to build the URL for the current request builder</summary>
        private string UrlTemplate { get; set; }
        /// <summary>Url template parameters for the request</summary>
        private Dictionary<string, string> UrlTemplateParameters { get; set; }
        /// <summary>
        /// Instantiates a new UserRequestBuilder and sets the default values.
        /// <param name="requestAdapter">The request adapter to use to execute the requests.</param>
        /// <param name="urlTemplateParameters">Url template parameters for the request</param>
        /// </summary>
        public UserRequestBuilder(Dictionary<string, string> urlTemplateParameters, IRequestAdapter requestAdapter) {
            _ = requestAdapter ?? throw new ArgumentNullException(nameof(requestAdapter));
            _ = urlTemplateParameters ?? throw new ArgumentNullException(nameof(urlTemplateParameters));
            UrlTemplate = "https://graph.microsoft.com/v1.0/users/{user_id}";
            var urlTplParams = new Dictionary<string, string>(urlTemplateParameters);
            UrlTemplateParameters = urlTplParams;
            RequestAdapter = requestAdapter;
        }
        /// <summary>
        /// Instantiates a new UserRequestBuilder and sets the default values.
        /// <param name="rawUrl">The raw URL to use for the request builder.</param>
        /// <param name="requestAdapter">The request adapter to use to execute the requests.</param>
        /// </summary>
        public UserRequestBuilder(string rawUrl, IRequestAdapter requestAdapter) {
            if(string.IsNullOrEmpty(rawUrl)) throw new ArgumentNullException(nameof(rawUrl));
            _ = requestAdapter ?? throw new ArgumentNullException(nameof(requestAdapter));
            UrlTemplate = "https://graph.microsoft.com/v1.0/users/{user_id}";
            var urlTplParams = new Dictionary<string, string>();
            urlTplParams.Add("request-raw-url", rawUrl);
            UrlTemplateParameters = urlTplParams;
            RequestAdapter = requestAdapter;
        }
    }
}
