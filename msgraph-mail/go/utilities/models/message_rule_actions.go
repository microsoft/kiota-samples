package models

import (
    ie51e765764484004639cdfca9cb330e622fbf042d77be5ba8e7a002ffc8f8ed7 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users"
)

type MessageRuleActions struct {
    additionalData map[string]interface{};
    assignCategories []*string;
    copyToFolder *string;
    delete *bool;
    forwardAsAttachmentTo []Recipient;
    forwardTo []Recipient;
    markAsRead *bool;
    markImportance Importance;
    moveToFolder *string;
    permanentDelete *bool;
    redirectTo []Recipient;
    stopProcessingRules *bool;
}
