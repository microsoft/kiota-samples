using Graphdotnetv4.Models.Microsoft.Graph;
using Graphdotnetv4.Users.Item.MailFolders.Item.Messages.Item.Attachments.Item;
using Microsoft.Kiota.Abstractions;
using Microsoft.Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Threading.Tasks;
namespace Graphdotnetv4.Users.Item.MailFolders.Item.Messages.Item.Attachments {
    /// <summary>Builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}\messages\{message-id}\attachments</summary>
    public class AttachmentsRequestBuilder {
        /// <summary>The request adapter to use to execute the requests.</summary>
        private IRequestAdapter RequestAdapter { get; set; }
        /// <summary>Url template to use to build the URL for the current request builder</summary>
        private string UrlTemplate { get; set; }
        /// <summary>Url template parameters for the request</summary>
        private Dictionary<string, string> UrlTemplateParameters { get; set; }
        /// <summary>Gets an item from the Graphdotnetv4.users.item.mailFolders.item.messages.item.attachments.item collection</summary>
        public AttachmentRequestBuilder this[string position] { get {
            var urlTplParams = new Dictionary<string, string>(UrlTemplateParameters);
            urlTplParams.Add("attachment_id", position);
            return new AttachmentRequestBuilder(urlTplParams, RequestAdapter);
        } }
        /// <summary>
        /// Instantiates a new AttachmentsRequestBuilder and sets the default values.
        /// <param name="requestAdapter">The request adapter to use to execute the requests.</param>
        /// <param name="urlTemplateParameters">Url template parameters for the request</param>
        /// </summary>
        public AttachmentsRequestBuilder(Dictionary<string, string> urlTemplateParameters, IRequestAdapter requestAdapter) {
            _ = requestAdapter ?? throw new ArgumentNullException(nameof(requestAdapter));
            _ = urlTemplateParameters ?? throw new ArgumentNullException(nameof(urlTemplateParameters));
            UrlTemplate = "https://graph.microsoft.com/v1.0/users/{user_id}/mailFolders/{mailFolder_id}/messages/{message_id}/attachments{?top,skip,search,filter,count,orderby,select,expand}";
            var urlTplParams = new Dictionary<string, string>(urlTemplateParameters);
            UrlTemplateParameters = urlTplParams;
            RequestAdapter = requestAdapter;
        }
        /// <summary>
        /// Instantiates a new AttachmentsRequestBuilder and sets the default values.
        /// <param name="rawUrl">The raw URL to use for the request builder.</param>
        /// <param name="requestAdapter">The request adapter to use to execute the requests.</param>
        /// </summary>
        public AttachmentsRequestBuilder(string rawUrl, IRequestAdapter requestAdapter) {
            if(string.IsNullOrEmpty(rawUrl)) throw new ArgumentNullException(nameof(rawUrl));
            _ = requestAdapter ?? throw new ArgumentNullException(nameof(requestAdapter));
            UrlTemplate = "https://graph.microsoft.com/v1.0/users/{user_id}/mailFolders/{mailFolder_id}/messages/{message_id}/attachments{?top,skip,search,filter,count,orderby,select,expand}";
            var urlTplParams = new Dictionary<string, string>();
            urlTplParams.Add("request-raw-url", rawUrl);
            UrlTemplateParameters = urlTplParams;
            RequestAdapter = requestAdapter;
        }
        /// <summary>
        /// The fileAttachment and itemAttachment attachments for the message.
        /// <param name="h">Request headers</param>
        /// <param name="o">Request options</param>
        /// <param name="q">Request query parameters</param>
        /// </summary>
        public RequestInformation CreateGetRequestInformation(Action<GetQueryParameters> q = default, Action<IDictionary<string, string>> h = default, IEnumerable<IRequestOption> o = default) {
            var requestInfo = new RequestInformation {
                HttpMethod = HttpMethod.GET,
                UrlTemplate = UrlTemplate,
                UrlTemplateParameters = UrlTemplateParameters,
            };
            if (q != null) {
                var qParams = new GetQueryParameters();
                q.Invoke(qParams);
                qParams.AddQueryParameters(requestInfo.QueryParameters);
            }
            h?.Invoke(requestInfo.Headers);
            requestInfo.AddRequestOptions(o?.ToArray());
            return requestInfo;
        }
        /// <summary>
        /// The fileAttachment and itemAttachment attachments for the message.
        /// <param name="body"></param>
        /// <param name="h">Request headers</param>
        /// <param name="o">Request options</param>
        /// </summary>
        public RequestInformation CreatePostRequestInformation(Attachment body, Action<IDictionary<string, string>> h = default, IEnumerable<IRequestOption> o = default) {
            _ = body ?? throw new ArgumentNullException(nameof(body));
            var requestInfo = new RequestInformation {
                HttpMethod = HttpMethod.POST,
                UrlTemplate = UrlTemplate,
                UrlTemplateParameters = UrlTemplateParameters,
            };
            requestInfo.SetContentFromParsable(RequestAdapter, "application/json", body);
            h?.Invoke(requestInfo.Headers);
            requestInfo.AddRequestOptions(o?.ToArray());
            return requestInfo;
        }
        /// <summary>
        /// The fileAttachment and itemAttachment attachments for the message.
        /// <param name="h">Request headers</param>
        /// <param name="o">Request options</param>
        /// <param name="q">Request query parameters</param>
        /// <param name="responseHandler">Response handler to use in place of the default response handling provided by the core service</param>
        /// </summary>
        public async Task<AttachmentsResponse> GetAsync(Action<GetQueryParameters> q = default, Action<IDictionary<string, string>> h = default, IEnumerable<IRequestOption> o = default, IResponseHandler responseHandler = default) {
            var requestInfo = CreateGetRequestInformation(q, h, o);
            return await RequestAdapter.SendAsync<AttachmentsResponse>(requestInfo, responseHandler);
        }
        /// <summary>
        /// The fileAttachment and itemAttachment attachments for the message.
        /// <param name="body"></param>
        /// <param name="h">Request headers</param>
        /// <param name="o">Request options</param>
        /// <param name="responseHandler">Response handler to use in place of the default response handling provided by the core service</param>
        /// </summary>
        public async Task<Attachment> PostAsync(Attachment body, Action<IDictionary<string, string>> h = default, IEnumerable<IRequestOption> o = default, IResponseHandler responseHandler = default) {
            _ = body ?? throw new ArgumentNullException(nameof(body));
            var requestInfo = CreatePostRequestInformation(body, h, o);
            return await RequestAdapter.SendAsync<Attachment>(requestInfo, responseHandler);
        }
        /// <summary>The fileAttachment and itemAttachment attachments for the message.</summary>
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
