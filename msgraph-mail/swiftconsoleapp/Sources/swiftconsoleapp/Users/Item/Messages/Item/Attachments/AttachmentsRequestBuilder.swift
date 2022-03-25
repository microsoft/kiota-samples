extension Swiftconsoleapp.Users.Item.Messages.Item.Attachments {
    /// <summary>Builds and executes requests for operations under \users\{user-id}\messages\{message-id}\attachments</summary>
    public class AttachmentsRequestBuilder {
        private var pathParameters: Dictionary<string, object>
        private var requestAdapter: IRequestAdapter
        private var urlTemplate: String
    }
}
