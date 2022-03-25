import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Users.Item.Messages.Item.SingleValueExtendedProperties {
    /// <summary>Builds and executes requests for operations under \users\{user-id}\messages\{message-id}\singleValueExtendedProperties</summary>
    public class SingleValueExtendedPropertiesRequestBuilder {
        private var pathParameters: [String:String]?
        private var requestAdapter: RequestAdapter?
        private var urlTemplate: String = "{+baseurl}/users/{user_id}/messages/{message_id}/singleValueExtendedProperties{?top,skip,search,filter,count,orderby,select,expand}"
    }
}
