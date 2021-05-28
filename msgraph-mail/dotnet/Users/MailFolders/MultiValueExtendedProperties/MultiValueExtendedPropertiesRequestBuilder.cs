using Kiota.Abstractions;
using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Threading.Tasks;
using Graphdotnetv4.Users.MailFolders.MultiValueExtendedProperties.Item;
namespace Graphdotnetv4.Users.MailFolders.MultiValueExtendedProperties {
    /// <summary>Builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}\multiValueExtendedProperties</summary>
    public class MultiValueExtendedPropertiesRequestBuilder {
        /// <summary>Current path for the request</summary>
        public string CurrentPath { get; set; }
        /// <summary>Core service to use to execute the requests</summary>
        public IHttpCore HttpCore { get; set; }
        /// <summary>Path segment to use to build the URL for the current request builder</summary>
        private string PathSegment { get; set; } = "/multiValueExtendedProperties";
        /// <summary>Factory to use to get a serializer for payload serialization</summary>
        public ISerializationWriterFactory SerializerFactory { get; set; }
        /// <summary>Gets an item from the Graphdotnetv4.users.mailFolders.multiValueExtendedProperties collection</summary>
        public MultiValueLegacyExtendedPropertyRequestBuilder this[string position] { get {
            return new MultiValueLegacyExtendedPropertyRequestBuilder { HttpCore = HttpCore, SerializerFactory = SerializerFactory, CurrentPath = CurrentPath + PathSegment  + "/" + position};
        } }
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
        /// Create new navigation property to multiValueExtendedProperties for users
        /// <param name="body">The request body</param>
        /// <param name="h">Request headers</param>
        /// <typeparam name="T">The generic type for the model, must inherit MultiValueLegacyExtendedProperty</typeparam>
        /// </summary>
        public RequestInfo CreatePostRequestInfo<T>(T body, Action<IDictionary<string, string>> h = default) where T : MultiValueLegacyExtendedProperty, IParsable<T>, new(){
            var requestInfo = new RequestInfo {
                HttpMethod = HttpMethod.POST,
                URI = new Uri(CurrentPath + PathSegment),
            };
            requestInfo.SetContentFromParsable(body, SerializerFactory, "application/json");
            h?.Invoke(requestInfo.Headers);
            return requestInfo;
        }
        /// <summary>
        /// Get multiValueExtendedProperties from users
        /// <param name="h">Request headers</param>
        /// <param name="q">Request query parameters</param>
        /// <param name="responseHandler">Response handler to use in place of the default response handling provided by the core service</param>
        /// <typeparam name="T">The generic type for the model, must inherit MultiValueExtendedPropertiesResponse</typeparam>
        /// </summary>
        public async Task<T> GetAsync<T>(Action<GetQueryParameters> q = default, Action<IDictionary<string, string>> h = default, IResponseHandler responseHandler = default) where T : MultiValueExtendedPropertiesResponse, IParsable<T>, new(){
            var requestInfo = CreateGetRequestInfo(
                q, h
            );
            return await HttpCore.SendAsync<T>(requestInfo, responseHandler);
        }
        /// <summary>
        /// Create new navigation property to multiValueExtendedProperties for users
        /// <param name="body">The request body</param>
        /// <param name="h">Request headers</param>
        /// <param name="responseHandler">Response handler to use in place of the default response handling provided by the core service</param>
        /// <typeparam name="T">The generic type for the model, must inherit MultiValueLegacyExtendedProperty</typeparam>
        /// </summary>
        public async Task<T> PostAsync<T>(T body, Action<IDictionary<string, string>> h = default, IResponseHandler responseHandler = default) where T : MultiValueLegacyExtendedProperty, IParsable<T>, new(){
            var requestInfo = CreatePostRequestInfo(
                body, h
            );
            return await HttpCore.SendAsync<T>(requestInfo, responseHandler);
        }
        /// <summary>Get multiValueExtendedProperties from users</summary>
        public class GetQueryParameters : QueryParametersBase {
            /// <summary>Include count of items</summary>
            public bool? Count { get; set; }
            /// <summary>Expand related entities</summary>
            public string[] Expand { get; set; }
            /// <summary>Filter items by property values</summary>
            public string Filter { get; set; }
            /// <summary>Order items by property values</summary>
            public string[] Orderby { get; set; }
            /// <summary>Search items by search phrases</summary>
            public string Search { get; set; }
            /// <summary>Select properties to be returned</summary>
            public string[] Select { get; set; }
            /// <summary>Skip the first n items</summary>
            public int? Skip { get; set; }
            /// <summary>Show only the first n items</summary>
            public int? Top { get; set; }
        }
    }
}
