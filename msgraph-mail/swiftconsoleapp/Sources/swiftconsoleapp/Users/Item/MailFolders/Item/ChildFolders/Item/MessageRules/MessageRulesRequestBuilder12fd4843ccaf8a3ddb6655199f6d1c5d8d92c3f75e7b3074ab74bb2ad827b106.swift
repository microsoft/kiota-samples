import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Users.Item.MailFolders.Item.ChildFolders.Item.MessageRules {
    /// <summary>Builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}\childFolders\{mailFolder-id1}\messageRules</summary>
    public class MessageRulesRequestBuilder {
        private var pathParameters: [String:String]?
        private var requestAdapter: RequestAdapter?
        private var urlTemplate: String = "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/childFolders/{mailFolder%2Did1}/messageRules{?%24top,%24skip,%24filter,%24count,%24orderby,%24select}"
        /// <summary>Get all the messageRule objects defined for the user's inbox.</summary>
        public class MessageRulesRequestBuilderGetQueryParameters {
            public var count: MessageRulesRequestBuilderGetQueryParametersboolean?
            public var filter: MessageRulesRequestBuilderGetQueryParametersstring?
            public var orderby: [MessageRulesRequestBuilderGetQueryParametersstring]?
            public var select: [MessageRulesRequestBuilderGetQueryParametersstring]?
            public var skip: MessageRulesRequestBuilderGetQueryParametersinteger?
            public var top: MessageRulesRequestBuilderGetQueryParametersinteger?
        }
        /// <summary>Configuration for the request such as headers, query parameters, and middleware options.</summary>
        public class MessageRulesRequestBuilderGetRequestConfiguration {
            public var headers: IDictionary<string, string>?
            public var options: [RequestOption]
            public var queryParameters: MessageRulesRequestBuilderGetQueryParameters?
        }
        /// <summary>Configuration for the request such as headers, query parameters, and middleware options.</summary>
        public class MessageRulesRequestBuilderPostRequestConfiguration {
            public var headers: IDictionary<string, string>?
            public var options: [RequestOption]
        }
    }
}
