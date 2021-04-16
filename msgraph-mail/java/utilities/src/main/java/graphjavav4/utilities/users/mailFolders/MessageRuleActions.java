package graphjavav4.utilities.users.mailFolders;

import graphjavav4.utilities.users.Recipient;
import graphjavav4.utilities.users.Importance;
import java.util.List;
public class MessageRuleActions {
    @javax.annotation.Nullable
    public List<Object> assignCategories;
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
}
