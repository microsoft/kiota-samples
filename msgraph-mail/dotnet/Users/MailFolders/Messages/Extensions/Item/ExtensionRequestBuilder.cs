using Kiota.Abstractions;
using System;
using System.Collections.Generic;
using System.IO;
using System.Threading.Tasks;
namespace Graphdotnetv4.Users.MailFolders.Messages.Extensions.Item {
    public class ExtensionRequestBuilder {
        public async Task<Extension> GetAsync(Action<GetQueryParameters> q = default, Action<IDictionary<string, string>> h = default, IResponseHandler responseHandler = default) {
            var requestInfo = CreateGetRequestInfo(
                q, h
            );
            return await HttpCore.SendAsync<Extension>(requestInfo, responseHandler);
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
        public async Task<object> PatchAsync(Extension body, Action<IDictionary<string, string>> h = default, IResponseHandler responseHandler = default) {
            var requestInfo = CreatePatchRequestInfo(
                body, h
            );
            return await HttpCore.SendAsync<object>(requestInfo, responseHandler);
        }
        public RequestInfo CreatePatchRequestInfo(Extension body, Action<IDictionary<string, string>> h = default) {
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
        private string PathSegment { get; } = "";
        public string CurrentPath { get; set; }
        public IHttpCore HttpCore { get; set; }
        public class GetQueryParameters : QueryParametersBase {
            public string[] Select { get; set; }
            public string[] Expand { get; set; }
        }
    }
}
