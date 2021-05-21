import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';
import {Importance} from '../importance';
import {Recipient} from '../recipient';
import {MessageActionFlag} from './messageActionFlag';
import {Sensitivity} from './sensitivity';
import {SizeRange} from './sizeRange';

export class MessageRulePredicates implements Parsable<MessageRulePredicates> {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.  */
    public readonly additionalData: Map<string, unknown> = new Map<string, unknown>();
    /** Represents the strings that should appear in the body of an incoming message in order for the condition or exception to apply.  */
    public bodyContains?: string[] | undefined;
    /** Represents the strings that should appear in the body or subject of an incoming message in order for the condition or exception to apply.  */
    public bodyOrSubjectContains?: string[] | undefined;
    /** Represents the categories that an incoming message should be labeled with in order for the condition or exception to apply.  */
    public categories?: string[] | undefined;
    /** Represents the specific sender email addresses of an incoming message in order for the condition or exception to apply.  */
    public fromAddresses?: Recipient[] | undefined;
    /** Indicates whether an incoming message must have attachments in order for the condition or exception to apply.  */
    public hasAttachments?: boolean | undefined;
    /** Represents the strings that appear in the headers of an incoming message in order for the condition or exception to apply.  */
    public headerContains?: string[] | undefined;
    public importance?: Importance | undefined;
    /** Indicates whether an incoming message must be an approval request in order for the condition or exception to apply.  */
    public isApprovalRequest?: boolean | undefined;
    /** Indicates whether an incoming message must be automatically forwarded in order for the condition or exception to apply.  */
    public isAutomaticForward?: boolean | undefined;
    /** Indicates whether an incoming message must be an auto reply in order for the condition or exception to apply.  */
    public isAutomaticReply?: boolean | undefined;
    /** Indicates whether an incoming message must be encrypted in order for the condition or exception to apply.  */
    public isEncrypted?: boolean | undefined;
    /** Indicates whether an incoming message must be a meeting request in order for the condition or exception to apply.  */
    public isMeetingRequest?: boolean | undefined;
    /** Indicates whether an incoming message must be a meeting response in order for the condition or exception to apply.  */
    public isMeetingResponse?: boolean | undefined;
    /** Indicates whether an incoming message must be a non-delivery report in order for the condition or exception to apply.  */
    public isNonDeliveryReport?: boolean | undefined;
    /** Indicates whether an incoming message must be permission controlled (RMS-protected) in order for the condition or exception to apply.  */
    public isPermissionControlled?: boolean | undefined;
    /** Indicates whether an incoming message must be a read receipt in order for the condition or exception to apply.  */
    public isReadReceipt?: boolean | undefined;
    /** Indicates whether an incoming message must be S/MIME-signed in order for the condition or exception to apply.  */
    public isSigned?: boolean | undefined;
    /** Indicates whether an incoming message must be a voice mail in order for the condition or exception to apply.  */
    public isVoicemail?: boolean | undefined;
    public messageActionFlag?: MessageActionFlag | undefined;
    /** Indicates whether the owner of the mailbox must not be a recipient of an incoming message in order for the condition or exception to apply.  */
    public notSentToMe?: boolean | undefined;
    /** Represents the strings that appear in either the toRecipients or ccRecipients properties of an incoming message in order for the condition or exception to apply.  */
    public recipientContains?: string[] | undefined;
    /** Represents the strings that appear in the from property of an incoming message in order for the condition or exception to apply.  */
    public senderContains?: string[] | undefined;
    public sensitivity?: Sensitivity | undefined;
    /** Indicates whether the owner of the mailbox must be in the ccRecipients property of an incoming message in order for the condition or exception to apply.  */
    public sentCcMe?: boolean | undefined;
    /** Indicates whether the owner of the mailbox must be the only recipient in an incoming message in order for the condition or exception to apply.  */
    public sentOnlyToMe?: boolean | undefined;
    /** Represents the email addresses that an incoming message must have been sent to in order for the condition or exception to apply.  */
    public sentToAddresses?: Recipient[] | undefined;
    /** Indicates whether the owner of the mailbox must be in the toRecipients property of an incoming message in order for the condition or exception to apply.  */
    public sentToMe?: boolean | undefined;
    /** Indicates whether the owner of the mailbox must be in either a toRecipients or ccRecipients property of an incoming message in order for the condition or exception to apply.  */
    public sentToOrCcMe?: boolean | undefined;
    /** Represents the strings that appear in the subject of an incoming message in order for the condition or exception to apply.  */
    public subjectContains?: string[] | undefined;
    public withinSizeRange?: SizeRange | undefined;
    /**
     * The serialization information for the current model
     * @returns a Map<string, (item: MessageRulePredicates, node: ParseNode) => void>
     */
    public deserializeFields () : Map<string, (item: MessageRulePredicates, node: ParseNode) => void> {
        return new Map<string, (item: MessageRulePredicates, node: ParseNode) => void>([
            ["bodyContains", (o, n) => { o.bodyContains = n.getCollectionOfPrimitiveValues<string>(); }],
            ["bodyOrSubjectContains", (o, n) => { o.bodyOrSubjectContains = n.getCollectionOfPrimitiveValues<string>(); }],
            ["categories", (o, n) => { o.categories = n.getCollectionOfPrimitiveValues<string>(); }],
            ["fromAddresses", (o, n) => { o.fromAddresses = n.getCollectionOfObjectValues<Recipient>(Recipient); }],
            ["hasAttachments", (o, n) => { o.hasAttachments = n.getBooleanValue(); }],
            ["headerContains", (o, n) => { o.headerContains = n.getCollectionOfPrimitiveValues<string>(); }],
            ["importance", (o, n) => { o.importance = n.getEnumValue<Importance>(Importance); }],
            ["isApprovalRequest", (o, n) => { o.isApprovalRequest = n.getBooleanValue(); }],
            ["isAutomaticForward", (o, n) => { o.isAutomaticForward = n.getBooleanValue(); }],
            ["isAutomaticReply", (o, n) => { o.isAutomaticReply = n.getBooleanValue(); }],
            ["isEncrypted", (o, n) => { o.isEncrypted = n.getBooleanValue(); }],
            ["isMeetingRequest", (o, n) => { o.isMeetingRequest = n.getBooleanValue(); }],
            ["isMeetingResponse", (o, n) => { o.isMeetingResponse = n.getBooleanValue(); }],
            ["isNonDeliveryReport", (o, n) => { o.isNonDeliveryReport = n.getBooleanValue(); }],
            ["isPermissionControlled", (o, n) => { o.isPermissionControlled = n.getBooleanValue(); }],
            ["isReadReceipt", (o, n) => { o.isReadReceipt = n.getBooleanValue(); }],
            ["isSigned", (o, n) => { o.isSigned = n.getBooleanValue(); }],
            ["isVoicemail", (o, n) => { o.isVoicemail = n.getBooleanValue(); }],
            ["messageActionFlag", (o, n) => { o.messageActionFlag = n.getEnumValue<MessageActionFlag>(MessageActionFlag); }],
            ["notSentToMe", (o, n) => { o.notSentToMe = n.getBooleanValue(); }],
            ["recipientContains", (o, n) => { o.recipientContains = n.getCollectionOfPrimitiveValues<string>(); }],
            ["senderContains", (o, n) => { o.senderContains = n.getCollectionOfPrimitiveValues<string>(); }],
            ["sensitivity", (o, n) => { o.sensitivity = n.getEnumValue<Sensitivity>(Sensitivity); }],
            ["sentCcMe", (o, n) => { o.sentCcMe = n.getBooleanValue(); }],
            ["sentOnlyToMe", (o, n) => { o.sentOnlyToMe = n.getBooleanValue(); }],
            ["sentToAddresses", (o, n) => { o.sentToAddresses = n.getCollectionOfObjectValues<Recipient>(Recipient); }],
            ["sentToMe", (o, n) => { o.sentToMe = n.getBooleanValue(); }],
            ["sentToOrCcMe", (o, n) => { o.sentToOrCcMe = n.getBooleanValue(); }],
            ["subjectContains", (o, n) => { o.subjectContains = n.getCollectionOfPrimitiveValues<string>(); }],
            ["withinSizeRange", (o, n) => { o.withinSizeRange = n.getObjectValue<SizeRange>(SizeRange); }],
        ]);
    };
    /**
     * Serialiazes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @returns a void
     */
    public serialize (writer: SerializationWriter) : void {
        writer.writeCollectionOfPrimitiveValues<string>("bodyContains", this.bodyContains);
        writer.writeCollectionOfPrimitiveValues<string>("bodyOrSubjectContains", this.bodyOrSubjectContains);
        writer.writeCollectionOfPrimitiveValues<string>("categories", this.categories);
        writer.writeCollectionOfObjectValues<Recipient>("fromAddresses", this.fromAddresses);
        writer.writeBooleanValue("hasAttachments", this.hasAttachments);
        writer.writeCollectionOfPrimitiveValues<string>("headerContains", this.headerContains);
        writer.writeEnumValue<Importance>("importance", this.importance);
        writer.writeBooleanValue("isApprovalRequest", this.isApprovalRequest);
        writer.writeBooleanValue("isAutomaticForward", this.isAutomaticForward);
        writer.writeBooleanValue("isAutomaticReply", this.isAutomaticReply);
        writer.writeBooleanValue("isEncrypted", this.isEncrypted);
        writer.writeBooleanValue("isMeetingRequest", this.isMeetingRequest);
        writer.writeBooleanValue("isMeetingResponse", this.isMeetingResponse);
        writer.writeBooleanValue("isNonDeliveryReport", this.isNonDeliveryReport);
        writer.writeBooleanValue("isPermissionControlled", this.isPermissionControlled);
        writer.writeBooleanValue("isReadReceipt", this.isReadReceipt);
        writer.writeBooleanValue("isSigned", this.isSigned);
        writer.writeBooleanValue("isVoicemail", this.isVoicemail);
        writer.writeEnumValue<MessageActionFlag>("messageActionFlag", this.messageActionFlag);
        writer.writeBooleanValue("notSentToMe", this.notSentToMe);
        writer.writeCollectionOfPrimitiveValues<string>("recipientContains", this.recipientContains);
        writer.writeCollectionOfPrimitiveValues<string>("senderContains", this.senderContains);
        writer.writeEnumValue<Sensitivity>("sensitivity", this.sensitivity);
        writer.writeBooleanValue("sentCcMe", this.sentCcMe);
        writer.writeBooleanValue("sentOnlyToMe", this.sentOnlyToMe);
        writer.writeCollectionOfObjectValues<Recipient>("sentToAddresses", this.sentToAddresses);
        writer.writeBooleanValue("sentToMe", this.sentToMe);
        writer.writeBooleanValue("sentToOrCcMe", this.sentToOrCcMe);
        writer.writeCollectionOfPrimitiveValues<string>("subjectContains", this.subjectContains);
        writer.writeObjectValue<SizeRange>("withinSizeRange", this.withinSizeRange);
        writer.writeAdditionalData(this.additionalData);
    };
}
