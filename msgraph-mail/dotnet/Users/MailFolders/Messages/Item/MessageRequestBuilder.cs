using Kiota.Abstractions;
using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Threading.Tasks;
using Graphdotnetv4.Users.MailFolders.Messages.Attachments;
using Graphdotnetv4.Users.MailFolders.Messages.Content;
using Graphdotnetv4.Users.MailFolders.Messages.Extensions;
using Graphdotnetv4.Users.MailFolders.Messages.MultiValueExtendedProperties;
using Graphdotnetv4.Users.MailFolders.Messages.SingleValueExtendedProperties;
namespace Graphdotnetv4.Users.MailFolders.Messages.Item {
    public class MessageRequestBuilder {
        public ContentRequestBuilder Content { get =>
            new ContentRequestBuilder { HttpCore = HttpCore, SerializerFactory = SerializerFactory, CurrentPath = CurrentPath + PathSegment };
        }
        public AttachmentsRequestBuilder Attachments { get =>
            new AttachmentsRequestBuilder { HttpCore = HttpCore, SerializerFactory = SerializerFactory, CurrentPath = CurrentPath + PathSegment };
        }
        public ExtensionsRequestBuilder Extensions { get =>
            new ExtensionsRequestBuilder { HttpCore = HttpCore, SerializerFactory = SerializerFactory, CurrentPath = CurrentPath + PathSegment };
        }
        public MultiValueExtendedPropertiesRequestBuilder MultiValueExtendedProperties { get =>
            new MultiValueExtendedPropertiesRequestBuilder { HttpCore = HttpCore, SerializerFactory = SerializerFactory, CurrentPath = CurrentPath + PathSegment };
        }
        public SingleValueExtendedPropertiesRequestBuilder SingleValueExtendedProperties { get =>
            new SingleValueExtendedPropertiesRequestBuilder { HttpCore = HttpCore, SerializerFactory = SerializerFactory, CurrentPath = CurrentPath + PathSegment };
        }
        public async Task<Message> GetAsync(Action<GetQueryParameters> q = default, Action<IDictionary<string, string>> h = default, IResponseHandler responseHandler = default) {
            var requestInfo = CreateGetRequestInfo(
                q, h
            );
            return await HttpCore.SendAsync<Message>(requestInfo, responseHandler);
        }
        public RequestInfo CreateGetRequestInfo(Action<GetQueryParameters> q = default, Action<IDictionary<string, string>> h = default) {
            var requestInfo = new RequestInfo {
                HttpMethod = HttpMethod.GET,
                URI = new Uri(CurrentPath + PathSegment),
            };
            if (q != null) {
                var qParams = new GetQueryParameters();
                q.Invoke(qParams);
                qParams.AddQueryParameters(requestInfo.QueryParameters);
            }
            h?.Invoke(requestInfo.Headers);
            return requestInfo;
        }
        public async Task PatchAsync(Message body, Action<IDictionary<string, string>> h = default, IResponseHandler responseHandler = default) {
            var requestInfo = CreatePatchRequestInfo(
                body, h
            );
            await HttpCore.SendNoContentAsync(requestInfo, responseHandler);
        }
        public RequestInfo CreatePatchRequestInfo(Message body, Action<IDictionary<string, string>> h = default) {
            var requestInfo = new RequestInfo {
                HttpMethod = HttpMethod.PATCH,
                URI = new Uri(CurrentPath + PathSegment),
            };
            requestInfo.SetJsonContentFromParsable(body, SerializerFactory);
            h?.Invoke(requestInfo.Headers);
            return requestInfo;
        }
        public async Task DeleteAsync(Action<IDictionary<string, string>> h = default, IResponseHandler responseHandler = default) {
            var requestInfo = CreateDeleteRequestInfo(
                h
            );
            await HttpCore.SendNoContentAsync(requestInfo, responseHandler);
        }
        public RequestInfo CreateDeleteRequestInfo(Action<IDictionary<string, string>> h = default) {
            var requestInfo = new RequestInfo {
                HttpMethod = HttpMethod.DELETE,
                URI = new Uri(CurrentPath + PathSegment),
            };
            h?.Invoke(requestInfo.Headers);
            return requestInfo;
        }
        private string PathSegment { get; } = "";
        public string CurrentPath { get; set; }
        public IHttpCore HttpCore { get; set; }
        public Func<string, ISerializationWriter> SerializerFactory { get; set; }
        public class GetQueryParameters : QueryParametersBase {
            public string[] Select { get; set; }
            public string[] Expand { get; set; }
        }
    }
}
