import Foundation
import MicrosoftKiotaAbstractions

extension Swiftconsoleapp.Models {
    public class Attachment : Entity {
        public var contentType: String?
        public var isInline: Bool?
        public var lastModifiedDateTime: Date?
        public var name: String?
        public var size: Int32?
    }
}
