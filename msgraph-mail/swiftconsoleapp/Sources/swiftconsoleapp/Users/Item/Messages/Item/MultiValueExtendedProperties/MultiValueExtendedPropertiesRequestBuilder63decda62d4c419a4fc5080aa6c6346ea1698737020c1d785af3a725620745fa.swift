import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Users.Item.Messages.Item.MultiValueExtendedProperties {
    /// <summary>Builds and executes requests for operations under \users\{user-id}\messages\{message-id}\multiValueExtendedProperties</summary>
    public class MultiValueExtendedPropertiesRequestBuilder {
        private var pathParameters: [String:String]?
        private var requestAdapter: RequestAdapter?
        private var urlTemplate: String = "{+baseurl}/users/{user_id}/messages/{message_id}/multiValueExtendedProperties{?top,skip,search,filter,count,orderby,select,expand}"
    }
}
