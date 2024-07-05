package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Pet struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The category property
    category Categoryable
    // The id property
    id *int64
    // The name property
    name *string
    // The photoUrls property
    photoUrls []string
    // pet status in the store
    status *Pet_status
    // The tags property
    tags []Tagable
}
// NewPet instantiates a new Pet and sets the default values.
func NewPet()(*Pet) {
    m := &Pet{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPet(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Pet) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCategory gets the category property value. The category property
// returns a Categoryable when successful
func (m *Pet) GetCategory()(Categoryable) {
    return m.category
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Pet) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["category"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory(val.(Categoryable))
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
    res["photoUrls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetPhotoUrls(res)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePet_status)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*Pet_status))
        }
        return nil
    }
    res["tags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTagFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Tagable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Tagable)
                }
            }
            m.SetTags(res)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *Pet) GetId()(*int64) {
    return m.id
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *Pet) GetName()(*string) {
    return m.name
}
// GetPhotoUrls gets the photoUrls property value. The photoUrls property
// returns a []string when successful
func (m *Pet) GetPhotoUrls()([]string) {
    return m.photoUrls
}
// GetStatus gets the status property value. pet status in the store
// returns a *Pet_status when successful
func (m *Pet) GetStatus()(*Pet_status) {
    return m.status
}
// GetTags gets the tags property value. The tags property
// returns a []Tagable when successful
func (m *Pet) GetTags()([]Tagable) {
    return m.tags
}
// Serialize serializes information the current object
func (m *Pet) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("category", m.GetCategory())
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
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    if m.GetPhotoUrls() != nil {
        err := writer.WriteCollectionOfStringValues("photoUrls", m.GetPhotoUrls())
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
    if m.GetTags() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTags()))
        for i, v := range m.GetTags() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("tags", cast)
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
func (m *Pet) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCategory sets the category property value. The category property
func (m *Pet) SetCategory(value Categoryable)() {
    m.category = value
}
// SetId sets the id property value. The id property
func (m *Pet) SetId(value *int64)() {
    m.id = value
}
// SetName sets the name property value. The name property
func (m *Pet) SetName(value *string)() {
    m.name = value
}
// SetPhotoUrls sets the photoUrls property value. The photoUrls property
func (m *Pet) SetPhotoUrls(value []string)() {
    m.photoUrls = value
}
// SetStatus sets the status property value. pet status in the store
func (m *Pet) SetStatus(value *Pet_status)() {
    m.status = value
}
// SetTags sets the tags property value. The tags property
func (m *Pet) SetTags(value []Tagable)() {
    m.tags = value
}
type Petable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCategory()(Categoryable)
    GetId()(*int64)
    GetName()(*string)
    GetPhotoUrls()([]string)
    GetStatus()(*Pet_status)
    GetTags()([]Tagable)
    SetCategory(value Categoryable)()
    SetId(value *int64)()
    SetName(value *string)()
    SetPhotoUrls(value []string)()
    SetStatus(value *Pet_status)()
    SetTags(value []Tagable)()
}
