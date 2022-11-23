import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Models {
    public class EmailAddress : AdditionalDataHolder {
        public var additionalData: [String:Any] = [String:Any]()
        public var address: String?
        public var name: String?
    }
}
