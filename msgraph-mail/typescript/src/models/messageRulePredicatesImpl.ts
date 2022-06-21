import {createRecipientFromDiscriminatorValue} from './createRecipientFromDiscriminatorValue';
import {createSizeRangeFromDiscriminatorValue} from './createSizeRangeFromDiscriminatorValue';
import {Importance} from './importance';
import {RecipientImpl, SizeRangeImpl} from './index';
import {MessageActionFlag} from './messageActionFlag';
import {MessageRulePredicates} from './messageRulePredicates';
import {Recipient} from './recipient';
import {Sensitivity} from './sensitivity';
import {SizeRange} from './sizeRange';
import {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export class MessageRulePredicatesImpl implements MessageRulePredicates {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well. */
    public additionalData: Record<string, unknown>;
    /** Represents the strings that should appear in the body of an incoming message in order for the condition or exception to apply. */
    public bodyContains?: string[] | undefined;
    /** Represents the strings that should appear in the body or subject of an incoming message in order for the condition or exception to apply. */
    public bodyOrSubjectContains?: string[] | undefined;
    /** Represents the categories that an incoming message should be labeled with in order for the condition or exception to apply. */
    public categories?: string[] | undefined;
    /** Represents the specific sender email addresses of an incoming message in order for the condition or exception to apply. */
    public fromAddresses?: Recipient[] | undefined;
    /** Indicates whether an incoming message must have attachments in order for the condition or exception to apply. */
    public hasAttachments?: boolean | undefined;
    /** Represents the strings that appear in the headers of an incoming message in order for the condition or exception to apply. */
    public headerContains?: string[] | undefined;
    /** The importance property */
    public importance?: Importance | undefined;
    /** Indicates whether an incoming message must be an approval request in order for the condition or exception to apply. */
    public isApprovalRequest?: boolean | undefined;
    /** Indicates whether an incoming message must be automatically forwarded in order for the condition or exception to apply. */
    public isAutomaticForward?: boolean | undefined;
    /** Indicates whether an incoming message must be an auto reply in order for the condition or exception to apply. */
    public isAutomaticReply?: boolean | undefined;
    /** Indicates whether an incoming message must be encrypted in order for the condition or exception to apply. */
    public isEncrypted?: boolean | undefined;
    /** Indicates whether an incoming message must be a meeting request in order for the condition or exception to apply. */
    public isMeetingRequest?: boolean | undefined;
    /** Indicates whether an incoming message must be a meeting response in order for the condition or exception to apply. */
    public isMeetingResponse?: boolean | undefined;
    /** Indicates whether an incoming message must be a non-delivery report in order for the condition or exception to apply. */
    public isNonDeliveryReport?: boolean | undefined;
    /** Indicates whether an incoming message must be permission controlled (RMS-protected) in order for the condition or exception to apply. */
    public isPermissionControlled?: boolean | undefined;
    /** Indicates whether an incoming message must be a read receipt in order for the condition or exception to apply. */
    public isReadReceipt?: boolean | undefined;
    /** Indicates whether an incoming message must be S/MIME-signed in order for the condition or exception to apply. */
    public isSigned?: boolean | undefined;
    /** Indicates whether an incoming message must be a voice mail in order for the condition or exception to apply. */
    public isVoicemail?: boolean | undefined;
    /** The messageActionFlag property */
    public messageActionFlag?: MessageActionFlag | undefined;
    /** Indicates whether the owner of the mailbox must not be a recipient of an incoming message in order for the condition or exception to apply. */
    public notSentToMe?: boolean | undefined;
    /** Represents the strings that appear in either the toRecipients or ccRecipients properties of an incoming message in order for the condition or exception to apply. */
    public recipientContains?: string[] | undefined;
    /** Represents the strings that appear in the from property of an incoming message in order for the condition or exception to apply. */
    public senderContains?: string[] | undefined;
    /** The sensitivity property */
    public sensitivity?: Sensitivity | undefined;
    /** Indicates whether the owner of the mailbox must be in the ccRecipients property of an incoming message in order for the condition or exception to apply. */
    public sentCcMe?: boolean | undefined;
    /** Indicates whether the owner of the mailbox must be the only recipient in an incoming message in order for the condition or exception to apply. */
    public sentOnlyToMe?: boolean | undefined;
    /** Represents the email addresses that an incoming message must have been sent to in order for the condition or exception to apply. */
    public sentToAddresses?: Recipient[] | undefined;
    /** Indicates whether the owner of the mailbox must be in the toRecipients property of an incoming message in order for the condition or exception to apply. */
    public sentToMe?: boolean | undefined;
    /** Indicates whether the owner of the mailbox must be in either a toRecipients or ccRecipients property of an incoming message in order for the condition or exception to apply. */
    public sentToOrCcMe?: boolean | undefined;
    /** Represents the strings that appear in the subject of an incoming message in order for the condition or exception to apply. */
    public subjectContains?: string[] | undefined;
    /** The withinSizeRange property */
    public withinSizeRange?: SizeRange | undefined;
    /**
     * Instantiates a new messageRulePredicates and sets the default values.
     * @param messageRulePredicatesParameterValue 
     */
    public constructor(messageRulePredicatesParameterValue?: MessageRulePredicates | undefined) {
        this.additionalData = messageRulePredicatesParameterValue?.additionalData ? messageRulePredicatesParameterValue?.additionalData! : {};
        this.bodyContains = messageRulePredicatesParameterValue?.bodyContains;
        this.bodyOrSubjectContains = messageRulePredicatesParameterValue?.bodyOrSubjectContains;
        this.categories = messageRulePredicatesParameterValue?.categories;
        this.fromAddresses = messageRulePredicatesParameterValue?.fromAddresses;
        this.hasAttachments = messageRulePredicatesParameterValue?.hasAttachments;
        this.headerContains = messageRulePredicatesParameterValue?.headerContains;
        this.importance = messageRulePredicatesParameterValue?.importance;
        this.isApprovalRequest = messageRulePredicatesParameterValue?.isApprovalRequest;
        this.isAutomaticForward = messageRulePredicatesParameterValue?.isAutomaticForward;
        this.isAutomaticReply = messageRulePredicatesParameterValue?.isAutomaticReply;
        this.isEncrypted = messageRulePredicatesParameterValue?.isEncrypted;
        this.isMeetingRequest = messageRulePredicatesParameterValue?.isMeetingRequest;
        this.isMeetingResponse = messageRulePredicatesParameterValue?.isMeetingResponse;
        this.isNonDeliveryReport = messageRulePredicatesParameterValue?.isNonDeliveryReport;
        this.isPermissionControlled = messageRulePredicatesParameterValue?.isPermissionControlled;
        this.isReadReceipt = messageRulePredicatesParameterValue?.isReadReceipt;
        this.isSigned = messageRulePredicatesParameterValue?.isSigned;
        this.isVoicemail = messageRulePredicatesParameterValue?.isVoicemail;
        this.messageActionFlag = messageRulePredicatesParameterValue?.messageActionFlag;
        this.notSentToMe = messageRulePredicatesParameterValue?.notSentToMe;
        this.recipientContains = messageRulePredicatesParameterValue?.recipientContains;
        this.senderContains = messageRulePredicatesParameterValue?.senderContains;
        this.sensitivity = messageRulePredicatesParameterValue?.sensitivity;
        this.sentCcMe = messageRulePredicatesParameterValue?.sentCcMe;
        this.sentOnlyToMe = messageRulePredicatesParameterValue?.sentOnlyToMe;
        this.sentToAddresses = messageRulePredicatesParameterValue?.sentToAddresses;
        this.sentToMe = messageRulePredicatesParameterValue?.sentToMe;
        this.sentToOrCcMe = messageRulePredicatesParameterValue?.sentToOrCcMe;
        this.subjectContains = messageRulePredicatesParameterValue?.subjectContains;
        this.withinSizeRange = messageRulePredicatesParameterValue?.withinSizeRange;
    };
    /**
     * The deserialization information for the current model
     * @returns a Record<string, (node: ParseNode) => void>
     */
    public getFieldDeserializers() : Record<string, (node: ParseNode) => void> {
        return {
            "bodyContains": n => { this.bodyContains = n.getCollectionOfPrimitiveValues<string>(); },
            "bodyOrSubjectContains": n => { this.bodyOrSubjectContains = n.getCollectionOfPrimitiveValues<string>(); },
            "categories": n => { this.categories = n.getCollectionOfPrimitiveValues<string>(); },
            "fromAddresses": n => { this.fromAddresses = n.getCollectionOfObjectValues<RecipientImpl>(createRecipientFromDiscriminatorValue); },
            "hasAttachments": n => { this.hasAttachments = n.getBooleanValue(); },
            "headerContains": n => { this.headerContains = n.getCollectionOfPrimitiveValues<string>(); },
            "importance": n => { this.importance = n.getEnumValue<Importance>(Importance); },
            "isApprovalRequest": n => { this.isApprovalRequest = n.getBooleanValue(); },
            "isAutomaticForward": n => { this.isAutomaticForward = n.getBooleanValue(); },
            "isAutomaticReply": n => { this.isAutomaticReply = n.getBooleanValue(); },
            "isEncrypted": n => { this.isEncrypted = n.getBooleanValue(); },
            "isMeetingRequest": n => { this.isMeetingRequest = n.getBooleanValue(); },
            "isMeetingResponse": n => { this.isMeetingResponse = n.getBooleanValue(); },
            "isNonDeliveryReport": n => { this.isNonDeliveryReport = n.getBooleanValue(); },
            "isPermissionControlled": n => { this.isPermissionControlled = n.getBooleanValue(); },
            "isReadReceipt": n => { this.isReadReceipt = n.getBooleanValue(); },
            "isSigned": n => { this.isSigned = n.getBooleanValue(); },
            "isVoicemail": n => { this.isVoicemail = n.getBooleanValue(); },
            "messageActionFlag": n => { this.messageActionFlag = n.getEnumValue<MessageActionFlag>(MessageActionFlag); },
            "notSentToMe": n => { this.notSentToMe = n.getBooleanValue(); },
            "recipientContains": n => { this.recipientContains = n.getCollectionOfPrimitiveValues<string>(); },
            "senderContains": n => { this.senderContains = n.getCollectionOfPrimitiveValues<string>(); },
            "sensitivity": n => { this.sensitivity = n.getEnumValue<Sensitivity>(Sensitivity); },
            "sentCcMe": n => { this.sentCcMe = n.getBooleanValue(); },
            "sentOnlyToMe": n => { this.sentOnlyToMe = n.getBooleanValue(); },
            "sentToAddresses": n => { this.sentToAddresses = n.getCollectionOfObjectValues<RecipientImpl>(createRecipientFromDiscriminatorValue); },
            "sentToMe": n => { this.sentToMe = n.getBooleanValue(); },
            "sentToOrCcMe": n => { this.sentToOrCcMe = n.getBooleanValue(); },
            "subjectContains": n => { this.subjectContains = n.getCollectionOfPrimitiveValues<string>(); },
            "withinSizeRange": n => { this.withinSizeRange = n.getObjectValue<SizeRangeImpl>(createSizeRangeFromDiscriminatorValue); },
        };
    };
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     */
    public serialize(writer: SerializationWriter) : void {
        if(!writer) throw new Error("writer cannot be undefined");
        if(this.bodyContains){
            writer.writeCollectionOfPrimitiveValues<string>("bodyContains", this.bodyContains);
        }
        if(this.bodyOrSubjectContains){
            writer.writeCollectionOfPrimitiveValues<string>("bodyOrSubjectContains", this.bodyOrSubjectContains);
        }
        if(this.categories){
            writer.writeCollectionOfPrimitiveValues<string>("categories", this.categories);
        }
        if(this.fromAddresses && this.fromAddresses.length != 0){        const fromAddressesArrValue: RecipientImpl[] = []; this.fromAddresses?.forEach(element => {fromAddressesArrValue.push(new RecipientImpl(element));});
            writer.writeCollectionOfObjectValues<RecipientImpl>("fromAddresses", fromAddressesArrValue);
        }
        if(this.hasAttachments){
            writer.writeBooleanValue("hasAttachments", this.hasAttachments);
        }
        if(this.headerContains){
            writer.writeCollectionOfPrimitiveValues<string>("headerContains", this.headerContains);
        }
        if(this.importance){
            writer.writeEnumValue<Importance>("importance", this.importance);
        }
        if(this.isApprovalRequest){
            writer.writeBooleanValue("isApprovalRequest", this.isApprovalRequest);
        }
        if(this.isAutomaticForward){
            writer.writeBooleanValue("isAutomaticForward", this.isAutomaticForward);
        }
        if(this.isAutomaticReply){
            writer.writeBooleanValue("isAutomaticReply", this.isAutomaticReply);
        }
        if(this.isEncrypted){
            writer.writeBooleanValue("isEncrypted", this.isEncrypted);
        }
        if(this.isMeetingRequest){
            writer.writeBooleanValue("isMeetingRequest", this.isMeetingRequest);
        }
        if(this.isMeetingResponse){
            writer.writeBooleanValue("isMeetingResponse", this.isMeetingResponse);
        }
        if(this.isNonDeliveryReport){
            writer.writeBooleanValue("isNonDeliveryReport", this.isNonDeliveryReport);
        }
        if(this.isPermissionControlled){
            writer.writeBooleanValue("isPermissionControlled", this.isPermissionControlled);
        }
        if(this.isReadReceipt){
            writer.writeBooleanValue("isReadReceipt", this.isReadReceipt);
        }
        if(this.isSigned){
            writer.writeBooleanValue("isSigned", this.isSigned);
        }
        if(this.isVoicemail){
            writer.writeBooleanValue("isVoicemail", this.isVoicemail);
        }
        if(this.messageActionFlag){
            writer.writeEnumValue<MessageActionFlag>("messageActionFlag", this.messageActionFlag);
        }
        if(this.notSentToMe){
            writer.writeBooleanValue("notSentToMe", this.notSentToMe);
        }
        if(this.recipientContains){
            writer.writeCollectionOfPrimitiveValues<string>("recipientContains", this.recipientContains);
        }
        if(this.senderContains){
            writer.writeCollectionOfPrimitiveValues<string>("senderContains", this.senderContains);
        }
        if(this.sensitivity){
            writer.writeEnumValue<Sensitivity>("sensitivity", this.sensitivity);
        }
        if(this.sentCcMe){
            writer.writeBooleanValue("sentCcMe", this.sentCcMe);
        }
        if(this.sentOnlyToMe){
            writer.writeBooleanValue("sentOnlyToMe", this.sentOnlyToMe);
        }
        if(this.sentToAddresses && this.sentToAddresses.length != 0){        const sentToAddressesArrValue: RecipientImpl[] = []; this.sentToAddresses?.forEach(element => {sentToAddressesArrValue.push(new RecipientImpl(element));});
            writer.writeCollectionOfObjectValues<RecipientImpl>("sentToAddresses", sentToAddressesArrValue);
        }
        if(this.sentToMe){
            writer.writeBooleanValue("sentToMe", this.sentToMe);
        }
        if(this.sentToOrCcMe){
            writer.writeBooleanValue("sentToOrCcMe", this.sentToOrCcMe);
        }
        if(this.subjectContains){
            writer.writeCollectionOfPrimitiveValues<string>("subjectContains", this.subjectContains);
        }
        if(this.withinSizeRange){
            writer.writeObjectValue<SizeRangeImpl>("withinSizeRange", new SizeRangeImpl(this.withinSizeRange));
        }
        writer.writeAdditionalData(this.additionalData);
    };
}
