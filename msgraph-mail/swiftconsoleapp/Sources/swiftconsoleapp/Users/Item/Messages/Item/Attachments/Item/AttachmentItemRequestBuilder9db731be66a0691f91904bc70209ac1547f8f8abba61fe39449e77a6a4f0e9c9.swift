import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Users.Item.Messages.Item.Attachments.Item {
    /// <summary>Builds and executes requests for operations under \users\{user-id}\messages\{message-id}\attachments\{attachment-id}</summary>
    public class AttachmentItemRequestBuilder {
        private var pathParameters: [String:String]?
        private var requestAdapter: RequestAdapter?
        private var urlTemplate: String = "{+baseurl}/users/{user_id}/messages/{message_id}/attachments/{attachment_id}{?select,expand}"
    }
}
