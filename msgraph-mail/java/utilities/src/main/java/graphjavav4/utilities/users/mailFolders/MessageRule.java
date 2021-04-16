package graphjavav4.utilities.users.mailFolders;

import graphjavav4.utilities.users.Entity;
public class MessageRule extends Entity {
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
}
