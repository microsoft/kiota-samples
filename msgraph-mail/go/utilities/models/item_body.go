package models

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ItemBody struct {
    additionalData map[string]interface{};
    content *string;
    contentType *BodyType;
}
func NewItemBody()(*ItemBody) {
    return nil
}
func (m *ItemBody) getAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
func (m *ItemBody) getContent()(*string) {
    return m.content
}
func (m *ItemBody) getContentType()(*BodyType) {
    return m.contentType
}
func (m *ItemBody) getFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error), error) {
    return nil, nil
}
func (m *ItemBody) serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    return nil
}
func (m *ItemBody) setAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ItemBody) setContent(value *string)() {
    m.content = value
}
func (m *ItemBody) setContentType(value *BodyType)() {
    m.contentType = value
}
