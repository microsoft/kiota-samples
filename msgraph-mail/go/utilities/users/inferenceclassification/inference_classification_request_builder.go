package inferenceclassification

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/models"
    id37525301f612e8f9f28658e2ceeb7e0dbca43868161cbacdbd261efc2c213df "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/inferenceclassification/overrides"
    ic84e53bbe0d89365c60cb7ce42953129b0a25b34d24d19255e83087a1152ecff "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/inferenceclassification/overrides/item"
)

type InferenceClassificationRequestBuilder struct {
    currentPath string;
    httpCore ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.HttpCore;
    isRawUrl bool;
    pathSegment string;
}
type InferenceClassificationRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select []string;
}
func NewInferenceClassificationRequestBuilder(currentPath string, httpCore ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.HttpCore, isRawUrl bool)(*InferenceClassificationRequestBuilder) {
    m := &InferenceClassificationRequestBuilder{
    }
    m.pathSegment = "/inferenceClassification";
    m.httpCore = httpCore;
    m.currentPath = currentPath;
    m.isRawUrl = isRawUrl;
    return m
}
func (m *InferenceClassificationRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.MiddlewareOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *InferenceClassificationRequestBuilder) CreateGetRequestInformation(q func (value *InferenceClassificationRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.MiddlewareOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    err := requestInfo.SetUri(m.currentPath, m.pathSegment, m.isRawUrl)
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if err != nil {
        return nil, err
    }
    if q != nil {
        qParams := new(InferenceClassificationRequestBuilderGetQueryParameters)
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
func (m *InferenceClassificationRequestBuilder) CreatePatchRequestInformation(body *ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.InferenceClassification, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.MiddlewareOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *InferenceClassificationRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.MiddlewareOption, responseHandler *ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(func() (error)) {
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
func (m *InferenceClassificationRequestBuilder) Get(q func (value *InferenceClassificationRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.MiddlewareOption, responseHandler *ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(func() (*ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.InferenceClassification, error)) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return func() (*ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.InferenceClassification, error) { return nil, err }
    }
    return func() (*ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.InferenceClassification, error) {
        res, err := m.httpCore.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return new(ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.InferenceClassification) }, *responseHandler)()
        if err != nil {
            return nil, err
        }
        return res.(*ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.InferenceClassification), nil
    }
}
func (m *InferenceClassificationRequestBuilder) Overrides()(id37525301f612e8f9f28658e2ceeb7e0dbca43868161cbacdbd261efc2c213df.OverridesRequestBuilder) {
    return *id37525301f612e8f9f28658e2ceeb7e0dbca43868161cbacdbd261efc2c213df.NewOverridesRequestBuilder(m.currentPath + m.pathSegment, m.httpCore, false);
}
func (m *InferenceClassificationRequestBuilder) OverridesById(id string)(ic84e53bbe0d89365c60cb7ce42953129b0a25b34d24d19255e83087a1152ecff.InferenceClassificationOverrideRequestBuilder) {
    return *ic84e53bbe0d89365c60cb7ce42953129b0a25b34d24d19255e83087a1152ecff.NewInferenceClassificationOverrideRequestBuilder(m.currentPath + m.pathSegment + "/overrides/" + id, m.httpCore, false);
}
func (m *InferenceClassificationRequestBuilder) Patch(body *ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.InferenceClassification, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.MiddlewareOption, responseHandler *ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(func() (error)) {
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
