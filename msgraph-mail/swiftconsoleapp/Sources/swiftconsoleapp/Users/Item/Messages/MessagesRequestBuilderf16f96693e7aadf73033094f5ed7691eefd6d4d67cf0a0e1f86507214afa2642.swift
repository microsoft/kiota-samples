import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Users.Item.Messages {
    /// <summary>Builds and executes requests for operations under \users\{user-id}\messages</summary>
    public class MessagesRequestBuilder {
        private var pathParameters: [String:String]?
        private var requestAdapter: RequestAdapter?
        private var urlTemplate: String = "{+baseurl}/users/{user_id}/messages{?top,skip,search,filter,count,orderby,select}"
    }
}
