package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// MessageRuleActionsable 
type MessageRuleActionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignCategories()([]string)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetCopyToFolder()(*string)
    GetDelete()(*bool)
    GetForwardAsAttachmentTo()([]Recipientable)
    GetForwardTo()([]Recipientable)
    GetMarkAsRead()(*bool)
    GetMarkImportance()(*Importance)
    GetMoveToFolder()(*string)
    GetPermanentDelete()(*bool)
    GetRedirectTo()([]Recipientable)
    GetStopProcessingRules()(*bool)
    SetAssignCategories(value []string)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetCopyToFolder(value *string)()
    SetDelete(value *bool)()
    SetForwardAsAttachmentTo(value []Recipientable)()
    SetForwardTo(value []Recipientable)()
    SetMarkAsRead(value *bool)()
    SetMarkImportance(value *Importance)()
    SetMoveToFolder(value *string)()
    SetPermanentDelete(value *bool)()
    SetRedirectTo(value []Recipientable)()
    SetStopProcessingRules(value *bool)()
}
