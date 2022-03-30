import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Users.Item.MailFolders {
    public class MailFoldersResponse : AdditionalDataHolder {
        public var additionalData: [String:Any] = [String:Any]()
        public var nextLink: String?
        public var value: [Swiftconsoleapp.Models.Microsoft.Graph.MailFolder]?
    }
}
