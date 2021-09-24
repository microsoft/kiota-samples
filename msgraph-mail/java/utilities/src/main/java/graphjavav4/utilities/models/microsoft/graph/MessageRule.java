package graphjavav4.utilities.models.microsoft.graph;

import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import java.util.function.BiConsumer;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
public class MessageRule extends Entity implements Parsable {
    private MessageRuleActions _actions;
    private MessageRulePredicates _conditions;
    /** The display name of the rule.  */
    private String _displayName;
    private MessageRulePredicates _exceptions;
    /** Indicates whether the rule is in an error condition. Read-only.  */
    private Boolean _hasError;
    /** Indicates whether the rule is enabled to be applied to messages.  */
    private Boolean _isEnabled;
    /** Indicates if the rule is read-only and cannot be modified or deleted by the rules REST API.  */
    private Boolean _isReadOnly;
    /** Indicates the order in which the rule is executed, among other rules.  */
    private Integer _sequence;
    /**
     * Instantiates a new messageRule and sets the default values.
     * @return a void
     */
    public MessageRule() {
        super();
    }
    /**
     * Gets the actions property value. 
     * @return a messageRuleActions
     */
    @javax.annotation.Nullable
    public MessageRuleActions getActions() {
        return this._actions;
    }
    /**
     * Gets the conditions property value. 
     * @return a messageRulePredicates
     */
    @javax.annotation.Nullable
    public MessageRulePredicates getConditions() {
        return this._conditions;
    }
    /**
     * Gets the displayName property value. The display name of the rule.
     * @return a string
     */
    @javax.annotation.Nullable
    public String getDisplayName() {
        return this._displayName;
    }
    /**
     * Gets the exceptions property value. 
     * @return a messageRulePredicates
     */
    @javax.annotation.Nullable
    public MessageRulePredicates getExceptions() {
        return this._exceptions;
    }
    /**
     * Gets the hasError property value. Indicates whether the rule is in an error condition. Read-only.
     * @return a boolean
     */
    @javax.annotation.Nullable
    public Boolean getHasError() {
        return this._hasError;
    }
    /**
     * Gets the isEnabled property value. Indicates whether the rule is enabled to be applied to messages.
     * @return a boolean
     */
    @javax.annotation.Nullable
    public Boolean getIsEnabled() {
        return this._isEnabled;
    }
    /**
     * Gets the isReadOnly property value. Indicates if the rule is read-only and cannot be modified or deleted by the rules REST API.
     * @return a boolean
     */
    @javax.annotation.Nullable
    public Boolean getIsReadOnly() {
        return this._isReadOnly;
    }
    /**
     * Gets the sequence property value. Indicates the order in which the rule is executed, among other rules.
     * @return a integer
     */
    @javax.annotation.Nullable
    public Integer getSequence() {
        return this._sequence;
    }
    /**
     * The deserialization information for the current model
     * @return a Map<String, BiConsumer<T, ParseNode>>
     */
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getFieldDeserializers() {
        return new HashMap<>(super.getFieldDeserializers()) {{
            this.put("actions", (o, n) -> { ((MessageRule)o).setActions(n.getObjectValue(MessageRuleActions.class)); });
            this.put("conditions", (o, n) -> { ((MessageRule)o).setConditions(n.getObjectValue(MessageRulePredicates.class)); });
            this.put("displayName", (o, n) -> { ((MessageRule)o).setDisplayName(n.getStringValue()); });
            this.put("exceptions", (o, n) -> { ((MessageRule)o).setExceptions(n.getObjectValue(MessageRulePredicates.class)); });
            this.put("hasError", (o, n) -> { ((MessageRule)o).setHasError(n.getBooleanValue()); });
            this.put("isEnabled", (o, n) -> { ((MessageRule)o).setIsEnabled(n.getBooleanValue()); });
            this.put("isReadOnly", (o, n) -> { ((MessageRule)o).setIsReadOnly(n.getBooleanValue()); });
            this.put("sequence", (o, n) -> { ((MessageRule)o).setSequence(n.getIntegerValue()); });
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
        writer.writeObjectValue("actions", this.getActions());
        writer.writeObjectValue("conditions", this.getConditions());
        writer.writeStringValue("displayName", this.getDisplayName());
        writer.writeObjectValue("exceptions", this.getExceptions());
        writer.writeBooleanValue("hasError", this.getHasError());
        writer.writeBooleanValue("isEnabled", this.getIsEnabled());
        writer.writeBooleanValue("isReadOnly", this.getIsReadOnly());
        writer.writeIntegerValue("sequence", this.getSequence());
    }
    /**
     * Sets the actions property value. 
     * @param value Value to set for the actions property.
     * @return a void
     */
    public void setActions(@javax.annotation.Nullable final MessageRuleActions value) {
        this._actions = value;
    }
    /**
     * Sets the conditions property value. 
     * @param value Value to set for the conditions property.
     * @return a void
     */
    public void setConditions(@javax.annotation.Nullable final MessageRulePredicates value) {
        this._conditions = value;
    }
    /**
     * Sets the displayName property value. The display name of the rule.
     * @param value Value to set for the displayName property.
     * @return a void
     */
    public void setDisplayName(@javax.annotation.Nullable final String value) {
        this._displayName = value;
    }
    /**
     * Sets the exceptions property value. 
     * @param value Value to set for the exceptions property.
     * @return a void
     */
    public void setExceptions(@javax.annotation.Nullable final MessageRulePredicates value) {
        this._exceptions = value;
    }
    /**
     * Sets the hasError property value. Indicates whether the rule is in an error condition. Read-only.
     * @param value Value to set for the hasError property.
     * @return a void
     */
    public void setHasError(@javax.annotation.Nullable final Boolean value) {
        this._hasError = value;
    }
    /**
     * Sets the isEnabled property value. Indicates whether the rule is enabled to be applied to messages.
     * @param value Value to set for the isEnabled property.
     * @return a void
     */
    public void setIsEnabled(@javax.annotation.Nullable final Boolean value) {
        this._isEnabled = value;
    }
    /**
     * Sets the isReadOnly property value. Indicates if the rule is read-only and cannot be modified or deleted by the rules REST API.
     * @param value Value to set for the isReadOnly property.
     * @return a void
     */
    public void setIsReadOnly(@javax.annotation.Nullable final Boolean value) {
        this._isReadOnly = value;
    }
    /**
     * Sets the sequence property value. Indicates the order in which the rule is executed, among other rules.
     * @param value Value to set for the sequence property.
     * @return a void
     */
    public void setSequence(@javax.annotation.Nullable final Integer value) {
        this._sequence = value;
    }
}
