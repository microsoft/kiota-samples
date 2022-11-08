import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Users.Item.MailFolders.Item.ChildFolders.Item.Messages {
    /// <summary>Builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}\childFolders\{mailFolder-id1}\messages</summary>
    public class MessagesRequestBuilder {
        private var pathParameters: [String:String]?
        private var requestAdapter: RequestAdapter?
        private var urlTemplate: String = "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/childFolders/{mailFolder%2Did1}/messages{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}"
        /// <summary>Get all the messages in the specified user's mailbox, or those messages in a specified folder in the mailbox.</summary>
        public class MessagesRequestBuilderGetQueryParameters {
            public var count: MessagesRequestBuilderGetQueryParametersboolean?
            public var expand: [MessagesRequestBuilderGetQueryParametersstring]?
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
