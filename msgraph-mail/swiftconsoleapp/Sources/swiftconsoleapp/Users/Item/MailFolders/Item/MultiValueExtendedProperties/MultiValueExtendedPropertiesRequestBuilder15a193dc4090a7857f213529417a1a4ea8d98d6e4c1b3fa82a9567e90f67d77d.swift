import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Users.Item.MailFolders.Item.MultiValueExtendedProperties {
    /// <summary>Builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}\multiValueExtendedProperties</summary>
    public class MultiValueExtendedPropertiesRequestBuilder {
        private var pathParameters: [String:String]?
        private var requestAdapter: RequestAdapter?
        private var urlTemplate: String = "{+baseurl}/users/{user_id}/mailFolders/{mailFolder_id}/multiValueExtendedProperties{?top,skip,search,filter,count,orderby,select,expand}"
    }
}
