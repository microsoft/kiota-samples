package item

import (
    url "net/url"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/models"
    i1e8f5e1b0ab12ae689692569cf00409a7ffcf8119e0c731d1aab8fd8c24b9c9a "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailfolders/messages/extensions"
    i3a8d456234d570904652042e0d4c0c5c4510aca0788f80c075c5daa8f534b477 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailfolders/messages/attachments"
    i5f4f3512bc76131a2bfb807e4ef99b24c83995d2d107aba0fd68cf3b07f8b750 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailfolders/messages/singlevalueextendedproperties"
    iaeab5810035e3d7ff3e4901893330c349ce3db065ba7b9e0dcaadaf7479be003 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailfolders/messages/content"
    ic1605e2a56328ce22ef4db9ad89f54800a8f837cf7ba24dd304d6ddd4f629093 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailfolders/messages/multivalueextendedproperties"
    i99bbed137b62cf613096d6395b85791513d9e404f617b07c9675049bc9df09f8 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailfolders/messages/multivalueextendedproperties/item"
    iace81820e7196eb3ec506adf50a909fd751a256f8329891ef2cbbdd1d8fe9b2d "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailfolders/messages/attachments/item"
    ibce1fc49ff1dfb26493978b2a56f0be047bae8336ff396cb9deb7e0715c92eca "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailfolders/messages/extensions/item"
    idf8de8c43ef15f0fdc414d98d3e0cf730d22735ef1baafa82a410f3c5cbf2e90 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailfolders/messages/singlevalueextendedproperties/item"
)

type MessageRequestBuilder struct {
    currentPath string;
    httpCore ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.HttpCore;
    pathSegment string;
}
type MessageRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select []string;
}
func (m *MessageRequestBuilder) Attachments()(i3a8d456234d570904652042e0d4c0c5c4510aca0788f80c075c5daa8f534b477.AttachmentsRequestBuilder) {
    return *i3a8d456234d570904652042e0d4c0c5c4510aca0788f80c075c5daa8f534b477.NewAttachmentsRequestBuilder(m.currentPath + m.pathSegment, m.httpCore);
}
func (m *MessageRequestBuilder) AttachmentsById(id string)(iace81820e7196eb3ec506adf50a909fd751a256f8329891ef2cbbdd1d8fe9b2d.AttachmentRequestBuilder) {
    return *iace81820e7196eb3ec506adf50a909fd751a256f8329891ef2cbbdd1d8fe9b2d.NewAttachmentRequestBuilder(m.currentPath + m.pathSegment + "/attachments/" + id, m.httpCore);
}
func NewMessageRequestBuilder(currentPath string, httpCore ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.HttpCore)(*MessageRequestBuilder) {
    m := &MessageRequestBuilder{
    }
    m.pathSegment = "";
    m.httpCore = httpCore;
    m.currentPath = currentPath;
    return m
}
func (m *MessageRequestBuilder) Content()(iaeab5810035e3d7ff3e4901893330c349ce3db065ba7b9e0dcaadaf7479be003.ContentRequestBuilder) {
    return *iaeab5810035e3d7ff3e4901893330c349ce3db065ba7b9e0dcaadaf7479be003.NewContentRequestBuilder(m.currentPath + m.pathSegment, m.httpCore);
}
func (m *MessageRequestBuilder) CreateDeleteRequestInfo(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.MiddlewareOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInfo, error) {
    requestInfo := new(ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInfo)
    uri, err := url.Parse(m.currentPath + m.pathSegment)
    requestInfo.URI = *uri
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
    return requestInfo, err
}
func (m *MessageRequestBuilder) CreateGetRequestInfo(q func (value *MessageRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.MiddlewareOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInfo, error) {
    requestInfo := new(ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInfo)
    uri, err := url.Parse(m.currentPath + m.pathSegment)
    requestInfo.URI = *uri
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
    return requestInfo, err
}
func (m *MessageRequestBuilder) CreatePatchRequestInfo(body *ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.Message, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.MiddlewareOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInfo, error) {
    requestInfo := new(ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInfo)
    uri, err := url.Parse(m.currentPath + m.pathSegment)
    requestInfo.URI = *uri
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    if err != nil {
        return nil, err
    }
    requestInfo.SetContentFromParsable(body, m.httpCore, "application/json")
    if h != nil {
        err = h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, err
}
func (m *MessageRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.MiddlewareOption, responseHandler *ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(func() (error)) {
    requestInfo, err := m.CreateDeleteRequestInfo(h, o);
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
func (m *MessageRequestBuilder) Extensions()(i1e8f5e1b0ab12ae689692569cf00409a7ffcf8119e0c731d1aab8fd8c24b9c9a.ExtensionsRequestBuilder) {
    return *i1e8f5e1b0ab12ae689692569cf00409a7ffcf8119e0c731d1aab8fd8c24b9c9a.NewExtensionsRequestBuilder(m.currentPath + m.pathSegment, m.httpCore);
}
func (m *MessageRequestBuilder) ExtensionsById(id string)(ibce1fc49ff1dfb26493978b2a56f0be047bae8336ff396cb9deb7e0715c92eca.ExtensionRequestBuilder) {
    return *ibce1fc49ff1dfb26493978b2a56f0be047bae8336ff396cb9deb7e0715c92eca.NewExtensionRequestBuilder(m.currentPath + m.pathSegment + "/extensions/" + id, m.httpCore);
}
func (m *MessageRequestBuilder) Get(q func (value *MessageRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.MiddlewareOption, responseHandler *ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(func() (*ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.Message, error)) {
    requestInfo, err := m.CreateGetRequestInfo(q, h, o);
    if err != nil {
        return func() (*ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.Message, error) { return nil, err }
    }
    return func() (*ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.Message, error) {
        res, err := m.httpCore.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return new(ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.Message) }, *responseHandler)()
        if err != nil {
            return nil, err
        }
        return res.(*ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.Message), nil
    }
}
func (m *MessageRequestBuilder) MultiValueExtendedProperties()(ic1605e2a56328ce22ef4db9ad89f54800a8f837cf7ba24dd304d6ddd4f629093.MultiValueExtendedPropertiesRequestBuilder) {
    return *ic1605e2a56328ce22ef4db9ad89f54800a8f837cf7ba24dd304d6ddd4f629093.NewMultiValueExtendedPropertiesRequestBuilder(m.currentPath + m.pathSegment, m.httpCore);
}
func (m *MessageRequestBuilder) MultiValueExtendedPropertiesById(id string)(i99bbed137b62cf613096d6395b85791513d9e404f617b07c9675049bc9df09f8.MultiValueLegacyExtendedPropertyRequestBuilder) {
    return *i99bbed137b62cf613096d6395b85791513d9e404f617b07c9675049bc9df09f8.NewMultiValueLegacyExtendedPropertyRequestBuilder(m.currentPath + m.pathSegment + "/multiValueExtendedProperties/" + id, m.httpCore);
}
func (m *MessageRequestBuilder) Patch(body *ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.Message, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.MiddlewareOption, responseHandler *ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(func() (error)) {
    requestInfo, err := m.CreatePatchRequestInfo(body, h, o);
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
func (m *MessageRequestBuilder) SingleValueExtendedProperties()(i5f4f3512bc76131a2bfb807e4ef99b24c83995d2d107aba0fd68cf3b07f8b750.SingleValueExtendedPropertiesRequestBuilder) {
    return *i5f4f3512bc76131a2bfb807e4ef99b24c83995d2d107aba0fd68cf3b07f8b750.NewSingleValueExtendedPropertiesRequestBuilder(m.currentPath + m.pathSegment, m.httpCore);
}
func (m *MessageRequestBuilder) SingleValueExtendedPropertiesById(id string)(idf8de8c43ef15f0fdc414d98d3e0cf730d22735ef1baafa82a410f3c5cbf2e90.SingleValueLegacyExtendedPropertyRequestBuilder) {
    return *idf8de8c43ef15f0fdc414d98d3e0cf730d22735ef1baafa82a410f3c5cbf2e90.NewSingleValueLegacyExtendedPropertyRequestBuilder(m.currentPath + m.pathSegment + "/singleValueExtendedProperties/" + id, m.httpCore);
}
