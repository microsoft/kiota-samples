import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Models {
    public class InferenceClassificationOverride : Entity {
        public var classifyAs: InferenceClassificationType?
        public var senderEmailAddress: EmailAddress?
    }
}
