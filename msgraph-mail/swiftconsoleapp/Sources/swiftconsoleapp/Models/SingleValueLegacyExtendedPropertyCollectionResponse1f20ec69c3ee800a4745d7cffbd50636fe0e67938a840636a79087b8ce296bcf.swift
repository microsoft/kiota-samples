import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Models {
    public class SingleValueLegacyExtendedPropertyCollectionResponse : AdditionalDataHolder {
        public var additionalData: [String:Any] = [String:Any]()
        public var odataNextLink: String?
        public var value: [SingleValueLegacyExtendedProperty]?
    }
}
