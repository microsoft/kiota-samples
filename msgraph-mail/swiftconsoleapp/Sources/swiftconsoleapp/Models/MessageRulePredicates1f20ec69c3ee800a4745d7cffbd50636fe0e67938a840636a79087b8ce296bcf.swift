import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Models {
    public class MessageRulePredicates : AdditionalDataHolder {
        public var additionalData: [String:Any] = [String:Any]()
        public var bodyContains: [String]?
        public var bodyOrSubjectContains: [String]?
        public var categories: [String]?
        public var fromAddresses: [Recipient]?
        public var hasAttachments: Bool?
        public var headerContains: [String]?
        public var importance: Importance?
        public var isApprovalRequest: Bool?
        public var isAutomaticForward: Bool?
        public var isAutomaticReply: Bool?
        public var isEncrypted: Bool?
        public var isMeetingRequest: Bool?
        public var isMeetingResponse: Bool?
        public var isNonDeliveryReport: Bool?
        public var isPermissionControlled: Bool?
        public var isReadReceipt: Bool?
        public var isSigned: Bool?
        public var isVoicemail: Bool?
        public var messageActionFlag: MessageActionFlag?
        public var notSentToMe: Bool?
        public var recipientContains: [String]?
        public var senderContains: [String]?
        public var sensitivity: Sensitivity?
        public var sentCcMe: Bool?
        public var sentOnlyToMe: Bool?
        public var sentToAddresses: [Recipient]?
        public var sentToMe: Bool?
        public var sentToOrCcMe: Bool?
        public var subjectContains: [String]?
        public var withinSizeRange: SizeRange?
    }
}
