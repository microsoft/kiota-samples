import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Users.Item.MailFolders.Item.SingleValueExtendedProperties.Item {
    /// <summary>Builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}\singleValueExtendedProperties\{singleValueLegacyExtendedProperty-id}</summary>
    public class SingleValueLegacyExtendedPropertyItemRequestBuilder {
        private var pathParameters: [String:String]?
        private var requestAdapter: RequestAdapter?
        private var urlTemplate: String = "{+baseurl}/users/{user_id}/mailFolders/{mailFolder_id}/singleValueExtendedProperties/{singleValueLegacyExtendedProperty_id}{?select,expand}"
    }
}
