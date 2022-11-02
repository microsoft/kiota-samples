import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Users.Item.MailFolders.Item.ChildFolders.Item.Messages.Item {
    /// <summary>Builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}\childFolders\{mailFolder-id1}\messages\{message-id}</summary>
    public class MessageItemRequestBuilder {
        public var attachments: Attachments.AttachmentsRequestBuilder?
        public var content: Value.ContentRequestBuilder?
        public var extensions: Extensions.ExtensionsRequestBuilder?
        public var multiValueExtendedProperties: MultiValueExtendedProperties.MultiValueExtendedPropertiesRequestBuilder?
        private var pathParameters: [String:String]?
        private var requestAdapter: RequestAdapter?
        public var singleValueExtendedProperties: SingleValueExtendedProperties.SingleValueExtendedPropertiesRequestBuilder?
        private var urlTemplate: String = "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/childFolders/{mailFolder%2Did1}/messages/{message%2Did}{?%24select,%24expand}"
        /// <summary>Configuration for the request such as headers, query parameters, and middleware options.</summary>
        public class MessageItemRequestBuilderDeleteRequestConfiguration {
            public var headers: IDictionary<string, string>?
            public var options: [RequestOption]
        }
        /// <summary>The collection of messages in the mailFolder.</summary>
        public class MessageItemRequestBuilderGetQueryParameters {
            public var expand: [MessageItemRequestBuilderGetQueryParametersstring]?
            public var select: [MessageItemRequestBuilderGetQueryParametersstring]?
        }
        /// <summary>Configuration for the request such as headers, query parameters, and middleware options.</summary>
        public class MessageItemRequestBuilderGetRequestConfiguration {
            public var headers: IDictionary<string, string>?
            public var options: [RequestOption]
            public var queryParameters: MessageItemRequestBuilderGetQueryParameters?
        }
        /// <summary>Configuration for the request such as headers, query parameters, and middleware options.</summary>
        public class MessageItemRequestBuilderPatchRequestConfiguration {
            public var headers: IDictionary<string, string>?
            public var options: [RequestOption]
        }
    }
}
