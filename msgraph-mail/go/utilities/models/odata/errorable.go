package odata

import (
    iecd558b5bae11a51f88fb38e0fb1e378f141a48ce05284655fe1ad361c3389cf "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/models/odata/error"
)

// Errorable 
type Errorable interface {
    GetError()(iecd558b5bae11a51f88fb38e0fb1e378f141a48ce05284655fe1ad361c3389cf.Mainable)
    SetError(value iecd558b5bae11a51f88fb38e0fb1e378f141a48ce05284655fe1ad361c3389cf.Mainable)()
}
