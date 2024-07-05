package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type User struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The email property
    email *string
    // The firstName property
    firstName *string
    // The id property
    id *int64
    // The lastName property
    lastName *string
    // The password property
    password *string
    // The phone property
    phone *string
    // The username property
    username *string
    // User Status
    userStatus *int32
}
// NewUser instantiates a new User and sets the default values.
func NewUser()(*User) {
    m := &User{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUserFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUser(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *User) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEmail gets the email property value. The email property
// returns a *string when successful
func (m *User) GetEmail()(*string) {
    return m.email
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *User) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
    res["firstName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirstName(val)
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
    res["lastName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastName(val)
        }
        return nil
    }
    res["password"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPassword(val)
        }
        return nil
    }
    res["phone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhone(val)
        }
        return nil
    }
    res["username"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsername(val)
        }
        return nil
    }
    res["userStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserStatus(val)
        }
        return nil
    }
    return res
}
// GetFirstName gets the firstName property value. The firstName property
// returns a *string when successful
func (m *User) GetFirstName()(*string) {
    return m.firstName
}
// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *User) GetId()(*int64) {
    return m.id
}
// GetLastName gets the lastName property value. The lastName property
// returns a *string when successful
func (m *User) GetLastName()(*string) {
    return m.lastName
}
// GetPassword gets the password property value. The password property
// returns a *string when successful
func (m *User) GetPassword()(*string) {
    return m.password
}
// GetPhone gets the phone property value. The phone property
// returns a *string when successful
func (m *User) GetPhone()(*string) {
    return m.phone
}
// GetUsername gets the username property value. The username property
// returns a *string when successful
func (m *User) GetUsername()(*string) {
    return m.username
}
// GetUserStatus gets the userStatus property value. User Status
// returns a *int32 when successful
func (m *User) GetUserStatus()(*int32) {
    return m.userStatus
}
// Serialize serializes information the current object
func (m *User) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("firstName", m.GetFirstName())
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
        err := writer.WriteStringValue("lastName", m.GetLastName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("password", m.GetPassword())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("phone", m.GetPhone())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("username", m.GetUsername())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("userStatus", m.GetUserStatus())
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
func (m *User) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEmail sets the email property value. The email property
func (m *User) SetEmail(value *string)() {
    m.email = value
}
// SetFirstName sets the firstName property value. The firstName property
func (m *User) SetFirstName(value *string)() {
    m.firstName = value
}
// SetId sets the id property value. The id property
func (m *User) SetId(value *int64)() {
    m.id = value
}
// SetLastName sets the lastName property value. The lastName property
func (m *User) SetLastName(value *string)() {
    m.lastName = value
}
// SetPassword sets the password property value. The password property
func (m *User) SetPassword(value *string)() {
    m.password = value
}
// SetPhone sets the phone property value. The phone property
func (m *User) SetPhone(value *string)() {
    m.phone = value
}
// SetUsername sets the username property value. The username property
func (m *User) SetUsername(value *string)() {
    m.username = value
}
// SetUserStatus sets the userStatus property value. User Status
func (m *User) SetUserStatus(value *int32)() {
    m.userStatus = value
}
type Userable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEmail()(*string)
    GetFirstName()(*string)
    GetId()(*int64)
    GetLastName()(*string)
    GetPassword()(*string)
    GetPhone()(*string)
    GetUsername()(*string)
    GetUserStatus()(*int32)
    SetEmail(value *string)()
    SetFirstName(value *string)()
    SetId(value *int64)()
    SetLastName(value *string)()
    SetPassword(value *string)()
    SetPhone(value *string)()
    SetUsername(value *string)()
    SetUserStatus(value *int32)()
}
