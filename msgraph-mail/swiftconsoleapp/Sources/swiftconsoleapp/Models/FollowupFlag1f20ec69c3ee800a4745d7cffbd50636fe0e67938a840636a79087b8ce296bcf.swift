import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Models {
    public class FollowupFlag : AdditionalDataHolder {
        public var additionalData: [String:Any] = [String:Any]()
        public var completedDateTime: DateTimeTimeZone?
        public var dueDateTime: DateTimeTimeZone?
        public var flagStatus: FollowupFlagStatus?
        public var startDateTime: DateTimeTimeZone?
    }
}
