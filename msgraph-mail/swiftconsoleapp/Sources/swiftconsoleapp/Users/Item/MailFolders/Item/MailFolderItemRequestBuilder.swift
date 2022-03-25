extension Swiftconsoleapp.Users.Item.MailFolders.Item {
    /// <summary>Builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}</summary>
    public class MailFolderItemRequestBuilder {
        public var childFolders: ChildFoldersRequestBuilder
        public var messageRules: MessageRulesRequestBuilder
        public var messages: MessagesRequestBuilder
        public var multiValueExtendedProperties: MultiValueExtendedPropertiesRequestBuilder
        private var pathParameters: Dictionary<string, object>
        private var requestAdapter: IRequestAdapter
        public var singleValueExtendedProperties: SingleValueExtendedPropertiesRequestBuilder
        private var urlTemplate: String
    }
}
