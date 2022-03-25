extension Swiftconsoleapp.Users.Item.MailFolders.Item.Messages {
    /// <summary>Builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}\messages</summary>
    public class MessagesRequestBuilder {
        private var pathParameters: Dictionary<string, object>
        private var requestAdapter: IRequestAdapter
        private var urlTemplate: String
    }
}
