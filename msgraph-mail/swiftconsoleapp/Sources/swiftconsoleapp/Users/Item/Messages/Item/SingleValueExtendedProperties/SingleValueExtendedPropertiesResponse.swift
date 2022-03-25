extension Swiftconsoleapp.Users.Item.Messages.Item.SingleValueExtendedProperties {
    public class SingleValueExtendedPropertiesResponse : IAdditionalDataHolder {
        public var additionalData: IDictionary<string, object>
        public var nextLink: String
        public var value: [SingleValueLegacyExtendedProperty]
    }
}
