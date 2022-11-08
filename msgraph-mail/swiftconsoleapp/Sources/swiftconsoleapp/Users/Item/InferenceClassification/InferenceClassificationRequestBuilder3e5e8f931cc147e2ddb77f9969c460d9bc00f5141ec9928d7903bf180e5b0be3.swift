import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Users.Item.InferenceClassification {
    /// <summary>Builds and executes requests for operations under \users\{user-id}\inferenceClassification</summary>
    public class InferenceClassificationRequestBuilder {
        public var overrides: Overrides.OverridesRequestBuilder?
        private var pathParameters: [String:String]?
        private var requestAdapter: RequestAdapter?
        private var urlTemplate: String = "{+baseurl}/users/{user%2Did}/inferenceClassification{?%24select}"
        /// <summary>Relevance classification of the user's messages based on explicit designations which override inferred relevance or importance.</summary>
        public class InferenceClassificationRequestBuilderGetQueryParameters {
            public var select: [InferenceClassificationRequestBuilderGetQueryParametersstring]?
        }
        /// <summary>Configuration for the request such as headers, query parameters, and middleware options.</summary>
        public class InferenceClassificationRequestBuilderGetRequestConfiguration {
            public var headers: IDictionary<string, string>?
            public var options: [RequestOption]
            public var queryParameters: InferenceClassificationRequestBuilderGetQueryParameters?
        }
        /// <summary>Configuration for the request such as headers, query parameters, and middleware options.</summary>
        public class InferenceClassificationRequestBuilderPatchRequestConfiguration {
            public var headers: IDictionary<string, string>?
            public var options: [RequestOption]
        }
    }
}
