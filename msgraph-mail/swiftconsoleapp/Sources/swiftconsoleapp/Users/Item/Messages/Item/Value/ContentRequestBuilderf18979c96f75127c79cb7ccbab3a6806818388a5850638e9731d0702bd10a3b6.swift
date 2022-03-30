import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Users.Item.Messages.Item.Value {
    /// <summary>Builds and executes requests for operations under \users\{user-id}\messages\{message-id}\$value</summary>
    public class ContentRequestBuilder {
        private var pathParameters: [String:String]?
        private var requestAdapter: RequestAdapter?
        private var urlTemplate: String = "{+baseurl}/users/{user_id}/messages/{message_id}/$value"
    }
}
