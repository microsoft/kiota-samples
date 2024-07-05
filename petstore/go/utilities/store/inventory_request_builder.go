package store

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// InventoryRequestBuilder builds and executes requests for operations under \store\inventory
type InventoryRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// InventoryRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type InventoryRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewInventoryRequestBuilderInternal instantiates a new InventoryRequestBuilder and sets the default values.
func NewInventoryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InventoryRequestBuilder) {
    m := &InventoryRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/store/inventory", pathParameters),
    }
    return m
}
// NewInventoryRequestBuilder instantiates a new InventoryRequestBuilder and sets the default values.
func NewInventoryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InventoryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewInventoryRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns a map of status codes to quantities
// Deprecated: This method is obsolete. Use GetAsInventoryGetResponse instead.
// returns a InventoryResponseable when successful
func (m *InventoryRequestBuilder) Get(ctx context.Context, requestConfiguration *InventoryRequestBuilderGetRequestConfiguration)(InventoryResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateInventoryResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(InventoryResponseable), nil
}
// GetAsInventoryGetResponse returns a map of status codes to quantities
// returns a InventoryGetResponseable when successful
func (m *InventoryRequestBuilder) GetAsInventoryGetResponse(ctx context.Context, requestConfiguration *InventoryRequestBuilderGetRequestConfiguration)(InventoryGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateInventoryGetResponseFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(InventoryGetResponseable), nil
}
// ToGetRequestInformation returns a map of status codes to quantities
// returns a *RequestInformation when successful
func (m *InventoryRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *InventoryRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *InventoryRequestBuilder when successful
func (m *InventoryRequestBuilder) WithUrl(rawUrl string)(*InventoryRequestBuilder) {
    return NewInventoryRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
