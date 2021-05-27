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
    /// <summary>Builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}\messages\{message-id}</summary>
    public class MessageRequestBuilder {
        public AttachmentsRequestBuilder Attachments { get =>
            new AttachmentsRequestBuilder { HttpCore = HttpCore, SerializerFactory = SerializerFactory, CurrentPath = CurrentPath + PathSegment };
        }
        public ContentRequestBuilder Content { get =>
            new ContentRequestBuilder { HttpCore = HttpCore, SerializerFactory = SerializerFactory, CurrentPath = CurrentPath + PathSegment };
        }
        /// <summary>Current path for the request</summary>
        public string CurrentPath { get; set; }
        public ExtensionsRequestBuilder Extensions { get =>
            new ExtensionsRequestBuilder { HttpCore = HttpCore, SerializerFactory = SerializerFactory, CurrentPath = CurrentPath + PathSegment };
        }
        /// <summary>Core service to use to execute the requests</summary>
        public IHttpCore HttpCore { get; set; }
        public MultiValueExtendedPropertiesRequestBuilder MultiValueExtendedProperties { get =>
            new MultiValueExtendedPropertiesRequestBuilder { HttpCore = HttpCore, SerializerFactory = SerializerFactory, CurrentPath = CurrentPath + PathSegment };
        }
        /// <summary>Path segment to use to build the URL for the current request builder</summary>
        private string PathSegment { get; set; } = "";
        /// <summary>Factory to use to get a serializer for payload serialization</summary>
        public ISerializationWriterFactory SerializerFactory { get; set; }
        public SingleValueExtendedPropertiesRequestBuilder SingleValueExtendedProperties { get =>
            new SingleValueExtendedPropertiesRequestBuilder { HttpCore = HttpCore, SerializerFactory = SerializerFactory, CurrentPath = CurrentPath + PathSegment };
        }
        /// <summary>
        /// Delete navigation property messages for users
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
        /// Get messages from users
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
        /// Update the navigation property messages in users
        /// <param name="body">The request body</param>
        /// <param name="h">Request headers</param>
        /// </summary>
        public RequestInfo CreatePatchRequestInfo<T>(T body, Action<IDictionary<string, string>> h = default) where T : Message, IParsable<T>, new(){
            var requestInfo = new RequestInfo {
                HttpMethod = HttpMethod.PATCH,
                URI = new Uri(CurrentPath + PathSegment),
            };
            requestInfo.SetContentFromParsable(body, SerializerFactory, "application/json");
            h?.Invoke(requestInfo.Headers);
            return requestInfo;
        }
        /// <summary>
        /// Delete navigation property messages for users
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
        /// Get messages from users
        /// <param name="h">Request headers</param>
        /// <param name="q">Request query parameters</param>
        /// <param name="responseHandler">Response handler to use in place of the default response handling provided by the core service</param>
        /// </summary>
        public async Task<T> GetAsync<T>(Action<GetQueryParameters> q = default, Action<IDictionary<string, string>> h = default, IResponseHandler responseHandler = default) where T : Message, IParsable<T>, new(){
            var requestInfo = CreateGetRequestInfo(
                q, h
            );
            return await HttpCore.SendAsync<T>(requestInfo, responseHandler);
        }
        /// <summary>
        /// Update the navigation property messages in users
        /// <param name="body">The request body</param>
        /// <param name="h">Request headers</param>
        /// <param name="responseHandler">Response handler to use in place of the default response handling provided by the core service</param>
        /// </summary>
        public async Task PatchAsync<T>(T body, Action<IDictionary<string, string>> h = default, IResponseHandler responseHandler = default) where T : Message, IParsable<T>, new(){
            var requestInfo = CreatePatchRequestInfo(
                body, h
            );
            await HttpCore.SendNoContentAsync(requestInfo, responseHandler);
        }
        /// <summary>Get messages from users</summary>
        public class GetQueryParameters : QueryParametersBase {
            /// <summary>Expand related entities</summary>
            public string[] Expand { get; set; }
            /// <summary>Select properties to be returned</summary>
            public string[] Select { get; set; }
        }
    }
}
