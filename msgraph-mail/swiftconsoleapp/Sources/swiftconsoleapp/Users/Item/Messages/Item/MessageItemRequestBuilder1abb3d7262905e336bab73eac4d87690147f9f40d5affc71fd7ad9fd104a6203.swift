import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Users.Item.Messages.Item {
    /// <summary>Builds and executes requests for operations under \users\{user-id}\messages\{message-id}</summary>
    public class MessageItemRequestBuilder {
        public var attachments: Attachments.AttachmentsRequestBuilder?
        public var content: Value.ContentRequestBuilder?
        public var extensions: Extensions.ExtensionsRequestBuilder?
        public var multiValueExtendedProperties: MultiValueExtendedProperties.MultiValueExtendedPropertiesRequestBuilder?
        private var pathParameters: [String:String]?
        private var requestAdapter: RequestAdapter?
        public var singleValueExtendedProperties: SingleValueExtendedProperties.SingleValueExtendedPropertiesRequestBuilder?
        private var urlTemplate: String = "{+baseurl}/users/{user_id}/messages/{message_id}{?select}"
    }
}
