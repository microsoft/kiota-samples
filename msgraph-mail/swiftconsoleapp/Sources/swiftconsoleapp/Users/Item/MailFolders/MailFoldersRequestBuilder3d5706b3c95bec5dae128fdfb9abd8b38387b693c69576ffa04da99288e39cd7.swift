import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Users.Item.MailFolders {
    /// <summary>Builds and executes requests for operations under \users\{user-id}\mailFolders</summary>
    public class MailFoldersRequestBuilder {
        private var pathParameters: [String:String]?
        private var requestAdapter: RequestAdapter?
        private var urlTemplate: String = "{+baseurl}/users/{user%2Did}/mailFolders{?%24top,%24skip,%24filter,%24count,%24orderby,%24select}"
        /// <summary>Get the mail folder collection directly under the root folder of the signed-in user. The returned collection includes any mail search folders directly under the root. By default, this operation does not return hidden folders. Use a query parameter _includeHiddenFolders_ to include them in the response.</summary>
        public class MailFoldersRequestBuilderGetQueryParameters {
            public var count: MailFoldersRequestBuilderGetQueryParametersboolean?
            public var filter: MailFoldersRequestBuilderGetQueryParametersstring?
            public var orderby: [MailFoldersRequestBuilderGetQueryParametersstring]?
            public var select: [MailFoldersRequestBuilderGetQueryParametersstring]?
            public var skip: MailFoldersRequestBuilderGetQueryParametersinteger?
            public var top: MailFoldersRequestBuilderGetQueryParametersinteger?
        }
        /// <summary>Configuration for the request such as headers, query parameters, and middleware options.</summary>
        public class MailFoldersRequestBuilderGetRequestConfiguration {
            public var headers: IDictionary<string, string>?
            public var options: [RequestOption]
            public var queryParameters: MailFoldersRequestBuilderGetQueryParameters?
        }
        /// <summary>Configuration for the request such as headers, query parameters, and middleware options.</summary>
        public class MailFoldersRequestBuilderPostRequestConfiguration {
            public var headers: IDictionary<string, string>?
            public var options: [RequestOption]
        }
    }
}
