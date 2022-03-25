import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Users.Item.InferenceClassification.Overrides.Item {
    /// <summary>Builds and executes requests for operations under \users\{user-id}\inferenceClassification\overrides\{inferenceClassificationOverride-id}</summary>
    public class InferenceClassificationOverrideItemRequestBuilder {
        private var pathParameters: [String:String]?
        private var requestAdapter: RequestAdapter?
        private var urlTemplate: String = "{+baseurl}/users/{user_id}/inferenceClassification/overrides/{inferenceClassificationOverride_id}{?select}"
    }
}
