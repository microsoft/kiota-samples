extension Swiftconsoleapp.Users.Item.Messages.Item.MultiValueExtendedProperties {
    public class MultiValueExtendedPropertiesResponse : IAdditionalDataHolder {
        public var additionalData: IDictionary<string, object>
        public var nextLink: String
        public var value: [MultiValueLegacyExtendedProperty]
    }
}
