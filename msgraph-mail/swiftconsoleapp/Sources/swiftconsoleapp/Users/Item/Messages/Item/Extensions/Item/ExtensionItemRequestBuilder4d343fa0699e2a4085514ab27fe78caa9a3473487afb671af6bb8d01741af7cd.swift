import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Users.Item.Messages.Item.Extensions.Item {
    /// <summary>Builds and executes requests for operations under \users\{user-id}\messages\{message-id}\extensions\{extension-id}</summary>
    public class ExtensionItemRequestBuilder {
        private var pathParameters: [String:String]?
        private var requestAdapter: RequestAdapter?
        private var urlTemplate: String = "{+baseurl}/users/{user%2Did}/messages/{message%2Did}/extensions/{extension%2Did}{?%24select,%24expand}"
        /// <summary>Configuration for the request such as headers, query parameters, and middleware options.</summary>
        public class ExtensionItemRequestBuilderDeleteRequestConfiguration {
            public var headers: IDictionary<string, string>?
            public var options: [RequestOption]
        }
        /// <summary>The collection of open extensions defined for the message. Nullable.</summary>
        public class ExtensionItemRequestBuilderGetQueryParameters {
            public var expand: [ExtensionItemRequestBuilderGetQueryParametersstring]?
            public var select: [ExtensionItemRequestBuilderGetQueryParametersstring]?
        }
        /// <summary>Configuration for the request such as headers, query parameters, and middleware options.</summary>
        public class ExtensionItemRequestBuilderGetRequestConfiguration {
            public var headers: IDictionary<string, string>?
            public var options: [RequestOption]
            public var queryParameters: ExtensionItemRequestBuilderGetQueryParameters?
        }
        /// <summary>Configuration for the request such as headers, query parameters, and middleware options.</summary>
        public class ExtensionItemRequestBuilderPatchRequestConfiguration {
            public var headers: IDictionary<string, string>?
            public var options: [RequestOption]
        }
    }
}
