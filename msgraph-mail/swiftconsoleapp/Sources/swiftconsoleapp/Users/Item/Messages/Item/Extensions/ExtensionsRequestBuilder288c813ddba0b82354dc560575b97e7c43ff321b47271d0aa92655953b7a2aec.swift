import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Users.Item.Messages.Item.Extensions {
    /// <summary>Builds and executes requests for operations under \users\{user-id}\messages\{message-id}\extensions</summary>
    public class ExtensionsRequestBuilder {
        private var pathParameters: [String:String]?
        private var requestAdapter: RequestAdapter?
        private var urlTemplate: String = "{+baseurl}/users/{user%2Did}/messages/{message%2Did}/extensions{?%24top,%24skip,%24filter,%24count,%24orderby,%24select,%24expand}"
        /// <summary>The collection of open extensions defined for the message. Nullable.</summary>
        public class ExtensionsRequestBuilderGetQueryParameters {
            public var count: ExtensionsRequestBuilderGetQueryParametersboolean?
            public var expand: [ExtensionsRequestBuilderGetQueryParametersstring]?
            public var filter: ExtensionsRequestBuilderGetQueryParametersstring?
            public var orderby: [ExtensionsRequestBuilderGetQueryParametersstring]?
            public var select: [ExtensionsRequestBuilderGetQueryParametersstring]?
            public var skip: ExtensionsRequestBuilderGetQueryParametersinteger?
            public var top: ExtensionsRequestBuilderGetQueryParametersinteger?
        }
        /// <summary>Configuration for the request such as headers, query parameters, and middleware options.</summary>
        public class ExtensionsRequestBuilderGetRequestConfiguration {
            public var headers: IDictionary<string, string>?
            public var options: [RequestOption]
            public var queryParameters: ExtensionsRequestBuilderGetQueryParameters?
        }
        /// <summary>Configuration for the request such as headers, query parameters, and middleware options.</summary>
        public class ExtensionsRequestBuilderPostRequestConfiguration {
            public var headers: IDictionary<string, string>?
            public var options: [RequestOption]
        }
    }
}
