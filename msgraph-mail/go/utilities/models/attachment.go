package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Attachment struct {
    contentType *string;
    isInline *bool;
    lastModifiedDateTime i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    name *string;
    size *int32;
}
func (m *Attachment) getContentType()(*string) {
    return nil
}
func (m *Attachment) getIsInline()(*bool) {
    return nil
}
func (m *Attachment) getLastModifiedDateTime()(i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return nil
}
func (m *Attachment) getName()(*string) {
    return nil
}
func (m *Attachment) getSize()(*int32) {
    return nil
}
func (m *Attachment) getFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error), error) {
    return nil
}
func (m *Attachment) serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    return nil
}
func (m *Attachment) setContentType(value *string)() {
    return nil
}
func (m *Attachment) setIsInline(value *bool)() {
    return nil
}
func (m *Attachment) setLastModifiedDateTime(value i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    return nil
}
func (m *Attachment) setName(value *string)() {
    return nil
}
func (m *Attachment) setSize(value *int32)() {
    return nil
}
