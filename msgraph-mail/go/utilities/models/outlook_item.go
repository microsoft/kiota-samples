package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type OutlookItem struct {
    categories []*string;
    changeKey *string;
    createdDateTime i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    lastModifiedDateTime i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
func (m *OutlookItem) getCategories()([]*string) {
    return nil
}
func (m *OutlookItem) getChangeKey()(*string) {
    return nil
}
func (m *OutlookItem) getCreatedDateTime()(i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return nil
}
func (m *OutlookItem) getLastModifiedDateTime()(i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return nil
}
func (m *OutlookItem) getFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error), error) {
    return nil
}
func (m *OutlookItem) serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    return nil
}
func (m *OutlookItem) setCategories(value []*string)() {
    return nil
}
func (m *OutlookItem) setChangeKey(value *string)() {
    return nil
}
func (m *OutlookItem) setCreatedDateTime(value i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    return nil
}
func (m *OutlookItem) setLastModifiedDateTime(value i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    return nil
}
