package graphjavav4.utilities.models.microsoft.graph;

import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import java.time.OffsetDateTime;
import java.util.function.BiConsumer;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
public class Message extends OutlookItem implements Parsable {
    /** The fileAttachment and itemAttachment attachments for the message.  */
    private java.util.List<Attachment> _attachments;
    /** The Bcc: recipients for the message.  */
    private java.util.List<Recipient> _bccRecipients;
    private ItemBody _body;
    /** The first 255 characters of the message body. It is in text format. If the message contains instances of mention, this property would contain a concatenation of these mentions as well.  */
    private String _bodyPreview;
    /** The Cc: recipients for the message.  */
    private java.util.List<Recipient> _ccRecipients;
    /** The ID of the conversation the email belongs to.  */
    private String _conversationId;
    /** Indicates the position of the message within the conversation.  */
    private byte[] _conversationIndex;
    /** The collection of open extensions defined for the message. Nullable.  */
    private java.util.List<Extension> _extensions;
    private FollowupFlag _flag;
    private Recipient _from;
    /** Indicates whether the message has attachments. This property doesn't include inline attachments, so if a message contains only inline attachments, this property is false. To verify the existence of inline attachments, parse the body property to look for a src attribute, such as <IMG src='cid:image001.jpg@01D26CD8.6C05F070'>.  */
    private Boolean _hasAttachments;
    private Importance _importance;
    private InferenceClassificationType _inferenceClassification;
    /** A collection of message headers defined by RFC5322. The set includes message headers indicating the network path taken by a message from the sender to the recipient. It can also contain custom message headers that hold app data for the message.  Returned only on applying a $select query option. Read-only.  */
    private java.util.List<InternetMessageHeader> _internetMessageHeaders;
    /** The message ID in the format specified by RFC2822.  */
    private String _internetMessageId;
    /** Indicates whether a read receipt is requested for the message.  */
    private Boolean _isDeliveryReceiptRequested;
    /** Indicates whether the message is a draft. A message is a draft if it hasn't been sent yet.  */
    private Boolean _isDraft;
    /** Indicates whether the message has been read.  */
    private Boolean _isRead;
    /** Indicates whether a read receipt is requested for the message.  */
    private Boolean _isReadReceiptRequested;
    /** The collection of multi-value extended properties defined for the message. Nullable.  */
    private java.util.List<MultiValueLegacyExtendedProperty> _multiValueExtendedProperties;
    /** The unique identifier for the message's parent mailFolder.  */
    private String _parentFolderId;
    /** The date and time the message was received.  The date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.  */
    private OffsetDateTime _receivedDateTime;
    /** The email addresses to use when replying.  */
    private java.util.List<Recipient> _replyTo;
    private Recipient _sender;
    /** The date and time the message was sent.  The date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.  */
    private OffsetDateTime _sentDateTime;
    /** The collection of single-value extended properties defined for the message. Nullable.  */
    private java.util.List<SingleValueLegacyExtendedProperty> _singleValueExtendedProperties;
    /** The subject of the message.  */
    private String _subject;
    /** The To: recipients for the message.  */
    private java.util.List<Recipient> _toRecipients;
    private ItemBody _uniqueBody;
    /** The URL to open the message in Outlook on the web.You can append an ispopout argument to the end of the URL to change how the message is displayed. If ispopout is not present or if it is set to 1, then the message is shown in a popout window. If ispopout is set to 0, then the browser will show the message in the Outlook on the web review pane.The message will open in the browser if you are logged in to your mailbox via Outlook on the web. You will be prompted to login if you are not already logged in with the browser.This URL cannot be accessed from within an iFrame.  */
    private String _webLink;
    /**
     * Instantiates a new message and sets the default values.
     * @return a void
     */
    public Message() {
        super();
    }
    /**
     * Gets the attachments property value. The fileAttachment and itemAttachment attachments for the message.
     * @return a attachment
     */
    @javax.annotation.Nullable
    public java.util.List<Attachment> getAttachments() {
        return this._attachments;
    }
    /**
     * Gets the bccRecipients property value. The Bcc: recipients for the message.
     * @return a recipient
     */
    @javax.annotation.Nullable
    public java.util.List<Recipient> getBccRecipients() {
        return this._bccRecipients;
    }
    /**
     * Gets the body property value. 
     * @return a itemBody
     */
    @javax.annotation.Nullable
    public ItemBody getBody() {
        return this._body;
    }
    /**
     * Gets the bodyPreview property value. The first 255 characters of the message body. It is in text format. If the message contains instances of mention, this property would contain a concatenation of these mentions as well.
     * @return a string
     */
    @javax.annotation.Nullable
    public String getBodyPreview() {
        return this._bodyPreview;
    }
    /**
     * Gets the ccRecipients property value. The Cc: recipients for the message.
     * @return a recipient
     */
    @javax.annotation.Nullable
    public java.util.List<Recipient> getCcRecipients() {
        return this._ccRecipients;
    }
    /**
     * Gets the conversationId property value. The ID of the conversation the email belongs to.
     * @return a string
     */
    @javax.annotation.Nullable
    public String getConversationId() {
        return this._conversationId;
    }
    /**
     * Gets the conversationIndex property value. Indicates the position of the message within the conversation.
     * @return a binary
     */
    @javax.annotation.Nullable
    public byte[] getConversationIndex() {
        return this._conversationIndex;
    }
    /**
     * Gets the extensions property value. The collection of open extensions defined for the message. Nullable.
     * @return a extension
     */
    @javax.annotation.Nullable
    public java.util.List<Extension> getExtensions() {
        return this._extensions;
    }
    /**
     * Gets the flag property value. 
     * @return a followupFlag
     */
    @javax.annotation.Nullable
    public FollowupFlag getFlag() {
        return this._flag;
    }
    /**
     * Gets the from property value. 
     * @return a recipient
     */
    @javax.annotation.Nullable
    public Recipient getFrom() {
        return this._from;
    }
    /**
     * Gets the hasAttachments property value. Indicates whether the message has attachments. This property doesn't include inline attachments, so if a message contains only inline attachments, this property is false. To verify the existence of inline attachments, parse the body property to look for a src attribute, such as <IMG src='cid:image001.jpg@01D26CD8.6C05F070'>.
     * @return a boolean
     */
    @javax.annotation.Nullable
    public Boolean getHasAttachments() {
        return this._hasAttachments;
    }
    /**
     * Gets the importance property value. 
     * @return a importance
     */
    @javax.annotation.Nullable
    public Importance getImportance() {
        return this._importance;
    }
    /**
     * Gets the inferenceClassification property value. 
     * @return a inferenceClassificationType
     */
    @javax.annotation.Nullable
    public InferenceClassificationType getInferenceClassification() {
        return this._inferenceClassification;
    }
    /**
     * Gets the internetMessageHeaders property value. A collection of message headers defined by RFC5322. The set includes message headers indicating the network path taken by a message from the sender to the recipient. It can also contain custom message headers that hold app data for the message.  Returned only on applying a $select query option. Read-only.
     * @return a internetMessageHeader
     */
    @javax.annotation.Nullable
    public java.util.List<InternetMessageHeader> getInternetMessageHeaders() {
        return this._internetMessageHeaders;
    }
    /**
     * Gets the internetMessageId property value. The message ID in the format specified by RFC2822.
     * @return a string
     */
    @javax.annotation.Nullable
    public String getInternetMessageId() {
        return this._internetMessageId;
    }
    /**
     * Gets the isDeliveryReceiptRequested property value. Indicates whether a read receipt is requested for the message.
     * @return a boolean
     */
    @javax.annotation.Nullable
    public Boolean getIsDeliveryReceiptRequested() {
        return this._isDeliveryReceiptRequested;
    }
    /**
     * Gets the isDraft property value. Indicates whether the message is a draft. A message is a draft if it hasn't been sent yet.
     * @return a boolean
     */
    @javax.annotation.Nullable
    public Boolean getIsDraft() {
        return this._isDraft;
    }
    /**
     * Gets the isRead property value. Indicates whether the message has been read.
     * @return a boolean
     */
    @javax.annotation.Nullable
    public Boolean getIsRead() {
        return this._isRead;
    }
    /**
     * Gets the isReadReceiptRequested property value. Indicates whether a read receipt is requested for the message.
     * @return a boolean
     */
    @javax.annotation.Nullable
    public Boolean getIsReadReceiptRequested() {
        return this._isReadReceiptRequested;
    }
    /**
     * Gets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the message. Nullable.
     * @return a multiValueLegacyExtendedProperty
     */
    @javax.annotation.Nullable
    public java.util.List<MultiValueLegacyExtendedProperty> getMultiValueExtendedProperties() {
        return this._multiValueExtendedProperties;
    }
    /**
     * Gets the parentFolderId property value. The unique identifier for the message's parent mailFolder.
     * @return a string
     */
    @javax.annotation.Nullable
    public String getParentFolderId() {
        return this._parentFolderId;
    }
    /**
     * Gets the receivedDateTime property value. The date and time the message was received.  The date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
     * @return a OffsetDateTime
     */
    @javax.annotation.Nullable
    public OffsetDateTime getReceivedDateTime() {
        return this._receivedDateTime;
    }
    /**
     * Gets the replyTo property value. The email addresses to use when replying.
     * @return a recipient
     */
    @javax.annotation.Nullable
    public java.util.List<Recipient> getReplyTo() {
        return this._replyTo;
    }
    /**
     * Gets the sender property value. 
     * @return a recipient
     */
    @javax.annotation.Nullable
    public Recipient getSender() {
        return this._sender;
    }
    /**
     * Gets the sentDateTime property value. The date and time the message was sent.  The date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
     * @return a OffsetDateTime
     */
    @javax.annotation.Nullable
    public OffsetDateTime getSentDateTime() {
        return this._sentDateTime;
    }
    /**
     * Gets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the message. Nullable.
     * @return a singleValueLegacyExtendedProperty
     */
    @javax.annotation.Nullable
    public java.util.List<SingleValueLegacyExtendedProperty> getSingleValueExtendedProperties() {
        return this._singleValueExtendedProperties;
    }
    /**
     * Gets the subject property value. The subject of the message.
     * @return a string
     */
    @javax.annotation.Nullable
    public String getSubject() {
        return this._subject;
    }
    /**
     * Gets the toRecipients property value. The To: recipients for the message.
     * @return a recipient
     */
    @javax.annotation.Nullable
    public java.util.List<Recipient> getToRecipients() {
        return this._toRecipients;
    }
    /**
     * Gets the uniqueBody property value. 
     * @return a itemBody
     */
    @javax.annotation.Nullable
    public ItemBody getUniqueBody() {
        return this._uniqueBody;
    }
    /**
     * Gets the webLink property value. The URL to open the message in Outlook on the web.You can append an ispopout argument to the end of the URL to change how the message is displayed. If ispopout is not present or if it is set to 1, then the message is shown in a popout window. If ispopout is set to 0, then the browser will show the message in the Outlook on the web review pane.The message will open in the browser if you are logged in to your mailbox via Outlook on the web. You will be prompted to login if you are not already logged in with the browser.This URL cannot be accessed from within an iFrame.
     * @return a string
     */
    @javax.annotation.Nullable
    public String getWebLink() {
        return this._webLink;
    }
    /**
     * The deserialization information for the current model
     * @return a Map<String, BiConsumer<T, ParseNode>>
     */
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getFieldDeserializers() {
        return new HashMap<>(super.getFieldDeserializers()) {{
            this.put("attachments", (o, n) -> { ((Message)o).setAttachments(n.getCollectionOfObjectValues(Attachment.class)); });
            this.put("bccRecipients", (o, n) -> { ((Message)o).setBccRecipients(n.getCollectionOfObjectValues(Recipient.class)); });
            this.put("body", (o, n) -> { ((Message)o).setBody(n.getObjectValue(ItemBody.class)); });
            this.put("bodyPreview", (o, n) -> { ((Message)o).setBodyPreview(n.getStringValue()); });
            this.put("ccRecipients", (o, n) -> { ((Message)o).setCcRecipients(n.getCollectionOfObjectValues(Recipient.class)); });
            this.put("conversationId", (o, n) -> { ((Message)o).setConversationId(n.getStringValue()); });
            this.put("conversationIndex", (o, n) -> { ((Message)o).setConversationIndex(n.getByteArrayValue()); });
            this.put("extensions", (o, n) -> { ((Message)o).setExtensions(n.getCollectionOfObjectValues(Extension.class)); });
            this.put("flag", (o, n) -> { ((Message)o).setFlag(n.getObjectValue(FollowupFlag.class)); });
            this.put("from", (o, n) -> { ((Message)o).setFrom(n.getObjectValue(Recipient.class)); });
            this.put("hasAttachments", (o, n) -> { ((Message)o).setHasAttachments(n.getBooleanValue()); });
            this.put("importance", (o, n) -> { ((Message)o).setImportance(n.getEnumValue(Importance.class)); });
            this.put("inferenceClassification", (o, n) -> { ((Message)o).setInferenceClassification(n.getEnumValue(InferenceClassificationType.class)); });
            this.put("internetMessageHeaders", (o, n) -> { ((Message)o).setInternetMessageHeaders(n.getCollectionOfObjectValues(InternetMessageHeader.class)); });
            this.put("internetMessageId", (o, n) -> { ((Message)o).setInternetMessageId(n.getStringValue()); });
            this.put("isDeliveryReceiptRequested", (o, n) -> { ((Message)o).setIsDeliveryReceiptRequested(n.getBooleanValue()); });
            this.put("isDraft", (o, n) -> { ((Message)o).setIsDraft(n.getBooleanValue()); });
            this.put("isRead", (o, n) -> { ((Message)o).setIsRead(n.getBooleanValue()); });
            this.put("isReadReceiptRequested", (o, n) -> { ((Message)o).setIsReadReceiptRequested(n.getBooleanValue()); });
            this.put("multiValueExtendedProperties", (o, n) -> { ((Message)o).setMultiValueExtendedProperties(n.getCollectionOfObjectValues(MultiValueLegacyExtendedProperty.class)); });
            this.put("parentFolderId", (o, n) -> { ((Message)o).setParentFolderId(n.getStringValue()); });
            this.put("receivedDateTime", (o, n) -> { ((Message)o).setReceivedDateTime(n.getOffsetDateTimeValue()); });
            this.put("replyTo", (o, n) -> { ((Message)o).setReplyTo(n.getCollectionOfObjectValues(Recipient.class)); });
            this.put("sender", (o, n) -> { ((Message)o).setSender(n.getObjectValue(Recipient.class)); });
            this.put("sentDateTime", (o, n) -> { ((Message)o).setSentDateTime(n.getOffsetDateTimeValue()); });
            this.put("singleValueExtendedProperties", (o, n) -> { ((Message)o).setSingleValueExtendedProperties(n.getCollectionOfObjectValues(SingleValueLegacyExtendedProperty.class)); });
            this.put("subject", (o, n) -> { ((Message)o).setSubject(n.getStringValue()); });
            this.put("toRecipients", (o, n) -> { ((Message)o).setToRecipients(n.getCollectionOfObjectValues(Recipient.class)); });
            this.put("uniqueBody", (o, n) -> { ((Message)o).setUniqueBody(n.getObjectValue(ItemBody.class)); });
            this.put("webLink", (o, n) -> { ((Message)o).setWebLink(n.getStringValue()); });
        }};
    }
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @return a void
     */
    public void serialize(@javax.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        super.serialize(writer);
        writer.writeCollectionOfObjectValues("attachments", this.getAttachments());
        writer.writeCollectionOfObjectValues("bccRecipients", this.getBccRecipients());
        writer.writeObjectValue("body", this.getBody());
        writer.writeStringValue("bodyPreview", this.getBodyPreview());
        writer.writeCollectionOfObjectValues("ccRecipients", this.getCcRecipients());
        writer.writeStringValue("conversationId", this.getConversationId());
        writer.writeByteArrayValue("conversationIndex", this.getConversationIndex());
        writer.writeCollectionOfObjectValues("extensions", this.getExtensions());
        writer.writeObjectValue("flag", this.getFlag());
        writer.writeObjectValue("from", this.getFrom());
        writer.writeBooleanValue("hasAttachments", this.getHasAttachments());
        writer.writeEnumValue("importance", this.getImportance());
        writer.writeEnumValue("inferenceClassification", this.getInferenceClassification());
        writer.writeCollectionOfObjectValues("internetMessageHeaders", this.getInternetMessageHeaders());
        writer.writeStringValue("internetMessageId", this.getInternetMessageId());
        writer.writeBooleanValue("isDeliveryReceiptRequested", this.getIsDeliveryReceiptRequested());
        writer.writeBooleanValue("isDraft", this.getIsDraft());
        writer.writeBooleanValue("isRead", this.getIsRead());
        writer.writeBooleanValue("isReadReceiptRequested", this.getIsReadReceiptRequested());
        writer.writeCollectionOfObjectValues("multiValueExtendedProperties", this.getMultiValueExtendedProperties());
        writer.writeStringValue("parentFolderId", this.getParentFolderId());
        writer.writeOffsetDateTimeValue("receivedDateTime", this.getReceivedDateTime());
        writer.writeCollectionOfObjectValues("replyTo", this.getReplyTo());
        writer.writeObjectValue("sender", this.getSender());
        writer.writeOffsetDateTimeValue("sentDateTime", this.getSentDateTime());
        writer.writeCollectionOfObjectValues("singleValueExtendedProperties", this.getSingleValueExtendedProperties());
        writer.writeStringValue("subject", this.getSubject());
        writer.writeCollectionOfObjectValues("toRecipients", this.getToRecipients());
        writer.writeObjectValue("uniqueBody", this.getUniqueBody());
        writer.writeStringValue("webLink", this.getWebLink());
    }
    /**
     * Sets the attachments property value. The fileAttachment and itemAttachment attachments for the message.
     * @param value Value to set for the attachments property.
     * @return a void
     */
    public void setAttachments(@javax.annotation.Nullable final java.util.List<Attachment> value) {
        this._attachments = value;
    }
    /**
     * Sets the bccRecipients property value. The Bcc: recipients for the message.
     * @param value Value to set for the bccRecipients property.
     * @return a void
     */
    public void setBccRecipients(@javax.annotation.Nullable final java.util.List<Recipient> value) {
        this._bccRecipients = value;
    }
    /**
     * Sets the body property value. 
     * @param value Value to set for the body property.
     * @return a void
     */
    public void setBody(@javax.annotation.Nullable final ItemBody value) {
        this._body = value;
    }
    /**
     * Sets the bodyPreview property value. The first 255 characters of the message body. It is in text format. If the message contains instances of mention, this property would contain a concatenation of these mentions as well.
     * @param value Value to set for the bodyPreview property.
     * @return a void
     */
    public void setBodyPreview(@javax.annotation.Nullable final String value) {
        this._bodyPreview = value;
    }
    /**
     * Sets the ccRecipients property value. The Cc: recipients for the message.
     * @param value Value to set for the ccRecipients property.
     * @return a void
     */
    public void setCcRecipients(@javax.annotation.Nullable final java.util.List<Recipient> value) {
        this._ccRecipients = value;
    }
    /**
     * Sets the conversationId property value. The ID of the conversation the email belongs to.
     * @param value Value to set for the conversationId property.
     * @return a void
     */
    public void setConversationId(@javax.annotation.Nullable final String value) {
        this._conversationId = value;
    }
    /**
     * Sets the conversationIndex property value. Indicates the position of the message within the conversation.
     * @param value Value to set for the conversationIndex property.
     * @return a void
     */
    public void setConversationIndex(@javax.annotation.Nullable final byte[] value) {
        this._conversationIndex = value;
    }
    /**
     * Sets the extensions property value. The collection of open extensions defined for the message. Nullable.
     * @param value Value to set for the extensions property.
     * @return a void
     */
    public void setExtensions(@javax.annotation.Nullable final java.util.List<Extension> value) {
        this._extensions = value;
    }
    /**
     * Sets the flag property value. 
     * @param value Value to set for the flag property.
     * @return a void
     */
    public void setFlag(@javax.annotation.Nullable final FollowupFlag value) {
        this._flag = value;
    }
    /**
     * Sets the from property value. 
     * @param value Value to set for the from property.
     * @return a void
     */
    public void setFrom(@javax.annotation.Nullable final Recipient value) {
        this._from = value;
    }
    /**
     * Sets the hasAttachments property value. Indicates whether the message has attachments. This property doesn't include inline attachments, so if a message contains only inline attachments, this property is false. To verify the existence of inline attachments, parse the body property to look for a src attribute, such as <IMG src='cid:image001.jpg@01D26CD8.6C05F070'>.
     * @param value Value to set for the hasAttachments property.
     * @return a void
     */
    public void setHasAttachments(@javax.annotation.Nullable final Boolean value) {
        this._hasAttachments = value;
    }
    /**
     * Sets the importance property value. 
     * @param value Value to set for the importance property.
     * @return a void
     */
    public void setImportance(@javax.annotation.Nullable final Importance value) {
        this._importance = value;
    }
    /**
     * Sets the inferenceClassification property value. 
     * @param value Value to set for the inferenceClassification property.
     * @return a void
     */
    public void setInferenceClassification(@javax.annotation.Nullable final InferenceClassificationType value) {
        this._inferenceClassification = value;
    }
    /**
     * Sets the internetMessageHeaders property value. A collection of message headers defined by RFC5322. The set includes message headers indicating the network path taken by a message from the sender to the recipient. It can also contain custom message headers that hold app data for the message.  Returned only on applying a $select query option. Read-only.
     * @param value Value to set for the internetMessageHeaders property.
     * @return a void
     */
    public void setInternetMessageHeaders(@javax.annotation.Nullable final java.util.List<InternetMessageHeader> value) {
        this._internetMessageHeaders = value;
    }
    /**
     * Sets the internetMessageId property value. The message ID in the format specified by RFC2822.
     * @param value Value to set for the internetMessageId property.
     * @return a void
     */
    public void setInternetMessageId(@javax.annotation.Nullable final String value) {
        this._internetMessageId = value;
    }
    /**
     * Sets the isDeliveryReceiptRequested property value. Indicates whether a read receipt is requested for the message.
     * @param value Value to set for the isDeliveryReceiptRequested property.
     * @return a void
     */
    public void setIsDeliveryReceiptRequested(@javax.annotation.Nullable final Boolean value) {
        this._isDeliveryReceiptRequested = value;
    }
    /**
     * Sets the isDraft property value. Indicates whether the message is a draft. A message is a draft if it hasn't been sent yet.
     * @param value Value to set for the isDraft property.
     * @return a void
     */
    public void setIsDraft(@javax.annotation.Nullable final Boolean value) {
        this._isDraft = value;
    }
    /**
     * Sets the isRead property value. Indicates whether the message has been read.
     * @param value Value to set for the isRead property.
     * @return a void
     */
    public void setIsRead(@javax.annotation.Nullable final Boolean value) {
        this._isRead = value;
    }
    /**
     * Sets the isReadReceiptRequested property value. Indicates whether a read receipt is requested for the message.
     * @param value Value to set for the isReadReceiptRequested property.
     * @return a void
     */
    public void setIsReadReceiptRequested(@javax.annotation.Nullable final Boolean value) {
        this._isReadReceiptRequested = value;
    }
    /**
     * Sets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the message. Nullable.
     * @param value Value to set for the multiValueExtendedProperties property.
     * @return a void
     */
    public void setMultiValueExtendedProperties(@javax.annotation.Nullable final java.util.List<MultiValueLegacyExtendedProperty> value) {
        this._multiValueExtendedProperties = value;
    }
    /**
     * Sets the parentFolderId property value. The unique identifier for the message's parent mailFolder.
     * @param value Value to set for the parentFolderId property.
     * @return a void
     */
    public void setParentFolderId(@javax.annotation.Nullable final String value) {
        this._parentFolderId = value;
    }
    /**
     * Sets the receivedDateTime property value. The date and time the message was received.  The date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
     * @param value Value to set for the receivedDateTime property.
     * @return a void
     */
    public void setReceivedDateTime(@javax.annotation.Nullable final OffsetDateTime value) {
        this._receivedDateTime = value;
    }
    /**
     * Sets the replyTo property value. The email addresses to use when replying.
     * @param value Value to set for the replyTo property.
     * @return a void
     */
    public void setReplyTo(@javax.annotation.Nullable final java.util.List<Recipient> value) {
        this._replyTo = value;
    }
    /**
     * Sets the sender property value. 
     * @param value Value to set for the sender property.
     * @return a void
     */
    public void setSender(@javax.annotation.Nullable final Recipient value) {
        this._sender = value;
    }
    /**
     * Sets the sentDateTime property value. The date and time the message was sent.  The date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
     * @param value Value to set for the sentDateTime property.
     * @return a void
     */
    public void setSentDateTime(@javax.annotation.Nullable final OffsetDateTime value) {
        this._sentDateTime = value;
    }
    /**
     * Sets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the message. Nullable.
     * @param value Value to set for the singleValueExtendedProperties property.
     * @return a void
     */
    public void setSingleValueExtendedProperties(@javax.annotation.Nullable final java.util.List<SingleValueLegacyExtendedProperty> value) {
        this._singleValueExtendedProperties = value;
    }
    /**
     * Sets the subject property value. The subject of the message.
     * @param value Value to set for the subject property.
     * @return a void
     */
    public void setSubject(@javax.annotation.Nullable final String value) {
        this._subject = value;
    }
    /**
     * Sets the toRecipients property value. The To: recipients for the message.
     * @param value Value to set for the toRecipients property.
     * @return a void
     */
    public void setToRecipients(@javax.annotation.Nullable final java.util.List<Recipient> value) {
        this._toRecipients = value;
    }
    /**
     * Sets the uniqueBody property value. 
     * @param value Value to set for the uniqueBody property.
     * @return a void
     */
    public void setUniqueBody(@javax.annotation.Nullable final ItemBody value) {
        this._uniqueBody = value;
    }
    /**
     * Sets the webLink property value. The URL to open the message in Outlook on the web.You can append an ispopout argument to the end of the URL to change how the message is displayed. If ispopout is not present or if it is set to 1, then the message is shown in a popout window. If ispopout is set to 0, then the browser will show the message in the Outlook on the web review pane.The message will open in the browser if you are logged in to your mailbox via Outlook on the web. You will be prompted to login if you are not already logged in with the browser.This URL cannot be accessed from within an iFrame.
     * @param value Value to set for the webLink property.
     * @return a void
     */
    public void setWebLink(@javax.annotation.Nullable final String value) {
        this._webLink = value;
    }
}
