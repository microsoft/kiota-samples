using Kiota.Abstractions;
using System;
using System.Collections.Generic;
using System.IO;
using System.Threading.Tasks;
namespace Graphdotnetv4.Users.MailFolders.Messages.Content {
    public class ContentRequestBuilder {
        public async Task<object> GetAsync(Action<IDictionary<string, string>> h = default, IResponseHandler responseHandler = default) {
            var requestInfo = CreateGetRequestInfo(
                h
            );
            return await HttpCore.SendAsync<object>(requestInfo, responseHandler);
        }
        public RequestInfo CreateGetRequestInfo(Action<IDictionary<string, string>> h = default) {
            var requestInfo = new RequestInfo {
                HttpMethod = HttpMethod.GET,
                URI = new Uri(CurrentPath + PathSegment),
            };
            h?.Invoke(requestInfo.Headers);
            return requestInfo;
        }
        public async Task<object> PutAsync(Action<IDictionary<string, string>> h = default, IResponseHandler responseHandler = default) {
            var requestInfo = CreatePutRequestInfo(
                h
            );
            return await HttpCore.SendAsync<object>(requestInfo, responseHandler);
        }
        public RequestInfo CreatePutRequestInfo(Action<IDictionary<string, string>> h = default) {
            var requestInfo = new RequestInfo {
                HttpMethod = HttpMethod.PUT,
                URI = new Uri(CurrentPath + PathSegment),
            };
            h?.Invoke(requestInfo.Headers);
            return requestInfo;
        }
        private string PathSegment { get; } = "/$value";
        public string CurrentPath { get; set; }
        public IHttpCore HttpCore { get; set; }
    }
}
