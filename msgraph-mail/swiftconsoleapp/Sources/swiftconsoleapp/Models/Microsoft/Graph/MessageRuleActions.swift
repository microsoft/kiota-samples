extension Swiftconsoleapp.Models.Microsoft.Graph {
public class MessageRuleActions : IAdditionalDataHolder {
public var additionalData: IDictionary<string, object>
public var assignCategories: [String]
public var copyToFolder: String
public var delete: Bool?
public var forwardAsAttachmentTo: [Recipient]
public var forwardTo: [Recipient]
public var markAsRead: Bool?
public var markImportance: Importance?
public var moveToFolder: String
public var permanentDelete: Bool?
public var redirectTo: [Recipient]
public var stopProcessingRules: Bool?
}
}
