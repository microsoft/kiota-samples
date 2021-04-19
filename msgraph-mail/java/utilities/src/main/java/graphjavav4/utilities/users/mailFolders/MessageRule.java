package graphjavav4.utilities.users.mailFolders;

import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import graphjavav4.utilities.users.Entity;
import java.util.function.BiConsumer;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
public class MessageRule extends Entity implements Parsable {
    @javax.annotation.Nullable
    public MessageRuleActions actions;
    @javax.annotation.Nullable
    public MessageRulePredicates conditions;
    @javax.annotation.Nullable
    public String displayName;
    @javax.annotation.Nullable
    public MessageRulePredicates exceptions;
    @javax.annotation.Nullable
    public Boolean hasError;
    @javax.annotation.Nullable
    public Boolean isEnabled;
    @javax.annotation.Nullable
    public Boolean isReadOnly;
    @javax.annotation.Nullable
    public Integer sequence;
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
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getDeserializeFields() {
        final Map<String, BiConsumer<T, ParseNode>> fields = new HashMap<>(super.getDeserializeFields());
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
}
