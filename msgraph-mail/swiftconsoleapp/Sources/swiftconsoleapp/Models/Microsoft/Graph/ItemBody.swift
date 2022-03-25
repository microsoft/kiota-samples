extension Swiftconsoleapp.Models.Microsoft.Graph {
public class ItemBody : IAdditionalDataHolder {
public var additionalData: IDictionary<string, object>
public var content: String
public var contentType: BodyType?
}
}
