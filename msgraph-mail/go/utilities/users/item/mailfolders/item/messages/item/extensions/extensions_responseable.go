package extensions

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/models/microsoft/graph"
)

// ExtensionsResponseable 
type ExtensionsResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdatanextLink()(*string)
    GetValue()([]i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.Extensionable)
    SetOdatanextLink(value *string)()
    SetValue(value []i2bf413bd639f9258700927995a2deeba4c8f0c1344d988e5d8e5959b0bb6f4ce.Extensionable)()
}
