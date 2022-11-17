import {Attachment} from './attachment';
<<<<<<< HEAD
import {Extension} from './extension';
import {FollowupFlag} from './followupFlag';
import {Importance} from './importance';
import {InferenceClassificationType} from './inferenceClassificationType';
import {InternetMessageHeader} from './internetMessageHeader';
import {ItemBody} from './itemBody';
import {MultiValueLegacyExtendedProperty} from './multiValueLegacyExtendedProperty';
import {OutlookItem} from './outlookItem';
import {Recipient} from './recipient';
import {SingleValueLegacyExtendedProperty} from './singleValueLegacyExtendedProperty';
import {Parsable} from '@microsoft/kiota-abstractions';

export interface Message extends OutlookItem, Partial<Parsable> {
    /** The fileAttachment and itemAttachment attachments for the message. */
    attachments?: Attachment[] | undefined;
    /** The Bcc: recipients for the message. */
    bccRecipients?: Recipient[] | undefined;
    /** The body property */
    body?: ItemBody | undefined;
    /** The first 255 characters of the message body. It is in text format. If the message contains instances of mention, this property would contain a concatenation of these mentions as well. */
    bodyPreview?: string | undefined;
    /** The Cc: recipients for the message. */
    ccRecipients?: Recipient[] | undefined;
    /** The ID of the conversation the email belongs to. */
    conversationId?: string | undefined;
    /** Indicates the position of the message within the conversation. */
    conversationIndex?: string | undefined;
    /** The collection of open extensions defined for the message. Nullable. */
    extensions?: Extension[] | undefined;
    /** The flag property */
    flag?: FollowupFlag | undefined;
    /** The from property */
    from?: Recipient | undefined;
    /** Indicates whether the message has attachments. This property doesn't include inline attachments, so if a message contains only inline attachments, this property is false. To verify the existence of inline attachments, parse the body property to look for a src attribute, such as <IMG src='cid:image001.jpg@01D26CD8.6C05F070'>. */
    hasAttachments?: boolean | undefined;
    /** The importance property */
    importance?: Importance | undefined;
    /** The inferenceClassification property */
    inferenceClassification?: InferenceClassificationType | undefined;
    /** The internetMessageHeaders property */
    internetMessageHeaders?: InternetMessageHeader[] | undefined;
    /** The internetMessageId property */
    internetMessageId?: string | undefined;
    /** The isDeliveryReceiptRequested property */
    isDeliveryReceiptRequested?: boolean | undefined;
    /** The isDraft property */
    isDraft?: boolean | undefined;
    /** The isRead property */
    isRead?: boolean | undefined;
    /** The isReadReceiptRequested property */
    isReadReceiptRequested?: boolean | undefined;
    /** The collection of multi-value extended properties defined for the message. Nullable. */
    multiValueExtendedProperties?: MultiValueLegacyExtendedProperty[] | undefined;
    /** The parentFolderId property */
    parentFolderId?: string | undefined;
    /** The receivedDateTime property */
    receivedDateTime?: Date | undefined;
    /** The replyTo property */
    replyTo?: Recipient[] | undefined;
    /** The sender property */
    sender?: Recipient | undefined;
    /** The sentDateTime property */
    sentDateTime?: Date | undefined;
    /** The collection of single-value extended properties defined for the message. Nullable. */
    singleValueExtendedProperties?: SingleValueLegacyExtendedProperty[] | undefined;
    /** The subject property */
    subject?: string | undefined;
    /** The toRecipients property */
    toRecipients?: Recipient[] | undefined;
    /** The uniqueBody property */
    uniqueBody?: ItemBody | undefined;
    /** The webLink property */
    webLink?: string | undefined;
=======
import {createAttachmentFromDiscriminatorValue} from './createAttachmentFromDiscriminatorValue';
import {createExtensionFromDiscriminatorValue} from './createExtensionFromDiscriminatorValue';
import {createFollowupFlagFromDiscriminatorValue} from './createFollowupFlagFromDiscriminatorValue';
import {createInternetMessageHeaderFromDiscriminatorValue} from './createInternetMessageHeaderFromDiscriminatorValue';
import {createItemBodyFromDiscriminatorValue} from './createItemBodyFromDiscriminatorValue';
import {createMultiValueLegacyExtendedPropertyFromDiscriminatorValue} from './createMultiValueLegacyExtendedPropertyFromDiscriminatorValue';
import {createRecipientFromDiscriminatorValue} from './createRecipientFromDiscriminatorValue';
import {createSingleValueLegacyExtendedPropertyFromDiscriminatorValue} from './createSingleValueLegacyExtendedPropertyFromDiscriminatorValue';
import {Extension} from './extension';
import {FollowupFlag} from './followupFlag';
import {Importance} from './importance';
import {AttachmentImpl, ExtensionImpl, FollowupFlagImpl, InternetMessageHeaderImpl, ItemBodyImpl, MultiValueLegacyExtendedPropertyImpl, OutlookItemImpl, RecipientImpl, SingleValueLegacyExtendedPropertyImpl} from './index';
import {InferenceClassificationType} from './inferenceClassificationType';
import {InternetMessageHeader} from './internetMessageHeader';
import {ItemBody} from './itemBody';
import {Message} from './message';
import {MultiValueLegacyExtendedProperty} from './multiValueLegacyExtendedProperty';
import {Recipient} from './recipient';
import {SingleValueLegacyExtendedProperty} from './singleValueLegacyExtendedProperty';
import {Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export class MessageImpl extends OutlookItemImpl implements Message {
    /** The fileAttachment and itemAttachment attachments for the message. */
    private _attachments?: Attachment[] | undefined;
    /** The Bcc: recipients for the message. */
    private _bccRecipients?: Recipient[] | undefined;
    /** The body property */
    private _body?: ItemBody | undefined;
    /** The first 255 characters of the message body. It is in text format. If the message contains instances of mention, this property would contain a concatenation of these mentions as well. */
    private _bodyPreview?: string | undefined;
    /** The Cc: recipients for the message. */
    private _ccRecipients?: Recipient[] | undefined;
    /** The ID of the conversation the email belongs to. */
    private _conversationId?: string | undefined;
    /** Indicates the position of the message within the conversation. */
    private _conversationIndex?: string | undefined;
    /** The collection of open extensions defined for the message. Nullable. */
    private _extensions?: Extension[] | undefined;
    /** The flag property */
    private _flag?: FollowupFlag | undefined;
    /** The from property */
    private _from?: Recipient | undefined;
    /** Indicates whether the message has attachments. This property doesn't include inline attachments, so if a message contains only inline attachments, this property is false. To verify the existence of inline attachments, parse the body property to look for a src attribute, such as <IMG src='cid:image001.jpg@01D26CD8.6C05F070'>. */
    private _hasAttachments?: boolean | undefined;
    /** The importance property */
    private _importance?: Importance | undefined;
    /** The inferenceClassification property */
    private _inferenceClassification?: InferenceClassificationType | undefined;
    /** The internetMessageHeaders property */
    private _internetMessageHeaders?: InternetMessageHeader[] | undefined;
    /** The internetMessageId property */
    private _internetMessageId?: string | undefined;
    /** The isDeliveryReceiptRequested property */
    private _isDeliveryReceiptRequested?: boolean | undefined;
    /** The isDraft property */
    private _isDraft?: boolean | undefined;
    /** The isRead property */
    private _isRead?: boolean | undefined;
    /** The isReadReceiptRequested property */
    private _isReadReceiptRequested?: boolean | undefined;
    /** The collection of multi-value extended properties defined for the message. Nullable. */
    private _multiValueExtendedProperties?: MultiValueLegacyExtendedProperty[] | undefined;
    /** The parentFolderId property */
    private _parentFolderId?: string | undefined;
    /** The receivedDateTime property */
    private _receivedDateTime?: Date | undefined;
    /** The replyTo property */
    private _replyTo?: Recipient[] | undefined;
    /** The sender property */
    private _sender?: Recipient | undefined;
    /** The sentDateTime property */
    private _sentDateTime?: Date | undefined;
    /** The collection of single-value extended properties defined for the message. Nullable. */
    private _singleValueExtendedProperties?: SingleValueLegacyExtendedProperty[] | undefined;
    /** The subject property */
    private _subject?: string | undefined;
    /** The toRecipients property */
    private _toRecipients?: Recipient[] | undefined;
    /** The uniqueBody property */
    private _uniqueBody?: ItemBody | undefined;
    /** The webLink property */
    private _webLink?: string | undefined;
    /**
     * Gets the attachments property value. The fileAttachment and itemAttachment attachments for the message.
     * @returns a AttachmentInterface
     */
    public get attachments() {
        return this._attachments;
    };
    /**
     * Sets the attachments property value. The fileAttachment and itemAttachment attachments for the message.
     * @param value Value to set for the attachments property.
     */
    public set attachments(value: Attachment[] | undefined) {
        if(value) {
            const attachmentsArrValue: AttachmentImpl[] = [];
            this.attachments?.forEach(element => {
                attachmentsArrValue.push((element instanceof AttachmentImpl? element as AttachmentImpl:new AttachmentImpl(element)));
            });
            this._attachments = attachmentsArrValue;
        }
    };
    /**
     * Gets the bccRecipients property value. The Bcc: recipients for the message.
     * @returns a RecipientInterface
     */
    public get bccRecipients() {
        return this._bccRecipients;
    };
    /**
     * Sets the bccRecipients property value. The Bcc: recipients for the message.
     * @param value Value to set for the bccRecipients property.
     */
    public set bccRecipients(value: Recipient[] | undefined) {
        if(value) {
            const bccRecipientsArrValue: RecipientImpl[] = [];
            this.bccRecipients?.forEach(element => {
                bccRecipientsArrValue.push((element instanceof RecipientImpl? element as RecipientImpl:new RecipientImpl(element)));
            });
            this._bccRecipients = bccRecipientsArrValue;
        }
    };
    /**
     * Gets the body property value. The body property
     * @returns a ItemBodyInterface
     */
    public get body() {
        return this._body;
    };
    /**
     * Sets the body property value. The body property
     * @param value Value to set for the body property.
     */
    public set body(value: ItemBody | undefined) {
        if(value) {
            this._body = value instanceof ItemBodyImpl? value as ItemBodyImpl: new ItemBodyImpl(value);
        }
    };
    /**
     * Gets the bodyPreview property value. The first 255 characters of the message body. It is in text format. If the message contains instances of mention, this property would contain a concatenation of these mentions as well.
     * @returns a string
     */
    public get bodyPreview() {
        return this._bodyPreview;
    };
    /**
     * Sets the bodyPreview property value. The first 255 characters of the message body. It is in text format. If the message contains instances of mention, this property would contain a concatenation of these mentions as well.
     * @param value Value to set for the bodyPreview property.
     */
    public set bodyPreview(value: string | undefined) {
        if(value) {
            this._bodyPreview = value;
        }
    };
    /**
     * Gets the ccRecipients property value. The Cc: recipients for the message.
     * @returns a RecipientInterface
     */
    public get ccRecipients() {
        return this._ccRecipients;
    };
    /**
     * Sets the ccRecipients property value. The Cc: recipients for the message.
     * @param value Value to set for the ccRecipients property.
     */
    public set ccRecipients(value: Recipient[] | undefined) {
        if(value) {
            const ccRecipientsArrValue: RecipientImpl[] = [];
            this.ccRecipients?.forEach(element => {
                ccRecipientsArrValue.push((element instanceof RecipientImpl? element as RecipientImpl:new RecipientImpl(element)));
            });
            this._ccRecipients = ccRecipientsArrValue;
        }
    };
    /**
     * Instantiates a new message and sets the default values.
     * @param messageParameterValue 
     */
    public constructor(messageParameterValue?: Message | undefined) {
        super(messageParameterValue);
        this._attachments = messageParameterValue?.attachments;
        this._bccRecipients = messageParameterValue?.bccRecipients;
        this._body = messageParameterValue?.body;
        this._bodyPreview = messageParameterValue?.bodyPreview;
        this._ccRecipients = messageParameterValue?.ccRecipients;
        this._conversationId = messageParameterValue?.conversationId;
        this._conversationIndex = messageParameterValue?.conversationIndex;
        this._extensions = messageParameterValue?.extensions;
        this._flag = messageParameterValue?.flag;
        this._from = messageParameterValue?.from;
        this._hasAttachments = messageParameterValue?.hasAttachments;
        this._importance = messageParameterValue?.importance;
        this._inferenceClassification = messageParameterValue?.inferenceClassification;
        this._internetMessageHeaders = messageParameterValue?.internetMessageHeaders;
        this._internetMessageId = messageParameterValue?.internetMessageId;
        this._isDeliveryReceiptRequested = messageParameterValue?.isDeliveryReceiptRequested;
        this._isDraft = messageParameterValue?.isDraft;
        this._isRead = messageParameterValue?.isRead;
        this._isReadReceiptRequested = messageParameterValue?.isReadReceiptRequested;
        this._multiValueExtendedProperties = messageParameterValue?.multiValueExtendedProperties;
        this._parentFolderId = messageParameterValue?.parentFolderId;
        this._receivedDateTime = messageParameterValue?.receivedDateTime;
        this._replyTo = messageParameterValue?.replyTo;
        this._sender = messageParameterValue?.sender;
        this._sentDateTime = messageParameterValue?.sentDateTime;
        this._singleValueExtendedProperties = messageParameterValue?.singleValueExtendedProperties;
        this._subject = messageParameterValue?.subject;
        this._toRecipients = messageParameterValue?.toRecipients;
        this._uniqueBody = messageParameterValue?.uniqueBody;
        this._webLink = messageParameterValue?.webLink;
    };
    /**
     * Gets the conversationId property value. The ID of the conversation the email belongs to.
     * @returns a string
     */
    public get conversationId() {
        return this._conversationId;
    };
    /**
     * Sets the conversationId property value. The ID of the conversation the email belongs to.
     * @param value Value to set for the conversationId property.
     */
    public set conversationId(value: string | undefined) {
        if(value) {
            this._conversationId = value;
        }
    };
    /**
     * Gets the conversationIndex property value. Indicates the position of the message within the conversation.
     * @returns a binary
     */
    public get conversationIndex() {
        return this._conversationIndex;
    };
    /**
     * Sets the conversationIndex property value. Indicates the position of the message within the conversation.
     * @param value Value to set for the conversationIndex property.
     */
    public set conversationIndex(value: string | undefined) {
        if(value) {
            this._conversationIndex = value;
        }
    };
    /**
     * Gets the extensions property value. The collection of open extensions defined for the message. Nullable.
     * @returns a ExtensionInterface
     */
    public get extensions() {
        return this._extensions;
    };
    /**
     * Sets the extensions property value. The collection of open extensions defined for the message. Nullable.
     * @param value Value to set for the extensions property.
     */
    public set extensions(value: Extension[] | undefined) {
        if(value) {
            const extensionsArrValue: ExtensionImpl[] = [];
            this.extensions?.forEach(element => {
                extensionsArrValue.push((element instanceof ExtensionImpl? element as ExtensionImpl:new ExtensionImpl(element)));
            });
            this._extensions = extensionsArrValue;
        }
    };
    /**
     * Gets the flag property value. The flag property
     * @returns a FollowupFlagInterface
     */
    public get flag() {
        return this._flag;
    };
    /**
     * Sets the flag property value. The flag property
     * @param value Value to set for the flag property.
     */
    public set flag(value: FollowupFlag | undefined) {
        if(value) {
            this._flag = value instanceof FollowupFlagImpl? value as FollowupFlagImpl: new FollowupFlagImpl(value);
        }
    };
    /**
     * Gets the from property value. The from property
     * @returns a RecipientInterface
     */
    public get from() {
        return this._from;
    };
    /**
     * Sets the from property value. The from property
     * @param value Value to set for the from property.
     */
    public set from(value: Recipient | undefined) {
        if(value) {
            this._from = value instanceof RecipientImpl? value as RecipientImpl: new RecipientImpl(value);
        }
    };
    /**
     * The deserialization information for the current model
     * @returns a Record<string, (node: ParseNode) => void>
     */
    public getFieldDeserializers() : Record<string, (node: ParseNode) => void> {
        return {...super.getFieldDeserializers(),
            "attachments": n => { this.attachments = n.getCollectionOfObjectValues<AttachmentImpl>(createAttachmentFromDiscriminatorValue); },
            "bccRecipients": n => { this.bccRecipients = n.getCollectionOfObjectValues<RecipientImpl>(createRecipientFromDiscriminatorValue); },
            "body": n => { this.body = n.getObjectValue<ItemBodyImpl>(createItemBodyFromDiscriminatorValue); },
            "bodyPreview": n => { this.bodyPreview = n.getStringValue(); },
            "ccRecipients": n => { this.ccRecipients = n.getCollectionOfObjectValues<RecipientImpl>(createRecipientFromDiscriminatorValue); },
            "conversationId": n => { this.conversationId = n.getStringValue(); },
            "conversationIndex": n => { this.conversationIndex = n.getStringValue(); },
            "extensions": n => { this.extensions = n.getCollectionOfObjectValues<ExtensionImpl>(createExtensionFromDiscriminatorValue); },
            "flag": n => { this.flag = n.getObjectValue<FollowupFlagImpl>(createFollowupFlagFromDiscriminatorValue); },
            "from": n => { this.from = n.getObjectValue<RecipientImpl>(createRecipientFromDiscriminatorValue); },
            "hasAttachments": n => { this.hasAttachments = n.getBooleanValue(); },
            "importance": n => { this.importance = n.getEnumValue<Importance>(Importance); },
            "inferenceClassification": n => { this.inferenceClassification = n.getEnumValue<InferenceClassificationType>(InferenceClassificationType); },
            "internetMessageHeaders": n => { this.internetMessageHeaders = n.getCollectionOfObjectValues<InternetMessageHeaderImpl>(createInternetMessageHeaderFromDiscriminatorValue); },
            "internetMessageId": n => { this.internetMessageId = n.getStringValue(); },
            "isDeliveryReceiptRequested": n => { this.isDeliveryReceiptRequested = n.getBooleanValue(); },
            "isDraft": n => { this.isDraft = n.getBooleanValue(); },
            "isRead": n => { this.isRead = n.getBooleanValue(); },
            "isReadReceiptRequested": n => { this.isReadReceiptRequested = n.getBooleanValue(); },
            "multiValueExtendedProperties": n => { this.multiValueExtendedProperties = n.getCollectionOfObjectValues<MultiValueLegacyExtendedPropertyImpl>(createMultiValueLegacyExtendedPropertyFromDiscriminatorValue); },
            "parentFolderId": n => { this.parentFolderId = n.getStringValue(); },
            "receivedDateTime": n => { this.receivedDateTime = n.getDateValue(); },
            "replyTo": n => { this.replyTo = n.getCollectionOfObjectValues<RecipientImpl>(createRecipientFromDiscriminatorValue); },
            "sender": n => { this.sender = n.getObjectValue<RecipientImpl>(createRecipientFromDiscriminatorValue); },
            "sentDateTime": n => { this.sentDateTime = n.getDateValue(); },
            "singleValueExtendedProperties": n => { this.singleValueExtendedProperties = n.getCollectionOfObjectValues<SingleValueLegacyExtendedPropertyImpl>(createSingleValueLegacyExtendedPropertyFromDiscriminatorValue); },
            "subject": n => { this.subject = n.getStringValue(); },
            "toRecipients": n => { this.toRecipients = n.getCollectionOfObjectValues<RecipientImpl>(createRecipientFromDiscriminatorValue); },
            "uniqueBody": n => { this.uniqueBody = n.getObjectValue<ItemBodyImpl>(createItemBodyFromDiscriminatorValue); },
            "webLink": n => { this.webLink = n.getStringValue(); },
        };
    };
    /**
     * Gets the hasAttachments property value. Indicates whether the message has attachments. This property doesn't include inline attachments, so if a message contains only inline attachments, this property is false. To verify the existence of inline attachments, parse the body property to look for a src attribute, such as <IMG src='cid:image001.jpg@01D26CD8.6C05F070'>.
     * @returns a boolean
     */
    public get hasAttachments() {
        return this._hasAttachments;
    };
    /**
     * Sets the hasAttachments property value. Indicates whether the message has attachments. This property doesn't include inline attachments, so if a message contains only inline attachments, this property is false. To verify the existence of inline attachments, parse the body property to look for a src attribute, such as <IMG src='cid:image001.jpg@01D26CD8.6C05F070'>.
     * @param value Value to set for the hasAttachments property.
     */
    public set hasAttachments(value: boolean | undefined) {
        if(value) {
            this._hasAttachments = value;
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
     * Gets the inferenceClassification property value. The inferenceClassification property
     * @returns a inferenceClassificationType
     */
    public get inferenceClassification() {
        return this._inferenceClassification;
    };
    /**
     * Sets the inferenceClassification property value. The inferenceClassification property
     * @param value Value to set for the inferenceClassification property.
     */
    public set inferenceClassification(value: InferenceClassificationType | undefined) {
        if(value) {
            this._inferenceClassification = value;
        }
    };
    /**
     * Gets the internetMessageHeaders property value. The internetMessageHeaders property
     * @returns a InternetMessageHeaderInterface
     */
    public get internetMessageHeaders() {
        return this._internetMessageHeaders;
    };
    /**
     * Sets the internetMessageHeaders property value. The internetMessageHeaders property
     * @param value Value to set for the internetMessageHeaders property.
     */
    public set internetMessageHeaders(value: InternetMessageHeader[] | undefined) {
        if(value) {
            const internetMessageHeadersArrValue: InternetMessageHeaderImpl[] = [];
            this.internetMessageHeaders?.forEach(element => {
                internetMessageHeadersArrValue.push((element instanceof InternetMessageHeaderImpl? element as InternetMessageHeaderImpl:new InternetMessageHeaderImpl(element)));
            });
            this._internetMessageHeaders = internetMessageHeadersArrValue;
        }
    };
    /**
     * Gets the internetMessageId property value. The internetMessageId property
     * @returns a string
     */
    public get internetMessageId() {
        return this._internetMessageId;
    };
    /**
     * Sets the internetMessageId property value. The internetMessageId property
     * @param value Value to set for the internetMessageId property.
     */
    public set internetMessageId(value: string | undefined) {
        if(value) {
            this._internetMessageId = value;
        }
    };
    /**
     * Gets the isDeliveryReceiptRequested property value. The isDeliveryReceiptRequested property
     * @returns a boolean
     */
    public get isDeliveryReceiptRequested() {
        return this._isDeliveryReceiptRequested;
    };
    /**
     * Sets the isDeliveryReceiptRequested property value. The isDeliveryReceiptRequested property
     * @param value Value to set for the isDeliveryReceiptRequested property.
     */
    public set isDeliveryReceiptRequested(value: boolean | undefined) {
        if(value) {
            this._isDeliveryReceiptRequested = value;
        }
    };
    /**
     * Gets the isDraft property value. The isDraft property
     * @returns a boolean
     */
    public get isDraft() {
        return this._isDraft;
    };
    /**
     * Sets the isDraft property value. The isDraft property
     * @param value Value to set for the isDraft property.
     */
    public set isDraft(value: boolean | undefined) {
        if(value) {
            this._isDraft = value;
        }
    };
    /**
     * Gets the isRead property value. The isRead property
     * @returns a boolean
     */
    public get isRead() {
        return this._isRead;
    };
    /**
     * Sets the isRead property value. The isRead property
     * @param value Value to set for the isRead property.
     */
    public set isRead(value: boolean | undefined) {
        if(value) {
            this._isRead = value;
        }
    };
    /**
     * Gets the isReadReceiptRequested property value. The isReadReceiptRequested property
     * @returns a boolean
     */
    public get isReadReceiptRequested() {
        return this._isReadReceiptRequested;
    };
    /**
     * Sets the isReadReceiptRequested property value. The isReadReceiptRequested property
     * @param value Value to set for the isReadReceiptRequested property.
     */
    public set isReadReceiptRequested(value: boolean | undefined) {
        if(value) {
            this._isReadReceiptRequested = value;
        }
    };
    /**
     * Gets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the message. Nullable.
     * @returns a MultiValueLegacyExtendedPropertyInterface
     */
    public get multiValueExtendedProperties() {
        return this._multiValueExtendedProperties;
    };
    /**
     * Sets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the message. Nullable.
     * @param value Value to set for the multiValueExtendedProperties property.
     */
    public set multiValueExtendedProperties(value: MultiValueLegacyExtendedProperty[] | undefined) {
        if(value) {
            const multiValueExtendedPropertiesArrValue: MultiValueLegacyExtendedPropertyImpl[] = [];
            this.multiValueExtendedProperties?.forEach(element => {
                multiValueExtendedPropertiesArrValue.push((element instanceof MultiValueLegacyExtendedPropertyImpl? element as MultiValueLegacyExtendedPropertyImpl:new MultiValueLegacyExtendedPropertyImpl(element)));
            });
            this._multiValueExtendedProperties = multiValueExtendedPropertiesArrValue;
        }
    };
    /**
     * Gets the parentFolderId property value. The parentFolderId property
     * @returns a string
     */
    public get parentFolderId() {
        return this._parentFolderId;
    };
    /**
     * Sets the parentFolderId property value. The parentFolderId property
     * @param value Value to set for the parentFolderId property.
     */
    public set parentFolderId(value: string | undefined) {
        if(value) {
            this._parentFolderId = value;
        }
    };
    /**
     * Gets the receivedDateTime property value. The receivedDateTime property
     * @returns a Date
     */
    public get receivedDateTime() {
        return this._receivedDateTime;
    };
    /**
     * Sets the receivedDateTime property value. The receivedDateTime property
     * @param value Value to set for the receivedDateTime property.
     */
    public set receivedDateTime(value: Date | undefined) {
        if(value) {
            this._receivedDateTime = value;
        }
    };
    /**
     * Gets the replyTo property value. The replyTo property
     * @returns a RecipientInterface
     */
    public get replyTo() {
        return this._replyTo;
    };
    /**
     * Sets the replyTo property value. The replyTo property
     * @param value Value to set for the replyTo property.
     */
    public set replyTo(value: Recipient[] | undefined) {
        if(value) {
            const replyToArrValue: RecipientImpl[] = [];
            this.replyTo?.forEach(element => {
                replyToArrValue.push((element instanceof RecipientImpl? element as RecipientImpl:new RecipientImpl(element)));
            });
            this._replyTo = replyToArrValue;
        }
    };
    /**
     * Gets the sender property value. The sender property
     * @returns a RecipientInterface
     */
    public get sender() {
        return this._sender;
    };
    /**
     * Sets the sender property value. The sender property
     * @param value Value to set for the sender property.
     */
    public set sender(value: Recipient | undefined) {
        if(value) {
            this._sender = value instanceof RecipientImpl? value as RecipientImpl: new RecipientImpl(value);
        }
    };
    /**
     * Gets the sentDateTime property value. The sentDateTime property
     * @returns a Date
     */
    public get sentDateTime() {
        return this._sentDateTime;
    };
    /**
     * Sets the sentDateTime property value. The sentDateTime property
     * @param value Value to set for the sentDateTime property.
     */
    public set sentDateTime(value: Date | undefined) {
        if(value) {
            this._sentDateTime = value;
        }
    };
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     */
    public serialize(writer: SerializationWriter) : void {
        if(!writer) throw new Error("writer cannot be undefined");
        super.serialize(writer);
        if(this.attachments && this.attachments.length != 0){        const attachmentsArrValue: AttachmentImpl[] = [];
        this.attachments?.forEach(element => {
            attachmentsArrValue.push((element instanceof AttachmentImpl? element as AttachmentImpl:new AttachmentImpl(element)));
        });
            writer.writeCollectionOfObjectValues<AttachmentImpl>("attachments", attachmentsArrValue);
        }
        if(this.bccRecipients && this.bccRecipients.length != 0){        const bccRecipientsArrValue: RecipientImpl[] = [];
        this.bccRecipients?.forEach(element => {
            bccRecipientsArrValue.push((element instanceof RecipientImpl? element as RecipientImpl:new RecipientImpl(element)));
        });
            writer.writeCollectionOfObjectValues<RecipientImpl>("bccRecipients", bccRecipientsArrValue);
        }
        if(this.body){
            writer.writeObjectValue<ItemBodyImpl>("body", (this.body instanceof ItemBodyImpl? this.body as ItemBodyImpl: new ItemBodyImpl(this.body)));
        }
        if(this.bodyPreview){
            writer.writeStringValue("bodyPreview", this.bodyPreview);
        }
        if(this.ccRecipients && this.ccRecipients.length != 0){        const ccRecipientsArrValue: RecipientImpl[] = [];
        this.ccRecipients?.forEach(element => {
            ccRecipientsArrValue.push((element instanceof RecipientImpl? element as RecipientImpl:new RecipientImpl(element)));
        });
            writer.writeCollectionOfObjectValues<RecipientImpl>("ccRecipients", ccRecipientsArrValue);
        }
        if(this.conversationId){
            writer.writeStringValue("conversationId", this.conversationId);
        }
        if(this.conversationIndex){
            writer.writeStringValue("conversationIndex", this.conversationIndex);
        }
        if(this.extensions && this.extensions.length != 0){        const extensionsArrValue: ExtensionImpl[] = [];
        this.extensions?.forEach(element => {
            extensionsArrValue.push((element instanceof ExtensionImpl? element as ExtensionImpl:new ExtensionImpl(element)));
        });
            writer.writeCollectionOfObjectValues<ExtensionImpl>("extensions", extensionsArrValue);
        }
        if(this.flag){
            writer.writeObjectValue<FollowupFlagImpl>("flag", (this.flag instanceof FollowupFlagImpl? this.flag as FollowupFlagImpl: new FollowupFlagImpl(this.flag)));
        }
        if(this.from){
            writer.writeObjectValue<RecipientImpl>("from", (this.from instanceof RecipientImpl? this.from as RecipientImpl: new RecipientImpl(this.from)));
        }
        if(this.hasAttachments){
            writer.writeBooleanValue("hasAttachments", this.hasAttachments);
        }
        if(this.importance){
            writer.writeEnumValue<Importance>("importance", this.importance);
        }
        if(this.inferenceClassification){
            writer.writeEnumValue<InferenceClassificationType>("inferenceClassification", this.inferenceClassification);
        }
        if(this.internetMessageHeaders && this.internetMessageHeaders.length != 0){        const internetMessageHeadersArrValue: InternetMessageHeaderImpl[] = [];
        this.internetMessageHeaders?.forEach(element => {
            internetMessageHeadersArrValue.push((element instanceof InternetMessageHeaderImpl? element as InternetMessageHeaderImpl:new InternetMessageHeaderImpl(element)));
        });
            writer.writeCollectionOfObjectValues<InternetMessageHeaderImpl>("internetMessageHeaders", internetMessageHeadersArrValue);
        }
        if(this.internetMessageId){
            writer.writeStringValue("internetMessageId", this.internetMessageId);
        }
        if(this.isDeliveryReceiptRequested){
            writer.writeBooleanValue("isDeliveryReceiptRequested", this.isDeliveryReceiptRequested);
        }
        if(this.isDraft){
            writer.writeBooleanValue("isDraft", this.isDraft);
        }
        if(this.isRead){
            writer.writeBooleanValue("isRead", this.isRead);
        }
        if(this.isReadReceiptRequested){
            writer.writeBooleanValue("isReadReceiptRequested", this.isReadReceiptRequested);
        }
        if(this.multiValueExtendedProperties && this.multiValueExtendedProperties.length != 0){        const multiValueExtendedPropertiesArrValue: MultiValueLegacyExtendedPropertyImpl[] = [];
        this.multiValueExtendedProperties?.forEach(element => {
            multiValueExtendedPropertiesArrValue.push((element instanceof MultiValueLegacyExtendedPropertyImpl? element as MultiValueLegacyExtendedPropertyImpl:new MultiValueLegacyExtendedPropertyImpl(element)));
        });
            writer.writeCollectionOfObjectValues<MultiValueLegacyExtendedPropertyImpl>("multiValueExtendedProperties", multiValueExtendedPropertiesArrValue);
        }
        if(this.parentFolderId){
            writer.writeStringValue("parentFolderId", this.parentFolderId);
        }
        if(this.receivedDateTime){
            writer.writeDateValue("receivedDateTime", this.receivedDateTime);
        }
        if(this.replyTo && this.replyTo.length != 0){        const replyToArrValue: RecipientImpl[] = [];
        this.replyTo?.forEach(element => {
            replyToArrValue.push((element instanceof RecipientImpl? element as RecipientImpl:new RecipientImpl(element)));
        });
            writer.writeCollectionOfObjectValues<RecipientImpl>("replyTo", replyToArrValue);
        }
        if(this.sender){
            writer.writeObjectValue<RecipientImpl>("sender", (this.sender instanceof RecipientImpl? this.sender as RecipientImpl: new RecipientImpl(this.sender)));
        }
        if(this.sentDateTime){
            writer.writeDateValue("sentDateTime", this.sentDateTime);
        }
        if(this.singleValueExtendedProperties && this.singleValueExtendedProperties.length != 0){        const singleValueExtendedPropertiesArrValue: SingleValueLegacyExtendedPropertyImpl[] = [];
        this.singleValueExtendedProperties?.forEach(element => {
            singleValueExtendedPropertiesArrValue.push((element instanceof SingleValueLegacyExtendedPropertyImpl? element as SingleValueLegacyExtendedPropertyImpl:new SingleValueLegacyExtendedPropertyImpl(element)));
        });
            writer.writeCollectionOfObjectValues<SingleValueLegacyExtendedPropertyImpl>("singleValueExtendedProperties", singleValueExtendedPropertiesArrValue);
        }
        if(this.subject){
            writer.writeStringValue("subject", this.subject);
        }
        if(this.toRecipients && this.toRecipients.length != 0){        const toRecipientsArrValue: RecipientImpl[] = [];
        this.toRecipients?.forEach(element => {
            toRecipientsArrValue.push((element instanceof RecipientImpl? element as RecipientImpl:new RecipientImpl(element)));
        });
            writer.writeCollectionOfObjectValues<RecipientImpl>("toRecipients", toRecipientsArrValue);
        }
        if(this.uniqueBody){
            writer.writeObjectValue<ItemBodyImpl>("uniqueBody", (this.uniqueBody instanceof ItemBodyImpl? this.uniqueBody as ItemBodyImpl: new ItemBodyImpl(this.uniqueBody)));
        }
        if(this.webLink){
            writer.writeStringValue("webLink", this.webLink);
        }
    };
    /**
     * Gets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the message. Nullable.
     * @returns a SingleValueLegacyExtendedPropertyInterface
     */
    public get singleValueExtendedProperties() {
        return this._singleValueExtendedProperties;
    };
    /**
     * Sets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the message. Nullable.
     * @param value Value to set for the singleValueExtendedProperties property.
     */
    public set singleValueExtendedProperties(value: SingleValueLegacyExtendedProperty[] | undefined) {
        if(value) {
            const singleValueExtendedPropertiesArrValue: SingleValueLegacyExtendedPropertyImpl[] = [];
            this.singleValueExtendedProperties?.forEach(element => {
                singleValueExtendedPropertiesArrValue.push((element instanceof SingleValueLegacyExtendedPropertyImpl? element as SingleValueLegacyExtendedPropertyImpl:new SingleValueLegacyExtendedPropertyImpl(element)));
            });
            this._singleValueExtendedProperties = singleValueExtendedPropertiesArrValue;
        }
    };
    /**
     * Gets the subject property value. The subject property
     * @returns a string
     */
    public get subject() {
        return this._subject;
    };
    /**
     * Sets the subject property value. The subject property
     * @param value Value to set for the subject property.
     */
    public set subject(value: string | undefined) {
        if(value) {
            this._subject = value;
        }
    };
    /**
     * Gets the toRecipients property value. The toRecipients property
     * @returns a RecipientInterface
     */
    public get toRecipients() {
        return this._toRecipients;
    };
    /**
     * Sets the toRecipients property value. The toRecipients property
     * @param value Value to set for the toRecipients property.
     */
    public set toRecipients(value: Recipient[] | undefined) {
        if(value) {
            const toRecipientsArrValue: RecipientImpl[] = [];
            this.toRecipients?.forEach(element => {
                toRecipientsArrValue.push((element instanceof RecipientImpl? element as RecipientImpl:new RecipientImpl(element)));
            });
            this._toRecipients = toRecipientsArrValue;
        }
    };
    /**
     * Gets the uniqueBody property value. The uniqueBody property
     * @returns a ItemBodyInterface
     */
    public get uniqueBody() {
        return this._uniqueBody;
    };
    /**
     * Sets the uniqueBody property value. The uniqueBody property
     * @param value Value to set for the uniqueBody property.
     */
    public set uniqueBody(value: ItemBody | undefined) {
        if(value) {
            this._uniqueBody = value instanceof ItemBodyImpl? value as ItemBodyImpl: new ItemBodyImpl(value);
        }
    };
    /**
     * Gets the webLink property value. The webLink property
     * @returns a string
     */
    public get webLink() {
        return this._webLink;
    };
    /**
     * Sets the webLink property value. The webLink property
     * @param value Value to set for the webLink property.
     */
    public set webLink(value: string | undefined) {
        if(value) {
            this._webLink = value;
        }
    };
>>>>>>> main
}
