package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ToDoItem 
type ToDoItem struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The dueDate property
    dueDate *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The id property
    id *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // The isComplete property
    isComplete *bool
    // The name property
    name *string
    // The priority property
    priority *int32
    // The type property
    typeEscaped *string
}
// NewToDoItem instantiates a new ToDoItem and sets the default values.
func NewToDoItem()(*ToDoItem) {
    m := &ToDoItem{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateToDoItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateToDoItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewToDoItem(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ToDoItem) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDueDate gets the dueDate property value. The dueDate property
func (m *ToDoItem) GetDueDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.dueDate
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ToDoItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["dueDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDueDate(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["isComplete"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsComplete(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["priority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriority(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The id property
func (m *ToDoItem) GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.id
}
// GetIsComplete gets the isComplete property value. The isComplete property
func (m *ToDoItem) GetIsComplete()(*bool) {
    return m.isComplete
}
// GetName gets the name property value. The name property
func (m *ToDoItem) GetName()(*string) {
    return m.name
}
// GetPriority gets the priority property value. The priority property
func (m *ToDoItem) GetPriority()(*int32) {
    return m.priority
}
// GetType gets the type property value. The type property
func (m *ToDoItem) GetType()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *ToDoItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("dueDate", m.GetDueDate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteUUIDValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isComplete", m.GetIsComplete())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("priority", m.GetPriority())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetType())
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
func (m *ToDoItem) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDueDate sets the dueDate property value. The dueDate property
func (m *ToDoItem) SetDueDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.dueDate = value
}
// SetId sets the id property value. The id property
func (m *ToDoItem) SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.id = value
}
// SetIsComplete sets the isComplete property value. The isComplete property
func (m *ToDoItem) SetIsComplete(value *bool)() {
    m.isComplete = value
}
// SetName sets the name property value. The name property
func (m *ToDoItem) SetName(value *string)() {
    m.name = value
}
// SetPriority sets the priority property value. The priority property
func (m *ToDoItem) SetPriority(value *int32)() {
    m.priority = value
}
// SetType sets the type property value. The type property
func (m *ToDoItem) SetType(value *string)() {
    m.typeEscaped = value
}
// ToDoItemable 
type ToDoItemable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDueDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetIsComplete()(*bool)
    GetName()(*string)
    GetPriority()(*int32)
    GetType()(*string)
    SetDueDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetIsComplete(value *bool)()
    SetName(value *string)()
    SetPriority(value *int32)()
    SetType(value *string)()
}
