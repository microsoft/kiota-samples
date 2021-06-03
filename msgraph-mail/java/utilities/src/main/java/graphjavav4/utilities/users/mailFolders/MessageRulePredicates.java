package graphjavav4.utilities.users.mailFolders;

import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import graphjavav4.utilities.users.Importance;
import graphjavav4.utilities.users.Recipient;
import java.util.function.BiConsumer;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Objects;
public class MessageRulePredicates implements Parsable {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.  */
    @javax.annotation.Nonnull
    private final Map<String, Object> additionalData = new HashMap<>();
    /** Represents the strings that should appear in the body of an incoming message in order for the condition or exception to apply.  */
    @javax.annotation.Nullable
    public List<String> bodyContains;
    /** Represents the strings that should appear in the body or subject of an incoming message in order for the condition or exception to apply.  */
    @javax.annotation.Nullable
    public List<String> bodyOrSubjectContains;
    /** Represents the categories that an incoming message should be labeled with in order for the condition or exception to apply.  */
    @javax.annotation.Nullable
    public List<String> categories;
    /** Represents the specific sender email addresses of an incoming message in order for the condition or exception to apply.  */
    @javax.annotation.Nullable
    public List<Recipient> fromAddresses;
    /** Indicates whether an incoming message must have attachments in order for the condition or exception to apply.  */
    @javax.annotation.Nullable
    public Boolean hasAttachments;
    /** Represents the strings that appear in the headers of an incoming message in order for the condition or exception to apply.  */
    @javax.annotation.Nullable
    public List<String> headerContains;
    @javax.annotation.Nullable
    public Importance importance;
    /** Indicates whether an incoming message must be an approval request in order for the condition or exception to apply.  */
    @javax.annotation.Nullable
    public Boolean isApprovalRequest;
    /** Indicates whether an incoming message must be automatically forwarded in order for the condition or exception to apply.  */
    @javax.annotation.Nullable
    public Boolean isAutomaticForward;
    /** Indicates whether an incoming message must be an auto reply in order for the condition or exception to apply.  */
    @javax.annotation.Nullable
    public Boolean isAutomaticReply;
    /** Indicates whether an incoming message must be encrypted in order for the condition or exception to apply.  */
    @javax.annotation.Nullable
    public Boolean isEncrypted;
    /** Indicates whether an incoming message must be a meeting request in order for the condition or exception to apply.  */
    @javax.annotation.Nullable
    public Boolean isMeetingRequest;
    /** Indicates whether an incoming message must be a meeting response in order for the condition or exception to apply.  */
    @javax.annotation.Nullable
    public Boolean isMeetingResponse;
    /** Indicates whether an incoming message must be a non-delivery report in order for the condition or exception to apply.  */
    @javax.annotation.Nullable
    public Boolean isNonDeliveryReport;
    /** Indicates whether an incoming message must be permission controlled (RMS-protected) in order for the condition or exception to apply.  */
    @javax.annotation.Nullable
    public Boolean isPermissionControlled;
    /** Indicates whether an incoming message must be a read receipt in order for the condition or exception to apply.  */
    @javax.annotation.Nullable
    public Boolean isReadReceipt;
    /** Indicates whether an incoming message must be S/MIME-signed in order for the condition or exception to apply.  */
    @javax.annotation.Nullable
    public Boolean isSigned;
    /** Indicates whether an incoming message must be a voice mail in order for the condition or exception to apply.  */
    @javax.annotation.Nullable
    public Boolean isVoicemail;
    @javax.annotation.Nullable
    public MessageActionFlag messageActionFlag;
    /** Indicates whether the owner of the mailbox must not be a recipient of an incoming message in order for the condition or exception to apply.  */
    @javax.annotation.Nullable
    public Boolean notSentToMe;
    /** Represents the strings that appear in either the toRecipients or ccRecipients properties of an incoming message in order for the condition or exception to apply.  */
    @javax.annotation.Nullable
    public List<String> recipientContains;
    /** Represents the strings that appear in the from property of an incoming message in order for the condition or exception to apply.  */
    @javax.annotation.Nullable
    public List<String> senderContains;
    @javax.annotation.Nullable
    public Sensitivity sensitivity;
    /** Indicates whether the owner of the mailbox must be in the ccRecipients property of an incoming message in order for the condition or exception to apply.  */
    @javax.annotation.Nullable
    public Boolean sentCcMe;
    /** Indicates whether the owner of the mailbox must be the only recipient in an incoming message in order for the condition or exception to apply.  */
    @javax.annotation.Nullable
    public Boolean sentOnlyToMe;
    /** Represents the email addresses that an incoming message must have been sent to in order for the condition or exception to apply.  */
    @javax.annotation.Nullable
    public List<Recipient> sentToAddresses;
    /** Indicates whether the owner of the mailbox must be in the toRecipients property of an incoming message in order for the condition or exception to apply.  */
    @javax.annotation.Nullable
    public Boolean sentToMe;
    /** Indicates whether the owner of the mailbox must be in either a toRecipients or ccRecipients property of an incoming message in order for the condition or exception to apply.  */
    @javax.annotation.Nullable
    public Boolean sentToOrCcMe;
    /** Represents the strings that appear in the subject of an incoming message in order for the condition or exception to apply.  */
    @javax.annotation.Nullable
    public List<String> subjectContains;
    @javax.annotation.Nullable
    public SizeRange withinSizeRange;
    /**
     * Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     * @return a Map<String, Object>
     */
    @javax.annotation.Nonnull
    public Map<String, Object> getAdditionalData() {
        return additionalData;
    }
    /**
     * The deserialization information for the current model
     * @return a Map<String, BiConsumer<T, ParseNode>>
     */
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getFieldDeserializers() {
        final Map<String, BiConsumer<T, ParseNode>> fields = new HashMap<>(30);
        fields.put("bodyContains", (o, n) -> { ((MessageRulePredicates)o).bodyContains = n.getCollectionOfPrimitiveValues(String.class); });
        fields.put("bodyOrSubjectContains", (o, n) -> { ((MessageRulePredicates)o).bodyOrSubjectContains = n.getCollectionOfPrimitiveValues(String.class); });
        fields.put("categories", (o, n) -> { ((MessageRulePredicates)o).categories = n.getCollectionOfPrimitiveValues(String.class); });
        fields.put("fromAddresses", (o, n) -> { ((MessageRulePredicates)o).fromAddresses = n.getCollectionOfObjectValues(Recipient.class); });
        fields.put("hasAttachments", (o, n) -> { ((MessageRulePredicates)o).hasAttachments = n.getBooleanValue(); });
        fields.put("headerContains", (o, n) -> { ((MessageRulePredicates)o).headerContains = n.getCollectionOfPrimitiveValues(String.class); });
        fields.put("importance", (o, n) -> { ((MessageRulePredicates)o).importance = n.getEnumValue(Importance.class); });
        fields.put("isApprovalRequest", (o, n) -> { ((MessageRulePredicates)o).isApprovalRequest = n.getBooleanValue(); });
        fields.put("isAutomaticForward", (o, n) -> { ((MessageRulePredicates)o).isAutomaticForward = n.getBooleanValue(); });
        fields.put("isAutomaticReply", (o, n) -> { ((MessageRulePredicates)o).isAutomaticReply = n.getBooleanValue(); });
        fields.put("isEncrypted", (o, n) -> { ((MessageRulePredicates)o).isEncrypted = n.getBooleanValue(); });
        fields.put("isMeetingRequest", (o, n) -> { ((MessageRulePredicates)o).isMeetingRequest = n.getBooleanValue(); });
        fields.put("isMeetingResponse", (o, n) -> { ((MessageRulePredicates)o).isMeetingResponse = n.getBooleanValue(); });
        fields.put("isNonDeliveryReport", (o, n) -> { ((MessageRulePredicates)o).isNonDeliveryReport = n.getBooleanValue(); });
        fields.put("isPermissionControlled", (o, n) -> { ((MessageRulePredicates)o).isPermissionControlled = n.getBooleanValue(); });
        fields.put("isReadReceipt", (o, n) -> { ((MessageRulePredicates)o).isReadReceipt = n.getBooleanValue(); });
        fields.put("isSigned", (o, n) -> { ((MessageRulePredicates)o).isSigned = n.getBooleanValue(); });
        fields.put("isVoicemail", (o, n) -> { ((MessageRulePredicates)o).isVoicemail = n.getBooleanValue(); });
        fields.put("messageActionFlag", (o, n) -> { ((MessageRulePredicates)o).messageActionFlag = n.getEnumValue(MessageActionFlag.class); });
        fields.put("notSentToMe", (o, n) -> { ((MessageRulePredicates)o).notSentToMe = n.getBooleanValue(); });
        fields.put("recipientContains", (o, n) -> { ((MessageRulePredicates)o).recipientContains = n.getCollectionOfPrimitiveValues(String.class); });
        fields.put("senderContains", (o, n) -> { ((MessageRulePredicates)o).senderContains = n.getCollectionOfPrimitiveValues(String.class); });
        fields.put("sensitivity", (o, n) -> { ((MessageRulePredicates)o).sensitivity = n.getEnumValue(Sensitivity.class); });
        fields.put("sentCcMe", (o, n) -> { ((MessageRulePredicates)o).sentCcMe = n.getBooleanValue(); });
        fields.put("sentOnlyToMe", (o, n) -> { ((MessageRulePredicates)o).sentOnlyToMe = n.getBooleanValue(); });
        fields.put("sentToAddresses", (o, n) -> { ((MessageRulePredicates)o).sentToAddresses = n.getCollectionOfObjectValues(Recipient.class); });
        fields.put("sentToMe", (o, n) -> { ((MessageRulePredicates)o).sentToMe = n.getBooleanValue(); });
        fields.put("sentToOrCcMe", (o, n) -> { ((MessageRulePredicates)o).sentToOrCcMe = n.getBooleanValue(); });
        fields.put("subjectContains", (o, n) -> { ((MessageRulePredicates)o).subjectContains = n.getCollectionOfPrimitiveValues(String.class); });
        fields.put("withinSizeRange", (o, n) -> { ((MessageRulePredicates)o).withinSizeRange = n.getObjectValue(SizeRange.class); });
        return fields;
    }
    /**
     * Serialiazes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @return a void
     */
    public void serialize(@javax.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        writer.writeCollectionOfPrimitiveValues("bodyContains", bodyContains);
        writer.writeCollectionOfPrimitiveValues("bodyOrSubjectContains", bodyOrSubjectContains);
        writer.writeCollectionOfPrimitiveValues("categories", categories);
        writer.writeCollectionOfObjectValues("fromAddresses", fromAddresses);
        writer.writeBooleanValue("hasAttachments", hasAttachments);
        writer.writeCollectionOfPrimitiveValues("headerContains", headerContains);
        writer.writeEnumValue("importance", importance);
        writer.writeBooleanValue("isApprovalRequest", isApprovalRequest);
        writer.writeBooleanValue("isAutomaticForward", isAutomaticForward);
        writer.writeBooleanValue("isAutomaticReply", isAutomaticReply);
        writer.writeBooleanValue("isEncrypted", isEncrypted);
        writer.writeBooleanValue("isMeetingRequest", isMeetingRequest);
        writer.writeBooleanValue("isMeetingResponse", isMeetingResponse);
        writer.writeBooleanValue("isNonDeliveryReport", isNonDeliveryReport);
        writer.writeBooleanValue("isPermissionControlled", isPermissionControlled);
        writer.writeBooleanValue("isReadReceipt", isReadReceipt);
        writer.writeBooleanValue("isSigned", isSigned);
        writer.writeBooleanValue("isVoicemail", isVoicemail);
        writer.writeEnumValue("messageActionFlag", messageActionFlag);
        writer.writeBooleanValue("notSentToMe", notSentToMe);
        writer.writeCollectionOfPrimitiveValues("recipientContains", recipientContains);
        writer.writeCollectionOfPrimitiveValues("senderContains", senderContains);
        writer.writeEnumValue("sensitivity", sensitivity);
        writer.writeBooleanValue("sentCcMe", sentCcMe);
        writer.writeBooleanValue("sentOnlyToMe", sentOnlyToMe);
        writer.writeCollectionOfObjectValues("sentToAddresses", sentToAddresses);
        writer.writeBooleanValue("sentToMe", sentToMe);
        writer.writeBooleanValue("sentToOrCcMe", sentToOrCcMe);
        writer.writeCollectionOfPrimitiveValues("subjectContains", subjectContains);
        writer.writeObjectValue("withinSizeRange", withinSizeRange);
        writer.writeAdditionalData(this.additionalData);
    }
}
