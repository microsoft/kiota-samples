package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Attachment struct {
    Entity
    contentType *string;
    isInline *bool;
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    name *string;
    size *int32;
}
func NewAttachment()(*Attachment) {
    m := &Attachment{
        Entity: *NewEntity(),
    }
    return m
}
func (m *Attachment) GetContentType()(*string) {
    return m.contentType
}
func (m *Attachment) GetIsInline()(*bool) {
    return m.isInline
}
func (m *Attachment) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
func (m *Attachment) GetName()(*string) {
    return m.name
}
func (m *Attachment) GetSize()(*int32) {
    return m.size
}
func (m *Attachment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error), error) {
    return nil, nil
}
func (m *Attachment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    return nil
}
func (m *Attachment) SetContentType(value *string)() {
    m.contentType = value
}
func (m *Attachment) SetIsInline(value *bool)() {
    m.isInline = value
}
func (m *Attachment) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
func (m *Attachment) SetName(value *string)() {
    m.name = value
}
func (m *Attachment) SetSize(value *int32)() {
    m.size = value
}
