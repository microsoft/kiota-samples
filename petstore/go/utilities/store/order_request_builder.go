package store

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    idf4cc4a16f466bc4d40254b5ab3d20d0f80e475a6630c5e138f6c79181a5d398 "github.com/microsoft/kiota-samples/petstore/go/utilities/models"
)

// OrderRequestBuilder builds and executes requests for operations under \store\order
type OrderRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// OrderRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OrderRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewOrderRequestBuilderInternal instantiates a new OrderRequestBuilder and sets the default values.
func NewOrderRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OrderRequestBuilder) {
    m := &OrderRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/store/order";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewOrderRequestBuilder instantiates a new OrderRequestBuilder and sets the default values.
func NewOrderRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OrderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOrderRequestBuilderInternal(urlParams, requestAdapter)
}
// Post place an order for a pet
func (m *OrderRequestBuilder) Post(ctx context.Context, body idf4cc4a16f466bc4d40254b5ab3d20d0f80e475a6630c5e138f6c79181a5d398.Orderable, requestConfiguration *OrderRequestBuilderPostRequestConfiguration)(idf4cc4a16f466bc4d40254b5ab3d20d0f80e475a6630c5e138f6c79181a5d398.Orderable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, idf4cc4a16f466bc4d40254b5ab3d20d0f80e475a6630c5e138f6c79181a5d398.CreateOrderFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idf4cc4a16f466bc4d40254b5ab3d20d0f80e475a6630c5e138f6c79181a5d398.Orderable), nil
}
func (m *OrderRequestBuilder) ToPostRequestInformation(ctx context.Context, body idf4cc4a16f466bc4d40254b5ab3d20d0f80e475a6630c5e138f6c79181a5d398.Orderable, requestConfiguration *OrderRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
