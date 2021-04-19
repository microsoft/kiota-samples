package graphjavav4.utilities.users;

import java.util.Objects;
import com.microsoft.kiota.serialization.SerializationWriter;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.Parsable;
import java.util.function.BiConsumer;
import java.util.Map;
import java.util.HashMap;
import java.util.List;
public class Message extends OutlookItem implements Parsable {
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
    public void serialize(@javax.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        super.serialize(writer);
        writer.writeCollectionOfObjectValues("bccRecipients", bccRecipients);
        writer.writeObjectValue("body", body);
        writer.writeStringValue("bodyPreview", bodyPreview);
        writer.writeCollectionOfObjectValues("ccRecipients", ccRecipients);
        writer.writeStringValue("conversationId", conversationId);
        writer.writeStringValue("conversationIndex", conversationIndex);
        writer.writeObjectValue("flag", flag);
        writer.writeObjectValue("from", from);
        writer.writeBooleanValue("hasAttachments", hasAttachments);
        writer.writeObjectValue("importance", importance);
        writer.writeObjectValue("inferenceClassification", inferenceClassification);
        writer.writeCollectionOfObjectValues("internetMessageHeaders", internetMessageHeaders);
        writer.writeStringValue("internetMessageId", internetMessageId);
        writer.writeBooleanValue("isDeliveryReceiptRequested", isDeliveryReceiptRequested);
        writer.writeBooleanValue("isDraft", isDraft);
        writer.writeBooleanValue("isRead", isRead);
        writer.writeBooleanValue("isReadReceiptRequested", isReadReceiptRequested);
        writer.writeStringValue("parentFolderId", parentFolderId);
        writer.writeStringValue("receivedDateTime", receivedDateTime);
        writer.writeCollectionOfObjectValues("replyTo", replyTo);
        writer.writeObjectValue("sender", sender);
        writer.writeStringValue("sentDateTime", sentDateTime);
        writer.writeStringValue("subject", subject);
        writer.writeCollectionOfObjectValues("toRecipients", toRecipients);
        writer.writeObjectValue("uniqueBody", uniqueBody);
        writer.writeStringValue("webLink", webLink);
        writer.writeCollectionOfObjectValues("attachments", attachments);
        writer.writeCollectionOfObjectValues("extensions", extensions);
        writer.writeCollectionOfObjectValues("multiValueExtendedProperties", multiValueExtendedProperties);
        writer.writeCollectionOfObjectValues("singleValueExtendedProperties", singleValueExtendedProperties);
    }
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getDeserializeFields() {
        final Map<String, BiConsumer<T, ParseNode>> fields = new HashMap<>(super.getDeserializeFields());
        fields.put("bccRecipients", (o, n) -> { ((Message)o).bccRecipients = n.getCollectionOfObjectValues(Recipient.class); });
        fields.put("body", (o, n) -> { ((Message)o).body = n.getObjectValue(ItemBody.class); });
        fields.put("bodyPreview", (o, n) -> { ((Message)o).bodyPreview = n.getStringValue(); });
        fields.put("ccRecipients", (o, n) -> { ((Message)o).ccRecipients = n.getCollectionOfObjectValues(Recipient.class); });
        fields.put("conversationId", (o, n) -> { ((Message)o).conversationId = n.getStringValue(); });
        fields.put("conversationIndex", (o, n) -> { ((Message)o).conversationIndex = n.getStringValue(); });
        fields.put("flag", (o, n) -> { ((Message)o).flag = n.getObjectValue(FollowupFlag.class); });
        fields.put("from", (o, n) -> { ((Message)o).from = n.getObjectValue(Recipient.class); });
        fields.put("hasAttachments", (o, n) -> { ((Message)o).hasAttachments = n.getBooleanValue(); });
        fields.put("importance", (o, n) -> { ((Message)o).importance = n.getObjectValue(Importance.class); });
        fields.put("inferenceClassification", (o, n) -> { ((Message)o).inferenceClassification = n.getObjectValue(InferenceClassificationType.class); });
        fields.put("internetMessageHeaders", (o, n) -> { ((Message)o).internetMessageHeaders = n.getCollectionOfObjectValues(InternetMessageHeader.class); });
        fields.put("internetMessageId", (o, n) -> { ((Message)o).internetMessageId = n.getStringValue(); });
        fields.put("isDeliveryReceiptRequested", (o, n) -> { ((Message)o).isDeliveryReceiptRequested = n.getBooleanValue(); });
        fields.put("isDraft", (o, n) -> { ((Message)o).isDraft = n.getBooleanValue(); });
        fields.put("isRead", (o, n) -> { ((Message)o).isRead = n.getBooleanValue(); });
        fields.put("isReadReceiptRequested", (o, n) -> { ((Message)o).isReadReceiptRequested = n.getBooleanValue(); });
        fields.put("parentFolderId", (o, n) -> { ((Message)o).parentFolderId = n.getStringValue(); });
        fields.put("receivedDateTime", (o, n) -> { ((Message)o).receivedDateTime = n.getStringValue(); });
        fields.put("replyTo", (o, n) -> { ((Message)o).replyTo = n.getCollectionOfObjectValues(Recipient.class); });
        fields.put("sender", (o, n) -> { ((Message)o).sender = n.getObjectValue(Recipient.class); });
        fields.put("sentDateTime", (o, n) -> { ((Message)o).sentDateTime = n.getStringValue(); });
        fields.put("subject", (o, n) -> { ((Message)o).subject = n.getStringValue(); });
        fields.put("toRecipients", (o, n) -> { ((Message)o).toRecipients = n.getCollectionOfObjectValues(Recipient.class); });
        fields.put("uniqueBody", (o, n) -> { ((Message)o).uniqueBody = n.getObjectValue(ItemBody.class); });
        fields.put("webLink", (o, n) -> { ((Message)o).webLink = n.getStringValue(); });
        fields.put("attachments", (o, n) -> { ((Message)o).attachments = n.getCollectionOfObjectValues(Attachment.class); });
        fields.put("extensions", (o, n) -> { ((Message)o).extensions = n.getCollectionOfObjectValues(Extension.class); });
        fields.put("multiValueExtendedProperties", (o, n) -> { ((Message)o).multiValueExtendedProperties = n.getCollectionOfObjectValues(MultiValueLegacyExtendedProperty.class); });
        fields.put("singleValueExtendedProperties", (o, n) -> { ((Message)o).singleValueExtendedProperties = n.getCollectionOfObjectValues(SingleValueLegacyExtendedProperty.class); });
        return fields;
    }
}
