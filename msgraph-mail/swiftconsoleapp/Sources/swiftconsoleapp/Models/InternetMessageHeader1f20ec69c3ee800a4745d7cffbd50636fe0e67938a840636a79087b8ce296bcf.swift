import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Models {
    public class InternetMessageHeader : AdditionalDataHolder {
        public var additionalData: [String:Any] = [String:Any]()
        public var name: String?
        public var value: String?
    }
}
