package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Recipient 
type Recipient struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The emailAddress property
    emailAddress EmailAddressable
}
// NewRecipient instantiates a new recipient and sets the default values.
func NewRecipient()(*Recipient) {
    m := &Recipient{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateRecipientFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRecipientFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRecipient(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Recipient) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetEmailAddress gets the emailAddress property value. The emailAddress property
func (m *Recipient) GetEmailAddress()(EmailAddressable) {
    return m.emailAddress
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Recipient) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["emailAddress"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateEmailAddressFromDiscriminatorValue , m.SetEmailAddress)
    return res
}
// Serialize serializes information the current object
func (m *Recipient) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("emailAddress", m.GetEmailAddress())
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
func (m *Recipient) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetEmailAddress sets the emailAddress property value. The emailAddress property
func (m *Recipient) SetEmailAddress(value EmailAddressable)() {
    m.emailAddress = value
}
