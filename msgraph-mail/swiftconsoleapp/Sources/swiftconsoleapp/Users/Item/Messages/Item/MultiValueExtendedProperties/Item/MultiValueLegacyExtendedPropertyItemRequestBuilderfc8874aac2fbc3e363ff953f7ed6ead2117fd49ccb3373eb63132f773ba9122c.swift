import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Users.Item.Messages.Item.MultiValueExtendedProperties.Item {
    /// <summary>Builds and executes requests for operations under \users\{user-id}\messages\{message-id}\multiValueExtendedProperties\{multiValueLegacyExtendedProperty-id}</summary>
    public class MultiValueLegacyExtendedPropertyItemRequestBuilder {
        private var pathParameters: [String:String]?
        private var requestAdapter: RequestAdapter?
        private var urlTemplate: String = "{+baseurl}/users/{user_id}/messages/{message_id}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty_id}{?select,expand}"
    }
}
