package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/models/microsoft/graph"
    i10075314d3b96cd3f0f2195562b0027395b41f6069ced9cc8d93624bbc206528 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/messages/item/extensions"
    i170c26461395ae79c8485bb09f6acf415c54c121851bca73450cfff498f75eb8 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/messages/item/singlevalueextendedproperties"
    i228962c2a18462bcfbde8557f6668f3fd7d4e68d5c7f3f1f3b675bc816866033 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/messages/item/attachments"
    i9550b442511c46b7a974613d5ee3238fb100cef30fff0520a9682fb3d38ad326 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/messages/item/value"
    iadaad8f92ba85254c1d639d8a8cb86def5e843af4bc1062bc836f08d062196d6 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/messages/item/multivalueextendedproperties"
    i361d73846d6331cbbe8151dfaef3455f34e63dfa07626b67576681683b148587 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/messages/item/extensions/item"
    i9f73052b0f535426cd146e9403e4e088995d624a331ea8cf19005e2c6ff46d7f "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/messages/item/singlevalueextendedproperties/item"
    idd4800ce8115a136eb3f42205c28105422c049d9e5b23796f3bdf18e2778f6d6 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/messages/item/attachments/item"
    if06dd322de6c7a1abe876d18d2ddc86a4b8bcf49f2fea78dcba12aff7dcb4994 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/messages/item/multivalueextendedproperties/item"
)

type MessageRequestBuilder struct {
    currentPath string;
    httpCore ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.HttpCore;
    isRawUrl bool;
    pathSegment string;
}
type MessageRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *MessageRequestBuilder) Attachments()(i228962c2a18462bcfbde8557f6668f3fd7d4e68d5c7f3f1f3b675bc816866033.AttachmentsRequestBuilder) {
    return *i228962c2a18462bcfbde8557f6668f3fd7d4e68d5c7f3f1f3b675bc816866033.NewAttachmentsRequestBuilder(m.currentPath + m.pathSegment, m.httpCore, false);
}
func (m *MessageRequestBuilder) AttachmentsById(id string)(idd4800ce8115a136eb3f42205c28105422c049d9e5b23796f3bdf18e2778f6d6.AttachmentRequestBuilder) {
    return *idd4800ce8115a136eb3f42205c28105422c049d9e5b23796f3bdf18e2778f6d6.NewAttachmentRequestBuilder(m.currentPath + m.pathSegment + "/attachments/" + id, m.httpCore, false);
}
func NewMessageRequestBuilder(currentPath string, httpCore ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.HttpCore, isRawUrl bool)(*MessageRequestBuilder) {
    m := &MessageRequestBuilder{
    }
    m.pathSegment = "";
    m.httpCore = httpCore;
    m.currentPath = currentPath;
    m.isRawUrl = isRawUrl;
    return m
}
func (m *MessageRequestBuilder) Content()(i9550b442511c46b7a974613d5ee3238fb100cef30fff0520a9682fb3d38ad326.ContentRequestBuilder) {
    return *i9550b442511c46b7a974613d5ee3238fb100cef30fff0520a9682fb3d38ad326.NewContentRequestBuilder(m.currentPath + m.pathSegment, m.httpCore, false);
}
func (m *MessageRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.MiddlewareOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    err := requestInfo.SetUri(m.currentPath, m.pathSegment, m.isRawUrl)
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if err != nil {
        return nil, err
    }
    if h != nil {
        err = h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err = requestInfo.AddMiddlewareOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, err
}
func (m *MessageRequestBuilder) CreateGetRequestInformation(q func (value *MessageRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.MiddlewareOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    err := requestInfo.SetUri(m.currentPath, m.pathSegment, m.isRawUrl)
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if err != nil {
        return nil, err
    }
    if q != nil {
        qParams := new(MessageRequestBuilderGetQueryParameters)
        err = q(qParams)
        if err != nil {
            return nil, err
        }
        err := qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err = h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err = requestInfo.AddMiddlewareOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, err
}
func (m *MessageRequestBuilder) CreatePatchRequestInformation(body *i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.Message, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.MiddlewareOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    err := requestInfo.SetUri(m.currentPath, m.pathSegment, m.isRawUrl)
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    if err != nil {
        return nil, err
    }
    requestInfo.SetContentFromParsable(m.httpCore, "application/json", body)
    if h != nil {
        err = h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err = requestInfo.AddMiddlewareOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, err
}
func (m *MessageRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.MiddlewareOption, responseHandler *ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(func() (error)) {
    requestInfo, err := m.CreateDeleteRequestInformation(h, o);
    if err != nil {
        return func() (error) { return err }
    }
    return func() (error) {
        err := m.httpCore.SendNoContentAsync(*requestInfo, *responseHandler)()
        if err != nil {
            return err
        }
        return nil
    }
}
func (m *MessageRequestBuilder) Extensions()(i10075314d3b96cd3f0f2195562b0027395b41f6069ced9cc8d93624bbc206528.ExtensionsRequestBuilder) {
    return *i10075314d3b96cd3f0f2195562b0027395b41f6069ced9cc8d93624bbc206528.NewExtensionsRequestBuilder(m.currentPath + m.pathSegment, m.httpCore, false);
}
func (m *MessageRequestBuilder) ExtensionsById(id string)(i361d73846d6331cbbe8151dfaef3455f34e63dfa07626b67576681683b148587.ExtensionRequestBuilder) {
    return *i361d73846d6331cbbe8151dfaef3455f34e63dfa07626b67576681683b148587.NewExtensionRequestBuilder(m.currentPath + m.pathSegment + "/extensions/" + id, m.httpCore, false);
}
func (m *MessageRequestBuilder) Get(q func (value *MessageRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.MiddlewareOption, responseHandler *ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(func() (*i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.Message, error)) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return func() (*i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.Message, error) { return nil, err }
    }
    return func() (*i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.Message, error) {
        res, err := m.httpCore.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return new(i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.Message) }, *responseHandler)()
        if err != nil {
            return nil, err
        }
        return res.(*i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.Message), nil
    }
}
func (m *MessageRequestBuilder) MultiValueExtendedProperties()(iadaad8f92ba85254c1d639d8a8cb86def5e843af4bc1062bc836f08d062196d6.MultiValueExtendedPropertiesRequestBuilder) {
    return *iadaad8f92ba85254c1d639d8a8cb86def5e843af4bc1062bc836f08d062196d6.NewMultiValueExtendedPropertiesRequestBuilder(m.currentPath + m.pathSegment, m.httpCore, false);
}
func (m *MessageRequestBuilder) MultiValueExtendedPropertiesById(id string)(if06dd322de6c7a1abe876d18d2ddc86a4b8bcf49f2fea78dcba12aff7dcb4994.MultiValueLegacyExtendedPropertyRequestBuilder) {
    return *if06dd322de6c7a1abe876d18d2ddc86a4b8bcf49f2fea78dcba12aff7dcb4994.NewMultiValueLegacyExtendedPropertyRequestBuilder(m.currentPath + m.pathSegment + "/multiValueExtendedProperties/" + id, m.httpCore, false);
}
func (m *MessageRequestBuilder) Patch(body *i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.Message, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.MiddlewareOption, responseHandler *ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(func() (error)) {
    requestInfo, err := m.CreatePatchRequestInformation(body, h, o);
    if err != nil {
        return func() (error) { return err }
    }
    return func() (error) {
        err := m.httpCore.SendNoContentAsync(*requestInfo, *responseHandler)()
        if err != nil {
            return err
        }
        return nil
    }
}
func (m *MessageRequestBuilder) SingleValueExtendedProperties()(i170c26461395ae79c8485bb09f6acf415c54c121851bca73450cfff498f75eb8.SingleValueExtendedPropertiesRequestBuilder) {
    return *i170c26461395ae79c8485bb09f6acf415c54c121851bca73450cfff498f75eb8.NewSingleValueExtendedPropertiesRequestBuilder(m.currentPath + m.pathSegment, m.httpCore, false);
}
func (m *MessageRequestBuilder) SingleValueExtendedPropertiesById(id string)(i9f73052b0f535426cd146e9403e4e088995d624a331ea8cf19005e2c6ff46d7f.SingleValueLegacyExtendedPropertyRequestBuilder) {
    return *i9f73052b0f535426cd146e9403e4e088995d624a331ea8cf19005e2c6ff46d7f.NewSingleValueLegacyExtendedPropertyRequestBuilder(m.currentPath + m.pathSegment + "/singleValueExtendedProperties/" + id, m.httpCore, false);
}
