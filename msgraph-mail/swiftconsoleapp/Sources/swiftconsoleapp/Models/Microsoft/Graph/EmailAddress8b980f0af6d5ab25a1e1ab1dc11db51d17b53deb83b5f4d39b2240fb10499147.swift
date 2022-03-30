import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Models.Microsoft.Graph {
    public class EmailAddress : AdditionalDataHolder {
        public var additionalData: [String:Any] = [String:Any]()
        public var address: String?
        public var name: String?
    }
}
