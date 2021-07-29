package models

import (
    ie51e765764484004639cdfca9cb330e622fbf042d77be5ba8e7a002ffc8f8ed7 "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users"
)

type MessageRulePredicates struct {
    additionalData map[string]interface{};
    bodyContains []*string;
    bodyOrSubjectContains []*string;
    categories []*string;
    fromAddresses []Recipient;
    hasAttachments *bool;
    headerContains []*string;
    importance Importance;
    isApprovalRequest *bool;
    isAutomaticForward *bool;
    isAutomaticReply *bool;
    isEncrypted *bool;
    isMeetingRequest *bool;
    isMeetingResponse *bool;
    isNonDeliveryReport *bool;
    isPermissionControlled *bool;
    isReadReceipt *bool;
    isSigned *bool;
    isVoicemail *bool;
    messageActionFlag MessageActionFlag;
    notSentToMe *bool;
    recipientContains []*string;
    senderContains []*string;
    sensitivity Sensitivity;
    sentCcMe *bool;
    sentOnlyToMe *bool;
    sentToAddresses []Recipient;
    sentToMe *bool;
    sentToOrCcMe *bool;
    subjectContains []*string;
    withinSizeRange SizeRange;
}
