package odata
// Errorable 
type Errorable interface {
    GetError()(Mainable)
    SetError(value Mainable)()
}
