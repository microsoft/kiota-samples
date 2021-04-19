package graphjavav4.utilities.users.mailFolders;

import java.util.Objects;
import graphjavav4.utilities.users.Recipient;
import graphjavav4.utilities.users.Importance;
import com.microsoft.kiota.serialization.SerializationWriter;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.Parsable;
import java.util.function.BiConsumer;
import java.util.Map;
import java.util.HashMap;
import java.util.List;
public class MessageRulePredicates implements Parsable {
    @javax.annotation.Nullable
    public List<Object> bodyContains;
    @javax.annotation.Nullable
    public List<Object> bodyOrSubjectContains;
    @javax.annotation.Nullable
    public List<Object> categories;
    @javax.annotation.Nullable
    public List<Recipient> fromAddresses;
    @javax.annotation.Nullable
    public Boolean hasAttachments;
    @javax.annotation.Nullable
    public List<Object> headerContains;
    @javax.annotation.Nullable
    public Importance importance;
    @javax.annotation.Nullable
    public Boolean isApprovalRequest;
    @javax.annotation.Nullable
    public Boolean isAutomaticForward;
    @javax.annotation.Nullable
    public Boolean isAutomaticReply;
    @javax.annotation.Nullable
    public Boolean isEncrypted;
    @javax.annotation.Nullable
    public Boolean isMeetingRequest;
    @javax.annotation.Nullable
    public Boolean isMeetingResponse;
    @javax.annotation.Nullable
    public Boolean isNonDeliveryReport;
    @javax.annotation.Nullable
    public Boolean isPermissionControlled;
    @javax.annotation.Nullable
    public Boolean isReadReceipt;
    @javax.annotation.Nullable
    public Boolean isSigned;
    @javax.annotation.Nullable
    public Boolean isVoicemail;
    @javax.annotation.Nullable
    public MessageActionFlag messageActionFlag;
    @javax.annotation.Nullable
    public Boolean notSentToMe;
    @javax.annotation.Nullable
    public List<Object> recipientContains;
    @javax.annotation.Nullable
    public List<Object> senderContains;
    @javax.annotation.Nullable
    public Sensitivity sensitivity;
    @javax.annotation.Nullable
    public Boolean sentCcMe;
    @javax.annotation.Nullable
    public Boolean sentOnlyToMe;
    @javax.annotation.Nullable
    public List<Recipient> sentToAddresses;
    @javax.annotation.Nullable
    public Boolean sentToMe;
    @javax.annotation.Nullable
    public Boolean sentToOrCcMe;
    @javax.annotation.Nullable
    public List<Object> subjectContains;
    @javax.annotation.Nullable
    public SizeRange withinSizeRange;
    public void serialize(@javax.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        writer.writeCollectionOfPrimitiveValues("bodyContains", bodyContains);
        writer.writeCollectionOfPrimitiveValues("bodyOrSubjectContains", bodyOrSubjectContains);
        writer.writeCollectionOfPrimitiveValues("categories", categories);
        writer.writeCollectionOfObjectValues("fromAddresses", fromAddresses);
        writer.writeBooleanValue("hasAttachments", hasAttachments);
        writer.writeCollectionOfPrimitiveValues("headerContains", headerContains);
        writer.writeObjectValue("importance", importance);
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
        writer.writeObjectValue("messageActionFlag", messageActionFlag);
        writer.writeBooleanValue("notSentToMe", notSentToMe);
        writer.writeCollectionOfPrimitiveValues("recipientContains", recipientContains);
        writer.writeCollectionOfPrimitiveValues("senderContains", senderContains);
        writer.writeObjectValue("sensitivity", sensitivity);
        writer.writeBooleanValue("sentCcMe", sentCcMe);
        writer.writeBooleanValue("sentOnlyToMe", sentOnlyToMe);
        writer.writeCollectionOfObjectValues("sentToAddresses", sentToAddresses);
        writer.writeBooleanValue("sentToMe", sentToMe);
        writer.writeBooleanValue("sentToOrCcMe", sentToOrCcMe);
        writer.writeCollectionOfPrimitiveValues("subjectContains", subjectContains);
        writer.writeObjectValue("withinSizeRange", withinSizeRange);
    }
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getDeserializeFields() {
        final Map<String, BiConsumer<T, ParseNode>> fields = new HashMap<>(30);
        fields.put("bodyContains", (o, n) -> { ((MessageRulePredicates)o).bodyContains = n.getCollectionOfPrimitiveValues(Object.class); });
        fields.put("bodyOrSubjectContains", (o, n) -> { ((MessageRulePredicates)o).bodyOrSubjectContains = n.getCollectionOfPrimitiveValues(Object.class); });
        fields.put("categories", (o, n) -> { ((MessageRulePredicates)o).categories = n.getCollectionOfPrimitiveValues(Object.class); });
        fields.put("fromAddresses", (o, n) -> { ((MessageRulePredicates)o).fromAddresses = n.getCollectionOfObjectValues(Recipient.class); });
        fields.put("hasAttachments", (o, n) -> { ((MessageRulePredicates)o).hasAttachments = n.getBooleanValue(); });
        fields.put("headerContains", (o, n) -> { ((MessageRulePredicates)o).headerContains = n.getCollectionOfPrimitiveValues(Object.class); });
        fields.put("importance", (o, n) -> { ((MessageRulePredicates)o).importance = n.getObjectValue(Importance.class); });
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
        fields.put("messageActionFlag", (o, n) -> { ((MessageRulePredicates)o).messageActionFlag = n.getObjectValue(MessageActionFlag.class); });
        fields.put("notSentToMe", (o, n) -> { ((MessageRulePredicates)o).notSentToMe = n.getBooleanValue(); });
        fields.put("recipientContains", (o, n) -> { ((MessageRulePredicates)o).recipientContains = n.getCollectionOfPrimitiveValues(Object.class); });
        fields.put("senderContains", (o, n) -> { ((MessageRulePredicates)o).senderContains = n.getCollectionOfPrimitiveValues(Object.class); });
        fields.put("sensitivity", (o, n) -> { ((MessageRulePredicates)o).sensitivity = n.getObjectValue(Sensitivity.class); });
        fields.put("sentCcMe", (o, n) -> { ((MessageRulePredicates)o).sentCcMe = n.getBooleanValue(); });
        fields.put("sentOnlyToMe", (o, n) -> { ((MessageRulePredicates)o).sentOnlyToMe = n.getBooleanValue(); });
        fields.put("sentToAddresses", (o, n) -> { ((MessageRulePredicates)o).sentToAddresses = n.getCollectionOfObjectValues(Recipient.class); });
        fields.put("sentToMe", (o, n) -> { ((MessageRulePredicates)o).sentToMe = n.getBooleanValue(); });
        fields.put("sentToOrCcMe", (o, n) -> { ((MessageRulePredicates)o).sentToOrCcMe = n.getBooleanValue(); });
        fields.put("subjectContains", (o, n) -> { ((MessageRulePredicates)o).subjectContains = n.getCollectionOfPrimitiveValues(Object.class); });
        fields.put("withinSizeRange", (o, n) -> { ((MessageRulePredicates)o).withinSizeRange = n.getObjectValue(SizeRange.class); });
        return fields;
    }
}
