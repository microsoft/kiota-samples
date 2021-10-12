package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/models/microsoft/graph"
)

type MultiValueLegacyExtendedPropertyRequestBuilder struct {
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
    urlTemplateParameters map[string]string;
}
type MultiValueLegacyExtendedPropertyRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTemplateParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MultiValueLegacyExtendedPropertyRequestBuilder) {
    m := &MultiValueLegacyExtendedPropertyRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/users/{user_id}/mailFolders/{mailFolder_id}/messages/{message_id}/multiValueExtendedProperties/{multiValueLegacyExtendedProperty_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range urlTemplateParameters {
        urlTplParams[idx] = item
    }
    m.urlTemplateParameters = urlTemplateParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewMultiValueLegacyExtendedPropertyRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["raw-request-url"] = rawUrl
    return NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *MultiValueLegacyExtendedPropertyRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.UrlTemplateParameters = m.urlTemplateParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *MultiValueLegacyExtendedPropertyRequestBuilder) CreateGetRequestInformation(q func (value *MultiValueLegacyExtendedPropertyRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.UrlTemplateParameters = m.urlTemplateParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(MultiValueLegacyExtendedPropertyRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *MultiValueLegacyExtendedPropertyRequestBuilder) CreatePatchRequestInformation(body *i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.MultiValueLegacyExtendedProperty, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.UrlTemplateParameters = m.urlTemplateParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *MultiValueLegacyExtendedPropertyRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler *ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(func() (error)) {
    requestInfo, err := m.CreateDeleteRequestInformation(h, o);
    if err != nil {
        return func() (error) { return err }
    }
    return func() (error) {
        err := m.requestAdapter.SendNoContentAsync(*requestInfo, *responseHandler)()
        if err != nil {
            return err
        }
        return nil
    }
}
func (m *MultiValueLegacyExtendedPropertyRequestBuilder) Get(q func (value *MultiValueLegacyExtendedPropertyRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler *ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(func() (*i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.MultiValueLegacyExtendedProperty, error)) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return func() (*i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.MultiValueLegacyExtendedProperty, error) { return nil, err }
    }
    return func() (*i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.MultiValueLegacyExtendedProperty, error) {
        res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return new(i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.MultiValueLegacyExtendedProperty) }, *responseHandler)()
        if err != nil {
            return nil, err
        }
        return res.(*i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.MultiValueLegacyExtendedProperty), nil
    }
}
func (m *MultiValueLegacyExtendedPropertyRequestBuilder) Patch(body *i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.MultiValueLegacyExtendedProperty, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler *ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(func() (error)) {
    requestInfo, err := m.CreatePatchRequestInformation(body, h, o);
    if err != nil {
        return func() (error) { return err }
    }
    return func() (error) {
        err := m.requestAdapter.SendNoContentAsync(*requestInfo, *responseHandler)()
        if err != nil {
            return err
        }
        return nil
    }
}
