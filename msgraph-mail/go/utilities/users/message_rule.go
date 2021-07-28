package users

import (
    ia254bc7eebe487216682ef10c3748ccc8d561219b8bde857ab0c71fc56c4a3a7 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users/mailFolders"
)

type MessageRule struct {
    actions ia254bc7eebe487216682ef10c3748ccc8d561219b8bde857ab0c71fc56c4a3a7.MessageRuleActions;
    conditions ia254bc7eebe487216682ef10c3748ccc8d561219b8bde857ab0c71fc56c4a3a7.MessageRulePredicates;
    displayName *string;
    exceptions ia254bc7eebe487216682ef10c3748ccc8d561219b8bde857ab0c71fc56c4a3a7.MessageRulePredicates;
    hasError *bool;
    isEnabled *bool;
    isReadOnly *bool;
    sequence *int32;
}
