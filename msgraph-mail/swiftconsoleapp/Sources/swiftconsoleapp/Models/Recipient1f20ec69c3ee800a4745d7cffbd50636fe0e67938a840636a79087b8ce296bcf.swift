import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Models {
    public class Recipient : AdditionalDataHolder {
        public var additionalData: [String:Any] = [String:Any]()
        public var emailAddress: EmailAddress?
    }
}
