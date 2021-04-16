using Kiota.Abstractions;
using System;
using System.Collections.Generic;
using System.IO;
using System.Threading.Tasks;
using Graphdotnetv4.Users.MailFolders.ChildFolders;
using Graphdotnetv4.Users.MailFolders.MessageRules;
using Graphdotnetv4.Users.MailFolders.Messages;
using Graphdotnetv4.Users.MailFolders.Messages.MultiValueExtendedProperties;
using Graphdotnetv4.Users.MailFolders.Messages.SingleValueExtendedProperties;
namespace Graphdotnetv4.Users.MailFolders.Item {
    public class MailFolderRequestBuilder {
        public ChildFoldersRequestBuilder ChildFolders { get =>
            new ChildFoldersRequestBuilder { HttpCore = HttpCore, CurrentPath = CurrentPath + PathSegment };
        }
        public MessageRulesRequestBuilder MessageRules { get =>
            new MessageRulesRequestBuilder { HttpCore = HttpCore, CurrentPath = CurrentPath + PathSegment };
        }
        public MessagesRequestBuilder Messages { get =>
            new MessagesRequestBuilder { HttpCore = HttpCore, CurrentPath = CurrentPath + PathSegment };
        }
        public MultiValueExtendedPropertiesRequestBuilder MultiValueExtendedProperties { get =>
            new MultiValueExtendedPropertiesRequestBuilder { HttpCore = HttpCore, CurrentPath = CurrentPath + PathSegment };
        }
        public SingleValueExtendedPropertiesRequestBuilder SingleValueExtendedProperties { get =>
            new SingleValueExtendedPropertiesRequestBuilder { HttpCore = HttpCore, CurrentPath = CurrentPath + PathSegment };
        }
        public async Task<MailFolder> GetAsync(Action<GetQueryParameters> q = default, Action<IDictionary<string, string>> h = default, IResponseHandler responseHandler = default) {
            var requestInfo = CreateGetRequestInfo(
                q, h
            );
            return await HttpCore.SendAsync<MailFolder>(requestInfo, responseHandler);
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
        public async Task<object> PatchAsync(MailFolder body, Action<IDictionary<string, string>> h = default, IResponseHandler responseHandler = default) {
            var requestInfo = CreatePatchRequestInfo(
                body, h
            );
            return await HttpCore.SendAsync<object>(requestInfo, responseHandler);
        }
        public RequestInfo CreatePatchRequestInfo(MailFolder body, Action<IDictionary<string, string>> h = default) {
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
