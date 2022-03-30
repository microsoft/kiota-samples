import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Users.Item.MailFolders {
    /// <summary>Builds and executes requests for operations under \users\{user-id}\mailFolders</summary>
    public class MailFoldersRequestBuilder {
        private var pathParameters: [String:String]?
        private var requestAdapter: RequestAdapter?
        private var urlTemplate: String = "{+baseurl}/users/{user_id}/mailFolders{?top,skip,filter,count,orderby,select}"
    }
}
