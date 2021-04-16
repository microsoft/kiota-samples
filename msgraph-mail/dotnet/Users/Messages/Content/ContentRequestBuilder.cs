using Kiota.Abstractions;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
namespace Graphdotnetv4.Users.Messages.Content {
    public class ContentRequestBuilder {
        public async Task<object> GetAsync(Action<IDictionary<string, string>> h = default, IResponseHandler responseHandler = default) {
            var requestInfo = new RequestInfo {
                HttpMethod = HttpMethod.GET,
                URI = new Uri(CurrentPath),
            };
            h?.Invoke(requestInfo.Headers);
            return await HttpCore.SendAsync<object>(requestInfo, responseHandler);
        }
        public async Task<object> PutAsync(Action<IDictionary<string, string>> h = default, IResponseHandler responseHandler = default) {
            var requestInfo = new RequestInfo {
                HttpMethod = HttpMethod.PUT,
                URI = new Uri(CurrentPath),
            };
            h?.Invoke(requestInfo.Headers);
            return await HttpCore.SendAsync<object>(requestInfo, responseHandler);
        }
        private string PathSegment { get; } = "/$value";
        public string CurrentPath { get; set; }
        public IHttpCore HttpCore { get; set; }
    }
}
