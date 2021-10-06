package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/models/microsoft/graph"
    i000de88af33510156072c580aff36ddc934a6dbf385120aee0662cecf92dfee6 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/mailfolders/item/messages/item/value"
    i165e60b62d06f92d5e6029daee4ae129b1c3d1780dce2631cc18fbe8c6d39c9d "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/mailfolders/item/messages/item/multivalueextendedproperties"
    i281b692d3f673c995c16b5cc69287c818497fc74562aed972acc168040629f74 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/mailfolders/item/messages/item/extensions"
    i833cf7d5250828d0b2288e87035b59c2900b4aef7c2eb2ff5494a38edea8268e "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/mailfolders/item/messages/item/singlevalueextendedproperties"
    i9b8e8e1d8c7f8c0fa45b5abe65e2541be081536bd338e75409f4b0ecba2188aa "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/mailfolders/item/messages/item/attachments"
    i314a2a68928bac08cd3678f4259eab0e5845a67017fe63526811793a0c7a4615 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/mailfolders/item/messages/item/multivalueextendedproperties/item"
    i92b891944e73393404194faa96a7587068f07492bd44347d827c18456b8a2bcf "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/mailfolders/item/messages/item/attachments/item"
    ia3b46812df8e43b8356e8bad54f90359cc966c930f9415ac7f61670490ce857f "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/mailfolders/item/messages/item/singlevalueextendedproperties/item"
    icb423aed2922bb0caf9a0353a5ec5bfeb474588dbd1218008894e86ab49c6e25 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/mailfolders/item/messages/item/extensions/item"
)

type MessageRequestBuilder struct {
    currentPath string;
    isRawUrl bool;
    pathSegment string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
}
type MessageRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *MessageRequestBuilder) Attachments()(i9b8e8e1d8c7f8c0fa45b5abe65e2541be081536bd338e75409f4b0ecba2188aa.AttachmentsRequestBuilder) {
    return *i9b8e8e1d8c7f8c0fa45b5abe65e2541be081536bd338e75409f4b0ecba2188aa.NewAttachmentsRequestBuilder(m.currentPath + m.pathSegment, m.requestAdapter, false);
}
func (m *MessageRequestBuilder) AttachmentsById(id string)(i92b891944e73393404194faa96a7587068f07492bd44347d827c18456b8a2bcf.AttachmentRequestBuilder) {
    return *i92b891944e73393404194faa96a7587068f07492bd44347d827c18456b8a2bcf.NewAttachmentRequestBuilder(m.currentPath + m.pathSegment + "/attachments/" + id, m.requestAdapter, false);
}
func NewMessageRequestBuilder(currentPath string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter, isRawUrl bool)(*MessageRequestBuilder) {
    m := &MessageRequestBuilder{
    }
    m.pathSegment = "";
    m.requestAdapter = requestAdapter;
    m.currentPath = currentPath;
    m.isRawUrl = isRawUrl;
    return m
}
func (m *MessageRequestBuilder) Content()(i000de88af33510156072c580aff36ddc934a6dbf385120aee0662cecf92dfee6.ContentRequestBuilder) {
    return *i000de88af33510156072c580aff36ddc934a6dbf385120aee0662cecf92dfee6.NewContentRequestBuilder(m.currentPath + m.pathSegment, m.requestAdapter, false);
}
func (m *MessageRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *MessageRequestBuilder) CreateGetRequestInformation(q func (value *MessageRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    err := requestInfo.SetUri(m.currentPath, m.pathSegment, m.isRawUrl)
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
        err = requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, err
}
func (m *MessageRequestBuilder) CreatePatchRequestInformation(body *i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.Message, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *MessageRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler *ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(func() (error)) {
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
func (m *MessageRequestBuilder) Extensions()(i281b692d3f673c995c16b5cc69287c818497fc74562aed972acc168040629f74.ExtensionsRequestBuilder) {
    return *i281b692d3f673c995c16b5cc69287c818497fc74562aed972acc168040629f74.NewExtensionsRequestBuilder(m.currentPath + m.pathSegment, m.requestAdapter, false);
}
func (m *MessageRequestBuilder) ExtensionsById(id string)(icb423aed2922bb0caf9a0353a5ec5bfeb474588dbd1218008894e86ab49c6e25.ExtensionRequestBuilder) {
    return *icb423aed2922bb0caf9a0353a5ec5bfeb474588dbd1218008894e86ab49c6e25.NewExtensionRequestBuilder(m.currentPath + m.pathSegment + "/extensions/" + id, m.requestAdapter, false);
}
func (m *MessageRequestBuilder) Get(q func (value *MessageRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler *ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(func() (*i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.Message, error)) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return func() (*i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.Message, error) { return nil, err }
    }
    return func() (*i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.Message, error) {
        res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return new(i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.Message) }, *responseHandler)()
        if err != nil {
            return nil, err
        }
        return res.(*i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.Message), nil
    }
}
func (m *MessageRequestBuilder) MultiValueExtendedProperties()(i165e60b62d06f92d5e6029daee4ae129b1c3d1780dce2631cc18fbe8c6d39c9d.MultiValueExtendedPropertiesRequestBuilder) {
    return *i165e60b62d06f92d5e6029daee4ae129b1c3d1780dce2631cc18fbe8c6d39c9d.NewMultiValueExtendedPropertiesRequestBuilder(m.currentPath + m.pathSegment, m.requestAdapter, false);
}
func (m *MessageRequestBuilder) MultiValueExtendedPropertiesById(id string)(i314a2a68928bac08cd3678f4259eab0e5845a67017fe63526811793a0c7a4615.MultiValueLegacyExtendedPropertyRequestBuilder) {
    return *i314a2a68928bac08cd3678f4259eab0e5845a67017fe63526811793a0c7a4615.NewMultiValueLegacyExtendedPropertyRequestBuilder(m.currentPath + m.pathSegment + "/multiValueExtendedProperties/" + id, m.requestAdapter, false);
}
func (m *MessageRequestBuilder) Patch(body *i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.Message, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler *ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(func() (error)) {
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
func (m *MessageRequestBuilder) SingleValueExtendedProperties()(i833cf7d5250828d0b2288e87035b59c2900b4aef7c2eb2ff5494a38edea8268e.SingleValueExtendedPropertiesRequestBuilder) {
    return *i833cf7d5250828d0b2288e87035b59c2900b4aef7c2eb2ff5494a38edea8268e.NewSingleValueExtendedPropertiesRequestBuilder(m.currentPath + m.pathSegment, m.requestAdapter, false);
}
func (m *MessageRequestBuilder) SingleValueExtendedPropertiesById(id string)(ia3b46812df8e43b8356e8bad54f90359cc966c930f9415ac7f61670490ce857f.SingleValueLegacyExtendedPropertyRequestBuilder) {
    return *ia3b46812df8e43b8356e8bad54f90359cc966c930f9415ac7f61670490ce857f.NewSingleValueLegacyExtendedPropertyRequestBuilder(m.currentPath + m.pathSegment + "/singleValueExtendedProperties/" + id, m.requestAdapter, false);
}
