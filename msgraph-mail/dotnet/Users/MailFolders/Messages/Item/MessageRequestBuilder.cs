using Graphdotnetv4.Users.MailFolders.Messages.Attachments;
using Graphdotnetv4.Users.MailFolders.Messages.Content;
using Graphdotnetv4.Users.MailFolders.Messages.Extensions;
using Graphdotnetv4.Users.MailFolders.Messages.MultiValueExtendedProperties;
using Graphdotnetv4.Users.MailFolders.Messages.SingleValueExtendedProperties;
using Microsoft.Kiota.Abstractions;
using Microsoft.Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Threading.Tasks;
namespace Graphdotnetv4.Users.MailFolders.Messages.Item {
    /// <summary>Builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}\messages\{message-id}</summary>
    public class MessageRequestBuilder {
        public AttachmentsRequestBuilder Attachments { get =>
            new AttachmentsRequestBuilder { HttpCore = HttpCore, CurrentPath = CurrentPath + PathSegment };
        }
        public ContentRequestBuilder Content { get =>
            new ContentRequestBuilder { HttpCore = HttpCore, CurrentPath = CurrentPath + PathSegment };
        }
        /// <summary>Current path for the request</summary>
        public string CurrentPath { get; set; }
        public ExtensionsRequestBuilder Extensions { get =>
            new ExtensionsRequestBuilder { HttpCore = HttpCore, CurrentPath = CurrentPath + PathSegment };
        }
        /// <summary>Core service to use to execute the requests</summary>
        public IHttpCore HttpCore { get; set; }
        public MultiValueExtendedPropertiesRequestBuilder MultiValueExtendedProperties { get =>
            new MultiValueExtendedPropertiesRequestBuilder { HttpCore = HttpCore, CurrentPath = CurrentPath + PathSegment };
        }
        /// <summary>Path segment to use to build the URL for the current request builder</summary>
        private string PathSegment { get; set; }
        public SingleValueExtendedPropertiesRequestBuilder SingleValueExtendedProperties { get =>
            new SingleValueExtendedPropertiesRequestBuilder { HttpCore = HttpCore, CurrentPath = CurrentPath + PathSegment };
        }
        /// <summary>
        /// Instantiates a new MessageRequestBuilder and sets the default values.
        /// </summary>
        public MessageRequestBuilder() {
            PathSegment = "";
        }
        /// <summary>
        /// The collection of messages in the mailFolder.
        /// <param name="h">Request headers</param>
        /// </summary>
        public RequestInfo CreateDeleteRequestInfo(Action<IDictionary<string, string>> h = default) {
            var requestInfo = new RequestInfo {
                HttpMethod = HttpMethod.DELETE,
                URI = new Uri(CurrentPath + PathSegment),
            };
            h?.Invoke(requestInfo.Headers);
            return requestInfo;
        }
        /// <summary>
        /// The collection of messages in the mailFolder.
        /// <param name="h">Request headers</param>
        /// <param name="q">Request query parameters</param>
        /// </summary>
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
        /// <summary>
        /// The collection of messages in the mailFolder.
        /// <param name="body"></param>
        /// <param name="h">Request headers</param>
        /// </summary>
        public RequestInfo CreatePatchRequestInfo(Message body, Action<IDictionary<string, string>> h = default) {
            _ = body ?? throw new ArgumentNullException(nameof(body));
            var requestInfo = new RequestInfo {
                HttpMethod = HttpMethod.PATCH,
                URI = new Uri(CurrentPath + PathSegment),
            };
            requestInfo.SetContentFromParsable(body, HttpCore, "application/json");
            h?.Invoke(requestInfo.Headers);
            return requestInfo;
        }
        /// <summary>
        /// The collection of messages in the mailFolder.
        /// <param name="h">Request headers</param>
        /// <param name="responseHandler">Response handler to use in place of the default response handling provided by the core service</param>
        /// </summary>
        public async Task DeleteAsync(Action<IDictionary<string, string>> h = default, IResponseHandler responseHandler = default) {
            var requestInfo = CreateDeleteRequestInfo(
                h
            );
            await HttpCore.SendNoContentAsync(requestInfo, responseHandler);
        }
        /// <summary>
        /// The collection of messages in the mailFolder.
        /// <param name="h">Request headers</param>
        /// <param name="q">Request query parameters</param>
        /// <param name="responseHandler">Response handler to use in place of the default response handling provided by the core service</param>
        /// </summary>
        public async Task<Message> GetAsync(Action<GetQueryParameters> q = default, Action<IDictionary<string, string>> h = default, IResponseHandler responseHandler = default) {
            var requestInfo = CreateGetRequestInfo(
                q, h
            );
            return await HttpCore.SendAsync<Message>(requestInfo, responseHandler);
        }
        /// <summary>
        /// The collection of messages in the mailFolder.
        /// <param name="body"></param>
        /// <param name="h">Request headers</param>
        /// <param name="responseHandler">Response handler to use in place of the default response handling provided by the core service</param>
        /// </summary>
        public async Task PatchAsync(Message body, Action<IDictionary<string, string>> h = default, IResponseHandler responseHandler = default) {
            _ = body ?? throw new ArgumentNullException(nameof(body));
            var requestInfo = CreatePatchRequestInfo(
                body, h
            );
            await HttpCore.SendNoContentAsync(requestInfo, responseHandler);
        }
        /// <summary>The collection of messages in the mailFolder.</summary>
        public class GetQueryParameters : QueryParametersBase {
            /// <summary>Expand related entities</summary>
            public string[] Expand { get; set; }
            /// <summary>Select properties to be returned</summary>
            public string[] Select { get; set; }
        }
    }
}
