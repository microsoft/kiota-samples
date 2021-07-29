package mailfolders

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/models"
)

type MailFoldersResponse struct {
    additionalData map[string]interface{};
    nextLink *string;
    value []ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MailFolder;
}
func NewMailFoldersResponse()(*MailFoldersResponse) {
    return nil
}
func (m *MailFoldersResponse) getAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
func (m *MailFoldersResponse) getNextLink()(*string) {
    return m.nextLink
}
func (m *MailFoldersResponse) getValue()([]ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MailFolder) {
    return m.value
}
func (m *MailFoldersResponse) getFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error), error) {
    return nil
}
func (m *MailFoldersResponse) serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    return nil
}
func (m *MailFoldersResponse) setAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *MailFoldersResponse) setNextLink(value *string)() {
    m.nextLink = value
}
func (m *MailFoldersResponse) setValue(value []ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MailFolder)() {
    m.value = value
}
