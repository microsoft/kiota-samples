extension Swiftconsoleapp.Users.Item.Messages {
    /// <summary>Builds and executes requests for operations under \users\{user-id}\messages</summary>
    public class MessagesRequestBuilder {
        private var pathParameters: Dictionary<string, object>
        private var requestAdapter: IRequestAdapter
        private var urlTemplate: String
    }
}
