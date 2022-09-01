package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/models"
    i13ad6e53b1a8f540fab96224bcb4d5881ade631021d055b7b09e69ae995b5f39 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/mailfolders/item/childfolders/item/singlevalueextendedproperties"
    ibe5596e904d8875047b37e601f443a745f97e000b552b472cf8bf3d99b19f739 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/mailfolders/item/childfolders/item/messagerules"
    idb63710df8e89ab2e37d18aa00dd694c4a90d87f59809f6edb965e840fb0f10f "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/mailfolders/item/childfolders/item/multivalueextendedproperties"
    ie9c27f72c6ede1a0f0d330da6eda2bb2a76d484807fc7aa5b1a453cef6652380 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/mailfolders/item/childfolders/item/messages"
    i24628f8898b3dea703cc83eb7f541bb3b983a98735a1220250a7f9f1b736ef35 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/mailfolders/item/childfolders/item/singlevalueextendedproperties/item"
    i9d7f60919f34b51af8023dd559cc1c8aa1d6e134951fd793963e942aac75b37d "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/mailfolders/item/childfolders/item/messages/item"
    iad17ef363295f0dcae113b634fc7785e1312ddf07f41900c1294b67159c4dae3 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/mailfolders/item/childfolders/item/multivalueextendedproperties/item"
    iadc28bafba493be7b466b720963724d92d420760689b1739176ef1f76369f264 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/mailfolders/item/childfolders/item/messagerules/item"
)

// MailFolderItemRequestBuilder builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}\childFolders\{mailFolder-id1}
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
}
// MailFolderItemRequestBuilderGetQueryParameters get childFolders from users
type MailFolderItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
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
}
// MailFolderItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MailFolderItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMailFolderItemRequestBuilderInternal instantiates a new MailFolderItemRequestBuilder and sets the default values.
func NewMailFolderItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MailFolderItemRequestBuilder) {
    m := &MailFolderItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/childFolders/{mailFolder%2Did1}{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property childFolders for users
func (m *MailFolderItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property childFolders for users
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
// CreateGetRequestInformation get childFolders from users
func (m *MailFolderItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get childFolders from users
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
// CreatePatchRequestInformation update the navigation property childFolders in users
func (m *MailFolderItemRequestBuilder) CreatePatchRequestInformation(body ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MailFolderable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property childFolders in users
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
// Delete delete navigation property childFolders for users
func (m *MailFolderItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MailFolderItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, requestInfo.ResponseHandler, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get get childFolders from users
func (m *MailFolderItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MailFolderItemRequestBuilderGetRequestConfiguration)(ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MailFolderable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.CreateMailFolderFromDiscriminatorValue, requestInfo.ResponseHandler, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MailFolderable), nil
}
// MessageRules the messageRules property
func (m *MailFolderItemRequestBuilder) MessageRules()(*ibe5596e904d8875047b37e601f443a745f97e000b552b472cf8bf3d99b19f739.MessageRulesRequestBuilder) {
    return ibe5596e904d8875047b37e601f443a745f97e000b552b472cf8bf3d99b19f739.NewMessageRulesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessageRulesById gets an item from the github.com/microsoft/kiota-samples/msgraph-mail/go/utilities.users.item.mailFolders.item.childFolders.item.messageRules.item collection
func (m *MailFolderItemRequestBuilder) MessageRulesById(id string)(*iadc28bafba493be7b466b720963724d92d420760689b1739176ef1f76369f264.MessageRuleItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["messageRule%2Did"] = id
    }
    return iadc28bafba493be7b466b720963724d92d420760689b1739176ef1f76369f264.NewMessageRuleItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Messages the messages property
func (m *MailFolderItemRequestBuilder) Messages()(*ie9c27f72c6ede1a0f0d330da6eda2bb2a76d484807fc7aa5b1a453cef6652380.MessagesRequestBuilder) {
    return ie9c27f72c6ede1a0f0d330da6eda2bb2a76d484807fc7aa5b1a453cef6652380.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById gets an item from the github.com/microsoft/kiota-samples/msgraph-mail/go/utilities.users.item.mailFolders.item.childFolders.item.messages.item collection
func (m *MailFolderItemRequestBuilder) MessagesById(id string)(*i9d7f60919f34b51af8023dd559cc1c8aa1d6e134951fd793963e942aac75b37d.MessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["message%2Did"] = id
    }
    return i9d7f60919f34b51af8023dd559cc1c8aa1d6e134951fd793963e942aac75b37d.NewMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *MailFolderItemRequestBuilder) MultiValueExtendedProperties()(*idb63710df8e89ab2e37d18aa00dd694c4a90d87f59809f6edb965e840fb0f10f.MultiValueExtendedPropertiesRequestBuilder) {
    return idb63710df8e89ab2e37d18aa00dd694c4a90d87f59809f6edb965e840fb0f10f.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoft/kiota-samples/msgraph-mail/go/utilities.users.item.mailFolders.item.childFolders.item.multiValueExtendedProperties.item collection
func (m *MailFolderItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*iad17ef363295f0dcae113b634fc7785e1312ddf07f41900c1294b67159c4dae3.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return iad17ef363295f0dcae113b634fc7785e1312ddf07f41900c1294b67159c4dae3.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property childFolders in users
func (m *MailFolderItemRequestBuilder) Patch(ctx context.Context, body ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MailFolderable, requestConfiguration *MailFolderItemRequestBuilderPatchRequestConfiguration)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, requestInfo.ResponseHandler, nil)
    if err != nil {
        return err
    }
    return nil
}
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *MailFolderItemRequestBuilder) SingleValueExtendedProperties()(*i13ad6e53b1a8f540fab96224bcb4d5881ade631021d055b7b09e69ae995b5f39.SingleValueExtendedPropertiesRequestBuilder) {
    return i13ad6e53b1a8f540fab96224bcb4d5881ade631021d055b7b09e69ae995b5f39.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoft/kiota-samples/msgraph-mail/go/utilities.users.item.mailFolders.item.childFolders.item.singleValueExtendedProperties.item collection
func (m *MailFolderItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i24628f8898b3dea703cc83eb7f541bb3b983a98735a1220250a7f9f1b736ef35.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i24628f8898b3dea703cc83eb7f541bb3b983a98735a1220250a7f9f1b736ef35.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
