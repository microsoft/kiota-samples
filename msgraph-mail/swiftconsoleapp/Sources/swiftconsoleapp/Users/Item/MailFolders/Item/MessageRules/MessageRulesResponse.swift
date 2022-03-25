extension Swiftconsoleapp.Users.Item.MailFolders.Item.MessageRules {
    public class MessageRulesResponse : IAdditionalDataHolder {
        public var additionalData: IDictionary<string, object>
        public var nextLink: String
        public var value: [MessageRule]
    }
}
