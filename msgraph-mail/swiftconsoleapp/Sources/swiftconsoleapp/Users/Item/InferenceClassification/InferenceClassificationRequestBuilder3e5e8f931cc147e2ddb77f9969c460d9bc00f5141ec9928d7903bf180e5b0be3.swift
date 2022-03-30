import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Users.Item.InferenceClassification {
    /// <summary>Builds and executes requests for operations under \users\{user-id}\inferenceClassification</summary>
    public class InferenceClassificationRequestBuilder {
        public var overrides: Overrides.OverridesRequestBuilder?
        private var pathParameters: [String:String]?
        private var requestAdapter: RequestAdapter?
        private var urlTemplate: String = "{+baseurl}/users/{user_id}/inferenceClassification{?select}"
    }
}
