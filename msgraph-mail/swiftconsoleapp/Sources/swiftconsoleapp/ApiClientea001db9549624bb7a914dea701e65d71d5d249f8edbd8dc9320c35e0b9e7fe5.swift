import MicrosoftKiotaAbstractions

extension Swiftconsoleapp {
    /// <summary>The main entry point of the SDK, exposes the configuration and the fluent API.</summary>
    public class ApiClient {
        private var pathParameters: [String:String]? = [String:String]()
        private var requestAdapter: RequestAdapter?
        private var urlTemplate: String = "{+baseurl}"
        public var users: Users.UsersRequestBuilder?
    }
}
