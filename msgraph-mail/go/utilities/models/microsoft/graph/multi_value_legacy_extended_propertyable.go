package graph
// MultiValueLegacyExtendedPropertyable 
type MultiValueLegacyExtendedPropertyable interface {
    Entityable
    GetValue()([]string)
    SetValue(value []string)()
}
