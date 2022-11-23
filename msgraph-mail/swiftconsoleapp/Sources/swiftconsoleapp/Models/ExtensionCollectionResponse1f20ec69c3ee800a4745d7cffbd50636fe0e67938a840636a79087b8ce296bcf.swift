import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Models {
    public class ExtensionCollectionResponse : AdditionalDataHolder {
        public var additionalData: [String:Any] = [String:Any]()
        public var odataNextLink: String?
        public var value: [Extension]?
    }
}
