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
public class MessageRuleActions implements Parsable {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.  */
    @javax.annotation.Nonnull
    private Map<String, Object> additionalData = new HashMap<>();
    /** A list of categories to be assigned to a message.  */
    @javax.annotation.Nullable
    public List<String> assignCategories;
    /** The ID of a folder that a message is to be copied to.  */
    @javax.annotation.Nullable
    public String copyToFolder;
    /** Indicates whether a message should be moved to the Deleted Items folder.  */
    @javax.annotation.Nullable
    public Boolean delete;
    /** The email addresses of the recipients to which a message should be forwarded as an attachment.  */
    @javax.annotation.Nullable
    public List<Recipient> forwardAsAttachmentTo;
    /** The email addresses of the recipients to which a message should be forwarded.  */
    @javax.annotation.Nullable
    public List<Recipient> forwardTo;
    /** Indicates whether a message should be marked as read.  */
    @javax.annotation.Nullable
    public Boolean markAsRead;
    @javax.annotation.Nullable
    public Importance markImportance;
    /** The ID of the folder that a message will be moved to.  */
    @javax.annotation.Nullable
    public String moveToFolder;
    /** Indicates whether a message should be permanently deleted and not saved to the Deleted Items folder.  */
    @javax.annotation.Nullable
    public Boolean permanentDelete;
    /** The email address to which a message should be redirected.  */
    @javax.annotation.Nullable
    public List<Recipient> redirectTo;
    /** Indicates whether subsequent rules should be evaluated.  */
    @javax.annotation.Nullable
    public Boolean stopProcessingRules;
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
        return new HashMap<>(11) {{
            this.put("assignCategories", (o, n) -> { ((MessageRuleActions)o).assignCategories = n.getCollectionOfPrimitiveValues(String.class); });
            this.put("copyToFolder", (o, n) -> { ((MessageRuleActions)o).copyToFolder = n.getStringValue(); });
            this.put("delete", (o, n) -> { ((MessageRuleActions)o).delete = n.getBooleanValue(); });
            this.put("forwardAsAttachmentTo", (o, n) -> { ((MessageRuleActions)o).forwardAsAttachmentTo = n.getCollectionOfObjectValues(Recipient.class); });
            this.put("forwardTo", (o, n) -> { ((MessageRuleActions)o).forwardTo = n.getCollectionOfObjectValues(Recipient.class); });
            this.put("markAsRead", (o, n) -> { ((MessageRuleActions)o).markAsRead = n.getBooleanValue(); });
            this.put("markImportance", (o, n) -> { ((MessageRuleActions)o).markImportance = n.getEnumValue(Importance.class); });
            this.put("moveToFolder", (o, n) -> { ((MessageRuleActions)o).moveToFolder = n.getStringValue(); });
            this.put("permanentDelete", (o, n) -> { ((MessageRuleActions)o).permanentDelete = n.getBooleanValue(); });
            this.put("redirectTo", (o, n) -> { ((MessageRuleActions)o).redirectTo = n.getCollectionOfObjectValues(Recipient.class); });
            this.put("stopProcessingRules", (o, n) -> { ((MessageRuleActions)o).stopProcessingRules = n.getBooleanValue(); });
        }};
    }
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @return a void
     */
    public void serialize(@javax.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        writer.writeCollectionOfPrimitiveValues("assignCategories", assignCategories);
        writer.writeStringValue("copyToFolder", copyToFolder);
        writer.writeBooleanValue("delete", delete);
        writer.writeCollectionOfObjectValues("forwardAsAttachmentTo", forwardAsAttachmentTo);
        writer.writeCollectionOfObjectValues("forwardTo", forwardTo);
        writer.writeBooleanValue("markAsRead", markAsRead);
        writer.writeEnumValue("markImportance", markImportance);
        writer.writeStringValue("moveToFolder", moveToFolder);
        writer.writeBooleanValue("permanentDelete", permanentDelete);
        writer.writeCollectionOfObjectValues("redirectTo", redirectTo);
        writer.writeBooleanValue("stopProcessingRules", stopProcessingRules);
        writer.writeAdditionalData(this.additionalData);
    }
}
