package models

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type FollowupFlag struct {
    additionalData map[string]interface{};
    completedDateTime *DateTimeTimeZone;
    dueDateTime *DateTimeTimeZone;
    flagStatus *FollowupFlagStatus;
    startDateTime *DateTimeTimeZone;
}
func NewFollowupFlag()(*FollowupFlag) {
    m := &FollowupFlag{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *FollowupFlag) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
func (m *FollowupFlag) GetCompletedDateTime()(*DateTimeTimeZone) {
    return m.completedDateTime
}
func (m *FollowupFlag) GetDueDateTime()(*DateTimeTimeZone) {
    return m.dueDateTime
}
func (m *FollowupFlag) GetFlagStatus()(*FollowupFlagStatus) {
    return m.flagStatus
}
func (m *FollowupFlag) GetStartDateTime()(*DateTimeTimeZone) {
    return m.startDateTime
}
func (m *FollowupFlag) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["completedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () interface{} { return NewDateTimeTimeZone() })
        if err != nil {
            return err
        }
        o.(*FollowupFlag).SetCompletedDateTime(val.(*DateTimeTimeZone))
        return nil
    }
    res["dueDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () interface{} { return NewDateTimeTimeZone() })
        if err != nil {
            return err
        }
        o.(*FollowupFlag).SetDueDateTime(val.(*DateTimeTimeZone))
        return nil
    }
    res["flagStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseFollowupFlagStatus)
        if err != nil {
            return err
        }
        o.(*FollowupFlag).SetFlagStatus(val.(*FollowupFlagStatus))
        return nil
    }
    res["startDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () interface{} { return NewDateTimeTimeZone() })
        if err != nil {
            return err
        }
        o.(*FollowupFlag).SetStartDateTime(val.(*DateTimeTimeZone))
        return nil
    }
    return res
}
func (m *FollowupFlag) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := writer.WriteObjectValue("completedDateTime", m.GetCompletedDateTime())
    if err != nil {
        return err
    }
    err = writer.WriteObjectValue("dueDateTime", m.GetDueDateTime())
    if err != nil {
        return err
    }
    if m.GetFlagStatus() != nil {
        err = writer.WritePrimitiveValue("flagStatus", m.GetFlagStatus().String())
        if err != nil {
            return err
        }
    }
    err = writer.WriteObjectValue("startDateTime", m.GetStartDateTime())
    if err != nil {
        return err
    }
    err = writer.WriteAdditionalData(m.GetAdditionalData())
    if err != nil {
        return err
    }
    return nil
}
func (m *FollowupFlag) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *FollowupFlag) SetCompletedDateTime(value *DateTimeTimeZone)() {
    m.completedDateTime = value
}
func (m *FollowupFlag) SetDueDateTime(value *DateTimeTimeZone)() {
    m.dueDateTime = value
}
func (m *FollowupFlag) SetFlagStatus(value *FollowupFlagStatus)() {
    m.flagStatus = value
}
func (m *FollowupFlag) SetStartDateTime(value *DateTimeTimeZone)() {
    m.startDateTime = value
}
