package graphjavav4.utilities.users.mailFolders;

import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import graphjavav4.utilities.users.Entity;
import graphjavav4.utilities.users.Message;
import graphjavav4.utilities.users.MultiValueLegacyExtendedProperty;
import graphjavav4.utilities.users.SingleValueLegacyExtendedProperty;
import java.util.function.BiConsumer;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Objects;
public class MailFolder extends Entity implements Parsable {
    /** The number of immediate child mailFolders in the current mailFolder.  */
    @javax.annotation.Nullable
    public Integer childFolderCount;
    /** The mailFolder's display name.  */
    @javax.annotation.Nullable
    public String displayName;
    /** The unique identifier for the mailFolder's parent mailFolder.  */
    @javax.annotation.Nullable
    public String parentFolderId;
    /** The number of items in the mailFolder.  */
    @javax.annotation.Nullable
    public Integer totalItemCount;
    /** The number of items in the mailFolder marked as unread.  */
    @javax.annotation.Nullable
    public Integer unreadItemCount;
    /** The collection of child folders in the mailFolder.  */
    @javax.annotation.Nullable
    public List<MailFolder> childFolders;
    /** The collection of rules that apply to the user's Inbox folder.  */
    @javax.annotation.Nullable
    public List<MessageRule> messageRules;
    /** The collection of messages in the mailFolder.  */
    @javax.annotation.Nullable
    public List<Message> messages;
    /** The collection of multi-value extended properties defined for the mailFolder. Read-only. Nullable.  */
    @javax.annotation.Nullable
    public List<MultiValueLegacyExtendedProperty> multiValueExtendedProperties;
    /** The collection of single-value extended properties defined for the mailFolder. Read-only. Nullable.  */
    @javax.annotation.Nullable
    public List<SingleValueLegacyExtendedProperty> singleValueExtendedProperties;
    /**
     * Serialiazes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @return a void
     */
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
    /**
     * The serialization information for the current model
     * @return a Map<String, BiConsumer<T, ParseNode>>
     */
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
