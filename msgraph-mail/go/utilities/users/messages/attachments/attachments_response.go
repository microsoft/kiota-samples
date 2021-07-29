package attachments

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/models"
)

type AttachmentsResponse struct {
    additionalData map[string]interface{};
    nextLink *string;
    value []ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.Attachment;
}
func (m *AttachmentsResponse) getAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
func (m *AttachmentsResponse) getNextLink()(*string) {
    return m.nextLink
}
func (m *AttachmentsResponse) getValue()([]ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.Attachment) {
    return m.value
}
func (m *AttachmentsResponse) getFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error), error) {
    return nil
}
func (m *AttachmentsResponse) serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    return nil
}
func (m *AttachmentsResponse) setAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AttachmentsResponse) setNextLink(value *string)() {
    m.nextLink = value
}
func (m *AttachmentsResponse) setValue(value []ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.Attachment)() {
    m.value = value
}
