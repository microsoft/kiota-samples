import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Users.Item.InferenceClassification.Overrides.Item {
    /// <summary>Builds and executes requests for operations under \users\{user-id}\inferenceClassification\overrides\{inferenceClassificationOverride-id}</summary>
    public class InferenceClassificationOverrideItemRequestBuilder {
        private var pathParameters: [String:String]?
        private var requestAdapter: RequestAdapter?
        private var urlTemplate: String = "{+baseurl}/users/{user%2Did}/inferenceClassification/overrides/{inferenceClassificationOverride%2Did}{?%24select}"
        /// <summary>Configuration for the request such as headers, query parameters, and middleware options.</summary>
        public class InferenceClassificationOverrideItemRequestBuilderDeleteRequestConfiguration {
            public var headers: IDictionary<string, string>?
            public var options: [RequestOption]
        }
        /// <summary>A set of overrides for a user to always classify messages from specific senders in certain ways: focused, or other. Read-only. Nullable.</summary>
        public class InferenceClassificationOverrideItemRequestBuilderGetQueryParameters {
            public var select: [InferenceClassificationOverrideItemRequestBuilderGetQueryParametersstring]?
        }
        /// <summary>Configuration for the request such as headers, query parameters, and middleware options.</summary>
        public class InferenceClassificationOverrideItemRequestBuilderGetRequestConfiguration {
            public var headers: IDictionary<string, string>?
            public var options: [RequestOption]
            public var queryParameters: InferenceClassificationOverrideItemRequestBuilderGetQueryParameters?
        }
        /// <summary>Configuration for the request such as headers, query parameters, and middleware options.</summary>
        public class InferenceClassificationOverrideItemRequestBuilderPatchRequestConfiguration {
            public var headers: IDictionary<string, string>?
            public var options: [RequestOption]
        }
    }
}
