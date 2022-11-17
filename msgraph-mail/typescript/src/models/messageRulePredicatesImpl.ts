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
    private _additionalData?: Record<string, unknown> | undefined;
    /** Represents the strings that should appear in the body of an incoming message in order for the condition or exception to apply. */
    private _bodyContains?: string[] | undefined;
    /** Represents the strings that should appear in the body or subject of an incoming message in order for the condition or exception to apply. */
    private _bodyOrSubjectContains?: string[] | undefined;
    /** Represents the categories that an incoming message should be labeled with in order for the condition or exception to apply. */
    private _categories?: string[] | undefined;
    /** Represents the specific sender email addresses of an incoming message in order for the condition or exception to apply. */
    private _fromAddresses?: Recipient[] | undefined;
    /** Indicates whether an incoming message must have attachments in order for the condition or exception to apply. */
    private _hasAttachments?: boolean | undefined;
    /** Represents the strings that appear in the headers of an incoming message in order for the condition or exception to apply. */
    private _headerContains?: string[] | undefined;
    /** The importance property */
    private _importance?: Importance | undefined;
    /** Indicates whether an incoming message must be an approval request in order for the condition or exception to apply. */
    private _isApprovalRequest?: boolean | undefined;
    /** Indicates whether an incoming message must be automatically forwarded in order for the condition or exception to apply. */
    private _isAutomaticForward?: boolean | undefined;
    /** Indicates whether an incoming message must be an auto reply in order for the condition or exception to apply. */
    private _isAutomaticReply?: boolean | undefined;
    /** Indicates whether an incoming message must be encrypted in order for the condition or exception to apply. */
    private _isEncrypted?: boolean | undefined;
    /** Indicates whether an incoming message must be a meeting request in order for the condition or exception to apply. */
    private _isMeetingRequest?: boolean | undefined;
    /** Indicates whether an incoming message must be a meeting response in order for the condition or exception to apply. */
    private _isMeetingResponse?: boolean | undefined;
    /** Indicates whether an incoming message must be a non-delivery report in order for the condition or exception to apply. */
    private _isNonDeliveryReport?: boolean | undefined;
    /** Indicates whether an incoming message must be permission controlled (RMS-protected) in order for the condition or exception to apply. */
    private _isPermissionControlled?: boolean | undefined;
    /** Indicates whether an incoming message must be a read receipt in order for the condition or exception to apply. */
    private _isReadReceipt?: boolean | undefined;
    /** Indicates whether an incoming message must be S/MIME-signed in order for the condition or exception to apply. */
    private _isSigned?: boolean | undefined;
    /** Indicates whether an incoming message must be a voice mail in order for the condition or exception to apply. */
    private _isVoicemail?: boolean | undefined;
    /** The messageActionFlag property */
    private _messageActionFlag?: MessageActionFlag | undefined;
    /** Indicates whether the owner of the mailbox must not be a recipient of an incoming message in order for the condition or exception to apply. */
    private _notSentToMe?: boolean | undefined;
    /** Represents the strings that appear in either the toRecipients or ccRecipients properties of an incoming message in order for the condition or exception to apply. */
    private _recipientContains?: string[] | undefined;
    /** Represents the strings that appear in the from property of an incoming message in order for the condition or exception to apply. */
    private _senderContains?: string[] | undefined;
    /** The sensitivity property */
    private _sensitivity?: Sensitivity | undefined;
    /** Indicates whether the owner of the mailbox must be in the ccRecipients property of an incoming message in order for the condition or exception to apply. */
    private _sentCcMe?: boolean | undefined;
    /** Indicates whether the owner of the mailbox must be the only recipient in an incoming message in order for the condition or exception to apply. */
    private _sentOnlyToMe?: boolean | undefined;
    /** Represents the email addresses that an incoming message must have been sent to in order for the condition or exception to apply. */
    private _sentToAddresses?: Recipient[] | undefined;
    /** Indicates whether the owner of the mailbox must be in the toRecipients property of an incoming message in order for the condition or exception to apply. */
    private _sentToMe?: boolean | undefined;
    /** Indicates whether the owner of the mailbox must be in either a toRecipients or ccRecipients property of an incoming message in order for the condition or exception to apply. */
    private _sentToOrCcMe?: boolean | undefined;
    /** Represents the strings that appear in the subject of an incoming message in order for the condition or exception to apply. */
    private _subjectContains?: string[] | undefined;
    /** The withinSizeRange property */
    private _withinSizeRange?: SizeRange | undefined;
    /**
     * Gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     * @returns a Record<string, unknown>
     */
    public get additionalData() {
        return this._additionalData;
    };
    /**
     * Sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     * @param value Value to set for the AdditionalData property.
     */
    public set additionalData(value: Record<string, unknown> | undefined) {
        if(value) {
            this._additionalData = value;
        }
    };
    /**
     * Gets the bodyContains property value. Represents the strings that should appear in the body of an incoming message in order for the condition or exception to apply.
     * @returns a string
     */
    public get bodyContains() {
        return this._bodyContains;
    };
    /**
     * Sets the bodyContains property value. Represents the strings that should appear in the body of an incoming message in order for the condition or exception to apply.
     * @param value Value to set for the bodyContains property.
     */
    public set bodyContains(value: string[] | undefined) {
        if(value) {
            this._bodyContains = value;
        }
    };
    /**
     * Gets the bodyOrSubjectContains property value. Represents the strings that should appear in the body or subject of an incoming message in order for the condition or exception to apply.
     * @returns a string
     */
    public get bodyOrSubjectContains() {
        return this._bodyOrSubjectContains;
    };
    /**
     * Sets the bodyOrSubjectContains property value. Represents the strings that should appear in the body or subject of an incoming message in order for the condition or exception to apply.
     * @param value Value to set for the bodyOrSubjectContains property.
     */
    public set bodyOrSubjectContains(value: string[] | undefined) {
        if(value) {
            this._bodyOrSubjectContains = value;
        }
    };
    /**
     * Gets the categories property value. Represents the categories that an incoming message should be labeled with in order for the condition or exception to apply.
     * @returns a string
     */
    public get categories() {
        return this._categories;
    };
    /**
     * Sets the categories property value. Represents the categories that an incoming message should be labeled with in order for the condition or exception to apply.
     * @param value Value to set for the categories property.
     */
    public set categories(value: string[] | undefined) {
        if(value) {
            this._categories = value;
        }
    };
    /**
     * Instantiates a new messageRulePredicates and sets the default values.
     * @param messageRulePredicatesParameterValue 
     */
    public constructor(messageRulePredicatesParameterValue?: MessageRulePredicates | undefined) {
        this._additionalData = {};
        this._additionalData = messageRulePredicatesParameterValue?.additionalData ? messageRulePredicatesParameterValue?.additionalData! : {};
        this._bodyContains = messageRulePredicatesParameterValue?.bodyContains;
        this._bodyOrSubjectContains = messageRulePredicatesParameterValue?.bodyOrSubjectContains;
        this._categories = messageRulePredicatesParameterValue?.categories;
        this._fromAddresses = messageRulePredicatesParameterValue?.fromAddresses;
        this._hasAttachments = messageRulePredicatesParameterValue?.hasAttachments;
        this._headerContains = messageRulePredicatesParameterValue?.headerContains;
        this._importance = messageRulePredicatesParameterValue?.importance;
        this._isApprovalRequest = messageRulePredicatesParameterValue?.isApprovalRequest;
        this._isAutomaticForward = messageRulePredicatesParameterValue?.isAutomaticForward;
        this._isAutomaticReply = messageRulePredicatesParameterValue?.isAutomaticReply;
        this._isEncrypted = messageRulePredicatesParameterValue?.isEncrypted;
        this._isMeetingRequest = messageRulePredicatesParameterValue?.isMeetingRequest;
        this._isMeetingResponse = messageRulePredicatesParameterValue?.isMeetingResponse;
        this._isNonDeliveryReport = messageRulePredicatesParameterValue?.isNonDeliveryReport;
        this._isPermissionControlled = messageRulePredicatesParameterValue?.isPermissionControlled;
        this._isReadReceipt = messageRulePredicatesParameterValue?.isReadReceipt;
        this._isSigned = messageRulePredicatesParameterValue?.isSigned;
        this._isVoicemail = messageRulePredicatesParameterValue?.isVoicemail;
        this._messageActionFlag = messageRulePredicatesParameterValue?.messageActionFlag;
        this._notSentToMe = messageRulePredicatesParameterValue?.notSentToMe;
        this._recipientContains = messageRulePredicatesParameterValue?.recipientContains;
        this._senderContains = messageRulePredicatesParameterValue?.senderContains;
        this._sensitivity = messageRulePredicatesParameterValue?.sensitivity;
        this._sentCcMe = messageRulePredicatesParameterValue?.sentCcMe;
        this._sentOnlyToMe = messageRulePredicatesParameterValue?.sentOnlyToMe;
        this._sentToAddresses = messageRulePredicatesParameterValue?.sentToAddresses;
        this._sentToMe = messageRulePredicatesParameterValue?.sentToMe;
        this._sentToOrCcMe = messageRulePredicatesParameterValue?.sentToOrCcMe;
        this._subjectContains = messageRulePredicatesParameterValue?.subjectContains;
        this._withinSizeRange = messageRulePredicatesParameterValue?.withinSizeRange;
    };
    /**
     * Gets the fromAddresses property value. Represents the specific sender email addresses of an incoming message in order for the condition or exception to apply.
     * @returns a RecipientInterface
     */
    public get fromAddresses() {
        return this._fromAddresses;
    };
    /**
     * Sets the fromAddresses property value. Represents the specific sender email addresses of an incoming message in order for the condition or exception to apply.
     * @param value Value to set for the fromAddresses property.
     */
    public set fromAddresses(value: Recipient[] | undefined) {
        if(value) {
            const fromAddressesArrValue: RecipientImpl[] = [];
            this.fromAddresses?.forEach(element => {
                fromAddressesArrValue.push((element instanceof RecipientImpl? element as RecipientImpl:new RecipientImpl(element)));
            });
            this._fromAddresses = fromAddressesArrValue;
        }
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
     * Gets the hasAttachments property value. Indicates whether an incoming message must have attachments in order for the condition or exception to apply.
     * @returns a boolean
     */
    public get hasAttachments() {
        return this._hasAttachments;
    };
    /**
     * Sets the hasAttachments property value. Indicates whether an incoming message must have attachments in order for the condition or exception to apply.
     * @param value Value to set for the hasAttachments property.
     */
    public set hasAttachments(value: boolean | undefined) {
        if(value) {
            this._hasAttachments = value;
        }
    };
    /**
     * Gets the headerContains property value. Represents the strings that appear in the headers of an incoming message in order for the condition or exception to apply.
     * @returns a string
     */
    public get headerContains() {
        return this._headerContains;
    };
    /**
     * Sets the headerContains property value. Represents the strings that appear in the headers of an incoming message in order for the condition or exception to apply.
     * @param value Value to set for the headerContains property.
     */
    public set headerContains(value: string[] | undefined) {
        if(value) {
            this._headerContains = value;
        }
    };
    /**
     * Gets the importance property value. The importance property
     * @returns a importance
     */
    public get importance() {
        return this._importance;
    };
    /**
     * Sets the importance property value. The importance property
     * @param value Value to set for the importance property.
     */
    public set importance(value: Importance | undefined) {
        if(value) {
            this._importance = value;
        }
    };
    /**
     * Gets the isApprovalRequest property value. Indicates whether an incoming message must be an approval request in order for the condition or exception to apply.
     * @returns a boolean
     */
    public get isApprovalRequest() {
        return this._isApprovalRequest;
    };
    /**
     * Sets the isApprovalRequest property value. Indicates whether an incoming message must be an approval request in order for the condition or exception to apply.
     * @param value Value to set for the isApprovalRequest property.
     */
    public set isApprovalRequest(value: boolean | undefined) {
        if(value) {
            this._isApprovalRequest = value;
        }
    };
    /**
     * Gets the isAutomaticForward property value. Indicates whether an incoming message must be automatically forwarded in order for the condition or exception to apply.
     * @returns a boolean
     */
    public get isAutomaticForward() {
        return this._isAutomaticForward;
    };
    /**
     * Sets the isAutomaticForward property value. Indicates whether an incoming message must be automatically forwarded in order for the condition or exception to apply.
     * @param value Value to set for the isAutomaticForward property.
     */
    public set isAutomaticForward(value: boolean | undefined) {
        if(value) {
            this._isAutomaticForward = value;
        }
    };
    /**
     * Gets the isAutomaticReply property value. Indicates whether an incoming message must be an auto reply in order for the condition or exception to apply.
     * @returns a boolean
     */
    public get isAutomaticReply() {
        return this._isAutomaticReply;
    };
    /**
     * Sets the isAutomaticReply property value. Indicates whether an incoming message must be an auto reply in order for the condition or exception to apply.
     * @param value Value to set for the isAutomaticReply property.
     */
    public set isAutomaticReply(value: boolean | undefined) {
        if(value) {
            this._isAutomaticReply = value;
        }
    };
    /**
     * Gets the isEncrypted property value. Indicates whether an incoming message must be encrypted in order for the condition or exception to apply.
     * @returns a boolean
     */
    public get isEncrypted() {
        return this._isEncrypted;
    };
    /**
     * Sets the isEncrypted property value. Indicates whether an incoming message must be encrypted in order for the condition or exception to apply.
     * @param value Value to set for the isEncrypted property.
     */
    public set isEncrypted(value: boolean | undefined) {
        if(value) {
            this._isEncrypted = value;
        }
    };
    /**
     * Gets the isMeetingRequest property value. Indicates whether an incoming message must be a meeting request in order for the condition or exception to apply.
     * @returns a boolean
     */
    public get isMeetingRequest() {
        return this._isMeetingRequest;
    };
    /**
     * Sets the isMeetingRequest property value. Indicates whether an incoming message must be a meeting request in order for the condition or exception to apply.
     * @param value Value to set for the isMeetingRequest property.
     */
    public set isMeetingRequest(value: boolean | undefined) {
        if(value) {
            this._isMeetingRequest = value;
        }
    };
    /**
     * Gets the isMeetingResponse property value. Indicates whether an incoming message must be a meeting response in order for the condition or exception to apply.
     * @returns a boolean
     */
    public get isMeetingResponse() {
        return this._isMeetingResponse;
    };
    /**
     * Sets the isMeetingResponse property value. Indicates whether an incoming message must be a meeting response in order for the condition or exception to apply.
     * @param value Value to set for the isMeetingResponse property.
     */
    public set isMeetingResponse(value: boolean | undefined) {
        if(value) {
            this._isMeetingResponse = value;
        }
    };
    /**
     * Gets the isNonDeliveryReport property value. Indicates whether an incoming message must be a non-delivery report in order for the condition or exception to apply.
     * @returns a boolean
     */
    public get isNonDeliveryReport() {
        return this._isNonDeliveryReport;
    };
    /**
     * Sets the isNonDeliveryReport property value. Indicates whether an incoming message must be a non-delivery report in order for the condition or exception to apply.
     * @param value Value to set for the isNonDeliveryReport property.
     */
    public set isNonDeliveryReport(value: boolean | undefined) {
        if(value) {
            this._isNonDeliveryReport = value;
        }
    };
    /**
     * Gets the isPermissionControlled property value. Indicates whether an incoming message must be permission controlled (RMS-protected) in order for the condition or exception to apply.
     * @returns a boolean
     */
    public get isPermissionControlled() {
        return this._isPermissionControlled;
    };
    /**
     * Sets the isPermissionControlled property value. Indicates whether an incoming message must be permission controlled (RMS-protected) in order for the condition or exception to apply.
     * @param value Value to set for the isPermissionControlled property.
     */
    public set isPermissionControlled(value: boolean | undefined) {
        if(value) {
            this._isPermissionControlled = value;
        }
    };
    /**
     * Gets the isReadReceipt property value. Indicates whether an incoming message must be a read receipt in order for the condition or exception to apply.
     * @returns a boolean
     */
    public get isReadReceipt() {
        return this._isReadReceipt;
    };
    /**
     * Sets the isReadReceipt property value. Indicates whether an incoming message must be a read receipt in order for the condition or exception to apply.
     * @param value Value to set for the isReadReceipt property.
     */
    public set isReadReceipt(value: boolean | undefined) {
        if(value) {
            this._isReadReceipt = value;
        }
    };
    /**
     * Gets the isSigned property value. Indicates whether an incoming message must be S/MIME-signed in order for the condition or exception to apply.
     * @returns a boolean
     */
    public get isSigned() {
        return this._isSigned;
    };
    /**
     * Sets the isSigned property value. Indicates whether an incoming message must be S/MIME-signed in order for the condition or exception to apply.
     * @param value Value to set for the isSigned property.
     */
    public set isSigned(value: boolean | undefined) {
        if(value) {
            this._isSigned = value;
        }
    };
    /**
     * Gets the isVoicemail property value. Indicates whether an incoming message must be a voice mail in order for the condition or exception to apply.
     * @returns a boolean
     */
    public get isVoicemail() {
        return this._isVoicemail;
    };
    /**
     * Sets the isVoicemail property value. Indicates whether an incoming message must be a voice mail in order for the condition or exception to apply.
     * @param value Value to set for the isVoicemail property.
     */
    public set isVoicemail(value: boolean | undefined) {
        if(value) {
            this._isVoicemail = value;
        }
    };
    /**
     * Gets the messageActionFlag property value. The messageActionFlag property
     * @returns a messageActionFlag
     */
    public get messageActionFlag() {
        return this._messageActionFlag;
    };
    /**
     * Sets the messageActionFlag property value. The messageActionFlag property
     * @param value Value to set for the messageActionFlag property.
     */
    public set messageActionFlag(value: MessageActionFlag | undefined) {
        if(value) {
            this._messageActionFlag = value;
        }
    };
    /**
     * Gets the notSentToMe property value. Indicates whether the owner of the mailbox must not be a recipient of an incoming message in order for the condition or exception to apply.
     * @returns a boolean
     */
    public get notSentToMe() {
        return this._notSentToMe;
    };
    /**
     * Sets the notSentToMe property value. Indicates whether the owner of the mailbox must not be a recipient of an incoming message in order for the condition or exception to apply.
     * @param value Value to set for the notSentToMe property.
     */
    public set notSentToMe(value: boolean | undefined) {
        if(value) {
            this._notSentToMe = value;
        }
    };
    /**
     * Gets the recipientContains property value. Represents the strings that appear in either the toRecipients or ccRecipients properties of an incoming message in order for the condition or exception to apply.
     * @returns a string
     */
    public get recipientContains() {
        return this._recipientContains;
    };
    /**
     * Sets the recipientContains property value. Represents the strings that appear in either the toRecipients or ccRecipients properties of an incoming message in order for the condition or exception to apply.
     * @param value Value to set for the recipientContains property.
     */
    public set recipientContains(value: string[] | undefined) {
        if(value) {
            this._recipientContains = value;
        }
    };
    /**
     * Gets the senderContains property value. Represents the strings that appear in the from property of an incoming message in order for the condition or exception to apply.
     * @returns a string
     */
    public get senderContains() {
        return this._senderContains;
    };
    /**
     * Sets the senderContains property value. Represents the strings that appear in the from property of an incoming message in order for the condition or exception to apply.
     * @param value Value to set for the senderContains property.
     */
    public set senderContains(value: string[] | undefined) {
        if(value) {
            this._senderContains = value;
        }
    };
    /**
     * Gets the sensitivity property value. The sensitivity property
     * @returns a sensitivity
     */
    public get sensitivity() {
        return this._sensitivity;
    };
    /**
     * Sets the sensitivity property value. The sensitivity property
     * @param value Value to set for the sensitivity property.
     */
    public set sensitivity(value: Sensitivity | undefined) {
        if(value) {
            this._sensitivity = value;
        }
    };
    /**
     * Gets the sentCcMe property value. Indicates whether the owner of the mailbox must be in the ccRecipients property of an incoming message in order for the condition or exception to apply.
     * @returns a boolean
     */
    public get sentCcMe() {
        return this._sentCcMe;
    };
    /**
     * Sets the sentCcMe property value. Indicates whether the owner of the mailbox must be in the ccRecipients property of an incoming message in order for the condition or exception to apply.
     * @param value Value to set for the sentCcMe property.
     */
    public set sentCcMe(value: boolean | undefined) {
        if(value) {
            this._sentCcMe = value;
        }
    };
    /**
     * Gets the sentOnlyToMe property value. Indicates whether the owner of the mailbox must be the only recipient in an incoming message in order for the condition or exception to apply.
     * @returns a boolean
     */
    public get sentOnlyToMe() {
        return this._sentOnlyToMe;
    };
    /**
     * Sets the sentOnlyToMe property value. Indicates whether the owner of the mailbox must be the only recipient in an incoming message in order for the condition or exception to apply.
     * @param value Value to set for the sentOnlyToMe property.
     */
    public set sentOnlyToMe(value: boolean | undefined) {
        if(value) {
            this._sentOnlyToMe = value;
        }
    };
    /**
     * Gets the sentToAddresses property value. Represents the email addresses that an incoming message must have been sent to in order for the condition or exception to apply.
     * @returns a RecipientInterface
     */
    public get sentToAddresses() {
        return this._sentToAddresses;
    };
    /**
     * Sets the sentToAddresses property value. Represents the email addresses that an incoming message must have been sent to in order for the condition or exception to apply.
     * @param value Value to set for the sentToAddresses property.
     */
    public set sentToAddresses(value: Recipient[] | undefined) {
        if(value) {
            const sentToAddressesArrValue: RecipientImpl[] = [];
            this.sentToAddresses?.forEach(element => {
                sentToAddressesArrValue.push((element instanceof RecipientImpl? element as RecipientImpl:new RecipientImpl(element)));
            });
            this._sentToAddresses = sentToAddressesArrValue;
        }
    };
    /**
     * Gets the sentToMe property value. Indicates whether the owner of the mailbox must be in the toRecipients property of an incoming message in order for the condition or exception to apply.
     * @returns a boolean
     */
    public get sentToMe() {
        return this._sentToMe;
    };
    /**
     * Sets the sentToMe property value. Indicates whether the owner of the mailbox must be in the toRecipients property of an incoming message in order for the condition or exception to apply.
     * @param value Value to set for the sentToMe property.
     */
    public set sentToMe(value: boolean | undefined) {
        if(value) {
            this._sentToMe = value;
        }
    };
    /**
     * Gets the sentToOrCcMe property value. Indicates whether the owner of the mailbox must be in either a toRecipients or ccRecipients property of an incoming message in order for the condition or exception to apply.
     * @returns a boolean
     */
    public get sentToOrCcMe() {
        return this._sentToOrCcMe;
    };
    /**
     * Sets the sentToOrCcMe property value. Indicates whether the owner of the mailbox must be in either a toRecipients or ccRecipients property of an incoming message in order for the condition or exception to apply.
     * @param value Value to set for the sentToOrCcMe property.
     */
    public set sentToOrCcMe(value: boolean | undefined) {
        if(value) {
            this._sentToOrCcMe = value;
        }
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
        if(this.fromAddresses && this.fromAddresses.length != 0){        const fromAddressesArrValue: RecipientImpl[] = [];
        this.fromAddresses?.forEach(element => {
            fromAddressesArrValue.push((element instanceof RecipientImpl? element as RecipientImpl:new RecipientImpl(element)));
        });
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
        if(this.sentToAddresses && this.sentToAddresses.length != 0){        const sentToAddressesArrValue: RecipientImpl[] = [];
        this.sentToAddresses?.forEach(element => {
            sentToAddressesArrValue.push((element instanceof RecipientImpl? element as RecipientImpl:new RecipientImpl(element)));
        });
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
            writer.writeObjectValue<SizeRangeImpl>("withinSizeRange", (this.withinSizeRange instanceof SizeRangeImpl? this.withinSizeRange as SizeRangeImpl: new SizeRangeImpl(this.withinSizeRange)));
        }
        writer.writeAdditionalData(this.additionalData!);
    };
    /**
     * Gets the subjectContains property value. Represents the strings that appear in the subject of an incoming message in order for the condition or exception to apply.
     * @returns a string
     */
    public get subjectContains() {
        return this._subjectContains;
    };
    /**
     * Sets the subjectContains property value. Represents the strings that appear in the subject of an incoming message in order for the condition or exception to apply.
     * @param value Value to set for the subjectContains property.
     */
    public set subjectContains(value: string[] | undefined) {
        if(value) {
            this._subjectContains = value;
        }
    };
    /**
     * Gets the withinSizeRange property value. The withinSizeRange property
     * @returns a SizeRangeInterface
     */
    public get withinSizeRange() {
        return this._withinSizeRange;
    };
    /**
     * Sets the withinSizeRange property value. The withinSizeRange property
     * @param value Value to set for the withinSizeRange property.
     */
    public set withinSizeRange(value: SizeRange | undefined) {
        if(value) {
            this._withinSizeRange = value instanceof SizeRangeImpl? value as SizeRangeImpl: new SizeRangeImpl(value);
        }
    };
}
