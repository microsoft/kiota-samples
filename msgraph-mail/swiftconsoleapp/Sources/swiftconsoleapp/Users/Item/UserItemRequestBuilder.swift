extension Swiftconsoleapp.Users.Item {
    /// <summary>Builds and executes requests for operations under \users\{user-id}</summary>
    public class UserItemRequestBuilder {
        public var inferenceClassification: InferenceClassificationRequestBuilder
        public var mailFolders: MailFoldersRequestBuilder
        public var messages: MessagesRequestBuilder
        private var pathParameters: Dictionary<string, object>
        private var requestAdapter: IRequestAdapter
        private var urlTemplate: String
    }
}
