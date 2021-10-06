using Graphdotnetv4.Models.Microsoft.Graph;
using Graphdotnetv4.Users.Item.InferenceClassification.Overrides;
using Microsoft.Kiota.Abstractions;
using Microsoft.Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Threading.Tasks;
namespace Graphdotnetv4.Users.Item.InferenceClassification {
    /// <summary>Builds and executes requests for operations under \users\{user-id}\inferenceClassification</summary>
    public class InferenceClassificationRequestBuilder {
        public OverridesRequestBuilder Overrides { get =>
            new OverridesRequestBuilder(UrlTemplateParameters, RequestAdapter);
        }
        /// <summary>The request adapter to use to execute the requests.</summary>
        private IRequestAdapter RequestAdapter { get; set; }
        /// <summary>Url template to use to build the URL for the current request builder</summary>
        private string UrlTemplate { get; set; }
        /// <summary>Url template parameters for the request</summary>
        private Dictionary<string, string> UrlTemplateParameters { get; set; }
        /// <summary>
        /// Instantiates a new InferenceClassificationRequestBuilder and sets the default values.
        /// <param name="requestAdapter">The request adapter to use to execute the requests.</param>
        /// <param name="urlTemplateParameters">Url template parameters for the request</param>
        /// </summary>
        public InferenceClassificationRequestBuilder(Dictionary<string, string> urlTemplateParameters, IRequestAdapter requestAdapter) {
            _ = requestAdapter ?? throw new ArgumentNullException(nameof(requestAdapter));
            _ = urlTemplateParameters ?? throw new ArgumentNullException(nameof(urlTemplateParameters));
            UrlTemplate = "https://graph.microsoft.com/v1.0/users/{user_id}/inferenceClassification{?select,expand}";
            RequestAdapter = requestAdapter;
            UrlTemplateParameters = urlTemplateParameters;
        }
        /// <summary>
        /// Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
        /// <param name="h">Request headers</param>
        /// <param name="o">Request options</param>
        /// </summary>
        public RequestInformation CreateDeleteRequestInformation(Action<IDictionary<string, string>> h = default, IEnumerable<IRequestOption> o = default) {
            var requestInfo = new RequestInformation {
                HttpMethod = HttpMethod.DELETE,
                UrlTemplate = UrlTemplate,
                UrlTemplateParameters = UrlTemplateParameters,
            };
            h?.Invoke(requestInfo.Headers);
            requestInfo.AddRequestOptions(o?.ToArray());
            return requestInfo;
        }
        /// <summary>
        /// Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
        /// <param name="h">Request headers</param>
        /// <param name="o">Request options</param>
        /// <param name="q">Request query parameters</param>
        /// </summary>
        public RequestInformation CreateGetRequestInformation(Action<GetQueryParameters> q = default, Action<IDictionary<string, string>> h = default, IEnumerable<IRequestOption> o = default) {
            var requestInfo = new RequestInformation {
                HttpMethod = HttpMethod.GET,
                UrlTemplate = UrlTemplate,
                UrlTemplateParameters = UrlTemplateParameters,
            };
            if (q != null) {
                var qParams = new GetQueryParameters();
                q.Invoke(qParams);
                qParams.AddQueryParameters(requestInfo.QueryParameters);
            }
            h?.Invoke(requestInfo.Headers);
            requestInfo.AddRequestOptions(o?.ToArray());
            return requestInfo;
        }
        /// <summary>
        /// Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
        /// <param name="body"></param>
        /// <param name="h">Request headers</param>
        /// <param name="o">Request options</param>
        /// </summary>
        public RequestInformation CreatePatchRequestInformation(Graphdotnetv4.Models.Microsoft.Graph.InferenceClassification body, Action<IDictionary<string, string>> h = default, IEnumerable<IRequestOption> o = default) {
            _ = body ?? throw new ArgumentNullException(nameof(body));
            var requestInfo = new RequestInformation {
                HttpMethod = HttpMethod.PATCH,
                UrlTemplate = UrlTemplate,
                UrlTemplateParameters = UrlTemplateParameters,
            };
            requestInfo.SetContentFromParsable(RequestAdapter, "application/json", body);
            h?.Invoke(requestInfo.Headers);
            requestInfo.AddRequestOptions(o?.ToArray());
            return requestInfo;
        }
        /// <summary>
        /// Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
        /// <param name="h">Request headers</param>
        /// <param name="o">Request options</param>
        /// <param name="responseHandler">Response handler to use in place of the default response handling provided by the core service</param>
        /// </summary>
        public async Task DeleteAsync(Action<IDictionary<string, string>> h = default, IEnumerable<IRequestOption> o = default, IResponseHandler responseHandler = default) {
            var requestInfo = CreateDeleteRequestInformation(h, o);
            await RequestAdapter.SendNoContentAsync(requestInfo, responseHandler);
        }
        /// <summary>
        /// Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
        /// <param name="h">Request headers</param>
        /// <param name="o">Request options</param>
        /// <param name="q">Request query parameters</param>
        /// <param name="responseHandler">Response handler to use in place of the default response handling provided by the core service</param>
        /// </summary>
        public async Task<Graphdotnetv4.Models.Microsoft.Graph.InferenceClassification> GetAsync(Action<GetQueryParameters> q = default, Action<IDictionary<string, string>> h = default, IEnumerable<IRequestOption> o = default, IResponseHandler responseHandler = default) {
            var requestInfo = CreateGetRequestInformation(q, h, o);
            return await RequestAdapter.SendAsync<Graphdotnetv4.Models.Microsoft.Graph.InferenceClassification>(requestInfo, responseHandler);
        }
        /// <summary>
        /// Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.
        /// <param name="body"></param>
        /// <param name="h">Request headers</param>
        /// <param name="o">Request options</param>
        /// <param name="responseHandler">Response handler to use in place of the default response handling provided by the core service</param>
        /// </summary>
        public async Task PatchAsync(Graphdotnetv4.Models.Microsoft.Graph.InferenceClassification body, Action<IDictionary<string, string>> h = default, IEnumerable<IRequestOption> o = default, IResponseHandler responseHandler = default) {
            _ = body ?? throw new ArgumentNullException(nameof(body));
            var requestInfo = CreatePatchRequestInformation(body, h, o);
            await RequestAdapter.SendNoContentAsync(requestInfo, responseHandler);
        }
        /// <summary>Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.</summary>
        public class GetQueryParameters : QueryParametersBase {
            /// <summary>Expand related entities</summary>
            public string[] Expand { get; set; }
            /// <summary>Select properties to be returned</summary>
            public string[] Select { get; set; }
        }
    }
}
