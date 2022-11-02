import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Users.Item.Messages.Item.MultiValueExtendedProperties {
    /// <summary>Builds and executes requests for operations under \users\{user-id}\messages\{message-id}\multiValueExtendedProperties</summary>
    public class MultiValueExtendedPropertiesRequestBuilder {
        private var pathParameters: [String:String]?
        private var requestAdapter: RequestAdapter?
        private var urlTemplate: String = "{+baseurl}/users/{user%2Did}/messages/{message%2Did}/multiValueExtendedProperties{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}"
        /// <summary>The collection of multi-value extended properties defined for the message. Nullable.</summary>
        public class MultiValueExtendedPropertiesRequestBuilderGetQueryParameters {
            public var count: MultiValueExtendedPropertiesRequestBuilderGetQueryParametersboolean?
            public var expand: [MultiValueExtendedPropertiesRequestBuilderGetQueryParametersstring]?
            public var filter: MultiValueExtendedPropertiesRequestBuilderGetQueryParametersstring?
            public var orderby: [MultiValueExtendedPropertiesRequestBuilderGetQueryParametersstring]?
            public var search: MultiValueExtendedPropertiesRequestBuilderGetQueryParametersstring?
            public var select: [MultiValueExtendedPropertiesRequestBuilderGetQueryParametersstring]?
            public var skip: MultiValueExtendedPropertiesRequestBuilderGetQueryParametersinteger?
            public var top: MultiValueExtendedPropertiesRequestBuilderGetQueryParametersinteger?
        }
        /// <summary>Configuration for the request such as headers, query parameters, and middleware options.</summary>
        public class MultiValueExtendedPropertiesRequestBuilderGetRequestConfiguration {
            public var headers: IDictionary<string, string>?
            public var options: [RequestOption]
            public var queryParameters: MultiValueExtendedPropertiesRequestBuilderGetQueryParameters?
        }
        /// <summary>Configuration for the request such as headers, query parameters, and middleware options.</summary>
        public class MultiValueExtendedPropertiesRequestBuilderPostRequestConfiguration {
            public var headers: IDictionary<string, string>?
            public var options: [RequestOption]
        }
    }
}
