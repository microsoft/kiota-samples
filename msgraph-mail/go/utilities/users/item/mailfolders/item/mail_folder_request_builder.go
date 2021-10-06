package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/models/microsoft/graph"
    i25eccaab35ae736780065dc9a8dc7482fdcb0caedd552ebe3bc32819574bdd1c "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/mailfolders/item/childfolders"
    i813340affbf9a598103dd11e02bd85608be12b91284a30d07f60a1aac1a263e6 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/mailfolders/item/singlevalueextendedproperties"
    i8292dd9b1a0fee9945f96685dd56cd53e49fea510defe9e9be5f718692b7780b "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/mailfolders/item/messagerules"
    i9edbc9e69c9374a312359d6ba60470f270b3741c34a0b8bf1d40a0b0997e5c7b "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/mailfolders/item/multivalueextendedproperties"
    ifde4ccdcab963c649b4bf69fc48707d15a301d0c5085720fb9c273d0b6fdab37 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/mailfolders/item/messages"
    i0fbca58a8d08af34163fee8e520aa4597b6c44f81ff3cfb5f35e29bcdca13621 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/mailfolders/item/messagerules/item"
    i181e312a167613000dccbb9ebecd1868a9fa4f41ee5a7ecb670740df9765d172 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/mailfolders/item/childfolders/item"
    i33ecd413f6751e2cbf414d6016f0af76590db7141528c4aac7ee62c4cb3c4bcc "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/mailfolders/item/multivalueextendedproperties/item"
    id9e66f56b6b4a00e301a4a3f0507d126e9c676567ee5351fbae1f29bdf7542a1 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/mailfolders/item/messages/item"
    idbb7dc5e3b94a6671fc2fde24821fbe6f90b0bfdd4d8781f45540e2e3d778e66 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/mailfolders/item/singlevalueextendedproperties/item"
)

type MailFolderRequestBuilder struct {
    currentPath string;
    isRawUrl bool;
    pathSegment string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
}
type MailFolderRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *MailFolderRequestBuilder) ChildFolders()(i25eccaab35ae736780065dc9a8dc7482fdcb0caedd552ebe3bc32819574bdd1c.ChildFoldersRequestBuilder) {
    return *i25eccaab35ae736780065dc9a8dc7482fdcb0caedd552ebe3bc32819574bdd1c.NewChildFoldersRequestBuilder(m.currentPath + m.pathSegment, m.requestAdapter, false);
}
func (m *MailFolderRequestBuilder) ChildFoldersById(id string)(i181e312a167613000dccbb9ebecd1868a9fa4f41ee5a7ecb670740df9765d172.MailFolderRequestBuilder) {
    return *i181e312a167613000dccbb9ebecd1868a9fa4f41ee5a7ecb670740df9765d172.NewMailFolderRequestBuilder(m.currentPath + m.pathSegment + "/childFolders/" + id, m.requestAdapter, false);
}
func NewMailFolderRequestBuilder(currentPath string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter, isRawUrl bool)(*MailFolderRequestBuilder) {
    m := &MailFolderRequestBuilder{
    }
    m.pathSegment = "";
    m.requestAdapter = requestAdapter;
    m.currentPath = currentPath;
    m.isRawUrl = isRawUrl;
    return m
}
func (m *MailFolderRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *MailFolderRequestBuilder) CreateGetRequestInformation(q func (value *MailFolderRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    err := requestInfo.SetUri(m.currentPath, m.pathSegment, m.isRawUrl)
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
    if o != nil {
        err = requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, err
}
func (m *MailFolderRequestBuilder) CreatePatchRequestInformation(body *i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.MailFolder, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *MailFolderRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler *ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(func() (error)) {
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
func (m *MailFolderRequestBuilder) Get(q func (value *MailFolderRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler *ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(func() (*i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.MailFolder, error)) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return func() (*i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.MailFolder, error) { return nil, err }
    }
    return func() (*i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.MailFolder, error) {
        res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return new(i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.MailFolder) }, *responseHandler)()
        if err != nil {
            return nil, err
        }
        return res.(*i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.MailFolder), nil
    }
}
func (m *MailFolderRequestBuilder) MessageRules()(i8292dd9b1a0fee9945f96685dd56cd53e49fea510defe9e9be5f718692b7780b.MessageRulesRequestBuilder) {
    return *i8292dd9b1a0fee9945f96685dd56cd53e49fea510defe9e9be5f718692b7780b.NewMessageRulesRequestBuilder(m.currentPath + m.pathSegment, m.requestAdapter, false);
}
func (m *MailFolderRequestBuilder) MessageRulesById(id string)(i0fbca58a8d08af34163fee8e520aa4597b6c44f81ff3cfb5f35e29bcdca13621.MessageRuleRequestBuilder) {
    return *i0fbca58a8d08af34163fee8e520aa4597b6c44f81ff3cfb5f35e29bcdca13621.NewMessageRuleRequestBuilder(m.currentPath + m.pathSegment + "/messageRules/" + id, m.requestAdapter, false);
}
func (m *MailFolderRequestBuilder) Messages()(ifde4ccdcab963c649b4bf69fc48707d15a301d0c5085720fb9c273d0b6fdab37.MessagesRequestBuilder) {
    return *ifde4ccdcab963c649b4bf69fc48707d15a301d0c5085720fb9c273d0b6fdab37.NewMessagesRequestBuilder(m.currentPath + m.pathSegment, m.requestAdapter, false);
}
func (m *MailFolderRequestBuilder) MessagesById(id string)(id9e66f56b6b4a00e301a4a3f0507d126e9c676567ee5351fbae1f29bdf7542a1.MessageRequestBuilder) {
    return *id9e66f56b6b4a00e301a4a3f0507d126e9c676567ee5351fbae1f29bdf7542a1.NewMessageRequestBuilder(m.currentPath + m.pathSegment + "/messages/" + id, m.requestAdapter, false);
}
func (m *MailFolderRequestBuilder) MultiValueExtendedProperties()(i9edbc9e69c9374a312359d6ba60470f270b3741c34a0b8bf1d40a0b0997e5c7b.MultiValueExtendedPropertiesRequestBuilder) {
    return *i9edbc9e69c9374a312359d6ba60470f270b3741c34a0b8bf1d40a0b0997e5c7b.NewMultiValueExtendedPropertiesRequestBuilder(m.currentPath + m.pathSegment, m.requestAdapter, false);
}
func (m *MailFolderRequestBuilder) MultiValueExtendedPropertiesById(id string)(i33ecd413f6751e2cbf414d6016f0af76590db7141528c4aac7ee62c4cb3c4bcc.MultiValueLegacyExtendedPropertyRequestBuilder) {
    return *i33ecd413f6751e2cbf414d6016f0af76590db7141528c4aac7ee62c4cb3c4bcc.NewMultiValueLegacyExtendedPropertyRequestBuilder(m.currentPath + m.pathSegment + "/multiValueExtendedProperties/" + id, m.requestAdapter, false);
}
func (m *MailFolderRequestBuilder) Patch(body *i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.MailFolder, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler *ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(func() (error)) {
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
func (m *MailFolderRequestBuilder) SingleValueExtendedProperties()(i813340affbf9a598103dd11e02bd85608be12b91284a30d07f60a1aac1a263e6.SingleValueExtendedPropertiesRequestBuilder) {
    return *i813340affbf9a598103dd11e02bd85608be12b91284a30d07f60a1aac1a263e6.NewSingleValueExtendedPropertiesRequestBuilder(m.currentPath + m.pathSegment, m.requestAdapter, false);
}
func (m *MailFolderRequestBuilder) SingleValueExtendedPropertiesById(id string)(idbb7dc5e3b94a6671fc2fde24821fbe6f90b0bfdd4d8781f45540e2e3d778e66.SingleValueLegacyExtendedPropertyRequestBuilder) {
    return *idbb7dc5e3b94a6671fc2fde24821fbe6f90b0bfdd4d8781f45540e2e3d778e66.NewSingleValueLegacyExtendedPropertyRequestBuilder(m.currentPath + m.pathSegment + "/singleValueExtendedProperties/" + id, m.requestAdapter, false);
}
