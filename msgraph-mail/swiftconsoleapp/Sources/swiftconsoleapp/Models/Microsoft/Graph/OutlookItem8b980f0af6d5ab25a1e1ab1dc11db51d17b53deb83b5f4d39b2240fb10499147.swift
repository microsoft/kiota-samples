import Foundation
import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Models.Microsoft.Graph {
    public class OutlookItem : Entity {
        public var categories: [String]?
        public var changeKey: String?
        public var createdDateTime: Date?
        public var lastModifiedDateTime: Date?
    }
}
