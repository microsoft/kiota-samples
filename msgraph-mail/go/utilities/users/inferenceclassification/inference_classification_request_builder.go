package inferenceclassification

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/models"
    id37525301f612e8f9f28658e2ceeb7e0dbca43868161cbacdbd261efc2c213df "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/inferenceclassification/overrides"
    ic84e53bbe0d89365c60cb7ce42953129b0a25b34d24d19255e83087a1152ecff "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/inferenceclassification/overrides/item"
)

type InferenceClassificationRequestBuilder struct {
    currentPath *string;
    httpCore ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.HttpCore;
    Overrides id37525301f612e8f9f28658e2ceeb7e0dbca43868161cbacdbd261efc2c213df.OverridesRequestBuilder;
    pathSegment *string;
}
func NewInferenceClassificationRequestBuilder(currentPath *string, httpCore ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.HttpCore)(*InferenceClassificationRequestBuilder) {
    return nil
}
func (m *InferenceClassificationRequestBuilder) createDeleteRequestInfo(h func (value map[string]string) (err error))(ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInfo, error) {
    return nil
}
func (m *InferenceClassificationRequestBuilder) createGetRequestInfo(q func (value GetQueryParameters) (err error), h func (value map[string]string) (err error))(ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInfo, error) {
    return nil
}
func (m *InferenceClassificationRequestBuilder) createPatchRequestInfo(body ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.InferenceClassification, h func (value map[string]string) (err error))(ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInfo, error) {
    return nil
}
func (m *InferenceClassificationRequestBuilder) delete(h func (value map[string]string) (err error), responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(func() (error)) {
    return nil
}
func (m *InferenceClassificationRequestBuilder) get(q func (value GetQueryParameters) (err error), h func (value map[string]string) (err error), responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(func() (ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.InferenceClassification, error)) {
    return nil
}
func (m *InferenceClassificationRequestBuilder) overrides(id *string)(ic84e53bbe0d89365c60cb7ce42953129b0a25b34d24d19255e83087a1152ecff.InferenceClassificationOverrideRequestBuilder) {
    return ic84e53bbe0d89365c60cb7ce42953129b0a25b34d24d19255e83087a1152ecff.InferenceClassificationOverrideRequestBuilder.NewInferenceClassificationOverrideRequestBuilder(m.currentPath + m.pathSegment + "/overrides/" + id, m.httpCore);
}
func (m *InferenceClassificationRequestBuilder) patch(body ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.InferenceClassification, h func (value map[string]string) (err error), responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(func() (error)) {
    return nil
}
