import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Users.Item.Messages.Item.Extensions {
    /// <summary>Builds and executes requests for operations under \users\{user-id}\messages\{message-id}\extensions</summary>
    public class ExtensionsRequestBuilder {
        private var pathParameters: [String:String]?
        private var requestAdapter: RequestAdapter?
        private var urlTemplate: String = "{+baseurl}/users/{user_id}/messages/{message_id}/extensions{?top,skip,filter,count,orderby,select,expand}"
    }
}
