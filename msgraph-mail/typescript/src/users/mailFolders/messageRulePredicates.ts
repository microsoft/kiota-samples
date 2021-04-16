import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';
import {Importance} from '../importance';
import {Recipient} from '../recipient';
import {MessageActionFlag} from './messageActionFlag';
import {Sensitivity} from './sensitivity';
import {SizeRange} from './sizeRange';

export class MessageRulePredicates implements Parsable<MessageRulePredicates> {
    public bodyContains?: object[] | undefined;
    public bodyOrSubjectContains?: object[] | undefined;
    public categories?: object[] | undefined;
    public fromAddresses?: Recipient[] | undefined;
    public hasAttachments?: boolean | undefined;
    public headerContains?: object[] | undefined;
    public importance?: Importance | undefined;
    public isApprovalRequest?: boolean | undefined;
    public isAutomaticForward?: boolean | undefined;
    public isAutomaticReply?: boolean | undefined;
    public isEncrypted?: boolean | undefined;
    public isMeetingRequest?: boolean | undefined;
    public isMeetingResponse?: boolean | undefined;
    public isNonDeliveryReport?: boolean | undefined;
    public isPermissionControlled?: boolean | undefined;
    public isReadReceipt?: boolean | undefined;
    public isSigned?: boolean | undefined;
    public isVoicemail?: boolean | undefined;
    public messageActionFlag?: MessageActionFlag | undefined;
    public notSentToMe?: boolean | undefined;
    public recipientContains?: object[] | undefined;
    public senderContains?: object[] | undefined;
    public sensitivity?: Sensitivity | undefined;
    public sentCcMe?: boolean | undefined;
    public sentOnlyToMe?: boolean | undefined;
    public sentToAddresses?: Recipient[] | undefined;
    public sentToMe?: boolean | undefined;
    public sentToOrCcMe?: boolean | undefined;
    public subjectContains?: object[] | undefined;
    public withinSizeRange?: SizeRange | undefined;
    public serialize (writer: SerializationWriter) : void {
        writer.writeCollectionOfPrimitiveValues<object>("bodyContains", this.bodyContains);
        writer.writeCollectionOfPrimitiveValues<object>("bodyOrSubjectContains", this.bodyOrSubjectContains);
        writer.writeCollectionOfPrimitiveValues<object>("categories", this.categories);
        writer.writeCollectionOfObjectValues<Recipient>("fromAddresses", this.fromAddresses);
        writer.writeBooleanValue("hasAttachments", this.hasAttachments);
        writer.writeCollectionOfPrimitiveValues<object>("headerContains", this.headerContains);
        writer.writeObjectValue<Importance>("importance", this.importance);
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
        writer.writeObjectValue<MessageActionFlag>("messageActionFlag", this.messageActionFlag);
        writer.writeBooleanValue("notSentToMe", this.notSentToMe);
        writer.writeCollectionOfPrimitiveValues<object>("recipientContains", this.recipientContains);
        writer.writeCollectionOfPrimitiveValues<object>("senderContains", this.senderContains);
        writer.writeObjectValue<Sensitivity>("sensitivity", this.sensitivity);
        writer.writeBooleanValue("sentCcMe", this.sentCcMe);
        writer.writeBooleanValue("sentOnlyToMe", this.sentOnlyToMe);
        writer.writeCollectionOfObjectValues<Recipient>("sentToAddresses", this.sentToAddresses);
        writer.writeBooleanValue("sentToMe", this.sentToMe);
        writer.writeBooleanValue("sentToOrCcMe", this.sentToOrCcMe);
        writer.writeCollectionOfPrimitiveValues<object>("subjectContains", this.subjectContains);
        writer.writeObjectValue<SizeRange>("withinSizeRange", this.withinSizeRange);
    };
    public deserializeFields () : Map<string, (item: MessageRulePredicates, node: ParseNode) => void> {
        return new Map<string, (item: MessageRulePredicates, node: ParseNode) => void>([
            ["bodyContains", (o, n) => { o.bodyContains = n.getCollectionOfPrimitiveValues<object>(); }],
            ["bodyOrSubjectContains", (o, n) => { o.bodyOrSubjectContains = n.getCollectionOfPrimitiveValues<object>(); }],
            ["categories", (o, n) => { o.categories = n.getCollectionOfPrimitiveValues<object>(); }],
            ["fromAddresses", (o, n) => { o.fromAddresses = n.getCollectionOfObjectValues<Recipient>(Recipient); }],
            ["hasAttachments", (o, n) => { o.hasAttachments = n.getBooleanValue(); }],
            ["headerContains", (o, n) => { o.headerContains = n.getCollectionOfPrimitiveValues<object>(); }],
            ["importance", (o, n) => { o.importance = n.getObjectValue<Importance>(Importance); }],
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
            ["messageActionFlag", (o, n) => { o.messageActionFlag = n.getObjectValue<MessageActionFlag>(MessageActionFlag); }],
            ["notSentToMe", (o, n) => { o.notSentToMe = n.getBooleanValue(); }],
            ["recipientContains", (o, n) => { o.recipientContains = n.getCollectionOfPrimitiveValues<object>(); }],
            ["senderContains", (o, n) => { o.senderContains = n.getCollectionOfPrimitiveValues<object>(); }],
            ["sensitivity", (o, n) => { o.sensitivity = n.getObjectValue<Sensitivity>(Sensitivity); }],
            ["sentCcMe", (o, n) => { o.sentCcMe = n.getBooleanValue(); }],
            ["sentOnlyToMe", (o, n) => { o.sentOnlyToMe = n.getBooleanValue(); }],
            ["sentToAddresses", (o, n) => { o.sentToAddresses = n.getCollectionOfObjectValues<Recipient>(Recipient); }],
            ["sentToMe", (o, n) => { o.sentToMe = n.getBooleanValue(); }],
            ["sentToOrCcMe", (o, n) => { o.sentToOrCcMe = n.getBooleanValue(); }],
            ["subjectContains", (o, n) => { o.subjectContains = n.getCollectionOfPrimitiveValues<object>(); }],
            ["withinSizeRange", (o, n) => { o.withinSizeRange = n.getObjectValue<SizeRange>(SizeRange); }],
        ]);
    };
}
