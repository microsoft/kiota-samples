package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/models"
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

// MailFolderItemRequestBuilder builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}
type MailFolderItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MailFolderItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MailFolderItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// MailFolderItemRequestBuilderGetQueryParameters get mailFolders from users
type MailFolderItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MailFolderItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MailFolderItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MailFolderItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// MailFolderItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MailFolderItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// ChildFolders the childFolders property
func (m *MailFolderItemRequestBuilder) ChildFolders()(*i25eccaab35ae736780065dc9a8dc7482fdcb0caedd552ebe3bc32819574bdd1c.ChildFoldersRequestBuilder) {
    return i25eccaab35ae736780065dc9a8dc7482fdcb0caedd552ebe3bc32819574bdd1c.NewChildFoldersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildFoldersById gets an item from the github.com/microsoft/kiota-samples/msgraph-mail/go/utilities.users.item.mailFolders.item.childFolders.item collection
func (m *MailFolderItemRequestBuilder) ChildFoldersById(id string)(*i181e312a167613000dccbb9ebecd1868a9fa4f41ee5a7ecb670740df9765d172.MailFolderItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mailFolder%2Did1"] = id
    }
    return i181e312a167613000dccbb9ebecd1868a9fa4f41ee5a7ecb670740df9765d172.NewMailFolderItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewMailFolderItemRequestBuilderInternal instantiates a new MailFolderItemRequestBuilder and sets the default values.
func NewMailFolderItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MailFolderItemRequestBuilder) {
    m := &MailFolderItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}{?%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMailFolderItemRequestBuilder instantiates a new MailFolderItemRequestBuilder and sets the default values.
func NewMailFolderItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MailFolderItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMailFolderItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property mailFolders for users
func (m *MailFolderItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property mailFolders for users
func (m *MailFolderItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *MailFolderItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get mailFolders from users
func (m *MailFolderItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get mailFolders from users
func (m *MailFolderItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *MailFolderItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property mailFolders in users
func (m *MailFolderItemRequestBuilder) CreatePatchRequestInformation(body ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MailFolderable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property mailFolders in users
func (m *MailFolderItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MailFolderable, requestConfiguration *MailFolderItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property mailFolders for users
func (m *MailFolderItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MailFolderItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    var responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler = nil
    if requestConfiguration != nil && requestConfiguration.ResponseHandler != nil {
        responseHandler = requestConfiguration.ResponseHandler
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, responseHandler, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get get mailFolders from users
func (m *MailFolderItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MailFolderItemRequestBuilderGetRequestConfiguration)(ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MailFolderable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    var responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler = nil
    if requestConfiguration != nil && requestConfiguration.ResponseHandler != nil {
        responseHandler = requestConfiguration.ResponseHandler
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.CreateMailFolderFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MailFolderable), nil
}
// MessageRules the messageRules property
func (m *MailFolderItemRequestBuilder) MessageRules()(*i8292dd9b1a0fee9945f96685dd56cd53e49fea510defe9e9be5f718692b7780b.MessageRulesRequestBuilder) {
    return i8292dd9b1a0fee9945f96685dd56cd53e49fea510defe9e9be5f718692b7780b.NewMessageRulesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessageRulesById gets an item from the github.com/microsoft/kiota-samples/msgraph-mail/go/utilities.users.item.mailFolders.item.messageRules.item collection
func (m *MailFolderItemRequestBuilder) MessageRulesById(id string)(*i0fbca58a8d08af34163fee8e520aa4597b6c44f81ff3cfb5f35e29bcdca13621.MessageRuleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["messageRule%2Did"] = id
    }
    return i0fbca58a8d08af34163fee8e520aa4597b6c44f81ff3cfb5f35e29bcdca13621.NewMessageRuleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Messages the messages property
func (m *MailFolderItemRequestBuilder) Messages()(*ifde4ccdcab963c649b4bf69fc48707d15a301d0c5085720fb9c273d0b6fdab37.MessagesRequestBuilder) {
    return ifde4ccdcab963c649b4bf69fc48707d15a301d0c5085720fb9c273d0b6fdab37.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById gets an item from the github.com/microsoft/kiota-samples/msgraph-mail/go/utilities.users.item.mailFolders.item.messages.item collection
func (m *MailFolderItemRequestBuilder) MessagesById(id string)(*id9e66f56b6b4a00e301a4a3f0507d126e9c676567ee5351fbae1f29bdf7542a1.MessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["message%2Did"] = id
    }
    return id9e66f56b6b4a00e301a4a3f0507d126e9c676567ee5351fbae1f29bdf7542a1.NewMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *MailFolderItemRequestBuilder) MultiValueExtendedProperties()(*i9edbc9e69c9374a312359d6ba60470f270b3741c34a0b8bf1d40a0b0997e5c7b.MultiValueExtendedPropertiesRequestBuilder) {
    return i9edbc9e69c9374a312359d6ba60470f270b3741c34a0b8bf1d40a0b0997e5c7b.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoft/kiota-samples/msgraph-mail/go/utilities.users.item.mailFolders.item.multiValueExtendedProperties.item collection
func (m *MailFolderItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i33ecd413f6751e2cbf414d6016f0af76590db7141528c4aac7ee62c4cb3c4bcc.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return i33ecd413f6751e2cbf414d6016f0af76590db7141528c4aac7ee62c4cb3c4bcc.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property mailFolders in users
func (m *MailFolderItemRequestBuilder) Patch(ctx context.Context, body ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MailFolderable, requestConfiguration *MailFolderItemRequestBuilderPatchRequestConfiguration)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    var responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler = nil
    if requestConfiguration != nil && requestConfiguration.ResponseHandler != nil {
        responseHandler = requestConfiguration.ResponseHandler
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, responseHandler, nil)
    if err != nil {
        return err
    }
    return nil
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *MailFolderItemRequestBuilder) SingleValueExtendedProperties()(*i813340affbf9a598103dd11e02bd85608be12b91284a30d07f60a1aac1a263e6.SingleValueExtendedPropertiesRequestBuilder) {
    return i813340affbf9a598103dd11e02bd85608be12b91284a30d07f60a1aac1a263e6.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoft/kiota-samples/msgraph-mail/go/utilities.users.item.mailFolders.item.singleValueExtendedProperties.item collection
func (m *MailFolderItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*idbb7dc5e3b94a6671fc2fde24821fbe6f90b0bfdd4d8781f45540e2e3d778e66.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return idbb7dc5e3b94a6671fc2fde24821fbe6f90b0bfdd4d8781f45540e2e3d778e66.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
