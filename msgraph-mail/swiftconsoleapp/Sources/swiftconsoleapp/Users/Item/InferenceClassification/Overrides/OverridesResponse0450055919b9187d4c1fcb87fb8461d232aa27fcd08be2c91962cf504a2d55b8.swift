import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Users.Item.InferenceClassification.Overrides {
    public class OverridesResponse : AdditionalDataHolder {
        public var additionalData: [String:Any] = [String:Any]()
        public var nextLink: String?
        public var value: [Swiftconsoleapp.Models.Microsoft.Graph.InferenceClassificationOverride]?
    }
}
