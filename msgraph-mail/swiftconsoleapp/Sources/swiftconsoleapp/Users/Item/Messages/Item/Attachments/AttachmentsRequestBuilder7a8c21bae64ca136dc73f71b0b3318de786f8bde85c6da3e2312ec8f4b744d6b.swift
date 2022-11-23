import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Users.Item.Messages.Item.Attachments {
    /// <summary>Builds and executes requests for operations under \users\{user-id}\messages\{message-id}\attachments</summary>
    public class AttachmentsRequestBuilder {
        private var pathParameters: [String:String]?
        private var requestAdapter: RequestAdapter?
        private var urlTemplate: String = "{+baseurl}/users/{user%2Did}/messages/{message%2Did}/attachments{?%24top,%24skip,%24filter,%24count,%24orderby,%24select,%24expand}"
        /// <summary>Retrieve a list of attachment objects.</summary>
        public class AttachmentsRequestBuilderGetQueryParameters {
            public var count: AttachmentsRequestBuilderGetQueryParametersboolean?
            public var expand: [AttachmentsRequestBuilderGetQueryParametersstring]?
            public var filter: AttachmentsRequestBuilderGetQueryParametersstring?
            public var orderby: [AttachmentsRequestBuilderGetQueryParametersstring]?
            public var select: [AttachmentsRequestBuilderGetQueryParametersstring]?
            public var skip: AttachmentsRequestBuilderGetQueryParametersinteger?
            public var top: AttachmentsRequestBuilderGetQueryParametersinteger?
        }
        /// <summary>Configuration for the request such as headers, query parameters, and middleware options.</summary>
        public class AttachmentsRequestBuilderGetRequestConfiguration {
            public var headers: IDictionary<string, string>?
            public var options: [RequestOption]
            public var queryParameters: AttachmentsRequestBuilderGetQueryParameters?
        }
        /// <summary>Configuration for the request such as headers, query parameters, and middleware options.</summary>
        public class AttachmentsRequestBuilderPostRequestConfiguration {
            public var headers: IDictionary<string, string>?
            public var options: [RequestOption]
        }
    }
}
