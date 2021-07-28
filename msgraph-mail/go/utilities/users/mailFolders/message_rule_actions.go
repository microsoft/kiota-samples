package mailfolders

import (
    ie51e765764484004639cdfca9cb330e622fbf042d77be5ba8e7a002ffc8f8ed7 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users"
)

type MessageRuleActions struct {
    additionalData map[string]interface{};
    assignCategories []*string;
    copyToFolder *string;
    delete *bool;
    forwardAsAttachmentTo []ie51e765764484004639cdfca9cb330e622fbf042d77be5ba8e7a002ffc8f8ed7.Recipient;
    forwardTo []ie51e765764484004639cdfca9cb330e622fbf042d77be5ba8e7a002ffc8f8ed7.Recipient;
    markAsRead *bool;
    markImportance ie51e765764484004639cdfca9cb330e622fbf042d77be5ba8e7a002ffc8f8ed7.Importance;
    moveToFolder *string;
    permanentDelete *bool;
    redirectTo []ie51e765764484004639cdfca9cb330e622fbf042d77be5ba8e7a002ffc8f8ed7.Recipient;
    stopProcessingRules *bool;
}
