package graphjavav4.utilities.users;

import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import java.time.OffsetDateTime;
import java.util.function.BiConsumer;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Objects;
public class Message extends OutlookItem implements Parsable {
    /** The fileAttachment and itemAttachment attachments for the message.  */
    @javax.annotation.Nullable
    public List<Attachment> attachments;
    /** The Bcc: recipients for the message.  */
    @javax.annotation.Nullable
    public List<Recipient> bccRecipients;
    @javax.annotation.Nullable
    public ItemBody body;
    /** The first 255 characters of the message body. It is in text format. If the message contains instances of mention, this property would contain a concatenation of these mentions as well.  */
    @javax.annotation.Nullable
    public String bodyPreview;
    /** The Cc: recipients for the message.  */
    @javax.annotation.Nullable
    public List<Recipient> ccRecipients;
    /** The ID of the conversation the email belongs to.  */
    @javax.annotation.Nullable
    public String conversationId;
    /** Indicates the position of the message within the conversation.  */
    @javax.annotation.Nullable
    public String conversationIndex;
    /** The collection of open extensions defined for the message. Nullable.  */
    @javax.annotation.Nullable
    public List<Extension> extensions;
    @javax.annotation.Nullable
    public FollowupFlag flag;
    @javax.annotation.Nullable
    public Recipient from;
    /** Indicates whether the message has attachments. This property doesn't include inline attachments, so if a message contains only inline attachments, this property is false. To verify the existence of inline attachments, parse the body property to look for a src attribute, such as <IMG src='cid:image001.jpg@01D26CD8.6C05F070'>.  */
    @javax.annotation.Nullable
    public Boolean hasAttachments;
    @javax.annotation.Nullable
    public Importance importance;
    @javax.annotation.Nullable
    public InferenceClassificationType inferenceClassification;
    /** A collection of message headers defined by RFC5322. The set includes message headers indicating the network path taken by a message from the sender to the recipient. It can also contain custom message headers that hold app data for the message.  Returned only on applying a $select query option. Read-only.  */
    @javax.annotation.Nullable
    public List<InternetMessageHeader> internetMessageHeaders;
    /** The message ID in the format specified by RFC2822.  */
    @javax.annotation.Nullable
    public String internetMessageId;
    /** Indicates whether a read receipt is requested for the message.  */
    @javax.annotation.Nullable
    public Boolean isDeliveryReceiptRequested;
    /** Indicates whether the message is a draft. A message is a draft if it hasn't been sent yet.  */
    @javax.annotation.Nullable
    public Boolean isDraft;
    /** Indicates whether the message has been read.  */
    @javax.annotation.Nullable
    public Boolean isRead;
    /** Indicates whether a read receipt is requested for the message.  */
    @javax.annotation.Nullable
    public Boolean isReadReceiptRequested;
    /** The collection of multi-value extended properties defined for the message. Nullable.  */
    @javax.annotation.Nullable
    public List<MultiValueLegacyExtendedProperty> multiValueExtendedProperties;
    /** The unique identifier for the message's parent mailFolder.  */
    @javax.annotation.Nullable
    public String parentFolderId;
    /** The date and time the message was received.  The date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.  */
    @javax.annotation.Nullable
    public OffsetDateTime receivedDateTime;
    /** The email addresses to use when replying.  */
    @javax.annotation.Nullable
    public List<Recipient> replyTo;
    @javax.annotation.Nullable
    public Recipient sender;
    /** The date and time the message was sent.  The date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.  */
    @javax.annotation.Nullable
    public OffsetDateTime sentDateTime;
    /** The collection of single-value extended properties defined for the message. Nullable.  */
    @javax.annotation.Nullable
    public List<SingleValueLegacyExtendedProperty> singleValueExtendedProperties;
    /** The subject of the message.  */
    @javax.annotation.Nullable
    public String subject;
    /** The To: recipients for the message.  */
    @javax.annotation.Nullable
    public List<Recipient> toRecipients;
    @javax.annotation.Nullable
    public ItemBody uniqueBody;
    /** The URL to open the message in Outlook on the web.You can append an ispopout argument to the end of the URL to change how the message is displayed. If ispopout is not present or if it is set to 1, then the message is shown in a popout window. If ispopout is set to 0, then the browser will show the message in the Outlook on the web review pane.The message will open in the browser if you are logged in to your mailbox via Outlook on the web. You will be prompted to login if you are not already logged in with the browser.This URL cannot be accessed from within an iFrame.  */
    @javax.annotation.Nullable
    public String webLink;
    /**
     * The serialization information for the current model
     * @return a Map<String, BiConsumer<T, ParseNode>>
     */
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getDeserializeFields() {
        final Map<String, BiConsumer<T, ParseNode>> fields = new HashMap<>(super.getDeserializeFields());
        fields.put("attachments", (o, n) -> { ((Message)o).attachments = n.getCollectionOfObjectValues(Attachment.class); });
        fields.put("bccRecipients", (o, n) -> { ((Message)o).bccRecipients = n.getCollectionOfObjectValues(Recipient.class); });
        fields.put("body", (o, n) -> { ((Message)o).body = n.getObjectValue(ItemBody.class); });
        fields.put("bodyPreview", (o, n) -> { ((Message)o).bodyPreview = n.getStringValue(); });
        fields.put("ccRecipients", (o, n) -> { ((Message)o).ccRecipients = n.getCollectionOfObjectValues(Recipient.class); });
        fields.put("conversationId", (o, n) -> { ((Message)o).conversationId = n.getStringValue(); });
        fields.put("conversationIndex", (o, n) -> { ((Message)o).conversationIndex = n.getStringValue(); });
        fields.put("extensions", (o, n) -> { ((Message)o).extensions = n.getCollectionOfObjectValues(Extension.class); });
        fields.put("flag", (o, n) -> { ((Message)o).flag = n.getObjectValue(FollowupFlag.class); });
        fields.put("from", (o, n) -> { ((Message)o).from = n.getObjectValue(Recipient.class); });
        fields.put("hasAttachments", (o, n) -> { ((Message)o).hasAttachments = n.getBooleanValue(); });
        fields.put("importance", (o, n) -> { ((Message)o).importance = n.getEnumValue(Importance.class); });
        fields.put("inferenceClassification", (o, n) -> { ((Message)o).inferenceClassification = n.getEnumValue(InferenceClassificationType.class); });
        fields.put("internetMessageHeaders", (o, n) -> { ((Message)o).internetMessageHeaders = n.getCollectionOfObjectValues(InternetMessageHeader.class); });
        fields.put("internetMessageId", (o, n) -> { ((Message)o).internetMessageId = n.getStringValue(); });
        fields.put("isDeliveryReceiptRequested", (o, n) -> { ((Message)o).isDeliveryReceiptRequested = n.getBooleanValue(); });
        fields.put("isDraft", (o, n) -> { ((Message)o).isDraft = n.getBooleanValue(); });
        fields.put("isRead", (o, n) -> { ((Message)o).isRead = n.getBooleanValue(); });
        fields.put("isReadReceiptRequested", (o, n) -> { ((Message)o).isReadReceiptRequested = n.getBooleanValue(); });
        fields.put("multiValueExtendedProperties", (o, n) -> { ((Message)o).multiValueExtendedProperties = n.getCollectionOfObjectValues(MultiValueLegacyExtendedProperty.class); });
        fields.put("parentFolderId", (o, n) -> { ((Message)o).parentFolderId = n.getStringValue(); });
        fields.put("receivedDateTime", (o, n) -> { ((Message)o).receivedDateTime = n.getOffsetDateTimeValue(); });
        fields.put("replyTo", (o, n) -> { ((Message)o).replyTo = n.getCollectionOfObjectValues(Recipient.class); });
        fields.put("sender", (o, n) -> { ((Message)o).sender = n.getObjectValue(Recipient.class); });
        fields.put("sentDateTime", (o, n) -> { ((Message)o).sentDateTime = n.getOffsetDateTimeValue(); });
        fields.put("singleValueExtendedProperties", (o, n) -> { ((Message)o).singleValueExtendedProperties = n.getCollectionOfObjectValues(SingleValueLegacyExtendedProperty.class); });
        fields.put("subject", (o, n) -> { ((Message)o).subject = n.getStringValue(); });
        fields.put("toRecipients", (o, n) -> { ((Message)o).toRecipients = n.getCollectionOfObjectValues(Recipient.class); });
        fields.put("uniqueBody", (o, n) -> { ((Message)o).uniqueBody = n.getObjectValue(ItemBody.class); });
        fields.put("webLink", (o, n) -> { ((Message)o).webLink = n.getStringValue(); });
        return fields;
    }
    /**
     * Serialiazes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @return a void
     */
    public void serialize(@javax.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        super.serialize(writer);
        writer.writeCollectionOfObjectValues("attachments", attachments);
        writer.writeCollectionOfObjectValues("bccRecipients", bccRecipients);
        writer.writeObjectValue("body", body);
        writer.writeStringValue("bodyPreview", bodyPreview);
        writer.writeCollectionOfObjectValues("ccRecipients", ccRecipients);
        writer.writeStringValue("conversationId", conversationId);
        writer.writeStringValue("conversationIndex", conversationIndex);
        writer.writeCollectionOfObjectValues("extensions", extensions);
        writer.writeObjectValue("flag", flag);
        writer.writeObjectValue("from", from);
        writer.writeBooleanValue("hasAttachments", hasAttachments);
        writer.writeEnumValue("importance", importance);
        writer.writeEnumValue("inferenceClassification", inferenceClassification);
        writer.writeCollectionOfObjectValues("internetMessageHeaders", internetMessageHeaders);
        writer.writeStringValue("internetMessageId", internetMessageId);
        writer.writeBooleanValue("isDeliveryReceiptRequested", isDeliveryReceiptRequested);
        writer.writeBooleanValue("isDraft", isDraft);
        writer.writeBooleanValue("isRead", isRead);
        writer.writeBooleanValue("isReadReceiptRequested", isReadReceiptRequested);
        writer.writeCollectionOfObjectValues("multiValueExtendedProperties", multiValueExtendedProperties);
        writer.writeStringValue("parentFolderId", parentFolderId);
        writer.writeOffsetDateTimeValue("receivedDateTime", receivedDateTime);
        writer.writeCollectionOfObjectValues("replyTo", replyTo);
        writer.writeObjectValue("sender", sender);
        writer.writeOffsetDateTimeValue("sentDateTime", sentDateTime);
        writer.writeCollectionOfObjectValues("singleValueExtendedProperties", singleValueExtendedProperties);
        writer.writeStringValue("subject", subject);
        writer.writeCollectionOfObjectValues("toRecipients", toRecipients);
        writer.writeObjectValue("uniqueBody", uniqueBody);
        writer.writeStringValue("webLink", webLink);
    }
}
