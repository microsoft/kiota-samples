using Kiota.Abstractions;
using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Threading.Tasks;
using Graphdotnetv4.Users.MailFolders.Messages.Attachments.Item;
namespace Graphdotnetv4.Users.Messages.Attachments {
    /// <summary>Builds and executes requests for operations under \users\{user-id}\messages\{message-id}\attachments</summary>
    public class AttachmentsRequestBuilder {
        /// <summary>Gets an item from the users.messages.attachments collection</summary>
        public AttachmentRequestBuilder this[string position] { get {
            return new AttachmentRequestBuilder { HttpCore = HttpCore, SerializerFactory = SerializerFactory, CurrentPath = CurrentPath + PathSegment  + "/" + position};
        } }
        /// <summary>
        /// Get attachments from users
        /// <param name="q">Request query parameters</param>
        /// <param name="h">Request headers</param>
        /// <param name="responseHandler">Response handler to use in place of the default response handling provided by the core service</param>
        /// </summary>
        public async Task<AttachmentsResponse> GetAsync(Action<GetQueryParameters> q = default, Action<IDictionary<string, string>> h = default, IResponseHandler responseHandler = default) {
            var requestInfo = CreateGetRequestInfo(
                q, h
            );
            return await HttpCore.SendAsync<AttachmentsResponse>(requestInfo, responseHandler);
        }
        /// <summary>
        /// Get attachments from users
        /// <param name="q">Request query parameters</param>
        /// <param name="h">Request headers</param>
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
        /// Create new navigation property to attachments for users
        /// <param name="body"></param>
        /// <param name="h">Request headers</param>
        /// <param name="responseHandler">Response handler to use in place of the default response handling provided by the core service</param>
        /// </summary>
        public async Task<Attachment> PostAsync(Attachment body, Action<IDictionary<string, string>> h = default, IResponseHandler responseHandler = default) {
            var requestInfo = CreatePostRequestInfo(
                body, h
            );
            return await HttpCore.SendAsync<Attachment>(requestInfo, responseHandler);
        }
        /// <summary>
        /// Create new navigation property to attachments for users
        /// <param name="body"></param>
        /// <param name="h">Request headers</param>
        /// </summary>
        public RequestInfo CreatePostRequestInfo(Attachment body, Action<IDictionary<string, string>> h = default) {
            var requestInfo = new RequestInfo {
                HttpMethod = HttpMethod.POST,
                URI = new Uri(CurrentPath + PathSegment),
            };
            requestInfo.SetJsonContentFromParsable(body, SerializerFactory);
            h?.Invoke(requestInfo.Headers);
            return requestInfo;
        }
        /// <summary>Path segment to use to build the URL for the current request builder</summary>
        private string PathSegment { get; } = "/attachments";
        /// <summary>Current path for the request</summary>
        public string CurrentPath { get; set; }
        /// <summary>Core service to use to execute the requests</summary>
        public IHttpCore HttpCore { get; set; }
        /// <summary>Factory to use to get a serializer for payload serialization</summary>
        public Func<string, ISerializationWriter> SerializerFactory { get; set; }
        /// <summary>Get attachments from users</summary>
        public class GetQueryParameters : QueryParametersBase {
            /// <summary>Show only the first n items</summary>
            public int? Top { get; set; }
            /// <summary>Skip the first n items</summary>
            public int? Skip { get; set; }
            /// <summary>Search items by search phrases</summary>
            public string Search { get; set; }
            /// <summary>Filter items by property values</summary>
            public string Filter { get; set; }
            /// <summary>Include count of items</summary>
            public bool? Count { get; set; }
            /// <summary>Order items by property values</summary>
            public string[] Orderby { get; set; }
            /// <summary>Select properties to be returned</summary>
            public string[] Select { get; set; }
            /// <summary>Expand related entities</summary>
            public string[] Expand { get; set; }
        }
    }
}
