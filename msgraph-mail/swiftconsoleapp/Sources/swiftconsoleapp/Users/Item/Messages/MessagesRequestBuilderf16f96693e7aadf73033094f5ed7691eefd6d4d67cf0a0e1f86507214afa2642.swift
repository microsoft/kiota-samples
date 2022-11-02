import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Users.Item.Messages {
    /// <summary>Builds and executes requests for operations under \users\{user-id}\messages</summary>
    public class MessagesRequestBuilder {
        private var pathParameters: [String:String]?
        private var requestAdapter: RequestAdapter?
        private var urlTemplate: String = "{+baseurl}/users/{user%2Did}/messages{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select}"
        /// <summary>Get the messages in the signed-in user's mailbox (including the Deleted Items and Clutter folders). Depending on the page size and mailbox data, getting messages from a mailbox can incur multiple requests. The default page size is 10 messages. Use `$top` to customize the page size, within the range of 1 and 1000. To improve the operation response time, use `$select` to specify the exact properties you need; see example 1 below. Fine-tune the values for `$select` and `$top`, especially when you must use a larger page size, as returning a page with hundreds of messages each with a full response payload may trigger the gateway timeout (HTTP 504). To get the next page of messages, simply apply the entire URL returned in `@odata.nextLink` to the next get-messages request. This URL includes any query parameters you may have specified in the initial request.  Do not try to extract the `$skip` value from the `@odata.nextLink` URL to manipulate responses. This API uses the `$skip` value to keep count of all the items it has gone through in the user's mailbox to return a page of message-type items. It's therefore possible that even in the initial response, the `$skip` value is larger than the page size. For more information, see Paging Microsoft Graph data in your app. Currently, this operation returns message bodies in only HTML format. There are two scenarios where an app can get messages in another user's mail folder:</summary>
        public class MessagesRequestBuilderGetQueryParameters {
            public var count: MessagesRequestBuilderGetQueryParametersboolean?
            public var filter: MessagesRequestBuilderGetQueryParametersstring?
            public var orderby: [MessagesRequestBuilderGetQueryParametersstring]?
            public var search: MessagesRequestBuilderGetQueryParametersstring?
            public var select: [MessagesRequestBuilderGetQueryParametersstring]?
            public var skip: MessagesRequestBuilderGetQueryParametersinteger?
            public var top: MessagesRequestBuilderGetQueryParametersinteger?
        }
        /// <summary>Configuration for the request such as headers, query parameters, and middleware options.</summary>
        public class MessagesRequestBuilderGetRequestConfiguration {
            public var headers: IDictionary<string, string>?
            public var options: [RequestOption]
            public var queryParameters: MessagesRequestBuilderGetQueryParameters?
        }
        /// <summary>Configuration for the request such as headers, query parameters, and middleware options.</summary>
        public class MessagesRequestBuilderPostRequestConfiguration {
            public var headers: IDictionary<string, string>?
            public var options: [RequestOption]
        }
    }
}
