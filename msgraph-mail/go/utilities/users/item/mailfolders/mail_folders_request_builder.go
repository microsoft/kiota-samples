package mailfolders

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/models/microsoft/graph"
)

// MailFoldersRequestBuilder builds and executes requests for operations under \users\{user-id}\mailFolders
type MailFoldersRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MailFoldersRequestBuilderGetQueryParameters the user's mail folders. Read-only. Nullable.
type MailFoldersRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// MailFoldersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MailFoldersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MailFoldersRequestBuilderGetQueryParameters
}
// MailFoldersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MailFoldersRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMailFoldersRequestBuilderInternal instantiates a new MailFoldersRequestBuilder and sets the default values.
func NewMailFoldersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MailFoldersRequestBuilder) {
    m := &MailFoldersRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/mailFolders{?%24top,%24skip,%24filter,%24count,%24orderby,%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMailFoldersRequestBuilder instantiates a new MailFoldersRequestBuilder and sets the default values.
func NewMailFoldersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MailFoldersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMailFoldersRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the user's mail folders. Read-only. Nullable.
func (m *MailFoldersRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the user's mail folders. Read-only. Nullable.
func (m *MailFoldersRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *MailFoldersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePostRequestInformation the user's mail folders. Read-only. Nullable.
func (m *MailFoldersRequestBuilder) CreatePostRequestInformation(body i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.MailFolderable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration the user's mail folders. Read-only. Nullable.
func (m *MailFoldersRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.MailFolderable, requestConfiguration *MailFoldersRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get the user's mail folders. Read-only. Nullable.
func (m *MailFoldersRequestBuilder) Get(requestConfiguration *MailFoldersRequestBuilderGetRequestConfiguration)(MailFoldersResponseable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler the user's mail folders. Read-only. Nullable.
func (m *MailFoldersRequestBuilder) GetWithResponseHandler(requestConfiguration *MailFoldersRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(MailFoldersResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateMailFoldersResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(MailFoldersResponseable), nil
}
// Post the user's mail folders. Read-only. Nullable.
func (m *MailFoldersRequestBuilder) Post(body i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.MailFolderable, requestConfiguration *MailFoldersRequestBuilderPostRequestConfiguration)(i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.MailFolderable, error) {
    return m.PostWithResponseHandler(body, requestConfiguration, nil);
}
// PostWithResponseHandler the user's mail folders. Read-only. Nullable.
func (m *MailFoldersRequestBuilder) PostWithResponseHandler(body i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.MailFolderable, requestConfiguration *MailFoldersRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.MailFolderable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.CreateMailFolderFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.MailFolderable), nil
}
