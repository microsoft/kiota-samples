extension Swiftconsoleapp.Users.Item.InferenceClassification.Overrides {
    public class OverridesResponse : IAdditionalDataHolder {
        public var additionalData: IDictionary<string, object>
        public var nextLink: String
        public var value: [InferenceClassificationOverride]
    }
}
