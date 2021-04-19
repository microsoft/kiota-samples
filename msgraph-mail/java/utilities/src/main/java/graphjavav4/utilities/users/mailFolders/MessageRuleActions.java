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
    @javax.annotation.Nullable
    public List<String> assignCategories;
    @javax.annotation.Nullable
    public String copyToFolder;
    @javax.annotation.Nullable
    public Boolean delete;
    @javax.annotation.Nullable
    public List<Recipient> forwardAsAttachmentTo;
    @javax.annotation.Nullable
    public List<Recipient> forwardTo;
    @javax.annotation.Nullable
    public Boolean markAsRead;
    @javax.annotation.Nullable
    public Importance markImportance;
    @javax.annotation.Nullable
    public String moveToFolder;
    @javax.annotation.Nullable
    public Boolean permanentDelete;
    @javax.annotation.Nullable
    public List<Recipient> redirectTo;
    @javax.annotation.Nullable
    public Boolean stopProcessingRules;
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
    }
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getDeserializeFields() {
        final Map<String, BiConsumer<T, ParseNode>> fields = new HashMap<>(11);
        fields.put("assignCategories", (o, n) -> { ((MessageRuleActions)o).assignCategories = n.getCollectionOfPrimitiveValues(String.class); });
        fields.put("copyToFolder", (o, n) -> { ((MessageRuleActions)o).copyToFolder = n.getStringValue(); });
        fields.put("delete", (o, n) -> { ((MessageRuleActions)o).delete = n.getBooleanValue(); });
        fields.put("forwardAsAttachmentTo", (o, n) -> { ((MessageRuleActions)o).forwardAsAttachmentTo = n.getCollectionOfObjectValues(Recipient.class); });
        fields.put("forwardTo", (o, n) -> { ((MessageRuleActions)o).forwardTo = n.getCollectionOfObjectValues(Recipient.class); });
        fields.put("markAsRead", (o, n) -> { ((MessageRuleActions)o).markAsRead = n.getBooleanValue(); });
        fields.put("markImportance", (o, n) -> { ((MessageRuleActions)o).markImportance = n.getEnumValue(Importance.class); });
        fields.put("moveToFolder", (o, n) -> { ((MessageRuleActions)o).moveToFolder = n.getStringValue(); });
        fields.put("permanentDelete", (o, n) -> { ((MessageRuleActions)o).permanentDelete = n.getBooleanValue(); });
        fields.put("redirectTo", (o, n) -> { ((MessageRuleActions)o).redirectTo = n.getCollectionOfObjectValues(Recipient.class); });
        fields.put("stopProcessingRules", (o, n) -> { ((MessageRuleActions)o).stopProcessingRules = n.getBooleanValue(); });
        return fields;
    }
}
