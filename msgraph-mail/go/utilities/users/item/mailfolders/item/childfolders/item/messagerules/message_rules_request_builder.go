package messagerules

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/models"
)

// MessageRulesRequestBuilder builds and executes requests for operations under \users\{user-id}\mailFolders\{mailFolder-id}\childFolders\{mailFolder-id1}\messageRules
type MessageRulesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MessageRulesRequestBuilderGetQueryParameters get messageRules from users
type MessageRulesRequestBuilderGetQueryParameters struct {
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
// MessageRulesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MessageRulesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MessageRulesRequestBuilderGetQueryParameters
}
// MessageRulesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MessageRulesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMessageRulesRequestBuilderInternal instantiates a new MessageRulesRequestBuilder and sets the default values.
func NewMessageRulesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MessageRulesRequestBuilder) {
    m := &MessageRulesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/childFolders/{mailFolder%2Did1}/messageRules{?%24top,%24skip,%24filter,%24count,%24orderby,%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMessageRulesRequestBuilder instantiates a new MessageRulesRequestBuilder and sets the default values.
func NewMessageRulesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MessageRulesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMessageRulesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get messageRules from users
func (m *MessageRulesRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get messageRules from users
func (m *MessageRulesRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *MessageRulesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to messageRules for users
func (m *MessageRulesRequestBuilder) CreatePostRequestInformation(body ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MessageRuleable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration create new navigation property to messageRules for users
func (m *MessageRulesRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MessageRuleable, requestConfiguration *MessageRulesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get get messageRules from users
func (m *MessageRulesRequestBuilder) Get(ctx context.Context, requestConfiguration *MessageRulesRequestBuilderGetRequestConfiguration)(ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MessageRuleCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.CreateMessageRuleCollectionResponseFromDiscriminatorValue, requestInfo.ResponseHandler, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MessageRuleCollectionResponseable), nil
}
// Post create new navigation property to messageRules for users
func (m *MessageRulesRequestBuilder) Post(ctx context.Context, body ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MessageRuleable, requestConfiguration *MessageRulesRequestBuilderPostRequestConfiguration)(ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MessageRuleable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.CreateMessageRuleFromDiscriminatorValue, requestInfo.ResponseHandler, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ieea96ea0706c7e10d110f01563f903230c17531f1ba4f5e7095035777bc8b5e5.MessageRuleable), nil
}
