package childfolders

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/models"
)

type ChildFoldersResponse struct {
    additionalData map[string]interface{};
    nextLink *string;
    value []ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MailFolder;
}
func NewChildFoldersResponse()(*ChildFoldersResponse) {
    return nil
}
func (m *ChildFoldersResponse) getAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
func (m *ChildFoldersResponse) getNextLink()(*string) {
    return m.nextLink
}
func (m *ChildFoldersResponse) getValue()([]ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MailFolder) {
    return m.value
}
func (m *ChildFoldersResponse) getFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error), error) {
    return nil, nil
}
func (m *ChildFoldersResponse) serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    return nil
}
func (m *ChildFoldersResponse) setAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ChildFoldersResponse) setNextLink(value *string)() {
    m.nextLink = value
}
func (m *ChildFoldersResponse) setValue(value []ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MailFolder)() {
    m.value = value
}
