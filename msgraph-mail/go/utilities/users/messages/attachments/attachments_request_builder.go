package attachments

import (
    url "net/url"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/models"
)

type AttachmentsRequestBuilder struct {
    currentPath *string;
    httpCore ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.HttpCore;
    pathSegment *string;
}
type AttachmentsRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Count *bool;
    Expand []*string;
    Filter *string;
    Orderby []*string;
    Search *string;
    Select []*string;
    Skip *int32;
    Top *int32;
}
func NewAttachmentsRequestBuilder(currentPath *string, httpCore ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.HttpCore)(*AttachmentsRequestBuilder) {
    return nil
}
func (m *AttachmentsRequestBuilder) createGetRequestInfo(q func (value *AttachmentsRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.MiddlewareOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInfo, error) {
    requestInfo := new(ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInfo)
    uri, err := url.Parse(*m.currentPath + *m.pathSegment)
    requestInfo.URI = *uri
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if err != nil {
        return nil, err
    }
    if q != nil {
        qParams := new(AttachmentsRequestBuilderGetQueryParameters)
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
func (m *AttachmentsRequestBuilder) createPostRequestInfo(body *ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.Attachment, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.MiddlewareOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInfo, error) {
    requestInfo := new(ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInfo)
    uri, err := url.Parse(*m.currentPath + *m.pathSegment)
    requestInfo.URI = *uri
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
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
func (m *AttachmentsRequestBuilder) get(q func (value *AttachmentsRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.MiddlewareOption, responseHandler *ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(func() (*AttachmentsResponse, error)) {
    return nil
}
func (m *AttachmentsRequestBuilder) post(body *ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.Attachment, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.MiddlewareOption, responseHandler *ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(func() (*ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.Attachment, error)) {
    return nil
}
