package todoitems

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia529d5fba26feec33c52ebc42b0bfa295e15ca84ffe68983636c172774ec4f04 "todoclient/client/models/todoapi/models"
)

// ToDoItemsRequestBuilder provides operations to manage the collection of ToDoItem entities.
type ToDoItemsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ToDoItemsRequestBuilderGetQueryParameters get entities from ToDoItems
type ToDoItemsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ToDoItemsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ToDoItemsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ToDoItemsRequestBuilderGetQueryParameters
}
// ToDoItemsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ToDoItemsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewToDoItemsRequestBuilderInternal instantiates a new ToDoItemsRequestBuilder and sets the default values.
func NewToDoItemsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ToDoItemsRequestBuilder) {
    m := &ToDoItemsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/ToDoItems{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewToDoItemsRequestBuilder instantiates a new ToDoItemsRequestBuilder and sets the default values.
func NewToDoItemsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ToDoItemsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewToDoItemsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *ToDoItemsRequestBuilder) Count()(*CountRequestBuilder) {
    return NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get get entities from ToDoItems
func (m *ToDoItemsRequestBuilder) Get(ctx context.Context, requestConfiguration *ToDoItemsRequestBuilderGetRequestConfiguration)(ia529d5fba26feec33c52ebc42b0bfa295e15ca84ffe68983636c172774ec4f04.ToDoItemCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ia529d5fba26feec33c52ebc42b0bfa295e15ca84ffe68983636c172774ec4f04.CreateToDoItemCollectionResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ia529d5fba26feec33c52ebc42b0bfa295e15ca84ffe68983636c172774ec4f04.ToDoItemCollectionResponseable), nil
}
// Post add new entity to ToDoItems
func (m *ToDoItemsRequestBuilder) Post(ctx context.Context, body ia529d5fba26feec33c52ebc42b0bfa295e15ca84ffe68983636c172774ec4f04.ToDoItemable, requestConfiguration *ToDoItemsRequestBuilderPostRequestConfiguration)(ia529d5fba26feec33c52ebc42b0bfa295e15ca84ffe68983636c172774ec4f04.ToDoItemable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ia529d5fba26feec33c52ebc42b0bfa295e15ca84ffe68983636c172774ec4f04.CreateToDoItemFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ia529d5fba26feec33c52ebc42b0bfa295e15ca84ffe68983636c172774ec4f04.ToDoItemable), nil
}
// ToGetRequestInformation get entities from ToDoItems
func (m *ToDoItemsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ToDoItemsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPostRequestInformation add new entity to ToDoItems
func (m *ToDoItemsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ia529d5fba26feec33c52ebc42b0bfa295e15ca84ffe68983636c172774ec4f04.ToDoItemable, requestConfiguration *ToDoItemsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
