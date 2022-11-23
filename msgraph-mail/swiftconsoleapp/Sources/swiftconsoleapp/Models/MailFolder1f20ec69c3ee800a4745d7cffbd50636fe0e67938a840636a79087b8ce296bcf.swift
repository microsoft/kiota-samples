import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Models {
    public class MailFolder : Entity {
        public var childFolderCount: Int32?
        public var childFolders: [MailFolder]?
        public var displayName: String?
        public var isHidden: Bool?
        public var messageRules: [MessageRule]?
        public var messages: [Message]?
        public var multiValueExtendedProperties: [MultiValueLegacyExtendedProperty]?
        public var parentFolderId: String?
        public var singleValueExtendedProperties: [SingleValueLegacyExtendedProperty]?
        public var totalItemCount: Int32?
        public var unreadItemCount: Int32?
    }
}
