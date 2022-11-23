import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Users.Item.MailFolders.Item.ChildFolders {
    /// <summary>Builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}\childFolders</summary>
    public class ChildFoldersRequestBuilder {
        private var pathParameters: [String:String]?
        private var requestAdapter: RequestAdapter?
        private var urlTemplate: String = "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/childFolders{?%24top,%24skip,%24filter,%24count,%24orderby,%24select,%24expand}"
        /// <summary>Get the folder collection under the specified folder. You can use the `.../me/mailFolders` shortcut to get the top-level folder collection and navigate to another folder. By default, this operation does not return hidden folders. Use a query parameter _includeHiddenFolders_ to include them in the response.</summary>
        public class ChildFoldersRequestBuilderGetQueryParameters {
            public var count: ChildFoldersRequestBuilderGetQueryParametersboolean?
            public var expand: [ChildFoldersRequestBuilderGetQueryParametersstring]?
            public var filter: ChildFoldersRequestBuilderGetQueryParametersstring?
            public var orderby: [ChildFoldersRequestBuilderGetQueryParametersstring]?
            public var select: [ChildFoldersRequestBuilderGetQueryParametersstring]?
            public var skip: ChildFoldersRequestBuilderGetQueryParametersinteger?
            public var top: ChildFoldersRequestBuilderGetQueryParametersinteger?
        }
        /// <summary>Configuration for the request such as headers, query parameters, and middleware options.</summary>
        public class ChildFoldersRequestBuilderGetRequestConfiguration {
            public var headers: IDictionary<string, string>?
            public var options: [RequestOption]
            public var queryParameters: ChildFoldersRequestBuilderGetQueryParameters?
        }
        /// <summary>Configuration for the request such as headers, query parameters, and middleware options.</summary>
        public class ChildFoldersRequestBuilderPostRequestConfiguration {
            public var headers: IDictionary<string, string>?
            public var options: [RequestOption]
        }
    }
}
