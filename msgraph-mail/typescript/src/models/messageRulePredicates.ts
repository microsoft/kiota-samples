import { Importance } from './importance';
import { MessageActionFlag } from './messageActionFlag';
import { createRecipientFromDiscriminatorValue, serializeRecipient, type Recipient } from './recipient';
import { Sensitivity } from './sensitivity';
import { createSizeRangeFromDiscriminatorValue, serializeSizeRange, type SizeRange } from './sizeRange';
import { type AdditionalDataHolder, type Parsable, type ParseNode, type SerializationWriter } from '@microsoft/kiota-abstractions';

export function createMessageRulePredicatesFromDiscriminatorValue(parseNode: ParseNode | undefined) {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return deserializeIntoMessageRulePredicates;
}
export function deserializeIntoMessageRulePredicates(messageRulePredicates: MessageRulePredicates | undefined = {} as MessageRulePredicates) : Record<string, (node: ParseNode) => void> {
    return {
        "bodyContains": n => { messageRulePredicates.bodyContains = n.getCollectionOfPrimitiveValues<string>(); },
        "bodyOrSubjectContains": n => { messageRulePredicates.bodyOrSubjectContains = n.getCollectionOfPrimitiveValues<string>(); },
        "categories": n => { messageRulePredicates.categories = n.getCollectionOfPrimitiveValues<string>(); },
        "fromAddresses": n => { messageRulePredicates.fromAddresses = n.getCollectionOfObjectValues<Recipient>(createRecipientFromDiscriminatorValue); },
        "hasAttachments": n => { messageRulePredicates.hasAttachments = n.getBooleanValue(); },
        "headerContains": n => { messageRulePredicates.headerContains = n.getCollectionOfPrimitiveValues<string>(); },
        "importance": n => { messageRulePredicates.importance = n.getEnumValue<Importance>(Importance); },
        "isApprovalRequest": n => { messageRulePredicates.isApprovalRequest = n.getBooleanValue(); },
        "isAutomaticForward": n => { messageRulePredicates.isAutomaticForward = n.getBooleanValue(); },
        "isAutomaticReply": n => { messageRulePredicates.isAutomaticReply = n.getBooleanValue(); },
        "isEncrypted": n => { messageRulePredicates.isEncrypted = n.getBooleanValue(); },
        "isMeetingRequest": n => { messageRulePredicates.isMeetingRequest = n.getBooleanValue(); },
        "isMeetingResponse": n => { messageRulePredicates.isMeetingResponse = n.getBooleanValue(); },
        "isNonDeliveryReport": n => { messageRulePredicates.isNonDeliveryReport = n.getBooleanValue(); },
        "isPermissionControlled": n => { messageRulePredicates.isPermissionControlled = n.getBooleanValue(); },
        "isReadReceipt": n => { messageRulePredicates.isReadReceipt = n.getBooleanValue(); },
        "isSigned": n => { messageRulePredicates.isSigned = n.getBooleanValue(); },
        "isVoicemail": n => { messageRulePredicates.isVoicemail = n.getBooleanValue(); },
        "messageActionFlag": n => { messageRulePredicates.messageActionFlag = n.getEnumValue<MessageActionFlag>(MessageActionFlag); },
        "notSentToMe": n => { messageRulePredicates.notSentToMe = n.getBooleanValue(); },
        "recipientContains": n => { messageRulePredicates.recipientContains = n.getCollectionOfPrimitiveValues<string>(); },
        "senderContains": n => { messageRulePredicates.senderContains = n.getCollectionOfPrimitiveValues<string>(); },
        "sensitivity": n => { messageRulePredicates.sensitivity = n.getEnumValue<Sensitivity>(Sensitivity); },
        "sentCcMe": n => { messageRulePredicates.sentCcMe = n.getBooleanValue(); },
        "sentOnlyToMe": n => { messageRulePredicates.sentOnlyToMe = n.getBooleanValue(); },
        "sentToAddresses": n => { messageRulePredicates.sentToAddresses = n.getCollectionOfObjectValues<Recipient>(createRecipientFromDiscriminatorValue); },
        "sentToMe": n => { messageRulePredicates.sentToMe = n.getBooleanValue(); },
        "sentToOrCcMe": n => { messageRulePredicates.sentToOrCcMe = n.getBooleanValue(); },
        "subjectContains": n => { messageRulePredicates.subjectContains = n.getCollectionOfPrimitiveValues<string>(); },
        "withinSizeRange": n => { messageRulePredicates.withinSizeRange = n.getObjectValue<SizeRange>(createSizeRangeFromDiscriminatorValue); },
    }
}
export interface MessageRulePredicates extends AdditionalDataHolder, Parsable {
    /**
     * Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     */
    additionalData?: Record<string, unknown>;
    /**
     * Represents the strings that should appear in the body of an incoming message in order for the condition or exception to apply.
     */
    bodyContains?: string[] | undefined;
    /**
     * Represents the strings that should appear in the body or subject of an incoming message in order for the condition or exception to apply.
     */
    bodyOrSubjectContains?: string[] | undefined;
    /**
     * Represents the categories that an incoming message should be labeled with in order for the condition or exception to apply.
     */
    categories?: string[] | undefined;
    /**
     * Represents the specific sender email addresses of an incoming message in order for the condition or exception to apply.
     */
    fromAddresses?: Recipient[] | undefined;
    /**
     * Indicates whether an incoming message must have attachments in order for the condition or exception to apply.
     */
    hasAttachments?: boolean | undefined;
    /**
     * Represents the strings that appear in the headers of an incoming message in order for the condition or exception to apply.
     */
    headerContains?: string[] | undefined;
    /**
     * The importance property
     */
    importance?: Importance | undefined;
    /**
     * Indicates whether an incoming message must be an approval request in order for the condition or exception to apply.
     */
    isApprovalRequest?: boolean | undefined;
    /**
     * Indicates whether an incoming message must be automatically forwarded in order for the condition or exception to apply.
     */
    isAutomaticForward?: boolean | undefined;
    /**
     * Indicates whether an incoming message must be an auto reply in order for the condition or exception to apply.
     */
    isAutomaticReply?: boolean | undefined;
    /**
     * Indicates whether an incoming message must be encrypted in order for the condition or exception to apply.
     */
    isEncrypted?: boolean | undefined;
    /**
     * Indicates whether an incoming message must be a meeting request in order for the condition or exception to apply.
     */
    isMeetingRequest?: boolean | undefined;
    /**
     * Indicates whether an incoming message must be a meeting response in order for the condition or exception to apply.
     */
    isMeetingResponse?: boolean | undefined;
    /**
     * Indicates whether an incoming message must be a non-delivery report in order for the condition or exception to apply.
     */
    isNonDeliveryReport?: boolean | undefined;
    /**
     * Indicates whether an incoming message must be permission controlled (RMS-protected) in order for the condition or exception to apply.
     */
    isPermissionControlled?: boolean | undefined;
    /**
     * Indicates whether an incoming message must be a read receipt in order for the condition or exception to apply.
     */
    isReadReceipt?: boolean | undefined;
    /**
     * Indicates whether an incoming message must be S/MIME-signed in order for the condition or exception to apply.
     */
    isSigned?: boolean | undefined;
    /**
     * Indicates whether an incoming message must be a voice mail in order for the condition or exception to apply.
     */
    isVoicemail?: boolean | undefined;
    /**
     * The messageActionFlag property
     */
    messageActionFlag?: MessageActionFlag | undefined;
    /**
     * Indicates whether the owner of the mailbox must not be a recipient of an incoming message in order for the condition or exception to apply.
     */
    notSentToMe?: boolean | undefined;
    /**
     * Represents the strings that appear in either the toRecipients or ccRecipients properties of an incoming message in order for the condition or exception to apply.
     */
    recipientContains?: string[] | undefined;
    /**
     * Represents the strings that appear in the from property of an incoming message in order for the condition or exception to apply.
     */
    senderContains?: string[] | undefined;
    /**
     * The sensitivity property
     */
    sensitivity?: Sensitivity | undefined;
    /**
     * Indicates whether the owner of the mailbox must be in the ccRecipients property of an incoming message in order for the condition or exception to apply.
     */
    sentCcMe?: boolean | undefined;
    /**
     * Indicates whether the owner of the mailbox must be the only recipient in an incoming message in order for the condition or exception to apply.
     */
    sentOnlyToMe?: boolean | undefined;
    /**
     * Represents the email addresses that an incoming message must have been sent to in order for the condition or exception to apply.
     */
    sentToAddresses?: Recipient[] | undefined;
    /**
     * Indicates whether the owner of the mailbox must be in the toRecipients property of an incoming message in order for the condition or exception to apply.
     */
    sentToMe?: boolean | undefined;
    /**
     * Indicates whether the owner of the mailbox must be in either a toRecipients or ccRecipients property of an incoming message in order for the condition or exception to apply.
     */
    sentToOrCcMe?: boolean | undefined;
    /**
     * Represents the strings that appear in the subject of an incoming message in order for the condition or exception to apply.
     */
    subjectContains?: string[] | undefined;
    /**
     * The withinSizeRange property
     */
    withinSizeRange?: SizeRange | undefined;
}
export function serializeMessageRulePredicates(writer: SerializationWriter, messageRulePredicates: MessageRulePredicates | undefined = {} as MessageRulePredicates) : void {
        writer.writeCollectionOfPrimitiveValues<string>("bodyContains", messageRulePredicates.bodyContains);
        writer.writeCollectionOfPrimitiveValues<string>("bodyOrSubjectContains", messageRulePredicates.bodyOrSubjectContains);
        writer.writeCollectionOfPrimitiveValues<string>("categories", messageRulePredicates.categories);
        writer.writeCollectionOfObjectValues<Recipient>("fromAddresses", messageRulePredicates.fromAddresses, );
        writer.writeBooleanValue("hasAttachments", messageRulePredicates.hasAttachments);
        writer.writeCollectionOfPrimitiveValues<string>("headerContains", messageRulePredicates.headerContains);
        writer.writeEnumValue<Importance>("importance", messageRulePredicates.importance);
        writer.writeBooleanValue("isApprovalRequest", messageRulePredicates.isApprovalRequest);
        writer.writeBooleanValue("isAutomaticForward", messageRulePredicates.isAutomaticForward);
        writer.writeBooleanValue("isAutomaticReply", messageRulePredicates.isAutomaticReply);
        writer.writeBooleanValue("isEncrypted", messageRulePredicates.isEncrypted);
        writer.writeBooleanValue("isMeetingRequest", messageRulePredicates.isMeetingRequest);
        writer.writeBooleanValue("isMeetingResponse", messageRulePredicates.isMeetingResponse);
        writer.writeBooleanValue("isNonDeliveryReport", messageRulePredicates.isNonDeliveryReport);
        writer.writeBooleanValue("isPermissionControlled", messageRulePredicates.isPermissionControlled);
        writer.writeBooleanValue("isReadReceipt", messageRulePredicates.isReadReceipt);
        writer.writeBooleanValue("isSigned", messageRulePredicates.isSigned);
        writer.writeBooleanValue("isVoicemail", messageRulePredicates.isVoicemail);
        writer.writeEnumValue<MessageActionFlag>("messageActionFlag", messageRulePredicates.messageActionFlag);
        writer.writeBooleanValue("notSentToMe", messageRulePredicates.notSentToMe);
        writer.writeCollectionOfPrimitiveValues<string>("recipientContains", messageRulePredicates.recipientContains);
        writer.writeCollectionOfPrimitiveValues<string>("senderContains", messageRulePredicates.senderContains);
        writer.writeEnumValue<Sensitivity>("sensitivity", messageRulePredicates.sensitivity);
        writer.writeBooleanValue("sentCcMe", messageRulePredicates.sentCcMe);
        writer.writeBooleanValue("sentOnlyToMe", messageRulePredicates.sentOnlyToMe);
        writer.writeCollectionOfObjectValues<Recipient>("sentToAddresses", messageRulePredicates.sentToAddresses, );
        writer.writeBooleanValue("sentToMe", messageRulePredicates.sentToMe);
        writer.writeBooleanValue("sentToOrCcMe", messageRulePredicates.sentToOrCcMe);
        writer.writeCollectionOfPrimitiveValues<string>("subjectContains", messageRulePredicates.subjectContains);
        writer.writeObjectValue<SizeRange>("withinSizeRange", messageRulePredicates.withinSizeRange, );
        writer.writeAdditionalData(messageRulePredicates.additionalData);
}
