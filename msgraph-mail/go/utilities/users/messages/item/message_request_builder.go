package item

import (
    url "net/url"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/models"
    i0686696817f5336ad23f5048aaf1973675e26317e928255ef44ff8f6cdd744ce "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/messages/attachments"
    i514af90a8d1241ee2aec5dae95b1b4028641b8511c05b3b7710d33898668b1c0 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/messages/singlevalueextendedproperties"
    i7a8f25a40ec18f53dd1fe13983249e861ca16f093020bf1d1f77bdea2b68285f "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/messages/content"
    ib2381cda6aa2598b763d17161ebaa5c941e140eebfa25805cfbb21dbd2ff0805 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/messages/extensions"
    id35afca47a5e5a75ec80ed3fe0e8a5bddf19271c23cfbd9f6678773404139756 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/messages/multivalueextendedproperties"
    i4c213381c08f42149bb0acad01e026154f1d08ef5829ececbd8ea51281fe708b "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/messages/extensions/item"
    i4dc594f64060f7179afa645e9dc31eac49b9bb535bf1df6fe899682087dc226c "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/messages/singlevalueextendedproperties/item"
    i54bf43874c83bde999c0160122d8c20bde3e0dfabeec784a4ae151188b4686c4 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/messages/multivalueextendedproperties/item"
    ib74bdf213c8ecff00ba896023527beb4a5cee0f1561b4ebb235535adb8fd01ba "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/messages/attachments/item"
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
func (m *MessageRequestBuilder) Attachments()(i0686696817f5336ad23f5048aaf1973675e26317e928255ef44ff8f6cdd744ce.AttachmentsRequestBuilder) {
    return *i0686696817f5336ad23f5048aaf1973675e26317e928255ef44ff8f6cdd744ce.NewAttachmentsRequestBuilder(m.currentPath + m.pathSegment, m.httpCore);
}
func (m *MessageRequestBuilder) AttachmentsById(id string)(ib74bdf213c8ecff00ba896023527beb4a5cee0f1561b4ebb235535adb8fd01ba.AttachmentRequestBuilder) {
    return *ib74bdf213c8ecff00ba896023527beb4a5cee0f1561b4ebb235535adb8fd01ba.NewAttachmentRequestBuilder(m.currentPath + m.pathSegment + "/attachments/" + id, m.httpCore);
}
func NewMessageRequestBuilder(currentPath string, httpCore ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.HttpCore)(*MessageRequestBuilder) {
    m := &MessageRequestBuilder{
    }
    m.pathSegment = "";
    m.httpCore = httpCore;
    m.currentPath = currentPath;
    return m
}
func (m *MessageRequestBuilder) Content()(i7a8f25a40ec18f53dd1fe13983249e861ca16f093020bf1d1f77bdea2b68285f.ContentRequestBuilder) {
    return *i7a8f25a40ec18f53dd1fe13983249e861ca16f093020bf1d1f77bdea2b68285f.NewContentRequestBuilder(m.currentPath + m.pathSegment, m.httpCore);
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
    if o != nil {
        err = requestInfo.AddMiddlewareOptions(o)
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
    if o != nil {
        err = requestInfo.AddMiddlewareOptions(o)
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
func (m *MessageRequestBuilder) Extensions()(ib2381cda6aa2598b763d17161ebaa5c941e140eebfa25805cfbb21dbd2ff0805.ExtensionsRequestBuilder) {
    return *ib2381cda6aa2598b763d17161ebaa5c941e140eebfa25805cfbb21dbd2ff0805.NewExtensionsRequestBuilder(m.currentPath + m.pathSegment, m.httpCore);
}
func (m *MessageRequestBuilder) ExtensionsById(id string)(i4c213381c08f42149bb0acad01e026154f1d08ef5829ececbd8ea51281fe708b.ExtensionRequestBuilder) {
    return *i4c213381c08f42149bb0acad01e026154f1d08ef5829ececbd8ea51281fe708b.NewExtensionRequestBuilder(m.currentPath + m.pathSegment + "/extensions/" + id, m.httpCore);
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
func (m *MessageRequestBuilder) MultiValueExtendedProperties()(id35afca47a5e5a75ec80ed3fe0e8a5bddf19271c23cfbd9f6678773404139756.MultiValueExtendedPropertiesRequestBuilder) {
    return *id35afca47a5e5a75ec80ed3fe0e8a5bddf19271c23cfbd9f6678773404139756.NewMultiValueExtendedPropertiesRequestBuilder(m.currentPath + m.pathSegment, m.httpCore);
}
func (m *MessageRequestBuilder) MultiValueExtendedPropertiesById(id string)(i54bf43874c83bde999c0160122d8c20bde3e0dfabeec784a4ae151188b4686c4.MultiValueLegacyExtendedPropertyRequestBuilder) {
    return *i54bf43874c83bde999c0160122d8c20bde3e0dfabeec784a4ae151188b4686c4.NewMultiValueLegacyExtendedPropertyRequestBuilder(m.currentPath + m.pathSegment + "/multiValueExtendedProperties/" + id, m.httpCore);
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
func (m *MessageRequestBuilder) SingleValueExtendedProperties()(i514af90a8d1241ee2aec5dae95b1b4028641b8511c05b3b7710d33898668b1c0.SingleValueExtendedPropertiesRequestBuilder) {
    return *i514af90a8d1241ee2aec5dae95b1b4028641b8511c05b3b7710d33898668b1c0.NewSingleValueExtendedPropertiesRequestBuilder(m.currentPath + m.pathSegment, m.httpCore);
}
func (m *MessageRequestBuilder) SingleValueExtendedPropertiesById(id string)(i4dc594f64060f7179afa645e9dc31eac49b9bb535bf1df6fe899682087dc226c.SingleValueLegacyExtendedPropertyRequestBuilder) {
    return *i4dc594f64060f7179afa645e9dc31eac49b9bb535bf1df6fe899682087dc226c.NewSingleValueLegacyExtendedPropertyRequestBuilder(m.currentPath + m.pathSegment + "/singleValueExtendedProperties/" + id, m.httpCore);
}
