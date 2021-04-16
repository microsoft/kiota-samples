using Kiota.Abstractions;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Graphdotnetv4.Users.MailFolders.Item;
namespace Graphdotnetv4.Users.MailFolders {
    public class MailFoldersRequestBuilder {
        public MailFolderRequestBuilder this[string position] { get {
            return new MailFolderRequestBuilder { HttpCore = HttpCore, CurrentPath = CurrentPath + PathSegment  + "/" + position};
        } }
        public async Task<MailFoldersResponse> GetAsync(Action<GetQueryParameters> q = default, Action<IDictionary<string, string>> h = default, IResponseHandler responseHandler = default) {
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
            return await HttpCore.SendAsync<MailFoldersResponse>(requestInfo, responseHandler);
        }
        public async Task<MailFolder> PostAsync(Action<IDictionary<string, string>> h = default, IResponseHandler responseHandler = default) {
            var requestInfo = new RequestInfo {
                HttpMethod = HttpMethod.POST,
                URI = new Uri(CurrentPath),
            };
            h?.Invoke(requestInfo.Headers);
            return await HttpCore.SendAsync<MailFolder>(requestInfo, responseHandler);
        }
        private string PathSegment { get; } = "/mailFolders";
        public string CurrentPath { get; set; }
        public IHttpCore HttpCore { get; set; }
        public class GetQueryParameters : QueryParametersBase {
            public int Top { get; set; }
            public int Skip { get; set; }
            public string Search { get; set; }
            public string Filter { get; set; }
            public bool Count { get; set; }
            public string[] Orderby { get; set; }
            public string[] Select { get; set; }
            public string[] Expand { get; set; }
        }
    }
}
