extension Swiftconsoleapp.Users.Item.MailFolders.Item.ChildFolders {
    public class ChildFoldersResponse : IAdditionalDataHolder {
        public var additionalData: IDictionary<string, object>
        public var nextLink: String
        public var value: [MailFolder]
    }
}
