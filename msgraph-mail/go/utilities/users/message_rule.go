package users

import (
    iac24930f553b5045ae4fd27ac1cd6de43babde38cdb777adf1d182559763183a "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailfolders"
)

type MessageRule struct {
    actions iac24930f553b5045ae4fd27ac1cd6de43babde38cdb777adf1d182559763183a.MessageRuleActions;
    conditions iac24930f553b5045ae4fd27ac1cd6de43babde38cdb777adf1d182559763183a.MessageRulePredicates;
    displayName *string;
    exceptions iac24930f553b5045ae4fd27ac1cd6de43babde38cdb777adf1d182559763183a.MessageRulePredicates;
    hasError *bool;
    isEnabled *bool;
    isReadOnly *bool;
    sequence *int32;
}
