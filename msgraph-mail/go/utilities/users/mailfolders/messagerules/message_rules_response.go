package messagerules

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/models"
)

type MessageRulesResponse struct {
    additionalData map[string]interface{};
    nextLink *string;
    value []ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MessageRule;
}
func NewMessageRulesResponse()(*MessageRulesResponse) {
    return nil
}
func (m *MessageRulesResponse) getAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
func (m *MessageRulesResponse) getNextLink()(*string) {
    return m.nextLink
}
func (m *MessageRulesResponse) getValue()([]ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MessageRule) {
    return m.value
}
func (m *MessageRulesResponse) getFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error), error) {
    return nil
}
func (m *MessageRulesResponse) serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    return nil
}
func (m *MessageRulesResponse) setAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *MessageRulesResponse) setNextLink(value *string)() {
    m.nextLink = value
}
func (m *MessageRulesResponse) setValue(value []ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MessageRule)() {
    m.value = value
}
