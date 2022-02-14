package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i645cf451e901c056b34ae87bfc90411c77b92a5c7b8ffa180f04f5e11c4d417d "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/models/odata"
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

// MailFolderRequestBuilder builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}
type MailFolderRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// MailFolderRequestBuilderDeleteOptions options for Delete
type MailFolderRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// MailFolderRequestBuilderGetOptions options for Get
type MailFolderRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *MailFolderRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// MailFolderRequestBuilderGetQueryParameters the user's mail folders. Read-only. Nullable.
type MailFolderRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string;
}
// MailFolderRequestBuilderPatchOptions options for Patch
type MailFolderRequestBuilderPatchOptions struct {
    // 
    Body *i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.MailFolder;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *MailFolderRequestBuilder) ChildFolders()(*i25eccaab35ae736780065dc9a8dc7482fdcb0caedd552ebe3bc32819574bdd1c.ChildFoldersRequestBuilder) {
    return i25eccaab35ae736780065dc9a8dc7482fdcb0caedd552ebe3bc32819574bdd1c.NewChildFoldersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildFoldersById gets an item from the github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/.users.item.mailFolders.item.childFolders.item collection
func (m *MailFolderRequestBuilder) ChildFoldersById(id string)(*i181e312a167613000dccbb9ebecd1868a9fa4f41ee5a7ecb670740df9765d172.MailFolderRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mailFolder_id1"] = id
    }
    return i181e312a167613000dccbb9ebecd1868a9fa4f41ee5a7ecb670740df9765d172.NewMailFolderRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewMailFolderRequestBuilderInternal instantiates a new MailFolderRequestBuilder and sets the default values.
func NewMailFolderRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MailFolderRequestBuilder) {
    m := &MailFolderRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/mailFolders/{mailFolder_id}{?select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMailFolderRequestBuilder instantiates a new MailFolderRequestBuilder and sets the default values.
func NewMailFolderRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MailFolderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMailFolderRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the user's mail folders. Read-only. Nullable.
func (m *MailFolderRequestBuilder) CreateDeleteRequestInformation(options *MailFolderRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the user's mail folders. Read-only. Nullable.
func (m *MailFolderRequestBuilder) CreateGetRequestInformation(options *MailFolderRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation the user's mail folders. Read-only. Nullable.
func (m *MailFolderRequestBuilder) CreatePatchRequestInformation(options *MailFolderRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete the user's mail folders. Read-only. Nullable.
func (m *MailFolderRequestBuilder) Delete(options *MailFolderRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i645cf451e901c056b34ae87bfc90411c77b92a5c7b8ffa180f04f5e11c4d417d.CreateErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the user's mail folders. Read-only. Nullable.
func (m *MailFolderRequestBuilder) Get(options *MailFolderRequestBuilderGetOptions)(*i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.MailFolder, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i645cf451e901c056b34ae87bfc90411c77b92a5c7b8ffa180f04f5e11c4d417d.CreateErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.CreateMailFolderFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(*i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.MailFolder), nil
}
func (m *MailFolderRequestBuilder) MessageRules()(*i8292dd9b1a0fee9945f96685dd56cd53e49fea510defe9e9be5f718692b7780b.MessageRulesRequestBuilder) {
    return i8292dd9b1a0fee9945f96685dd56cd53e49fea510defe9e9be5f718692b7780b.NewMessageRulesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessageRulesById gets an item from the github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/.users.item.mailFolders.item.messageRules.item collection
func (m *MailFolderRequestBuilder) MessageRulesById(id string)(*i0fbca58a8d08af34163fee8e520aa4597b6c44f81ff3cfb5f35e29bcdca13621.MessageRuleRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["messageRule_id"] = id
    }
    return i0fbca58a8d08af34163fee8e520aa4597b6c44f81ff3cfb5f35e29bcdca13621.NewMessageRuleRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MailFolderRequestBuilder) Messages()(*ifde4ccdcab963c649b4bf69fc48707d15a301d0c5085720fb9c273d0b6fdab37.MessagesRequestBuilder) {
    return ifde4ccdcab963c649b4bf69fc48707d15a301d0c5085720fb9c273d0b6fdab37.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById gets an item from the github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/.users.item.mailFolders.item.messages.item collection
func (m *MailFolderRequestBuilder) MessagesById(id string)(*id9e66f56b6b4a00e301a4a3f0507d126e9c676567ee5351fbae1f29bdf7542a1.MessageRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["message_id"] = id
    }
    return id9e66f56b6b4a00e301a4a3f0507d126e9c676567ee5351fbae1f29bdf7542a1.NewMessageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MailFolderRequestBuilder) MultiValueExtendedProperties()(*i9edbc9e69c9374a312359d6ba60470f270b3741c34a0b8bf1d40a0b0997e5c7b.MultiValueExtendedPropertiesRequestBuilder) {
    return i9edbc9e69c9374a312359d6ba60470f270b3741c34a0b8bf1d40a0b0997e5c7b.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/.users.item.mailFolders.item.multiValueExtendedProperties.item collection
func (m *MailFolderRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i33ecd413f6751e2cbf414d6016f0af76590db7141528c4aac7ee62c4cb3c4bcc.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i33ecd413f6751e2cbf414d6016f0af76590db7141528c4aac7ee62c4cb3c4bcc.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the user's mail folders. Read-only. Nullable.
func (m *MailFolderRequestBuilder) Patch(options *MailFolderRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i645cf451e901c056b34ae87bfc90411c77b92a5c7b8ffa180f04f5e11c4d417d.CreateErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *MailFolderRequestBuilder) SingleValueExtendedProperties()(*i813340affbf9a598103dd11e02bd85608be12b91284a30d07f60a1aac1a263e6.SingleValueExtendedPropertiesRequestBuilder) {
    return i813340affbf9a598103dd11e02bd85608be12b91284a30d07f60a1aac1a263e6.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/.users.item.mailFolders.item.singleValueExtendedProperties.item collection
func (m *MailFolderRequestBuilder) SingleValueExtendedPropertiesById(id string)(*idbb7dc5e3b94a6671fc2fde24821fbe6f90b0bfdd4d8781f45540e2e3d778e66.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return idbb7dc5e3b94a6671fc2fde24821fbe6f90b0bfdd4d8781f45540e2e3d778e66.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
