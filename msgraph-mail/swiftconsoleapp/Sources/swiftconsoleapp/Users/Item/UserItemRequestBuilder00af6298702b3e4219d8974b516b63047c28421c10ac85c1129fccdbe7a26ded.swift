import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Users.Item {
    /// <summary>Builds and executes requests for operations under \users\{user-id}</summary>
    public class UserItemRequestBuilder {
        public var inferenceClassification: InferenceClassification.InferenceClassificationRequestBuilder?
        public var mailFolders: MailFolders.MailFoldersRequestBuilder?
        public var messages: Messages.MessagesRequestBuilder?
        private var pathParameters: [String:String]?
        private var requestAdapter: RequestAdapter?
        private var urlTemplate: String = "{+baseurl}/users/{user%2Did}"
    }
}
