import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Users.Item.Messages.Item.Value {
    /// <summary>Builds and executes requests for operations under \users\{user-id}\messages\{message-id}\$value</summary>
    public class ContentRequestBuilder {
        private var pathParameters: [String:String]?
        private var requestAdapter: RequestAdapter?
        private var urlTemplate: String = "{+baseurl}/users/{user%2Did}/messages/{message%2Did}/$value"
        /// <summary>Configuration for the request such as headers, query parameters, and middleware options.</summary>
        public class ContentRequestBuilderGetRequestConfiguration {
            public var headers: IDictionary<string, string>?
            public var options: [RequestOption]
        }
        /// <summary>Configuration for the request such as headers, query parameters, and middleware options.</summary>
        public class ContentRequestBuilderPutRequestConfiguration {
            public var headers: IDictionary<string, string>?
            public var options: [RequestOption]
        }
    }
}
