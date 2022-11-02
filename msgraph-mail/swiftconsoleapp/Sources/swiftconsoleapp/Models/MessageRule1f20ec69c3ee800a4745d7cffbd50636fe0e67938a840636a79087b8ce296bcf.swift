import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Models {
    public class MessageRule : Entity {
        public var actions: MessageRuleActions?
        public var conditions: MessageRulePredicates?
        public var displayName: String?
        public var exceptions: MessageRulePredicates?
        public var hasError: Bool?
        public var isEnabled: Bool?
        public var isReadOnly: Bool?
        public var sequence: Int32?
    }
}
