import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Models.Microsoft.Graph {
    public class ItemBody : AdditionalDataHolder {
        public var additionalData: [String:Any] = [String:Any]()
        public var content: String?
        public var contentType: BodyType?
    }
}
