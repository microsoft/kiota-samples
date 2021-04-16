package graphjavav4.utilities.users.mailFolders;

import java.util.Objects;
import graphjavav4.utilities.users.Message;
import graphjavav4.utilities.users.MultiValueLegacyExtendedProperty;
import graphjavav4.utilities.users.SingleValueLegacyExtendedProperty;
import graphjavav4.utilities.users.Entity;
import com.microsoft.kiota.serialization.SerializationWriter;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.Parsable;
import java.util.function.BiConsumer;
import java.util.Map;
import java.util.HashMap;
import java.util.List;
public class MailFolder extends Entity implements Parsable {
    @javax.annotation.Nullable
    public Integer childFolderCount;
    @javax.annotation.Nullable
    public String displayName;
    @javax.annotation.Nullable
    public String parentFolderId;
    @javax.annotation.Nullable
    public Integer totalItemCount;
    @javax.annotation.Nullable
    public Integer unreadItemCount;
    @javax.annotation.Nullable
    public List<MailFolder> childFolders;
    @javax.annotation.Nullable
    public List<MessageRule> messageRules;
    @javax.annotation.Nullable
    public List<Message> messages;
    @javax.annotation.Nullable
    public List<MultiValueLegacyExtendedProperty> multiValueExtendedProperties;
    @javax.annotation.Nullable
    public List<SingleValueLegacyExtendedProperty> singleValueExtendedProperties;
    public void serialize(@javax.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        super.serialize(writer);
        writer.writeIntegerValue("childFolderCount", childFolderCount);
        writer.writeStringValue("displayName", displayName);
        writer.writeStringValue("parentFolderId", parentFolderId);
        writer.writeIntegerValue("totalItemCount", totalItemCount);
        writer.writeIntegerValue("unreadItemCount", unreadItemCount);
        writer.writeCollectionOfObjectValues("childFolders", childFolders);
        writer.writeCollectionOfObjectValues("messageRules", messageRules);
        writer.writeCollectionOfObjectValues("messages", messages);
        writer.writeCollectionOfObjectValues("multiValueExtendedProperties", multiValueExtendedProperties);
        writer.writeCollectionOfObjectValues("singleValueExtendedProperties", singleValueExtendedProperties);
    }
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getDeserializeFields() {
        final Map<String, BiConsumer<T, ParseNode>> fields = new HashMap<>(super.getDeserializeFields());
        fields.put("childFolderCount", (o, n) -> { ((MailFolder)o).childFolderCount = n.getIntegerValue(); });
        fields.put("displayName", (o, n) -> { ((MailFolder)o).displayName = n.getStringValue(); });
        fields.put("parentFolderId", (o, n) -> { ((MailFolder)o).parentFolderId = n.getStringValue(); });
        fields.put("totalItemCount", (o, n) -> { ((MailFolder)o).totalItemCount = n.getIntegerValue(); });
        fields.put("unreadItemCount", (o, n) -> { ((MailFolder)o).unreadItemCount = n.getIntegerValue(); });
        fields.put("childFolders", (o, n) -> { ((MailFolder)o).childFolders = n.getCollectionOfObjectValues(MailFolder.class); });
        fields.put("messageRules", (o, n) -> { ((MailFolder)o).messageRules = n.getCollectionOfObjectValues(MessageRule.class); });
        fields.put("messages", (o, n) -> { ((MailFolder)o).messages = n.getCollectionOfObjectValues(Message.class); });
        fields.put("multiValueExtendedProperties", (o, n) -> { ((MailFolder)o).multiValueExtendedProperties = n.getCollectionOfObjectValues(MultiValueLegacyExtendedProperty.class); });
        fields.put("singleValueExtendedProperties", (o, n) -> { ((MailFolder)o).singleValueExtendedProperties = n.getCollectionOfObjectValues(SingleValueLegacyExtendedProperty.class); });
        return fields;
    }
}
