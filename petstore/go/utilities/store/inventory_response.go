package store

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InventoryResponse 
// Deprecated: This class is obsolete. Use inventoryGetResponse instead.
type InventoryResponse struct {
    InventoryGetResponse
}
// NewInventoryResponse instantiates a new InventoryResponse and sets the default values.
func NewInventoryResponse()(*InventoryResponse) {
    m := &InventoryResponse{
        InventoryGetResponse: *NewInventoryGetResponse(),
    }
    return m
}
// CreateInventoryResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInventoryResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInventoryResponse(), nil
}
// InventoryResponseable 
// Deprecated: This class is obsolete. Use inventoryGetResponse instead.
type InventoryResponseable interface {
    InventoryGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
