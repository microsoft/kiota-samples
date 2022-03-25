extension Swiftconsoleapp.Users.Item.MailFolders.Item.Messages {
    public class MessagesResponse : IAdditionalDataHolder {
        public var additionalData: IDictionary<string, object>
        public var nextLink: String
        public var value: [Message]
    }
}
