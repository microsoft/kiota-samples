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
        private var urlTemplate: String = "{+baseurl}/users/{user_id}/mailFolders/{mailFolder_id}{?select}"
    }
}
