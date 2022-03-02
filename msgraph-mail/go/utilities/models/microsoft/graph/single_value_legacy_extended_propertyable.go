package graph
// SingleValueLegacyExtendedPropertyable 
type SingleValueLegacyExtendedPropertyable interface {
    Entityable
    GetValue()(*string)
    SetValue(value *string)()
}
