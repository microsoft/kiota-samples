package multivalueextendedproperties

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/models"
)

type MultiValueExtendedPropertiesRequestBuilder struct {
    currentPath string;
    httpCore ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.HttpCore;
    isRawUrl bool;
    pathSegment string;
}
type MultiValueExtendedPropertiesRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Count *bool;
    Expand []string;
    Filter *string;
    Orderby []string;
    Search *string;
    Select []string;
    Skip *int32;
    Top *int32;
}
func NewMultiValueExtendedPropertiesRequestBuilder(currentPath string, httpCore ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.HttpCore, isRawUrl bool)(*MultiValueExtendedPropertiesRequestBuilder) {
    m := &MultiValueExtendedPropertiesRequestBuilder{
    }
    m.pathSegment = "/multiValueExtendedProperties";
    m.httpCore = httpCore;
    m.currentPath = currentPath;
    m.isRawUrl = isRawUrl;
    return m
}
func (m *MultiValueExtendedPropertiesRequestBuilder) CreateGetRequestInfo(q func (value *MultiValueExtendedPropertiesRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.MiddlewareOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInfo, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInfo()
    err := requestInfo.SetUri(m.currentPath, m.pathSegment, m.isRawUrl)
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if err != nil {
        return nil, err
    }
    if q != nil {
        qParams := new(MultiValueExtendedPropertiesRequestBuilderGetQueryParameters)
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
func (m *MultiValueExtendedPropertiesRequestBuilder) CreatePostRequestInfo(body *ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MultiValueLegacyExtendedProperty, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.MiddlewareOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInfo, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInfo()
    err := requestInfo.SetUri(m.currentPath, m.pathSegment, m.isRawUrl)
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
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
func (m *MultiValueExtendedPropertiesRequestBuilder) Get(q func (value *MultiValueExtendedPropertiesRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.MiddlewareOption, responseHandler *ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(func() (*MultiValueExtendedPropertiesResponse, error)) {
    requestInfo, err := m.CreateGetRequestInfo(q, h, o);
    if err != nil {
        return func() (*MultiValueExtendedPropertiesResponse, error) { return nil, err }
    }
    return func() (*MultiValueExtendedPropertiesResponse, error) {
        res, err := m.httpCore.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return new(MultiValueExtendedPropertiesResponse) }, *responseHandler)()
        if err != nil {
            return nil, err
        }
        return res.(*MultiValueExtendedPropertiesResponse), nil
    }
}
func (m *MultiValueExtendedPropertiesRequestBuilder) Post(body *ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MultiValueLegacyExtendedProperty, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.MiddlewareOption, responseHandler *ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(func() (*ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MultiValueLegacyExtendedProperty, error)) {
    requestInfo, err := m.CreatePostRequestInfo(body, h, o);
    if err != nil {
        return func() (*ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MultiValueLegacyExtendedProperty, error) { return nil, err }
    }
    return func() (*ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MultiValueLegacyExtendedProperty, error) {
        res, err := m.httpCore.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return new(ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MultiValueLegacyExtendedProperty) }, *responseHandler)()
        if err != nil {
            return nil, err
        }
        return res.(*ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MultiValueLegacyExtendedProperty), nil
    }
}
