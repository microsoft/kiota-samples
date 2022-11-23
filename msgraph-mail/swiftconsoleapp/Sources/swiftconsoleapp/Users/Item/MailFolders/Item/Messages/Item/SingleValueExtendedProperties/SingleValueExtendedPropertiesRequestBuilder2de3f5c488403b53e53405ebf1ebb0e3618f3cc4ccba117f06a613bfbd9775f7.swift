import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Users.Item.MailFolders.Item.Messages.Item.SingleValueExtendedProperties {
    /// <summary>Builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}\messages\{message-id}\singleValueExtendedProperties</summary>
    public class SingleValueExtendedPropertiesRequestBuilder {
        private var pathParameters: [String:String]?
        private var requestAdapter: RequestAdapter?
        private var urlTemplate: String = "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/messages/{message%2Did}/singleValueExtendedProperties{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}"
        /// <summary>The collection of single-value extended properties defined for the message. Nullable.</summary>
        public class SingleValueExtendedPropertiesRequestBuilderGetQueryParameters {
            public var count: SingleValueExtendedPropertiesRequestBuilderGetQueryParametersboolean?
            public var expand: [SingleValueExtendedPropertiesRequestBuilderGetQueryParametersstring]?
            public var filter: SingleValueExtendedPropertiesRequestBuilderGetQueryParametersstring?
            public var orderby: [SingleValueExtendedPropertiesRequestBuilderGetQueryParametersstring]?
            public var search: SingleValueExtendedPropertiesRequestBuilderGetQueryParametersstring?
            public var select: [SingleValueExtendedPropertiesRequestBuilderGetQueryParametersstring]?
            public var skip: SingleValueExtendedPropertiesRequestBuilderGetQueryParametersinteger?
            public var top: SingleValueExtendedPropertiesRequestBuilderGetQueryParametersinteger?
        }
        /// <summary>Configuration for the request such as headers, query parameters, and middleware options.</summary>
        public class SingleValueExtendedPropertiesRequestBuilderGetRequestConfiguration {
            public var headers: IDictionary<string, string>?
            public var options: [RequestOption]
            public var queryParameters: SingleValueExtendedPropertiesRequestBuilderGetQueryParameters?
        }
        /// <summary>Configuration for the request such as headers, query parameters, and middleware options.</summary>
        public class SingleValueExtendedPropertiesRequestBuilderPostRequestConfiguration {
            public var headers: IDictionary<string, string>?
            public var options: [RequestOption]
        }
    }
}
