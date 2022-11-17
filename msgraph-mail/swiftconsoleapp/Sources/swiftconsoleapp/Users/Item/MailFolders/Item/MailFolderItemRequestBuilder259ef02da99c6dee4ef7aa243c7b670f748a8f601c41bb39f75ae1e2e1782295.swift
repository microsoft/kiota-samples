import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Users.Item.MailFolders.Item {
    /// <summary>Builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}</summary>
    public class MailFolderItemRequestBuilder {
        public var childFolders: ChildFolders.ChildFoldersRequestBuilder?
        public var messageRules: MessageRules.MessageRulesRequestBuilder?
        public var messages: Messages.MessagesRequestBuilder?
        public var multiValueExtendedProperties: MultiValueExtendedProperties.MultiValueExtendedPropertiesRequestBuilder?
        private var pathParameters: [String:String]?
        private var requestAdapter: RequestAdapter?
        public var singleValueExtendedProperties: SingleValueExtendedProperties.SingleValueExtendedPropertiesRequestBuilder?
        private var urlTemplate: String = "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}{?%24select}"
        /// <summary>Configuration for the request such as headers, query parameters, and middleware options.</summary>
        public class MailFolderItemRequestBuilderDeleteRequestConfiguration {
            public var headers: IDictionary<string, string>?
            public var options: [RequestOption]
        }
        /// <summary>The user's mail folders. Read-only. Nullable.</summary>
        public class MailFolderItemRequestBuilderGetQueryParameters {
            public var select: [MailFolderItemRequestBuilderGetQueryParametersstring]?
        }
        /// <summary>Configuration for the request such as headers, query parameters, and middleware options.</summary>
        public class MailFolderItemRequestBuilderGetRequestConfiguration {
            public var headers: IDictionary<string, string>?
            public var options: [RequestOption]
            public var queryParameters: MailFolderItemRequestBuilderGetQueryParameters?
        }
        /// <summary>Configuration for the request such as headers, query parameters, and middleware options.</summary>
        public class MailFolderItemRequestBuilderPatchRequestConfiguration {
            public var headers: IDictionary<string, string>?
            public var options: [RequestOption]
        }
    }
}
