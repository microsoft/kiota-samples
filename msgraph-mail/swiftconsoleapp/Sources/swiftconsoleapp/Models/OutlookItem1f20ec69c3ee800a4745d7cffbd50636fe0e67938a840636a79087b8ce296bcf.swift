import Foundation
import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Models {
    public class OutlookItem : Entity {
        public var categories: [String]?
        public var changeKey: String?
        public var createdDateTime: Date?
        public var lastModifiedDateTime: Date?
    }
}
