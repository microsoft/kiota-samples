package graph
// MessageRuleable 
type MessageRuleable interface {
    Entityable
    GetActions()(MessageRuleActionsable)
    GetConditions()(MessageRulePredicatesable)
    GetDisplayName()(*string)
    GetExceptions()(MessageRulePredicatesable)
    GetHasError()(*bool)
    GetIsEnabled()(*bool)
    GetIsReadOnly()(*bool)
    GetSequence()(*int32)
    SetActions(value MessageRuleActionsable)()
    SetConditions(value MessageRulePredicatesable)()
    SetDisplayName(value *string)()
    SetExceptions(value MessageRulePredicatesable)()
    SetHasError(value *bool)()
    SetIsEnabled(value *bool)()
    SetIsReadOnly(value *bool)()
    SetSequence(value *int32)()
}
