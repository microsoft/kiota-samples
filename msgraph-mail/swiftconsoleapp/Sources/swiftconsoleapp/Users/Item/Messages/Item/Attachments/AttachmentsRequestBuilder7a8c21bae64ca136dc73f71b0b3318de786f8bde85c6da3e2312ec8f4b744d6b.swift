import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Users.Item.Messages.Item.Attachments {
    /// <summary>Builds and executes requests for operations under \users\{user-id}\messages\{message-id}\attachments</summary>
    public class AttachmentsRequestBuilder {
        private var pathParameters: [String:String]?
        private var requestAdapter: RequestAdapter?
        private var urlTemplate: String = "{+baseurl}/users/{user_id}/messages/{message_id}/attachments{?top,skip,filter,count,orderby,select,expand}"
    }
}
