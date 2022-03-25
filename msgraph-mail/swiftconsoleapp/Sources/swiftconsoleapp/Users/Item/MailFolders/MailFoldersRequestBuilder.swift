extension Swiftconsoleapp.Users.Item.MailFolders {
    /// <summary>Builds and executes requests for operations under \users\{user-id}\mailFolders</summary>
    public class MailFoldersRequestBuilder {
        private var pathParameters: Dictionary<string, object>
        private var requestAdapter: IRequestAdapter
        private var urlTemplate: String
    }
}
