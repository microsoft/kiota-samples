package content

import (
	url "net/url"

	ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

type ContentRequestBuilder struct {
	currentPath string
	httpCore    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.HttpCore
	pathSegment string
}

func NewContentRequestBuilder(currentPath string, httpCore ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.HttpCore) *ContentRequestBuilder {
	m := &ContentRequestBuilder{}
	m.pathSegment = "/$value"
	m.httpCore = httpCore
	m.currentPath = currentPath
	return m
}
func (m *ContentRequestBuilder) CreateGetRequestInfo(h func(value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.MiddlewareOption) (*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInfo, error) {
	requestInfo := new(ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInfo)
	uri, err := url.Parse(m.currentPath + m.pathSegment)
	requestInfo.URI = *uri
	requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
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
func (m *ContentRequestBuilder) CreatePutRequestInfo(body []byte, h func(value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.MiddlewareOption) (*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInfo, error) {
	requestInfo := new(ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInfo)
	uri, err := url.Parse(m.currentPath + m.pathSegment)
	requestInfo.URI = *uri
	requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PUT
	if err != nil {
		return nil, err
	}
	requestInfo.SetStreamContent(body)
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
func (m *ContentRequestBuilder) Get(h func(value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.MiddlewareOption, responseHandler *ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler) func() ([]byte, error) {
	requestInfo, err := m.CreateGetRequestInfo(h, o)
	if err != nil {
		return func() ([]byte, error) { return nil, err }
	}
	return func() ([]byte, error) {
		res, err := m.httpCore.SendPrimitiveAsync(*requestInfo, "i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable", *responseHandler)()
		if err != nil {
			return nil, err
		}
		return res.([]byte), nil
	}
}
func (m *ContentRequestBuilder) Put(body []byte, h func(value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.MiddlewareOption, responseHandler *ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler) func() error {
	requestInfo, err := m.CreatePutRequestInfo(body, h, o)
	if err != nil {
		return func() error { return err }
	}
	return func() error {
		err := m.httpCore.SendNoContentAsync(*requestInfo, *responseHandler)()
		if err != nil {
			return err
		}
		return nil
	}
}
