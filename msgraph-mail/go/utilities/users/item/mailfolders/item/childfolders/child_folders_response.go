package childfolders

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/models/microsoft/graph"
)

// ChildFoldersResponse 
type ChildFoldersResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The nextLink property
    nextLink *string
    // The value property
    value []i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.MailFolderable
}
// NewChildFoldersResponse instantiates a new childFoldersResponse and sets the default values.
func NewChildFoldersResponse()(*ChildFoldersResponse) {
    m := &ChildFoldersResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateChildFoldersResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChildFoldersResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChildFoldersResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ChildFoldersResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChildFoldersResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["@odata.nextLink"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdatanextLink(val)
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.CreateMailFolderFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.MailFolderable, len(val))
            for i, v := range val {
                res[i] = v.(i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.MailFolderable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetOdatanextLink gets the @odata.nextLink property value. The nextLink property
func (m *ChildFoldersResponse) GetOdatanextLink()(*string) {
    if m == nil {
        return nil
    } else {
        return m.nextLink
    }
}
// GetValue gets the value property value. The value property
func (m *ChildFoldersResponse) GetValue()([]i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.MailFolderable) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
// Serialize serializes information the current object
func (m *ChildFoldersResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.nextLink", m.GetOdatanextLink())
        if err != nil {
            return err
        }
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
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
func (m *ChildFoldersResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetOdatanextLink sets the @odata.nextLink property value. The nextLink property
func (m *ChildFoldersResponse) SetOdatanextLink(value *string)() {
    if m != nil {
        m.nextLink = value
    }
}
// SetValue sets the value property value. The value property
func (m *ChildFoldersResponse) SetValue(value []i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.MailFolderable)() {
    if m != nil {
        m.value = value
    }
}
