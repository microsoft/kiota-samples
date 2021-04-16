package graphjavav4.utilities.users;

import java.util.List;
public class Message extends OutlookItem {
    @javax.annotation.Nullable
    public List<Recipient> bccRecipients;
    @javax.annotation.Nullable
    public ItemBody body;
    @javax.annotation.Nullable
    public String bodyPreview;
    @javax.annotation.Nullable
    public List<Recipient> ccRecipients;
    @javax.annotation.Nullable
    public String conversationId;
    @javax.annotation.Nullable
    public String conversationIndex;
    @javax.annotation.Nullable
    public FollowupFlag flag;
    @javax.annotation.Nullable
    public Recipient from;
    @javax.annotation.Nullable
    public Boolean hasAttachments;
    @javax.annotation.Nullable
    public Importance importance;
    @javax.annotation.Nullable
    public InferenceClassificationType inferenceClassification;
    @javax.annotation.Nullable
    public List<InternetMessageHeader> internetMessageHeaders;
    @javax.annotation.Nullable
    public String internetMessageId;
    @javax.annotation.Nullable
    public Boolean isDeliveryReceiptRequested;
    @javax.annotation.Nullable
    public Boolean isDraft;
    @javax.annotation.Nullable
    public Boolean isRead;
    @javax.annotation.Nullable
    public Boolean isReadReceiptRequested;
    @javax.annotation.Nullable
    public String parentFolderId;
    @javax.annotation.Nullable
    public String receivedDateTime;
    @javax.annotation.Nullable
    public List<Recipient> replyTo;
    @javax.annotation.Nullable
    public Recipient sender;
    @javax.annotation.Nullable
    public String sentDateTime;
    @javax.annotation.Nullable
    public String subject;
    @javax.annotation.Nullable
    public List<Recipient> toRecipients;
    @javax.annotation.Nullable
    public ItemBody uniqueBody;
    @javax.annotation.Nullable
    public String webLink;
    @javax.annotation.Nullable
    public List<Attachment> attachments;
    @javax.annotation.Nullable
    public List<Extension> extensions;
    @javax.annotation.Nullable
    public List<MultiValueLegacyExtendedProperty> multiValueExtendedProperties;
    @javax.annotation.Nullable
    public List<SingleValueLegacyExtendedProperty> singleValueExtendedProperties;
}
