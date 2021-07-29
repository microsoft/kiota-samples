package models

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type FollowupFlag struct {
    additionalData map[string]interface{};
    completedDateTime DateTimeTimeZone;
    dueDateTime DateTimeTimeZone;
    flagStatus FollowupFlagStatus;
    startDateTime DateTimeTimeZone;
}
func (m *FollowupFlag) getAdditionalData()(map[string]interface{}) {
    return nil
}
func (m *FollowupFlag) getCompletedDateTime()(DateTimeTimeZone) {
    return nil
}
func (m *FollowupFlag) getDueDateTime()(DateTimeTimeZone) {
    return nil
}
func (m *FollowupFlag) getFlagStatus()(FollowupFlagStatus) {
    return nil
}
func (m *FollowupFlag) getStartDateTime()(DateTimeTimeZone) {
    return nil
}
func (m *FollowupFlag) getFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error), error) {
    return nil
}
func (m *FollowupFlag) serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    return nil
}
func (m *FollowupFlag) setAdditionalData(value map[string]interface{})() {
    return nil
}
func (m *FollowupFlag) setCompletedDateTime(value DateTimeTimeZone)() {
    return nil
}
func (m *FollowupFlag) setDueDateTime(value DateTimeTimeZone)() {
    return nil
}
func (m *FollowupFlag) setFlagStatus(value FollowupFlagStatus)() {
    return nil
}
func (m *FollowupFlag) setStartDateTime(value DateTimeTimeZone)() {
    return nil
}
