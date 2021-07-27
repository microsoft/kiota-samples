package users
import (
    "github.com/microsoft/kiota/abstractions/go/serialization"
    "time"
)
type Message struct {
    attachments []Attachment;
    bccRecipients []Recipient;
    body ItemBody;
    bodyPreview *string;
    ccRecipients []Recipient;
    conversationId *string;
    conversationIndex *string;
    extensions []Extension;
    flag FollowupFlag;
    from Recipient;
    hasAttachments *bool;
    importance Importance;
    inferenceClassification InferenceClassificationType;
    internetMessageHeaders []InternetMessageHeader;
    internetMessageId *string;
    isDeliveryReceiptRequested *bool;
    isDraft *bool;
    isRead *bool;
    isReadReceiptRequested *bool;
    multiValueExtendedProperties []MultiValueLegacyExtendedProperty;
    parentFolderId *string;
    receivedDateTime Time;
    replyTo []Recipient;
    sender Recipient;
    sentDateTime Time;
    singleValueExtendedProperties []SingleValueLegacyExtendedProperty;
    subject *string;
    toRecipients []Recipient;
    uniqueBody ItemBody;
    webLink *string;
}
