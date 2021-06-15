package graphjavav4.utilities.users;

import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import java.util.function.BiConsumer;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Objects;
public class MailFolder extends Entity implements Parsable {
    /** The number of immediate child mailFolders in the current mailFolder.  */
    @javax.annotation.Nullable
    public Integer childFolderCount;
    /** The collection of child folders in the mailFolder.  */
    @javax.annotation.Nullable
    public List<MailFolder> childFolders;
    /** The mailFolder's display name.  */
    @javax.annotation.Nullable
    public String displayName;
    /** Indicates whether the mailFolder is hidden. This property can be set only when creating the folder. Find more information in Hidden mail folders.  */
    @javax.annotation.Nullable
    public Boolean isHidden;
    /** The collection of rules that apply to the user's Inbox folder.  */
    @javax.annotation.Nullable
    public List<MessageRule> messageRules;
    /** The collection of messages in the mailFolder.  */
    @javax.annotation.Nullable
    public List<Message> messages;
    /** The collection of multi-value extended properties defined for the mailFolder. Read-only. Nullable.  */
    @javax.annotation.Nullable
    public List<MultiValueLegacyExtendedProperty> multiValueExtendedProperties;
    /** The unique identifier for the mailFolder's parent mailFolder.  */
    @javax.annotation.Nullable
    public String parentFolderId;
    /** The collection of single-value extended properties defined for the mailFolder. Read-only. Nullable.  */
    @javax.annotation.Nullable
    public List<SingleValueLegacyExtendedProperty> singleValueExtendedProperties;
    /** The number of items in the mailFolder.  */
    @javax.annotation.Nullable
    public Integer totalItemCount;
    /** The number of items in the mailFolder marked as unread.  */
    @javax.annotation.Nullable
    public Integer unreadItemCount;
    /**
     * The deserialization information for the current model
     * @return a Map<String, BiConsumer<T, ParseNode>>
     */
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getFieldDeserializers() {
        return new HashMap<>(super.getFieldDeserializers()) {{
            this.put("childFolderCount", (o, n) -> { ((MailFolder)o).childFolderCount = n.getIntegerValue(); });
            this.put("childFolders", (o, n) -> { ((MailFolder)o).childFolders = n.getCollectionOfObjectValues(MailFolder.class); });
            this.put("displayName", (o, n) -> { ((MailFolder)o).displayName = n.getStringValue(); });
            this.put("isHidden", (o, n) -> { ((MailFolder)o).isHidden = n.getBooleanValue(); });
            this.put("messageRules", (o, n) -> { ((MailFolder)o).messageRules = n.getCollectionOfObjectValues(MessageRule.class); });
            this.put("messages", (o, n) -> { ((MailFolder)o).messages = n.getCollectionOfObjectValues(Message.class); });
            this.put("multiValueExtendedProperties", (o, n) -> { ((MailFolder)o).multiValueExtendedProperties = n.getCollectionOfObjectValues(MultiValueLegacyExtendedProperty.class); });
            this.put("parentFolderId", (o, n) -> { ((MailFolder)o).parentFolderId = n.getStringValue(); });
            this.put("singleValueExtendedProperties", (o, n) -> { ((MailFolder)o).singleValueExtendedProperties = n.getCollectionOfObjectValues(SingleValueLegacyExtendedProperty.class); });
            this.put("totalItemCount", (o, n) -> { ((MailFolder)o).totalItemCount = n.getIntegerValue(); });
            this.put("unreadItemCount", (o, n) -> { ((MailFolder)o).unreadItemCount = n.getIntegerValue(); });
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
        writer.writeIntegerValue("childFolderCount", childFolderCount);
        writer.writeCollectionOfObjectValues("childFolders", childFolders);
        writer.writeStringValue("displayName", displayName);
        writer.writeBooleanValue("isHidden", isHidden);
        writer.writeCollectionOfObjectValues("messageRules", messageRules);
        writer.writeCollectionOfObjectValues("messages", messages);
        writer.writeCollectionOfObjectValues("multiValueExtendedProperties", multiValueExtendedProperties);
        writer.writeStringValue("parentFolderId", parentFolderId);
        writer.writeCollectionOfObjectValues("singleValueExtendedProperties", singleValueExtendedProperties);
        writer.writeIntegerValue("totalItemCount", totalItemCount);
        writer.writeIntegerValue("unreadItemCount", unreadItemCount);
    }
}
