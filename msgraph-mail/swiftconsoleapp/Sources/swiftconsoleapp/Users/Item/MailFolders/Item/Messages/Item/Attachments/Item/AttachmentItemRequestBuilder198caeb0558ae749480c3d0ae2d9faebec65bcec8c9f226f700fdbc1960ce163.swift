import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Users.Item.MailFolders.Item.Messages.Item.Attachments.Item {
    /// <summary>Builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}\messages\{message-id}\attachments\{attachment-id}</summary>
    public class AttachmentItemRequestBuilder {
        private var pathParameters: [String:String]?
        private var requestAdapter: RequestAdapter?
        private var urlTemplate: String = "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/messages/{message%2Did}/attachments/{attachment%2Did}{?%24select,%24expand}"
        /// <summary>Configuration for the request such as headers, query parameters, and middleware options.</summary>
        public class AttachmentItemRequestBuilderDeleteRequestConfiguration {
            public var headers: IDictionary<string, string>?
            public var options: [RequestOption]
        }
        /// <summary>The fileAttachment and itemAttachment attachments for the message.</summary>
        public class AttachmentItemRequestBuilderGetQueryParameters {
            public var expand: [AttachmentItemRequestBuilderGetQueryParametersstring]?
            public var select: [AttachmentItemRequestBuilderGetQueryParametersstring]?
        }
        /// <summary>Configuration for the request such as headers, query parameters, and middleware options.</summary>
        public class AttachmentItemRequestBuilderGetRequestConfiguration {
            public var headers: IDictionary<string, string>?
            public var options: [RequestOption]
            public var queryParameters: AttachmentItemRequestBuilderGetQueryParameters?
        }
    }
}
