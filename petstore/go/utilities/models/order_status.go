package models
// Order Status
type Order_status int

const (
    PLACED_ORDER_STATUS Order_status = iota
    APPROVED_ORDER_STATUS
    DELIVERED_ORDER_STATUS
)

func (i Order_status) String() string {
    return []string{"placed", "approved", "delivered"}[i]
}
func ParseOrder_status(v string) (any, error) {
    result := PLACED_ORDER_STATUS
    switch v {
        case "placed":
            result = PLACED_ORDER_STATUS
        case "approved":
            result = APPROVED_ORDER_STATUS
        case "delivered":
            result = DELIVERED_ORDER_STATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeOrder_status(values []Order_status) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Order_status) isMultiValue() bool {
    return false
}
