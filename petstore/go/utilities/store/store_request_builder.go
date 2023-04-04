package store

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// StoreRequestBuilder builds and executes requests for operations under \store
type StoreRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewStoreRequestBuilderInternal instantiates a new StoreRequestBuilder and sets the default values.
func NewStoreRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StoreRequestBuilder) {
    m := &StoreRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/store";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewStoreRequestBuilder instantiates a new StoreRequestBuilder and sets the default values.
func NewStoreRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StoreRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewStoreRequestBuilderInternal(urlParams, requestAdapter)
}
// Inventory the inventory property
func (m *StoreRequestBuilder) Inventory()(*InventoryRequestBuilder) {
    return NewInventoryRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Order the order property
func (m *StoreRequestBuilder) Order()(*OrderRequestBuilder) {
    return NewOrderRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// OrderById gets an item from the github.com/microsoft/kiota-samples/petstore/go/utilities/.store.order.item collection
func (m *StoreRequestBuilder) OrderById(id string)(*OrderWithOrderItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["orderId"] = id
    }
    return NewOrderWithOrderItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
