package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/models/microsoft/graph"
)

type MessageRuleRequestBuilder struct {
    currentPath string;
    isRawUrl bool;
    pathSegment string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
}
type MessageRuleRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func NewMessageRuleRequestBuilder(currentPath string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter, isRawUrl bool)(*MessageRuleRequestBuilder) {
    m := &MessageRuleRequestBuilder{
    }
    m.pathSegment = "";
    m.requestAdapter = requestAdapter;
    m.currentPath = currentPath;
    m.isRawUrl = isRawUrl;
    return m
}
func (m *MessageRuleRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
        err = requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, err
}
func (m *MessageRuleRequestBuilder) CreateGetRequestInformation(q func (value *MessageRuleRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    err := requestInfo.SetUri(m.currentPath, m.pathSegment, m.isRawUrl)
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if err != nil {
        return nil, err
    }
    if q != nil {
        qParams := new(MessageRuleRequestBuilderGetQueryParameters)
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
        err = requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, err
}
func (m *MessageRuleRequestBuilder) CreatePatchRequestInformation(body *i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.MessageRule, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    err := requestInfo.SetUri(m.currentPath, m.pathSegment, m.isRawUrl)
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    if err != nil {
        return nil, err
    }
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err = h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err = requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, err
}
func (m *MessageRuleRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler *ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(func() (error)) {
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
func (m *MessageRuleRequestBuilder) Get(q func (value *MessageRuleRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler *ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(func() (*i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.MessageRule, error)) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return func() (*i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.MessageRule, error) { return nil, err }
    }
    return func() (*i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.MessageRule, error) {
        res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return new(i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.MessageRule) }, *responseHandler)()
        if err != nil {
            return nil, err
        }
        return res.(*i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.MessageRule), nil
    }
}
func (m *MessageRuleRequestBuilder) Patch(body *i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.MessageRule, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler *ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(func() (error)) {
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
