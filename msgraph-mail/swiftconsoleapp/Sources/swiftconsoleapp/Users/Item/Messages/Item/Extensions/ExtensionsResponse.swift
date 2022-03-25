extension Swiftconsoleapp.Users.Item.Messages.Item.Extensions {
    public class ExtensionsResponse : IAdditionalDataHolder {
        public var additionalData: IDictionary<string, object>
        public var nextLink: String
        public var value: [Extension]
    }
}
