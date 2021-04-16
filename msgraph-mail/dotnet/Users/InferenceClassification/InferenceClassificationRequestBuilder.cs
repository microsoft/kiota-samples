using Kiota.Abstractions;
using System;
using System.Collections.Generic;
using System.IO;
using System.Threading.Tasks;
using Graphdotnetv4.Users.InferenceClassification.Overrides;
namespace Graphdotnetv4.Users.InferenceClassification {
    public class InferenceClassificationRequestBuilder {
        public OverridesRequestBuilder Overrides { get =>
            new OverridesRequestBuilder { HttpCore = HttpCore, CurrentPath = CurrentPath + PathSegment };
        }
        public async Task<InferenceClassification> GetAsync(Action<GetQueryParameters> q = default, Action<IDictionary<string, string>> h = default, IResponseHandler responseHandler = default) {
            var requestInfo = CreateGetRequestInfo(
                q, h
            );
            return await HttpCore.SendAsync<InferenceClassification>(requestInfo, responseHandler);
        }
        public RequestInfo CreateGetRequestInfo(Action<GetQueryParameters> q = default, Action<IDictionary<string, string>> h = default) {
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
            return requestInfo;
        }
        public async Task<object> PatchAsync(InferenceClassification body, Action<IDictionary<string, string>> h = default, IResponseHandler responseHandler = default) {
            var requestInfo = CreatePatchRequestInfo(
                body, h
            );
            return await HttpCore.SendAsync<object>(requestInfo, responseHandler);
        }
        public RequestInfo CreatePatchRequestInfo(InferenceClassification body, Action<IDictionary<string, string>> h = default) {
            var requestInfo = new RequestInfo {
                HttpMethod = HttpMethod.PATCH,
                URI = new Uri(CurrentPath),
                Content = body as object as Stream
            };
            h?.Invoke(requestInfo.Headers);
            return requestInfo;
        }
        public async Task<object> DeleteAsync(Action<IDictionary<string, string>> h = default, IResponseHandler responseHandler = default) {
            var requestInfo = CreateDeleteRequestInfo(
                h
            );
            return await HttpCore.SendAsync<object>(requestInfo, responseHandler);
        }
        public RequestInfo CreateDeleteRequestInfo(Action<IDictionary<string, string>> h = default) {
            var requestInfo = new RequestInfo {
                HttpMethod = HttpMethod.DELETE,
                URI = new Uri(CurrentPath),
            };
            h?.Invoke(requestInfo.Headers);
            return requestInfo;
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
