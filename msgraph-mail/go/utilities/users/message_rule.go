package users
import (
    "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailFolders"
    "github.com/microsoft/kiota/abstractions/go/serialization"
)
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
