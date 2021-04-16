package graphjavav4.utilities.users.mailFolders;

import graphjavav4.utilities.users.Message;
import graphjavav4.utilities.users.MultiValueLegacyExtendedProperty;
import graphjavav4.utilities.users.SingleValueLegacyExtendedProperty;
import graphjavav4.utilities.users.Entity;
import java.util.List;
public class MailFolder extends Entity {
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
}
