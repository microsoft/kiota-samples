package users
import (
    "github.com/microsoft/kiota/abstractions/go/serialization"
)
type MailFolder struct {
    childFolderCount *int32;
    childFolders []MailFolder;
    displayName *string;
    isHidden *bool;
    messageRules []MessageRule;
    messages []Message;
    multiValueExtendedProperties []MultiValueLegacyExtendedProperty;
    parentFolderId *string;
    singleValueExtendedProperties []SingleValueLegacyExtendedProperty;
    totalItemCount *int32;
    unreadItemCount *int32;
}
