package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MessageRuleActions 
type MessageRuleActions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // A list of categories to be assigned to a message.
    assignCategories []string
    // The ID of a folder that a message is to be copied to.
    copyToFolder *string
    // Indicates whether a message should be moved to the Deleted Items folder.
    delete *bool
    // The email addresses of the recipients to which a message should be forwarded as an attachment.
    forwardAsAttachmentTo []Recipientable
    // The email addresses of the recipients to which a message should be forwarded.
    forwardTo []Recipientable
    // Indicates whether a message should be marked as read.
    markAsRead *bool
    // The markImportance property
    markImportance *Importance
    // The ID of the folder that a message will be moved to.
    moveToFolder *string
    // Indicates whether a message should be permanently deleted and not saved to the Deleted Items folder.
    permanentDelete *bool
    // The email address to which a message should be redirected.
    redirectTo []Recipientable
    // Indicates whether subsequent rules should be evaluated.
    stopProcessingRules *bool
}
// NewMessageRuleActions instantiates a new messageRuleActions and sets the default values.
func NewMessageRuleActions()(*MessageRuleActions) {
    m := &MessageRuleActions{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMessageRuleActionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMessageRuleActionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMessageRuleActions(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MessageRuleActions) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAssignCategories gets the assignCategories property value. A list of categories to be assigned to a message.
func (m *MessageRuleActions) GetAssignCategories()([]string) {
    return m.assignCategories
}
// GetCopyToFolder gets the copyToFolder property value. The ID of a folder that a message is to be copied to.
func (m *MessageRuleActions) GetCopyToFolder()(*string) {
    return m.copyToFolder
}
// GetDelete gets the delete property value. Indicates whether a message should be moved to the Deleted Items folder.
func (m *MessageRuleActions) GetDelete()(*bool) {
    return m.delete
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MessageRuleActions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assignCategories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        return i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfReferencedPrimitiveValue(n.GetCollectionOfPrimitiveValues , string , m.SetAssignCategories)
    }
    res["copyToFolder"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        return i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetValue(n.GetStringValue , m.SetCopyToFolder)
    }
    res["delete"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        return i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetValue(n.GetBoolValue , m.SetDelete)
    }
    res["forwardAsAttachmentTo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        return i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionValue(n.GetCollectionOfObjectValues , CreateRecipientFromDiscriminatorValue , m.SetForwardAsAttachmentTo)
    }
    res["forwardTo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        return i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionValue(n.GetCollectionOfObjectValues , CreateRecipientFromDiscriminatorValue , m.SetForwardTo)
    }
    res["markAsRead"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        return i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetValue(n.GetBoolValue , m.SetMarkAsRead)
    }
    res["markImportance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        return i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetReferencedEnumValue(n.GetEnumValue , ParseImportance , m.SetMarkImportance)
    }
    res["moveToFolder"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        return i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetValue(n.GetStringValue , m.SetMoveToFolder)
    }
    res["permanentDelete"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        return i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetValue(n.GetBoolValue , m.SetPermanentDelete)
    }
    res["redirectTo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        return i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionValue(n.GetCollectionOfObjectValues , CreateRecipientFromDiscriminatorValue , m.SetRedirectTo)
    }
    res["stopProcessingRules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        return i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetValue(n.GetBoolValue , m.SetStopProcessingRules)
    }
    return res
}
// GetForwardAsAttachmentTo gets the forwardAsAttachmentTo property value. The email addresses of the recipients to which a message should be forwarded as an attachment.
func (m *MessageRuleActions) GetForwardAsAttachmentTo()([]Recipientable) {
    return m.forwardAsAttachmentTo
}
// GetForwardTo gets the forwardTo property value. The email addresses of the recipients to which a message should be forwarded.
func (m *MessageRuleActions) GetForwardTo()([]Recipientable) {
    return m.forwardTo
}
// GetMarkAsRead gets the markAsRead property value. Indicates whether a message should be marked as read.
func (m *MessageRuleActions) GetMarkAsRead()(*bool) {
    return m.markAsRead
}
// GetMarkImportance gets the markImportance property value. The markImportance property
func (m *MessageRuleActions) GetMarkImportance()(*Importance) {
    return m.markImportance
}
// GetMoveToFolder gets the moveToFolder property value. The ID of the folder that a message will be moved to.
func (m *MessageRuleActions) GetMoveToFolder()(*string) {
    return m.moveToFolder
}
// GetPermanentDelete gets the permanentDelete property value. Indicates whether a message should be permanently deleted and not saved to the Deleted Items folder.
func (m *MessageRuleActions) GetPermanentDelete()(*bool) {
    return m.permanentDelete
}
// GetRedirectTo gets the redirectTo property value. The email address to which a message should be redirected.
func (m *MessageRuleActions) GetRedirectTo()([]Recipientable) {
    return m.redirectTo
}
// GetStopProcessingRules gets the stopProcessingRules property value. Indicates whether subsequent rules should be evaluated.
func (m *MessageRuleActions) GetStopProcessingRules()(*bool) {
    return m.stopProcessingRules
}
// Serialize serializes information the current object
func (m *MessageRuleActions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAssignCategories() != nil {
        err := writer.WriteCollectionOfStringValues("assignCategories", m.GetAssignCategories())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("copyToFolder", m.GetCopyToFolder())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("delete", m.GetDelete())
        if err != nil {
            return err
        }
    }
    if m.GetForwardAsAttachmentTo() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetForwardAsAttachmentTo()))
        for i, v := range m.GetForwardAsAttachmentTo() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("forwardAsAttachmentTo", cast)
        if err != nil {
            return err
        }
    }
    if m.GetForwardTo() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetForwardTo()))
        for i, v := range m.GetForwardTo() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("forwardTo", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("markAsRead", m.GetMarkAsRead())
        if err != nil {
            return err
        }
    }
    if m.GetMarkImportance() != nil {
        cast := (*m.GetMarkImportance()).String()
        err := writer.WriteStringValue("markImportance", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("moveToFolder", m.GetMoveToFolder())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("permanentDelete", m.GetPermanentDelete())
        if err != nil {
            return err
        }
    }
    if m.GetRedirectTo() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRedirectTo()))
        for i, v := range m.GetRedirectTo() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("redirectTo", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("stopProcessingRules", m.GetStopProcessingRules())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MessageRuleActions) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAssignCategories sets the assignCategories property value. A list of categories to be assigned to a message.
func (m *MessageRuleActions) SetAssignCategories(value []string)() {
    m.assignCategories = value
}
// SetCopyToFolder sets the copyToFolder property value. The ID of a folder that a message is to be copied to.
func (m *MessageRuleActions) SetCopyToFolder(value *string)() {
    m.copyToFolder = value
}
// SetDelete sets the delete property value. Indicates whether a message should be moved to the Deleted Items folder.
func (m *MessageRuleActions) SetDelete(value *bool)() {
    m.delete = value
}
// SetForwardAsAttachmentTo sets the forwardAsAttachmentTo property value. The email addresses of the recipients to which a message should be forwarded as an attachment.
func (m *MessageRuleActions) SetForwardAsAttachmentTo(value []Recipientable)() {
    m.forwardAsAttachmentTo = value
}
// SetForwardTo sets the forwardTo property value. The email addresses of the recipients to which a message should be forwarded.
func (m *MessageRuleActions) SetForwardTo(value []Recipientable)() {
    m.forwardTo = value
}
// SetMarkAsRead sets the markAsRead property value. Indicates whether a message should be marked as read.
func (m *MessageRuleActions) SetMarkAsRead(value *bool)() {
    m.markAsRead = value
}
// SetMarkImportance sets the markImportance property value. The markImportance property
func (m *MessageRuleActions) SetMarkImportance(value *Importance)() {
    m.markImportance = value
}
// SetMoveToFolder sets the moveToFolder property value. The ID of the folder that a message will be moved to.
func (m *MessageRuleActions) SetMoveToFolder(value *string)() {
    m.moveToFolder = value
}
// SetPermanentDelete sets the permanentDelete property value. Indicates whether a message should be permanently deleted and not saved to the Deleted Items folder.
func (m *MessageRuleActions) SetPermanentDelete(value *bool)() {
    m.permanentDelete = value
}
// SetRedirectTo sets the redirectTo property value. The email address to which a message should be redirected.
func (m *MessageRuleActions) SetRedirectTo(value []Recipientable)() {
    m.redirectTo = value
}
// SetStopProcessingRules sets the stopProcessingRules property value. Indicates whether subsequent rules should be evaluated.
func (m *MessageRuleActions) SetStopProcessingRules(value *bool)() {
    m.stopProcessingRules = value
}
