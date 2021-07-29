package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/models"
)

type ExtensionRequestBuilder struct {
    currentPath *string;
    httpCore ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.HttpCore;
    pathSegment *string;
}
func NewExtensionRequestBuilder(currentPath *string, httpCore ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.HttpCore)(*ExtensionRequestBuilder) {
    return nil
}
func (m *ExtensionRequestBuilder) createDeleteRequestInfo(h func (value map[string]string) (err error))(ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInfo, error) {
    requestInfo := new(ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInfo)
    requestInfo.URI = new URI(m.currentPath + m.pathSegment);
    requestInfo.httpMethod = HttpMethod.DELETE;
    if h != nil {
        h.accept(requestInfo.headers);
    }
    return requestInfo
}
func (m *ExtensionRequestBuilder) createGetRequestInfo(q func (value GetQueryParameters) (err error), h func (value map[string]string) (err error))(ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInfo, error) {
    requestInfo := new(ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInfo)
    requestInfo.URI = new URI(m.currentPath + m.pathSegment);
    requestInfo.httpMethod = HttpMethod.GET;
    if q != nil {
        qParams := new GetQueryParameters();
        q.accept(qParams);
        qParams.AddQueryParameters(requestInfo.queryParameters);
    }
    if h != nil {
        h.accept(requestInfo.headers);
    }
    return requestInfo
}
func (m *ExtensionRequestBuilder) createPatchRequestInfo(body ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.Extension, h func (value map[string]string) (err error))(ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInfo, error) {
    requestInfo := new(ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInfo)
    requestInfo.URI = new URI(m.currentPath + m.pathSegment);
    requestInfo.httpMethod = HttpMethod.PATCH;
    requestInfo.SetContentFromParsable(body, m.httpCore, "application/json");
    if h != nil {
        h.accept(requestInfo.headers);
    }
    return requestInfo
}
func (m *ExtensionRequestBuilder) delete(h func (value map[string]string) (err error), responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(func() (error)) {
    return nil
}
func (m *ExtensionRequestBuilder) get(q func (value GetQueryParameters) (err error), h func (value map[string]string) (err error), responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(func() (ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.Extension, error)) {
    return nil
}
func (m *ExtensionRequestBuilder) patch(body ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.Extension, h func (value map[string]string) (err error), responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(func() (error)) {
    return nil
}
