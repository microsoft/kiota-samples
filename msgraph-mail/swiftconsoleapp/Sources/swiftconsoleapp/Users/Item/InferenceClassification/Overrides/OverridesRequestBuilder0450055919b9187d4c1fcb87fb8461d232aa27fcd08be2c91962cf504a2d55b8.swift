import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Users.Item.InferenceClassification.Overrides {
    /// <summary>Builds and executes requests for operations under \users\{user-id}\inferenceClassification\overrides</summary>
    public class OverridesRequestBuilder {
        private var pathParameters: [String:String]?
        private var requestAdapter: RequestAdapter?
        private var urlTemplate: String = "{+baseurl}/users/{user_id}/inferenceClassification/overrides{?top,skip,filter,count,orderby,select}"
    }
}
