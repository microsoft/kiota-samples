package graph
// MailFolderable 
type MailFolderable interface {
    Entityable
    GetChildFolderCount()(*int32)
    GetChildFolders()([]MailFolderable)
    GetDisplayName()(*string)
    GetIsHidden()(*bool)
    GetMessageRules()([]MessageRuleable)
    GetMessages()([]Messageable)
    GetMultiValueExtendedProperties()([]MultiValueLegacyExtendedPropertyable)
    GetParentFolderId()(*string)
    GetSingleValueExtendedProperties()([]SingleValueLegacyExtendedPropertyable)
    GetTotalItemCount()(*int32)
    GetUnreadItemCount()(*int32)
    SetChildFolderCount(value *int32)()
    SetChildFolders(value []MailFolderable)()
    SetDisplayName(value *string)()
    SetIsHidden(value *bool)()
    SetMessageRules(value []MessageRuleable)()
    SetMessages(value []Messageable)()
    SetMultiValueExtendedProperties(value []MultiValueLegacyExtendedPropertyable)()
    SetParentFolderId(value *string)()
    SetSingleValueExtendedProperties(value []SingleValueLegacyExtendedPropertyable)()
    SetTotalItemCount(value *int32)()
    SetUnreadItemCount(value *int32)()
}
