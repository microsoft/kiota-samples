package findbystatus
type GetStatusQueryParameterType int

const (
    AVAILABLE_GETSTATUSQUERYPARAMETERTYPE GetStatusQueryParameterType = iota
    PENDING_GETSTATUSQUERYPARAMETERTYPE
    SOLD_GETSTATUSQUERYPARAMETERTYPE
)

func (i GetStatusQueryParameterType) String() string {
    return []string{"available", "pending", "sold"}[i]
}
func ParseGetStatusQueryParameterType(v string) (any, error) {
    result := AVAILABLE_GETSTATUSQUERYPARAMETERTYPE
    switch v {
        case "available":
            result = AVAILABLE_GETSTATUSQUERYPARAMETERTYPE
        case "pending":
            result = PENDING_GETSTATUSQUERYPARAMETERTYPE
        case "sold":
            result = SOLD_GETSTATUSQUERYPARAMETERTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeGetStatusQueryParameterType(values []GetStatusQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetStatusQueryParameterType) isMultiValue() bool {
    return false
}
