package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SingleValueLegacyExtendedPropertyCollectionResponse 
type SingleValueLegacyExtendedPropertyCollectionResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataNextLink property
    odataNextLink *string
    // The value property
    value []SingleValueLegacyExtendedPropertyable
}
// NewSingleValueLegacyExtendedPropertyCollectionResponse instantiates a new SingleValueLegacyExtendedPropertyCollectionResponse and sets the default values.
func NewSingleValueLegacyExtendedPropertyCollectionResponse()(*SingleValueLegacyExtendedPropertyCollectionResponse) {
    m := &SingleValueLegacyExtendedPropertyCollectionResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSingleValueLegacyExtendedPropertyCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSingleValueLegacyExtendedPropertyCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSingleValueLegacyExtendedPropertyCollectionResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SingleValueLegacyExtendedPropertyCollectionResponse) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SingleValueLegacyExtendedPropertyCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["@odata.nextLink"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataNextLink)
    res["value"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSingleValueLegacyExtendedPropertyFromDiscriminatorValue , m.SetValue)
    return res
}
// GetOdataNextLink gets the @odata.nextLink property value. The OdataNextLink property
func (m *SingleValueLegacyExtendedPropertyCollectionResponse) GetOdataNextLink()(*string) {
    return m.odataNextLink
}
// GetValue gets the value property value. The value property
func (m *SingleValueLegacyExtendedPropertyCollectionResponse) GetValue()([]SingleValueLegacyExtendedPropertyable) {
    return m.value
}
// Serialize serializes information the current object
func (m *SingleValueLegacyExtendedPropertyCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.nextLink", m.GetOdataNextLink())
        if err != nil {
            return err
        }
    }
    if m.GetValue() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetValue())
        err := writer.WriteCollectionOfObjectValues("value", cast)
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
func (m *SingleValueLegacyExtendedPropertyCollectionResponse) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataNextLink sets the @odata.nextLink property value. The OdataNextLink property
func (m *SingleValueLegacyExtendedPropertyCollectionResponse) SetOdataNextLink(value *string)() {
    m.odataNextLink = value
}
// SetValue sets the value property value. The value property
func (m *SingleValueLegacyExtendedPropertyCollectionResponse) SetValue(value []SingleValueLegacyExtendedPropertyable)() {
    m.value = value
}
