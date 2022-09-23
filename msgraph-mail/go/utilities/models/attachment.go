package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    utils "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Attachment 
type Attachment struct {
    Entity
    // The MIME type.
    contentType *string
    // true if the attachment is an inline attachment; otherwise, false.
    isInline *bool
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The attachment's file name.
    name *string
    // The length of the attachment in bytes.
    size *int32
}
// NewAttachment instantiates a new Attachment and sets the default values.
func NewAttachment()(*Attachment) {
    m := &Attachment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAttachmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAttachmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAttachment(), nil
}
// GetContentType gets the contentType property value. The MIME type.
func (m *Attachment) GetContentType()(*string) {
    return m.contentType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Attachment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["contentType"] = utils.SetStringValue(m.SetContentType)
    res["isInline"] = utils.SetBoolValue(m.SetIsInline)
    res["lastModifiedDateTime"] = utils.SetTimeValue(m.SetLastModifiedDateTime)
    res["name"] = utils.SetStringValue(m.SetName)
    res["size"] = utils.SetInt32Value(m.SetSize)
    return res
}
// GetIsInline gets the isInline property value. true if the attachment is an inline attachment; otherwise, false.
func (m *Attachment) GetIsInline()(*bool) {
    return m.isInline
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *Attachment) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetName gets the name property value. The attachment's file name.
func (m *Attachment) GetName()(*string) {
    return m.name
}
// GetSize gets the size property value. The length of the attachment in bytes.
func (m *Attachment) GetSize()(*int32) {
    return m.size
}
// Serialize serializes information the current object
func (m *Attachment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("contentType", m.GetContentType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isInline", m.GetIsInline())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("size", m.GetSize())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContentType sets the contentType property value. The MIME type.
func (m *Attachment) SetContentType(value *string)() {
    m.contentType = value
}
// SetIsInline sets the isInline property value. true if the attachment is an inline attachment; otherwise, false.
func (m *Attachment) SetIsInline(value *bool)() {
    m.isInline = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *Attachment) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetName sets the name property value. The attachment's file name.
func (m *Attachment) SetName(value *string)() {
    m.name = value
}
// SetSize sets the size property value. The length of the attachment in bytes.
func (m *Attachment) SetSize(value *int32)() {
    m.size = value
}
