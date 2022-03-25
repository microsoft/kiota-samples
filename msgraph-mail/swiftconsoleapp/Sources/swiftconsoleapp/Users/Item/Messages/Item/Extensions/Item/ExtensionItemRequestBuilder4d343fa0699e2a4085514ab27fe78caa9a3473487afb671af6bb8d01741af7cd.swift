import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Users.Item.Messages.Item.Extensions.Item {
    /// <summary>Builds and executes requests for operations under \users\{user-id}\messages\{message-id}\extensions\{extension-id}</summary>
    public class ExtensionItemRequestBuilder {
        private var pathParameters: [String:String]?
        private var requestAdapter: RequestAdapter?
        private var urlTemplate: String = "{+baseurl}/users/{user_id}/messages/{message_id}/extensions/{extension_id}{?select,expand}"
    }
}
