package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/models"
    i10075314d3b96cd3f0f2195562b0027395b41f6069ced9cc8d93624bbc206528 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/messages/item/extensions"
    i170c26461395ae79c8485bb09f6acf415c54c121851bca73450cfff498f75eb8 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/messages/item/singlevalueextendedproperties"
    i228962c2a18462bcfbde8557f6668f3fd7d4e68d5c7f3f1f3b675bc816866033 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/messages/item/attachments"
    i9550b442511c46b7a974613d5ee3238fb100cef30fff0520a9682fb3d38ad326 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/messages/item/value"
    iadaad8f92ba85254c1d639d8a8cb86def5e843af4bc1062bc836f08d062196d6 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/messages/item/multivalueextendedproperties"
    i361d73846d6331cbbe8151dfaef3455f34e63dfa07626b67576681683b148587 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/messages/item/extensions/item"
    i9f73052b0f535426cd146e9403e4e088995d624a331ea8cf19005e2c6ff46d7f "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/messages/item/singlevalueextendedproperties/item"
    idd4800ce8115a136eb3f42205c28105422c049d9e5b23796f3bdf18e2778f6d6 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/messages/item/attachments/item"
    if06dd322de6c7a1abe876d18d2ddc86a4b8bcf49f2fea78dcba12aff7dcb4994 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/item/messages/item/multivalueextendedproperties/item"
)

// MessageItemRequestBuilder builds and executes requests for operations under \users\{user-id}\messages\{message-id}
type MessageItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MessageItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MessageItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MessageItemRequestBuilderGetQueryParameters get messages from users
type MessageItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MessageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MessageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MessageItemRequestBuilderGetQueryParameters
}
// MessageItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MessageItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Attachments the attachments property
func (m *MessageItemRequestBuilder) Attachments()(*i228962c2a18462bcfbde8557f6668f3fd7d4e68d5c7f3f1f3b675bc816866033.AttachmentsRequestBuilder) {
    return i228962c2a18462bcfbde8557f6668f3fd7d4e68d5c7f3f1f3b675bc816866033.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AttachmentsById gets an item from the github.com/microsoft/kiota-samples/msgraph-mail/go/utilities.users.item.messages.item.attachments.item collection
func (m *MessageItemRequestBuilder) AttachmentsById(id string)(*idd4800ce8115a136eb3f42205c28105422c049d9e5b23796f3bdf18e2778f6d6.AttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return idd4800ce8115a136eb3f42205c28105422c049d9e5b23796f3bdf18e2778f6d6.NewAttachmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewMessageItemRequestBuilderInternal instantiates a new MessageItemRequestBuilder and sets the default values.
func NewMessageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MessageItemRequestBuilder) {
    m := &MessageItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/messages/{message%2Did}{?%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMessageItemRequestBuilder instantiates a new MessageItemRequestBuilder and sets the default values.
func NewMessageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MessageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMessageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content the Content property
func (m *MessageItemRequestBuilder) Content()(*i9550b442511c46b7a974613d5ee3238fb100cef30fff0520a9682fb3d38ad326.ContentRequestBuilder) {
    return i9550b442511c46b7a974613d5ee3238fb100cef30fff0520a9682fb3d38ad326.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property messages for users
func (m *MessageItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property messages for users
func (m *MessageItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *MessageItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get messages from users
func (m *MessageItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get messages from users
func (m *MessageItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *MessageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property messages in users
func (m *MessageItemRequestBuilder) CreatePatchRequestInformation(body ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.Messageable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property messages in users
func (m *MessageItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.Messageable, requestConfiguration *MessageItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property messages for users
func (m *MessageItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MessageItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Extensions the extensions property
func (m *MessageItemRequestBuilder) Extensions()(*i10075314d3b96cd3f0f2195562b0027395b41f6069ced9cc8d93624bbc206528.ExtensionsRequestBuilder) {
    return i10075314d3b96cd3f0f2195562b0027395b41f6069ced9cc8d93624bbc206528.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoft/kiota-samples/msgraph-mail/go/utilities.users.item.messages.item.extensions.item collection
func (m *MessageItemRequestBuilder) ExtensionsById(id string)(*i361d73846d6331cbbe8151dfaef3455f34e63dfa07626b67576681683b148587.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return i361d73846d6331cbbe8151dfaef3455f34e63dfa07626b67576681683b148587.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get messages from users
func (m *MessageItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MessageItemRequestBuilderGetRequestConfiguration)(ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.Messageable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.CreateMessageFromDiscriminatorValue, requestInfo.ResponseHandler, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.Messageable), nil
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *MessageItemRequestBuilder) MultiValueExtendedProperties()(*iadaad8f92ba85254c1d639d8a8cb86def5e843af4bc1062bc836f08d062196d6.MultiValueExtendedPropertiesRequestBuilder) {
    return iadaad8f92ba85254c1d639d8a8cb86def5e843af4bc1062bc836f08d062196d6.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoft/kiota-samples/msgraph-mail/go/utilities.users.item.messages.item.multiValueExtendedProperties.item collection
func (m *MessageItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*if06dd322de6c7a1abe876d18d2ddc86a4b8bcf49f2fea78dcba12aff7dcb4994.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return if06dd322de6c7a1abe876d18d2ddc86a4b8bcf49f2fea78dcba12aff7dcb4994.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property messages in users
func (m *MessageItemRequestBuilder) Patch(ctx context.Context, body ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.Messageable, requestConfiguration *MessageItemRequestBuilderPatchRequestConfiguration)(error) {
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
func (m *MessageItemRequestBuilder) SingleValueExtendedProperties()(*i170c26461395ae79c8485bb09f6acf415c54c121851bca73450cfff498f75eb8.SingleValueExtendedPropertiesRequestBuilder) {
    return i170c26461395ae79c8485bb09f6acf415c54c121851bca73450cfff498f75eb8.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoft/kiota-samples/msgraph-mail/go/utilities.users.item.messages.item.singleValueExtendedProperties.item collection
func (m *MessageItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i9f73052b0f535426cd146e9403e4e088995d624a331ea8cf19005e2c6ff46d7f.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i9f73052b0f535426cd146e9403e4e088995d624a331ea8cf19005e2c6ff46d7f.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
