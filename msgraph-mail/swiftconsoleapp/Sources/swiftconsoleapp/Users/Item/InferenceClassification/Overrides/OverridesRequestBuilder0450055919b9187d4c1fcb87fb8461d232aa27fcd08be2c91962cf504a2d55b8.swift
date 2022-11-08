import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Users.Item.InferenceClassification.Overrides {
    /// <summary>Builds and executes requests for operations under \users\{user-id}\inferenceClassification\overrides</summary>
    public class OverridesRequestBuilder {
        private var pathParameters: [String:String]?
        private var requestAdapter: RequestAdapter?
        private var urlTemplate: String = "{+baseurl}/users/{user%2Did}/inferenceClassification/overrides{?%24top,%24skip,%24filter,%24count,%24orderby,%24select}"
        /// <summary>Get the overrides that a user has set up to always classify messages from certain senders in specific ways. Each override corresponds to an SMTP address of a sender. Initially, a user does not have any overrides.</summary>
        public class OverridesRequestBuilderGetQueryParameters {
            public var count: OverridesRequestBuilderGetQueryParametersboolean?
            public var filter: OverridesRequestBuilderGetQueryParametersstring?
            public var orderby: [OverridesRequestBuilderGetQueryParametersstring]?
            public var select: [OverridesRequestBuilderGetQueryParametersstring]?
            public var skip: OverridesRequestBuilderGetQueryParametersinteger?
            public var top: OverridesRequestBuilderGetQueryParametersinteger?
        }
        /// <summary>Configuration for the request such as headers, query parameters, and middleware options.</summary>
        public class OverridesRequestBuilderGetRequestConfiguration {
            public var headers: IDictionary<string, string>?
            public var options: [RequestOption]
            public var queryParameters: OverridesRequestBuilderGetQueryParameters?
        }
        /// <summary>Configuration for the request such as headers, query parameters, and middleware options.</summary>
        public class OverridesRequestBuilderPostRequestConfiguration {
            public var headers: IDictionary<string, string>?
            public var options: [RequestOption]
        }
    }
}
