package store

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// StoreRequestBuilder builds and executes requests for operations under \store
type StoreRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewStoreRequestBuilderInternal instantiates a new StoreRequestBuilder and sets the default values.
func NewStoreRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StoreRequestBuilder) {
    m := &StoreRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/store", pathParameters),
    }
    return m
}
// NewStoreRequestBuilder instantiates a new StoreRequestBuilder and sets the default values.
func NewStoreRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StoreRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewStoreRequestBuilderInternal(urlParams, requestAdapter)
}
// Inventory the inventory property
// returns a *InventoryRequestBuilder when successful
func (m *StoreRequestBuilder) Inventory()(*InventoryRequestBuilder) {
    return NewInventoryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Order the order property
// returns a *OrderRequestBuilder when successful
func (m *StoreRequestBuilder) Order()(*OrderRequestBuilder) {
    return NewOrderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
