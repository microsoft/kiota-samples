import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Users.Item.Messages.Item.SingleValueExtendedProperties.Item {
    /// <summary>Builds and executes requests for operations under \users\{user-id}\messages\{message-id}\singleValueExtendedProperties\{singleValueLegacyExtendedProperty-id}</summary>
    public class SingleValueLegacyExtendedPropertyItemRequestBuilder {
        private var pathParameters: [String:String]?
        private var requestAdapter: RequestAdapter?
        private var urlTemplate: String = "{+baseurl}/users/{user_id}/messages/{message_id}/singleValueExtendedProperties/{singleValueLegacyExtendedProperty_id}{?select,expand}"
    }
}
