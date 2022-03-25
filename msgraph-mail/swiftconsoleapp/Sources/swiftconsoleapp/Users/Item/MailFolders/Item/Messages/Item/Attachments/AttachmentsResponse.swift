extension Swiftconsoleapp.Users.Item.MailFolders.Item.Messages.Item.Attachments {
    public class AttachmentsResponse : IAdditionalDataHolder {
        public var additionalData: IDictionary<string, object>
        public var nextLink: String
        public var value: [Attachment]
    }
}
