package mailfolders
import (
    "github.com/microsoft/kiota-samples/msgraph-mail/go/utilities/users"
    "github.com/microsoft/kiota/abstractions/go/serialization"
)
type MessageRulePredicates struct {
    additionalData Map<String, Object>;
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
