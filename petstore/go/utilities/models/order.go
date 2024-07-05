package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Order struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The complete property
    complete *bool
    // The id property
    id *int64
    // The petId property
    petId *int64
    // The quantity property
    quantity *int32
    // The shipDate property
    shipDate *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Order Status
    status *Order_status
}
// NewOrder instantiates a new Order and sets the default values.
func NewOrder()(*Order) {
    m := &Order{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOrderFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrderFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrder(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Order) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetComplete gets the complete property value. The complete property
// returns a *bool when successful
func (m *Order) GetComplete()(*bool) {
    return m.complete
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Order) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["complete"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComplete(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["petId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPetId(val)
        }
        return nil
    }
    res["quantity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuantity(val)
        }
        return nil
    }
    res["shipDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShipDate(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOrder_status)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*Order_status))
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *Order) GetId()(*int64) {
    return m.id
}
// GetPetId gets the petId property value. The petId property
// returns a *int64 when successful
func (m *Order) GetPetId()(*int64) {
    return m.petId
}
// GetQuantity gets the quantity property value. The quantity property
// returns a *int32 when successful
func (m *Order) GetQuantity()(*int32) {
    return m.quantity
}
// GetShipDate gets the shipDate property value. The shipDate property
// returns a *Time when successful
func (m *Order) GetShipDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.shipDate
}
// GetStatus gets the status property value. Order Status
// returns a *Order_status when successful
func (m *Order) GetStatus()(*Order_status) {
    return m.status
}
// Serialize serializes information the current object
func (m *Order) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("complete", m.GetComplete())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("petId", m.GetPetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("quantity", m.GetQuantity())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("shipDate", m.GetShipDate())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Order) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetComplete sets the complete property value. The complete property
func (m *Order) SetComplete(value *bool)() {
    m.complete = value
}
// SetId sets the id property value. The id property
func (m *Order) SetId(value *int64)() {
    m.id = value
}
// SetPetId sets the petId property value. The petId property
func (m *Order) SetPetId(value *int64)() {
    m.petId = value
}
// SetQuantity sets the quantity property value. The quantity property
func (m *Order) SetQuantity(value *int32)() {
    m.quantity = value
}
// SetShipDate sets the shipDate property value. The shipDate property
func (m *Order) SetShipDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.shipDate = value
}
// SetStatus sets the status property value. Order Status
func (m *Order) SetStatus(value *Order_status)() {
    m.status = value
}
type Orderable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetComplete()(*bool)
    GetId()(*int64)
    GetPetId()(*int64)
    GetQuantity()(*int32)
    GetShipDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetStatus()(*Order_status)
    SetComplete(value *bool)()
    SetId(value *int64)()
    SetPetId(value *int64)()
    SetQuantity(value *int32)()
    SetShipDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetStatus(value *Order_status)()
}
