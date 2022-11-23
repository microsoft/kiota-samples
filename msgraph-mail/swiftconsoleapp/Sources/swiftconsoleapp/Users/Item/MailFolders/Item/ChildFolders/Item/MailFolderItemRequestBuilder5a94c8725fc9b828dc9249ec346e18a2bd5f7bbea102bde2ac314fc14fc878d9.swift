import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Users.Item.MailFolders.Item.ChildFolders.Item {
    /// <summary>Builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}\childFolders\{mailFolder-id1}</summary>
    public class MailFolderItemRequestBuilder {
        public var messageRules: MessageRules.MessageRulesRequestBuilder?
        public var messages: Messages.MessagesRequestBuilder?
        public var multiValueExtendedProperties: MultiValueExtendedProperties.MultiValueExtendedPropertiesRequestBuilder?
        private var pathParameters: [String:String]?
        private var requestAdapter: RequestAdapter?
        public var singleValueExtendedProperties: SingleValueExtendedProperties.SingleValueExtendedPropertiesRequestBuilder?
        private var urlTemplate: String = "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/childFolders/{mailFolder%2Did1}{?%24select,%24expand}"
        /// <summary>Configuration for the request such as headers, query parameters, and middleware options.</summary>
        public class MailFolderItemRequestBuilderDeleteRequestConfiguration {
            public var headers: IDictionary<string, string>?
            public var options: [RequestOption]
        }
        /// <summary>The collection of child folders in the mailFolder.</summary>
        public class MailFolderItemRequestBuilderGetQueryParameters {
            public var expand: [MailFolderItemRequestBuilderGetQueryParametersstring]?
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
