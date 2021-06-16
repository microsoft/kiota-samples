using Microsoft.Kiota.Abstractions;
using Microsoft.Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Threading.Tasks;
namespace Graphdotnetv4.Users.MailFolders.Messages.MultiValueExtendedProperties.Item {
    /// <summary>Builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}\messages\{message-id}\multiValueExtendedProperties\{multiValueLegacyExtendedProperty-id}</summary>
    public class MultiValueLegacyExtendedPropertyRequestBuilder {
        /// <summary>Current path for the request</summary>
        public string CurrentPath { get; set; }
        /// <summary>Core service to use to execute the requests</summary>
        public IHttpCore HttpCore { get; set; }
        /// <summary>Path segment to use to build the URL for the current request builder</summary>
        private string PathSegment { get; set; }
        /// <summary>Factory to use to get a serializer for payload serialization</summary>
        public ISerializationWriterFactory SerializerFactory { get; set; }
        /// <summary>
        /// Instantiates a new MultiValueLegacyExtendedPropertyRequestBuilder and sets the default values.
        /// </summary>
        public MultiValueLegacyExtendedPropertyRequestBuilder() {
            PathSegment = "";
        }
        /// <summary>
        /// Delete navigation property multiValueExtendedProperties for users
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
        /// Get multiValueExtendedProperties from users
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
        /// Update the navigation property multiValueExtendedProperties in users
        /// <param name="body"></param>
        /// <param name="h">Request headers</param>
        /// </summary>
        public RequestInfo CreatePatchRequestInfo(MultiValueLegacyExtendedProperty body, Action<IDictionary<string, string>> h = default) {
            var requestInfo = new RequestInfo {
                HttpMethod = HttpMethod.PATCH,
                URI = new Uri(CurrentPath + PathSegment),
            };
            requestInfo.SetContentFromParsable(body, SerializerFactory, "application/json");
            h?.Invoke(requestInfo.Headers);
            return requestInfo;
        }
        /// <summary>
        /// Delete navigation property multiValueExtendedProperties for users
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
        /// Get multiValueExtendedProperties from users
        /// <param name="h">Request headers</param>
        /// <param name="q">Request query parameters</param>
        /// <param name="responseHandler">Response handler to use in place of the default response handling provided by the core service</param>
        /// </summary>
        public async Task<MultiValueLegacyExtendedProperty> GetAsync(Action<GetQueryParameters> q = default, Action<IDictionary<string, string>> h = default, IResponseHandler responseHandler = default) {
            var requestInfo = CreateGetRequestInfo(
                q, h
            );
            return await HttpCore.SendAsync<MultiValueLegacyExtendedProperty>(requestInfo, responseHandler);
        }
        /// <summary>
        /// Update the navigation property multiValueExtendedProperties in users
        /// <param name="body"></param>
        /// <param name="h">Request headers</param>
        /// <param name="responseHandler">Response handler to use in place of the default response handling provided by the core service</param>
        /// </summary>
        public async Task PatchAsync(MultiValueLegacyExtendedProperty body, Action<IDictionary<string, string>> h = default, IResponseHandler responseHandler = default) {
            var requestInfo = CreatePatchRequestInfo(
                body, h
            );
            await HttpCore.SendNoContentAsync(requestInfo, responseHandler);
        }
        /// <summary>Get multiValueExtendedProperties from users</summary>
        public class GetQueryParameters : QueryParametersBase {
            /// <summary>Expand related entities</summary>
            public string[] Expand { get; set; }
            /// <summary>Select properties to be returned</summary>
            public string[] Select { get; set; }
        }
    }
}
