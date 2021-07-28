package users

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
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
    receivedDateTime i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    replyTo []Recipient;
    sender Recipient;
    sentDateTime i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    singleValueExtendedProperties []SingleValueLegacyExtendedProperty;
    subject *string;
    toRecipients []Recipient;
    uniqueBody ItemBody;
    webLink *string;
}
