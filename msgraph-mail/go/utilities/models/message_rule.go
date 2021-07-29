package models
type MessageRule struct {
    actions MessageRuleActions;
    conditions MessageRulePredicates;
    displayName *string;
    exceptions MessageRulePredicates;
    hasError *bool;
    isEnabled *bool;
    isReadOnly *bool;
    sequence *int32;
}
