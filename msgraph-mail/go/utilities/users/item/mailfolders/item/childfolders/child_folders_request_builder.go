package childfolders

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/models/microsoft/graph"
)

// ChildFoldersRequestBuilder builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}\childFolders
type ChildFoldersRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ChildFoldersRequestBuilderGetQueryParameters the collection of child folders in the mailFolder.
type ChildFoldersRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
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
// ChildFoldersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChildFoldersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ChildFoldersRequestBuilderGetQueryParameters
}
// ChildFoldersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChildFoldersRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewChildFoldersRequestBuilderInternal instantiates a new ChildFoldersRequestBuilder and sets the default values.
func NewChildFoldersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChildFoldersRequestBuilder) {
    m := &ChildFoldersRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/childFolders{?%24top,%24skip,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewChildFoldersRequestBuilder instantiates a new ChildFoldersRequestBuilder and sets the default values.
func NewChildFoldersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChildFoldersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChildFoldersRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the collection of child folders in the mailFolder.
func (m *ChildFoldersRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the collection of child folders in the mailFolder.
func (m *ChildFoldersRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *ChildFoldersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation the collection of child folders in the mailFolder.
func (m *ChildFoldersRequestBuilder) CreatePostRequestInformation(body i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.MailFolderable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration the collection of child folders in the mailFolder.
func (m *ChildFoldersRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.MailFolderable, requestConfiguration *ChildFoldersRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get the collection of child folders in the mailFolder.
func (m *ChildFoldersRequestBuilder) Get(requestConfiguration *ChildFoldersRequestBuilderGetRequestConfiguration)(ChildFoldersResponseable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler the collection of child folders in the mailFolder.
func (m *ChildFoldersRequestBuilder) GetWithResponseHandler(requestConfiguration *ChildFoldersRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ChildFoldersResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateChildFoldersResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(ChildFoldersResponseable), nil
}
// Post the collection of child folders in the mailFolder.
func (m *ChildFoldersRequestBuilder) Post(body i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.MailFolderable, requestConfiguration *ChildFoldersRequestBuilderPostRequestConfiguration)(i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.MailFolderable, error) {
    return m.PostWithResponseHandler(body, requestConfiguration, nil);
}
// PostWithResponseHandler the collection of child folders in the mailFolder.
func (m *ChildFoldersRequestBuilder) PostWithResponseHandler(body i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.MailFolderable, requestConfiguration *ChildFoldersRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.MailFolderable, error) {
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
