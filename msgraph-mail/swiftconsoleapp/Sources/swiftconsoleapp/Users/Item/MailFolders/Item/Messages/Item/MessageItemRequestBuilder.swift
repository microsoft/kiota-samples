extension Swiftconsoleapp.Users.Item.MailFolders.Item.Messages.Item {
    /// <summary>Builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}\messages\{message-id}</summary>
    public class MessageItemRequestBuilder {
        public var attachments: AttachmentsRequestBuilder
        public var content: ContentRequestBuilder
        public var extensions: ExtensionsRequestBuilder
        public var multiValueExtendedProperties: MultiValueExtendedPropertiesRequestBuilder
        private var pathParameters: Dictionary<string, object>
        private var requestAdapter: IRequestAdapter
        public var singleValueExtendedProperties: SingleValueExtendedPropertiesRequestBuilder
        private var urlTemplate: String
    }
}
