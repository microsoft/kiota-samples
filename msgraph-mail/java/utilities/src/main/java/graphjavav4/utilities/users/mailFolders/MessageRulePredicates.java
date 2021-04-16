package graphjavav4.utilities.users.mailFolders;

import graphjavav4.utilities.users.Recipient;
import graphjavav4.utilities.users.Importance;
import java.util.List;
public class MessageRulePredicates {
    @javax.annotation.Nullable
    public List<Object> bodyContains;
    @javax.annotation.Nullable
    public List<Object> bodyOrSubjectContains;
    @javax.annotation.Nullable
    public List<Object> categories;
    @javax.annotation.Nullable
    public List<Recipient> fromAddresses;
    @javax.annotation.Nullable
    public Boolean hasAttachments;
    @javax.annotation.Nullable
    public List<Object> headerContains;
    @javax.annotation.Nullable
    public Importance importance;
    @javax.annotation.Nullable
    public Boolean isApprovalRequest;
    @javax.annotation.Nullable
    public Boolean isAutomaticForward;
    @javax.annotation.Nullable
    public Boolean isAutomaticReply;
    @javax.annotation.Nullable
    public Boolean isEncrypted;
    @javax.annotation.Nullable
    public Boolean isMeetingRequest;
    @javax.annotation.Nullable
    public Boolean isMeetingResponse;
    @javax.annotation.Nullable
    public Boolean isNonDeliveryReport;
    @javax.annotation.Nullable
    public Boolean isPermissionControlled;
    @javax.annotation.Nullable
    public Boolean isReadReceipt;
    @javax.annotation.Nullable
    public Boolean isSigned;
    @javax.annotation.Nullable
    public Boolean isVoicemail;
    @javax.annotation.Nullable
    public MessageActionFlag messageActionFlag;
    @javax.annotation.Nullable
    public Boolean notSentToMe;
    @javax.annotation.Nullable
    public List<Object> recipientContains;
    @javax.annotation.Nullable
    public List<Object> senderContains;
    @javax.annotation.Nullable
    public Sensitivity sensitivity;
    @javax.annotation.Nullable
    public Boolean sentCcMe;
    @javax.annotation.Nullable
    public Boolean sentOnlyToMe;
    @javax.annotation.Nullable
    public List<Recipient> sentToAddresses;
    @javax.annotation.Nullable
    public Boolean sentToMe;
    @javax.annotation.Nullable
    public Boolean sentToOrCcMe;
    @javax.annotation.Nullable
    public List<Object> subjectContains;
    @javax.annotation.Nullable
    public SizeRange withinSizeRange;
}
