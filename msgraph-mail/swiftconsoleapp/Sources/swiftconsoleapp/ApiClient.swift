extension Swiftconsoleapp {
    /// <summary>The main entry point of the SDK, exposes the configuration and the fluent API.</summary>
    public class ApiClient {
        private var pathParameters: Dictionary<string, object>
        private var requestAdapter: IRequestAdapter
        private var urlTemplate: String
        public var users: UsersRequestBuilder
    }
}
