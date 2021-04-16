using Kiota.Abstractions;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Graphdotnetv4.Users.MailFolders.Messages.Attachments;
using Graphdotnetv4.Users.MailFolders.Messages.Content;
using Graphdotnetv4.Users.MailFolders.Messages.Extensions;
using Graphdotnetv4.Users.MailFolders.Messages.MultiValueExtendedProperties;
using Graphdotnetv4.Users.MailFolders.Messages.SingleValueExtendedProperties;
namespace Graphdotnetv4.Users.Messages.Item {
    public class MessageRequestBuilder {
        public ContentRequestBuilder Content { get =>
            new ContentRequestBuilder { HttpCore = HttpCore, CurrentPath = CurrentPath + PathSegment };
        }
        public AttachmentsRequestBuilder Attachments { get =>
            new AttachmentsRequestBuilder { HttpCore = HttpCore, CurrentPath = CurrentPath + PathSegment };
        }
        public ExtensionsRequestBuilder Extensions { get =>
            new ExtensionsRequestBuilder { HttpCore = HttpCore, CurrentPath = CurrentPath + PathSegment };
        }
        public MultiValueExtendedPropertiesRequestBuilder MultiValueExtendedProperties { get =>
            new MultiValueExtendedPropertiesRequestBuilder { HttpCore = HttpCore, CurrentPath = CurrentPath + PathSegment };
        }
        public SingleValueExtendedPropertiesRequestBuilder SingleValueExtendedProperties { get =>
            new SingleValueExtendedPropertiesRequestBuilder { HttpCore = HttpCore, CurrentPath = CurrentPath + PathSegment };
        }
        public async Task<Message> GetAsync(Action<GetQueryParameters> q = default, Action<IDictionary<string, string>> h = default, IResponseHandler responseHandler = default) {
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
            return await HttpCore.SendAsync<Message>(requestInfo, responseHandler);
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
        private string PathSegment { get; } = "";
        public string CurrentPath { get; set; }
        public IHttpCore HttpCore { get; set; }
        public class GetQueryParameters : QueryParametersBase {
            public string[] Select { get; set; }
            public string[] Expand { get; set; }
        }
    }
}
