package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InferenceClassificationOverride 
type InferenceClassificationOverride struct {
    Entity
    // The classifyAs property
    classifyAs *InferenceClassificationType
    // The senderEmailAddress property
    senderEmailAddress EmailAddressable
}
// NewInferenceClassificationOverride instantiates a new InferenceClassificationOverride and sets the default values.
func NewInferenceClassificationOverride()(*InferenceClassificationOverride) {
    m := &InferenceClassificationOverride{
        Entity: *NewEntity(),
    }
    return m
}
// CreateInferenceClassificationOverrideFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInferenceClassificationOverrideFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInferenceClassificationOverride(), nil
}
// GetClassifyAs gets the classifyAs property value. The classifyAs property
func (m *InferenceClassificationOverride) GetClassifyAs()(*InferenceClassificationType) {
    return m.classifyAs
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InferenceClassificationOverride) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["classifyAs"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseInferenceClassificationType , m.SetClassifyAs)
    res["senderEmailAddress"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateEmailAddressFromDiscriminatorValue , m.SetSenderEmailAddress)
    return res
}
// GetSenderEmailAddress gets the senderEmailAddress property value. The senderEmailAddress property
func (m *InferenceClassificationOverride) GetSenderEmailAddress()(EmailAddressable) {
    return m.senderEmailAddress
}
// Serialize serializes information the current object
func (m *InferenceClassificationOverride) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetClassifyAs() != nil {
        cast := (*m.GetClassifyAs()).String()
        err = writer.WriteStringValue("classifyAs", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("senderEmailAddress", m.GetSenderEmailAddress())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetClassifyAs sets the classifyAs property value. The classifyAs property
func (m *InferenceClassificationOverride) SetClassifyAs(value *InferenceClassificationType)() {
    m.classifyAs = value
}
// SetSenderEmailAddress sets the senderEmailAddress property value. The senderEmailAddress property
func (m *InferenceClassificationOverride) SetSenderEmailAddress(value EmailAddressable)() {
    m.senderEmailAddress = value
}
