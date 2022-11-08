import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Users.Item.MailFolders.Item.ChildFolders.Item.MultiValueExtendedProperties.Item {
    /// <summary>Builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}\childFolders\{mailFolder-id1}\multiValueExtendedProperties\{multiValueLegacyExtendedProperty-id}</summary>
    public class MultiValueLegacyExtendedPropertyItemRequestBuilder {
        private var pathParameters: [String:String]?
        private var requestAdapter: RequestAdapter?
        private var urlTemplate: String = "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/childFolders/{mailFolder%2Did1}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty%2Did}{?%24select,%24expand}"
        /// <summary>Configuration for the request such as headers, query parameters, and middleware options.</summary>
        public class MultiValueLegacyExtendedPropertyItemRequestBuilderDeleteRequestConfiguration {
            public var headers: IDictionary<string, string>?
            public var options: [RequestOption]
        }
        /// <summary>The collection of multi-value extended properties defined for the mailFolder. Read-only. Nullable.</summary>
        public class MultiValueLegacyExtendedPropertyItemRequestBuilderGetQueryParameters {
            public var expand: [MultiValueLegacyExtendedPropertyItemRequestBuilderGetQueryParametersstring]?
            public var select: [MultiValueLegacyExtendedPropertyItemRequestBuilderGetQueryParametersstring]?
        }
        /// <summary>Configuration for the request such as headers, query parameters, and middleware options.</summary>
        public class MultiValueLegacyExtendedPropertyItemRequestBuilderGetRequestConfiguration {
            public var headers: IDictionary<string, string>?
            public var options: [RequestOption]
            public var queryParameters: MultiValueLegacyExtendedPropertyItemRequestBuilderGetQueryParameters?
        }
        /// <summary>Configuration for the request such as headers, query parameters, and middleware options.</summary>
        public class MultiValueLegacyExtendedPropertyItemRequestBuilderPatchRequestConfiguration {
            public var headers: IDictionary<string, string>?
            public var options: [RequestOption]
        }
    }
}
