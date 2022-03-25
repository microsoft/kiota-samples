extension Swiftconsoleapp.Models.Microsoft.Graph {
public class FollowupFlag : IAdditionalDataHolder {
public var additionalData: IDictionary<string, object>
public var completedDateTime: DateTimeTimeZone
public var dueDateTime: DateTimeTimeZone
public var flagStatus: FollowupFlagStatus?
public var startDateTime: DateTimeTimeZone
}
}
