package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Attachment struct {
    contentType *string;
    isInline *bool;
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    name *string;
    size *int32;
}
func NewAttachment()(*Attachment) {
    return nil
}
func (m *Attachment) getContentType()(*string) {
    return m.contentType
}
func (m *Attachment) getIsInline()(*bool) {
    return m.isInline
}
func (m *Attachment) getLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
func (m *Attachment) getName()(*string) {
    return m.name
}
func (m *Attachment) getSize()(*int32) {
    return m.size
}
func (m *Attachment) getFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error), error) {
    return nil, nil
}
func (m *Attachment) serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    return nil
}
func (m *Attachment) setContentType(value *string)() {
    m.contentType = value
}
func (m *Attachment) setIsInline(value *bool)() {
    m.isInline = value
}
func (m *Attachment) setLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
func (m *Attachment) setName(value *string)() {
    m.name = value
}
func (m *Attachment) setSize(value *int32)() {
    m.size = value
}
