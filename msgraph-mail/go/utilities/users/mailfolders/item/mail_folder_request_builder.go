package item

import (
    url "net/url"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/models"
    i17b176b6a1414dd086eba0883aee3fc49c21b71511308a01a6732762ead9b8e8 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailfolders/messages"
    i660191d0d3ef7133b6f3e0a4532d18c73c2f52a7222a034c4492d6ec6ef3fc3e "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailfolders/multivalueextendedproperties"
    iab1a0dd9a491da948f96274b48b251c29156f5ab5cdef57a533c314762a96cba "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailfolders/childfolders"
    ib4b1a269396ed151935b7d4ec590d14cf3ed072057ab6475f11863830766d145 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailfolders/singlevalueextendedproperties"
    id4ccc2d90f7e823c6e1c9e4a9a1f0e784a26e3fd6a68c038bda15d61728d0e46 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailfolders/messagerules"
    i0dcb0d8c6a911f67209ea659b254ba6713191c318f2799cf8a9383d821d87061 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailfolders/childfolders/item"
    i5d2f606000ccb668c7a688148beea132a38f81fbf9cf32e175421151435d1b29 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailfolders/messagerules/item"
    ic93bf2d2dca930088dbc375148726cf159640b4b7220bf5ee92920b4e2022faa "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailfolders/singlevalueextendedproperties/item"
    id2dc360a4b1db0a33958a93a105ee47a087c7d3853efa23c73dbf41c05cd738e "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailfolders/multivalueextendedproperties/item"
    ifdebf055cf7726bc6ca41ec4be0e381235fd73f507038a2ea27cd08a6fa09506 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailfolders/messages/item"
)

type MailFolderRequestBuilder struct {
    ChildFolders *iab1a0dd9a491da948f96274b48b251c29156f5ab5cdef57a533c314762a96cba.ChildFoldersRequestBuilder;
    currentPath *string;
    httpCore ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.HttpCore;
    MessageRules *id4ccc2d90f7e823c6e1c9e4a9a1f0e784a26e3fd6a68c038bda15d61728d0e46.MessageRulesRequestBuilder;
    Messages *i17b176b6a1414dd086eba0883aee3fc49c21b71511308a01a6732762ead9b8e8.MessagesRequestBuilder;
    MultiValueExtendedProperties *i660191d0d3ef7133b6f3e0a4532d18c73c2f52a7222a034c4492d6ec6ef3fc3e.MultiValueExtendedPropertiesRequestBuilder;
    pathSegment *string;
    SingleValueExtendedProperties *ib4b1a269396ed151935b7d4ec590d14cf3ed072057ab6475f11863830766d145.SingleValueExtendedPropertiesRequestBuilder;
}
type MailFolderRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []*string;
    Select []*string;
}
func (m *MailFolderRequestBuilder) ChildFoldersById(id *string)(i0dcb0d8c6a911f67209ea659b254ba6713191c318f2799cf8a9383d821d87061.MailFolderRequestBuilder) {
    return i0dcb0d8c6a911f67209ea659b254ba6713191c318f2799cf8a9383d821d87061.MailFolderRequestBuilder.NewMailFolderRequestBuilder(*m.currentPath + *m.pathSegment + "/childFolders/" + id, m.httpCore);
}
func NewMailFolderRequestBuilder(currentPath *string, httpCore ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.HttpCore)(*MailFolderRequestBuilder) {
    return nil
}
func (m *MailFolderRequestBuilder) CreateDeleteRequestInfo(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.MiddlewareOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInfo, error) {
    requestInfo := new(ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInfo)
    uri, err := url.Parse(*m.currentPath + *m.pathSegment)
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
    return requestInfo, err
}
func (m *MailFolderRequestBuilder) CreateGetRequestInfo(q func (value *MailFolderRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.MiddlewareOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInfo, error) {
    requestInfo := new(ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInfo)
    uri, err := url.Parse(*m.currentPath + *m.pathSegment)
    requestInfo.URI = *uri
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if err != nil {
        return nil, err
    }
    if q != nil {
        qParams := new(MailFolderRequestBuilderGetQueryParameters)
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
func (m *MailFolderRequestBuilder) CreatePatchRequestInfo(body *ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MailFolder, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.MiddlewareOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInfo, error) {
    requestInfo := new(ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInfo)
    uri, err := url.Parse(*m.currentPath + *m.pathSegment)
    requestInfo.URI = *uri
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
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
func (m *MailFolderRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.MiddlewareOption, responseHandler *ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(func() (error)) {
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
func (m *MailFolderRequestBuilder) Get(q func (value *MailFolderRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.MiddlewareOption, responseHandler *ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(func() (*ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MailFolder, error)) {
    requestInfo, err := m.CreateGetRequestInfo(q, h, o);
    if err != nil {
        return func() (*ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MailFolder, error) { return nil, err }
    }
    return func() (*ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MailFolder, error) {
        res, err := m.httpCore.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return new(ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MailFolder) }, *responseHandler)()
        if err != nil {
            return nil, err
        }
        return res.(*ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MailFolder), nil
    }
}
func (m *MailFolderRequestBuilder) MessageRulesById(id *string)(i5d2f606000ccb668c7a688148beea132a38f81fbf9cf32e175421151435d1b29.MessageRuleRequestBuilder) {
    return i5d2f606000ccb668c7a688148beea132a38f81fbf9cf32e175421151435d1b29.MessageRuleRequestBuilder.NewMessageRuleRequestBuilder(*m.currentPath + *m.pathSegment + "/messageRules/" + id, m.httpCore);
}
func (m *MailFolderRequestBuilder) MessagesById(id *string)(ifdebf055cf7726bc6ca41ec4be0e381235fd73f507038a2ea27cd08a6fa09506.MessageRequestBuilder) {
    return ifdebf055cf7726bc6ca41ec4be0e381235fd73f507038a2ea27cd08a6fa09506.MessageRequestBuilder.NewMessageRequestBuilder(*m.currentPath + *m.pathSegment + "/messages/" + id, m.httpCore);
}
func (m *MailFolderRequestBuilder) MultiValueExtendedPropertiesById(id *string)(id2dc360a4b1db0a33958a93a105ee47a087c7d3853efa23c73dbf41c05cd738e.MultiValueLegacyExtendedPropertyRequestBuilder) {
    return id2dc360a4b1db0a33958a93a105ee47a087c7d3853efa23c73dbf41c05cd738e.MultiValueLegacyExtendedPropertyRequestBuilder.NewMultiValueLegacyExtendedPropertyRequestBuilder(*m.currentPath + *m.pathSegment + "/multiValueExtendedProperties/" + id, m.httpCore);
}
func (m *MailFolderRequestBuilder) Patch(body *ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MailFolder, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.MiddlewareOption, responseHandler *ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(func() (error)) {
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
func (m *MailFolderRequestBuilder) SingleValueExtendedPropertiesById(id *string)(ic93bf2d2dca930088dbc375148726cf159640b4b7220bf5ee92920b4e2022faa.SingleValueLegacyExtendedPropertyRequestBuilder) {
    return ic93bf2d2dca930088dbc375148726cf159640b4b7220bf5ee92920b4e2022faa.SingleValueLegacyExtendedPropertyRequestBuilder.NewSingleValueLegacyExtendedPropertyRequestBuilder(*m.currentPath + *m.pathSegment + "/singleValueExtendedProperties/" + id, m.httpCore);
}
