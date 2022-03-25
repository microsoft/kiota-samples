extension Swiftconsoleapp.Users.Item.MailFolders {
    public class MailFoldersResponse : IAdditionalDataHolder {
        public var additionalData: IDictionary<string, object>
        public var nextLink: String
        public var value: [MailFolder]
    }
}
