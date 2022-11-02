import Foundation
import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Models {
    public class Message : OutlookItem {
        public var attachments: [Attachment]?
        public var bccRecipients: [Recipient]?
        public var body: ItemBody?
        public var bodyPreview: String?
        public var ccRecipients: [Recipient]?
        public var conversationId: String?
        public var conversationIndex: [UInt8]?
        public var extensions: [Extension]?
        public var flag: FollowupFlag?
        public var from: Recipient?
        public var hasAttachments: Bool?
        public var importance: Importance?
        public var inferenceClassification: InferenceClassificationType?
        public var internetMessageHeaders: [InternetMessageHeader]?
        public var internetMessageId: String?
        public var isDeliveryReceiptRequested: Bool?
        public var isDraft: Bool?
        public var isRead: Bool?
        public var isReadReceiptRequested: Bool?
        public var multiValueExtendedProperties: [MultiValueLegacyExtendedProperty]?
        public var parentFolderId: String?
        public var receivedDateTime: Date?
        public var replyTo: [Recipient]?
        public var sender: Recipient?
        public var sentDateTime: Date?
        public var singleValueExtendedProperties: [SingleValueLegacyExtendedProperty]?
        public var subject: String?
        public var toRecipients: [Recipient]?
        public var uniqueBody: ItemBody?
        public var webLink: String?
    }
}
