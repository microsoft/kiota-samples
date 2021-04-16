using Kiota.Abstractions;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Graphdotnetv4.Users.InferenceClassification.Overrides;
namespace Graphdotnetv4.Users.InferenceClassification {
    public class InferenceClassificationRequestBuilder {
        public OverridesRequestBuilder Overrides { get =>
            new OverridesRequestBuilder { HttpCore = HttpCore, CurrentPath = CurrentPath + PathSegment };
        }
        public async Task<InferenceClassification> GetAsync(Action<GetQueryParameters> q = default, Action<IDictionary<string, string>> h = default, IResponseHandler responseHandler = default) {
            var requestInfo = new RequestInfo {
                HttpMethod = HttpMethod.GET,
                URI = new Uri(CurrentPath),
            };
            if (q != null) {
                var qParams = new GetQueryParameters();
                q.Invoke(qParams);
                qParams.AddQueryParameters(requestInfo.QueryParameters);
            }
            h?.Invoke(requestInfo.Headers);
            return await HttpCore.SendAsync<InferenceClassification>(requestInfo, responseHandler);
        }
        public async Task<object> PatchAsync(Action<IDictionary<string, string>> h = default, IResponseHandler responseHandler = default) {
            var requestInfo = new RequestInfo {
                HttpMethod = HttpMethod.PATCH,
                URI = new Uri(CurrentPath),
            };
            h?.Invoke(requestInfo.Headers);
            return await HttpCore.SendAsync<object>(requestInfo, responseHandler);
        }
        public async Task<object> DeleteAsync(Action<IDictionary<string, string>> h = default, IResponseHandler responseHandler = default) {
            var requestInfo = new RequestInfo {
                HttpMethod = HttpMethod.DELETE,
                URI = new Uri(CurrentPath),
            };
            h?.Invoke(requestInfo.Headers);
            return await HttpCore.SendAsync<object>(requestInfo, responseHandler);
        }
        private string PathSegment { get; } = "/inferenceClassification";
        public string CurrentPath { get; set; }
        public IHttpCore HttpCore { get; set; }
        public class GetQueryParameters : QueryParametersBase {
            public string[] Select { get; set; }
            public string[] Expand { get; set; }
        }
    }
}
