import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Models {
    public class SizeRange : AdditionalDataHolder {
        public var additionalData: [String:Any] = [String:Any]()
        public var maximumSize: Int32?
        public var minimumSize: Int32?
    }
}
