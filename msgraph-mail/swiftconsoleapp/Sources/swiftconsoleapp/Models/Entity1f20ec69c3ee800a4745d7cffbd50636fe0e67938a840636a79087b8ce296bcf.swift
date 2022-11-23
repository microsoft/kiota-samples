import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Models {
    public class Entity : AdditionalDataHolder {
        public var additionalData: [String:Any] = [String:Any]()
        public var id: String?
    }
}
