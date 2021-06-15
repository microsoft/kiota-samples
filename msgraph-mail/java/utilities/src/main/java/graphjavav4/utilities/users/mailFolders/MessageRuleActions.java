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
    private Map<String, Object> _additionalData;
    /** A list of categories to be assigned to a message.  */
    private List<String> _assignCategories;
    /** The ID of a folder that a message is to be copied to.  */
    private String _copyToFolder;
    /** Indicates whether a message should be moved to the Deleted Items folder.  */
    private Boolean _delete;
    /** The email addresses of the recipients to which a message should be forwarded as an attachment.  */
    private List<Recipient> _forwardAsAttachmentTo;
    /** The email addresses of the recipients to which a message should be forwarded.  */
    private List<Recipient> _forwardTo;
    /** Indicates whether a message should be marked as read.  */
    private Boolean _markAsRead;
    private Importance _markImportance;
    /** The ID of the folder that a message will be moved to.  */
    private String _moveToFolder;
    /** Indicates whether a message should be permanently deleted and not saved to the Deleted Items folder.  */
    private Boolean _permanentDelete;
    /** The email address to which a message should be redirected.  */
    private List<Recipient> _redirectTo;
    /** Indicates whether subsequent rules should be evaluated.  */
    private Boolean _stopProcessingRules;
    /**
     * Instantiates a new messageRuleActions and sets the default values.
     * @return a void
     */
    public MessageRuleActions() {
        this._additionalData = new HashMap<>();
    }
    /**
     * Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     * @return a Map<String, Object>
     */
    @javax.annotation.Nonnull
    public Map<String, Object> getAdditionalData() {
        return this._additionalData;
    }
    /**
     * Gets the assignCategories property value. A list of categories to be assigned to a message.
     * @return a string
     */
    @javax.annotation.Nullable
    public List<String> getAssignCategories() {
        return this._assignCategories;
    }
    /**
     * Gets the copyToFolder property value. The ID of a folder that a message is to be copied to.
     * @return a string
     */
    @javax.annotation.Nullable
    public String getCopyToFolder() {
        return this._copyToFolder;
    }
    /**
     * Gets the delete property value. Indicates whether a message should be moved to the Deleted Items folder.
     * @return a boolean
     */
    @javax.annotation.Nullable
    public Boolean getDelete() {
        return this._delete;
    }
    /**
     * Gets the forwardAsAttachmentTo property value. The email addresses of the recipients to which a message should be forwarded as an attachment.
     * @return a recipient
     */
    @javax.annotation.Nullable
    public List<Recipient> getForwardAsAttachmentTo() {
        return this._forwardAsAttachmentTo;
    }
    /**
     * Gets the forwardTo property value. The email addresses of the recipients to which a message should be forwarded.
     * @return a recipient
     */
    @javax.annotation.Nullable
    public List<Recipient> getForwardTo() {
        return this._forwardTo;
    }
    /**
     * Gets the markAsRead property value. Indicates whether a message should be marked as read.
     * @return a boolean
     */
    @javax.annotation.Nullable
    public Boolean getMarkAsRead() {
        return this._markAsRead;
    }
    /**
     * Gets the markImportance property value. 
     * @return a importance
     */
    @javax.annotation.Nullable
    public Importance getMarkImportance() {
        return this._markImportance;
    }
    /**
     * Gets the moveToFolder property value. The ID of the folder that a message will be moved to.
     * @return a string
     */
    @javax.annotation.Nullable
    public String getMoveToFolder() {
        return this._moveToFolder;
    }
    /**
     * Gets the permanentDelete property value. Indicates whether a message should be permanently deleted and not saved to the Deleted Items folder.
     * @return a boolean
     */
    @javax.annotation.Nullable
    public Boolean getPermanentDelete() {
        return this._permanentDelete;
    }
    /**
     * Gets the redirectTo property value. The email address to which a message should be redirected.
     * @return a recipient
     */
    @javax.annotation.Nullable
    public List<Recipient> getRedirectTo() {
        return this._redirectTo;
    }
    /**
     * Gets the stopProcessingRules property value. Indicates whether subsequent rules should be evaluated.
     * @return a boolean
     */
    @javax.annotation.Nullable
    public Boolean getStopProcessingRules() {
        return this._stopProcessingRules;
    }
    /**
     * The deserialization information for the current model
     * @return a Map<String, BiConsumer<T, ParseNode>>
     */
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getFieldDeserializers() {
        return new HashMap<>(11) {{
            this.put("assignCategories", (o, n) -> { ((MessageRuleActions)o).setAssignCategories(n.getCollectionOfPrimitiveValues(String.class)); });
            this.put("copyToFolder", (o, n) -> { ((MessageRuleActions)o).setCopyToFolder(n.getStringValue()); });
            this.put("delete", (o, n) -> { ((MessageRuleActions)o).setDelete(n.getBooleanValue()); });
            this.put("forwardAsAttachmentTo", (o, n) -> { ((MessageRuleActions)o).setForwardAsAttachmentTo(n.getCollectionOfObjectValues(Recipient.class)); });
            this.put("forwardTo", (o, n) -> { ((MessageRuleActions)o).setForwardTo(n.getCollectionOfObjectValues(Recipient.class)); });
            this.put("markAsRead", (o, n) -> { ((MessageRuleActions)o).setMarkAsRead(n.getBooleanValue()); });
            this.put("markImportance", (o, n) -> { ((MessageRuleActions)o).setMarkImportance(n.getEnumValue(Importance.class)); });
            this.put("moveToFolder", (o, n) -> { ((MessageRuleActions)o).setMoveToFolder(n.getStringValue()); });
            this.put("permanentDelete", (o, n) -> { ((MessageRuleActions)o).setPermanentDelete(n.getBooleanValue()); });
            this.put("redirectTo", (o, n) -> { ((MessageRuleActions)o).setRedirectTo(n.getCollectionOfObjectValues(Recipient.class)); });
            this.put("stopProcessingRules", (o, n) -> { ((MessageRuleActions)o).setStopProcessingRules(n.getBooleanValue()); });
        }};
    }
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @return a void
     */
    public void serialize(@javax.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        writer.writeCollectionOfPrimitiveValues("assignCategories", this.getAssignCategories());
        writer.writeStringValue("copyToFolder", this.getCopyToFolder());
        writer.writeBooleanValue("delete", this.getDelete());
        writer.writeCollectionOfObjectValues("forwardAsAttachmentTo", this.getForwardAsAttachmentTo());
        writer.writeCollectionOfObjectValues("forwardTo", this.getForwardTo());
        writer.writeBooleanValue("markAsRead", this.getMarkAsRead());
        writer.writeEnumValue("markImportance", this.getMarkImportance());
        writer.writeStringValue("moveToFolder", this.getMoveToFolder());
        writer.writeBooleanValue("permanentDelete", this.getPermanentDelete());
        writer.writeCollectionOfObjectValues("redirectTo", this.getRedirectTo());
        writer.writeBooleanValue("stopProcessingRules", this.getStopProcessingRules());
        writer.writeAdditionalData(this.getAdditionalData());
    }
    /**
     * Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     * @param value Value to set for the AdditionalData property.
     * @return a void
     */
    public void setAdditionalData(@javax.annotation.Nullable final Map<String, Object> value) {
        this._additionalData = value;
    }
    /**
     * Sets the assignCategories property value. A list of categories to be assigned to a message.
     * @param value Value to set for the assignCategories property.
     * @return a void
     */
    public void setAssignCategories(@javax.annotation.Nullable final List<String> value) {
        this._assignCategories = value;
    }
    /**
     * Sets the copyToFolder property value. The ID of a folder that a message is to be copied to.
     * @param value Value to set for the copyToFolder property.
     * @return a void
     */
    public void setCopyToFolder(@javax.annotation.Nullable final String value) {
        this._copyToFolder = value;
    }
    /**
     * Sets the delete property value. Indicates whether a message should be moved to the Deleted Items folder.
     * @param value Value to set for the delete property.
     * @return a void
     */
    public void setDelete(@javax.annotation.Nullable final Boolean value) {
        this._delete = value;
    }
    /**
     * Sets the forwardAsAttachmentTo property value. The email addresses of the recipients to which a message should be forwarded as an attachment.
     * @param value Value to set for the forwardAsAttachmentTo property.
     * @return a void
     */
    public void setForwardAsAttachmentTo(@javax.annotation.Nullable final List<Recipient> value) {
        this._forwardAsAttachmentTo = value;
    }
    /**
     * Sets the forwardTo property value. The email addresses of the recipients to which a message should be forwarded.
     * @param value Value to set for the forwardTo property.
     * @return a void
     */
    public void setForwardTo(@javax.annotation.Nullable final List<Recipient> value) {
        this._forwardTo = value;
    }
    /**
     * Sets the markAsRead property value. Indicates whether a message should be marked as read.
     * @param value Value to set for the markAsRead property.
     * @return a void
     */
    public void setMarkAsRead(@javax.annotation.Nullable final Boolean value) {
        this._markAsRead = value;
    }
    /**
     * Sets the markImportance property value. 
     * @param value Value to set for the markImportance property.
     * @return a void
     */
    public void setMarkImportance(@javax.annotation.Nullable final Importance value) {
        this._markImportance = value;
    }
    /**
     * Sets the moveToFolder property value. The ID of the folder that a message will be moved to.
     * @param value Value to set for the moveToFolder property.
     * @return a void
     */
    public void setMoveToFolder(@javax.annotation.Nullable final String value) {
        this._moveToFolder = value;
    }
    /**
     * Sets the permanentDelete property value. Indicates whether a message should be permanently deleted and not saved to the Deleted Items folder.
     * @param value Value to set for the permanentDelete property.
     * @return a void
     */
    public void setPermanentDelete(@javax.annotation.Nullable final Boolean value) {
        this._permanentDelete = value;
    }
    /**
     * Sets the redirectTo property value. The email address to which a message should be redirected.
     * @param value Value to set for the redirectTo property.
     * @return a void
     */
    public void setRedirectTo(@javax.annotation.Nullable final List<Recipient> value) {
        this._redirectTo = value;
    }
    /**
     * Sets the stopProcessingRules property value. Indicates whether subsequent rules should be evaluated.
     * @param value Value to set for the stopProcessingRules property.
     * @return a void
     */
    public void setStopProcessingRules(@javax.annotation.Nullable final Boolean value) {
        this._stopProcessingRules = value;
    }
}
