using Kiota.Abstractions;
using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Threading.Tasks;
namespace Graphdotnetv4.Users.Messages.Content {
    /// <summary>Builds and executes requests for operations under \users\{user-id}\messages\{message-id}\$value</summary>
    public class ContentRequestBuilder {
        /// <summary>Current path for the request</summary>
        public string CurrentPath { get; set; }
        /// <summary>Core service to use to execute the requests</summary>
        public IHttpCore HttpCore { get; set; }
        /// <summary>Path segment to use to build the URL for the current request builder</summary>
        private string PathSegment { get; set; } = "/$value";
        /// <summary>Factory to use to get a serializer for payload serialization</summary>
        public ISerializationWriterFactory SerializerFactory { get; set; }
        /// <summary>
        /// Get media content for the navigation property messages from users
        /// <param name="h">Request headers</param>
        /// </summary>
        public RequestInfo CreateGetRequestInfo(Action<IDictionary<string, string>> h = default) {
            var requestInfo = new RequestInfo {
                HttpMethod = HttpMethod.GET,
                URI = new Uri(CurrentPath + PathSegment),
            };
            h?.Invoke(requestInfo.Headers);
            return requestInfo;
        }
        /// <summary>
        /// Update media content for the navigation property messages in users
        /// <param name="body">Binary request body</param>
        /// <param name="h">Request headers</param>
        /// <typeparam name="T">The generic type for the model, must inherit Stream</typeparam>
        /// </summary>
        public RequestInfo CreatePutRequestInfo<T>(T body, Action<IDictionary<string, string>> h = default) where T : Stream, IParsable<T>, new(){
            var requestInfo = new RequestInfo {
                HttpMethod = HttpMethod.PUT,
                URI = new Uri(CurrentPath + PathSegment),
            };
            requestInfo.SetContentFromParsable(body, SerializerFactory, "");
            h?.Invoke(requestInfo.Headers);
            return requestInfo;
        }
        /// <summary>
        /// Get media content for the navigation property messages from users
        /// <param name="h">Request headers</param>
        /// <param name="responseHandler">Response handler to use in place of the default response handling provided by the core service</param>
        /// <typeparam name="T">The generic type for the model, must inherit Stream</typeparam>
        /// </summary>
        public async Task<T> GetAsync<T>(Action<IDictionary<string, string>> h = default, IResponseHandler responseHandler = default) where T : Stream, IParsable<T>, new(){
            var requestInfo = CreateGetRequestInfo(
                h
            );
            return await HttpCore.SendPrimitiveAsync<T>(requestInfo, responseHandler);
        }
        /// <summary>
        /// Update media content for the navigation property messages in users
        /// <param name="body">Binary request body</param>
        /// <param name="h">Request headers</param>
        /// <param name="responseHandler">Response handler to use in place of the default response handling provided by the core service</param>
        /// <typeparam name="T">The generic type for the model, must inherit Stream</typeparam>
        /// </summary>
        public async Task PutAsync<T>(T body, Action<IDictionary<string, string>> h = default, IResponseHandler responseHandler = default) where T : Stream, IParsable<T>, new(){
            var requestInfo = CreatePutRequestInfo(
                body, h
            );
            await HttpCore.SendNoContentAsync(requestInfo, responseHandler);
        }
    }
}
