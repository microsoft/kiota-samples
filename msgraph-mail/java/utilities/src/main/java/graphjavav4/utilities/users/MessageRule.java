package graphjavav4.utilities.users;

import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import graphjavav4.utilities.users.mailFolders.MessageRuleActions;
import graphjavav4.utilities.users.mailFolders.MessageRulePredicates;
import java.util.function.BiConsumer;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
public class MessageRule extends Entity implements Parsable {
    @javax.annotation.Nullable
    public MessageRuleActions actions;
    @javax.annotation.Nullable
    public MessageRulePredicates conditions;
    /** The display name of the rule.  */
    @javax.annotation.Nullable
    public String displayName;
    @javax.annotation.Nullable
    public MessageRulePredicates exceptions;
    /** Indicates whether the rule is in an error condition. Read-only.  */
    @javax.annotation.Nullable
    public Boolean hasError;
    /** Indicates whether the rule is enabled to be applied to messages.  */
    @javax.annotation.Nullable
    public Boolean isEnabled;
    /** Indicates if the rule is read-only and cannot be modified or deleted by the rules REST API.  */
    @javax.annotation.Nullable
    public Boolean isReadOnly;
    /** Indicates the order in which the rule is executed, among other rules.  */
    @javax.annotation.Nullable
    public Integer sequence;
    /**
     * The deserialization information for the current model
     * @return a Map<String, BiConsumer<T, ParseNode>>
     */
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getFieldDeserializers() {
        final Map<String, BiConsumer<T, ParseNode>> fields = new HashMap<>(super.getFieldDeserializers());
        fields.put("actions", (o, n) -> { ((MessageRule)o).actions = n.getObjectValue(MessageRuleActions.class); });
        fields.put("conditions", (o, n) -> { ((MessageRule)o).conditions = n.getObjectValue(MessageRulePredicates.class); });
        fields.put("displayName", (o, n) -> { ((MessageRule)o).displayName = n.getStringValue(); });
        fields.put("exceptions", (o, n) -> { ((MessageRule)o).exceptions = n.getObjectValue(MessageRulePredicates.class); });
        fields.put("hasError", (o, n) -> { ((MessageRule)o).hasError = n.getBooleanValue(); });
        fields.put("isEnabled", (o, n) -> { ((MessageRule)o).isEnabled = n.getBooleanValue(); });
        fields.put("isReadOnly", (o, n) -> { ((MessageRule)o).isReadOnly = n.getBooleanValue(); });
        fields.put("sequence", (o, n) -> { ((MessageRule)o).sequence = n.getIntegerValue(); });
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
        writer.writeObjectValue("actions", actions);
        writer.writeObjectValue("conditions", conditions);
        writer.writeStringValue("displayName", displayName);
        writer.writeObjectValue("exceptions", exceptions);
        writer.writeBooleanValue("hasError", hasError);
        writer.writeBooleanValue("isEnabled", isEnabled);
        writer.writeBooleanValue("isReadOnly", isReadOnly);
        writer.writeIntegerValue("sequence", sequence);
    }
}
