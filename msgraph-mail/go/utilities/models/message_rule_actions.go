package models
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
