import {Attachment} from './attachment';
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
import {Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export class Message extends OutlookItem implements Parsable {
    /** The fileAttachment and itemAttachment attachments for the message.  */
    private _attachments?: Attachment[] | undefined;
    /** The Bcc: recipients for the message.  */
    private _bccRecipients?: Recipient[] | undefined;
    private _body?: ItemBody | undefined;
    /** The first 255 characters of the message body. It is in text format. If the message contains instances of mention, this property would contain a concatenation of these mentions as well.  */
    private _bodyPreview?: string | undefined;
    /** The Cc: recipients for the message.  */
    private _ccRecipients?: Recipient[] | undefined;
    /** The ID of the conversation the email belongs to.  */
    private _conversationId?: string | undefined;
    /** Indicates the position of the message within the conversation.  */
    private _conversationIndex?: string | undefined;
    /** The collection of open extensions defined for the message. Nullable.  */
    private _extensions?: Extension[] | undefined;
    private _flag?: FollowupFlag | undefined;
    private _from?: Recipient | undefined;
    /** Indicates whether the message has attachments. This property doesn't include inline attachments, so if a message contains only inline attachments, this property is false. To verify the existence of inline attachments, parse the body property to look for a src attribute, such as <IMG src='cid:image001.jpg@01D26CD8.6C05F070'>.  */
    private _hasAttachments?: boolean | undefined;
    private _importance?: Importance | undefined;
    private _inferenceClassification?: InferenceClassificationType | undefined;
    /** A collection of message headers defined by RFC5322. The set includes message headers indicating the network path taken by a message from the sender to the recipient. It can also contain custom message headers that hold app data for the message.  Returned only on applying a $select query option. Read-only.  */
    private _internetMessageHeaders?: InternetMessageHeader[] | undefined;
    /** The message ID in the format specified by RFC2822.  */
    private _internetMessageId?: string | undefined;
    /** Indicates whether a read receipt is requested for the message.  */
    private _isDeliveryReceiptRequested?: boolean | undefined;
    /** Indicates whether the message is a draft. A message is a draft if it hasn't been sent yet.  */
    private _isDraft?: boolean | undefined;
    /** Indicates whether the message has been read.  */
    private _isRead?: boolean | undefined;
    /** Indicates whether a read receipt is requested for the message.  */
    private _isReadReceiptRequested?: boolean | undefined;
    /** The collection of multi-value extended properties defined for the message. Nullable.  */
    private _multiValueExtendedProperties?: MultiValueLegacyExtendedProperty[] | undefined;
    /** The unique identifier for the message's parent mailFolder.  */
    private _parentFolderId?: string | undefined;
    /** The date and time the message was received.  The date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.  */
    private _receivedDateTime?: Date | undefined;
    /** The email addresses to use when replying.  */
    private _replyTo?: Recipient[] | undefined;
    private _sender?: Recipient | undefined;
    /** The date and time the message was sent.  The date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.  */
    private _sentDateTime?: Date | undefined;
    /** The collection of single-value extended properties defined for the message. Nullable.  */
    private _singleValueExtendedProperties?: SingleValueLegacyExtendedProperty[] | undefined;
    /** The subject of the message.  */
    private _subject?: string | undefined;
    /** The To: recipients for the message.  */
    private _toRecipients?: Recipient[] | undefined;
    private _uniqueBody?: ItemBody | undefined;
    /** The URL to open the message in Outlook on the web.You can append an ispopout argument to the end of the URL to change how the message is displayed. If ispopout is not present or if it is set to 1, then the message is shown in a popout window. If ispopout is set to 0, then the browser will show the message in the Outlook on the web review pane.The message will open in the browser if you are logged in to your mailbox via Outlook on the web. You will be prompted to login if you are not already logged in with the browser.This URL cannot be accessed from within an iFrame.  */
    private _webLink?: string | undefined;
    /**
     * Instantiates a new message and sets the default values.
     */
    public constructor() {
        super();
    };
    /**
     * Gets the attachments property value. The fileAttachment and itemAttachment attachments for the message.
     * @returns a attachment
     */
    public get attachments() {
        return this._attachments;
    };
    /**
     * Gets the bccRecipients property value. The Bcc: recipients for the message.
     * @returns a recipient
     */
    public get bccRecipients() {
        return this._bccRecipients;
    };
    /**
     * Gets the body property value. 
     * @returns a itemBody
     */
    public get body() {
        return this._body;
    };
    /**
     * Gets the bodyPreview property value. The first 255 characters of the message body. It is in text format. If the message contains instances of mention, this property would contain a concatenation of these mentions as well.
     * @returns a string
     */
    public get bodyPreview() {
        return this._bodyPreview;
    };
    /**
     * Gets the ccRecipients property value. The Cc: recipients for the message.
     * @returns a recipient
     */
    public get ccRecipients() {
        return this._ccRecipients;
    };
    /**
     * Gets the conversationId property value. The ID of the conversation the email belongs to.
     * @returns a string
     */
    public get conversationId() {
        return this._conversationId;
    };
    /**
     * Gets the conversationIndex property value. Indicates the position of the message within the conversation.
     * @returns a binary
     */
    public get conversationIndex() {
        return this._conversationIndex;
    };
    /**
     * Gets the extensions property value. The collection of open extensions defined for the message. Nullable.
     * @returns a extension
     */
    public get extensions() {
        return this._extensions;
    };
    /**
     * Gets the flag property value. 
     * @returns a followupFlag
     */
    public get flag() {
        return this._flag;
    };
    /**
     * Gets the from property value. 
     * @returns a recipient
     */
    public get from() {
        return this._from;
    };
    /**
     * Gets the hasAttachments property value. Indicates whether the message has attachments. This property doesn't include inline attachments, so if a message contains only inline attachments, this property is false. To verify the existence of inline attachments, parse the body property to look for a src attribute, such as <IMG src='cid:image001.jpg@01D26CD8.6C05F070'>.
     * @returns a boolean
     */
    public get hasAttachments() {
        return this._hasAttachments;
    };
    /**
     * Gets the importance property value. 
     * @returns a importance
     */
    public get importance() {
        return this._importance;
    };
    /**
     * Gets the inferenceClassification property value. 
     * @returns a inferenceClassificationType
     */
    public get inferenceClassification() {
        return this._inferenceClassification;
    };
    /**
     * Gets the internetMessageHeaders property value. A collection of message headers defined by RFC5322. The set includes message headers indicating the network path taken by a message from the sender to the recipient. It can also contain custom message headers that hold app data for the message.  Returned only on applying a $select query option. Read-only.
     * @returns a internetMessageHeader
     */
    public get internetMessageHeaders() {
        return this._internetMessageHeaders;
    };
    /**
     * Gets the internetMessageId property value. The message ID in the format specified by RFC2822.
     * @returns a string
     */
    public get internetMessageId() {
        return this._internetMessageId;
    };
    /**
     * Gets the isDeliveryReceiptRequested property value. Indicates whether a read receipt is requested for the message.
     * @returns a boolean
     */
    public get isDeliveryReceiptRequested() {
        return this._isDeliveryReceiptRequested;
    };
    /**
     * Gets the isDraft property value. Indicates whether the message is a draft. A message is a draft if it hasn't been sent yet.
     * @returns a boolean
     */
    public get isDraft() {
        return this._isDraft;
    };
    /**
     * Gets the isRead property value. Indicates whether the message has been read.
     * @returns a boolean
     */
    public get isRead() {
        return this._isRead;
    };
    /**
     * Gets the isReadReceiptRequested property value. Indicates whether a read receipt is requested for the message.
     * @returns a boolean
     */
    public get isReadReceiptRequested() {
        return this._isReadReceiptRequested;
    };
    /**
     * Gets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the message. Nullable.
     * @returns a multiValueLegacyExtendedProperty
     */
    public get multiValueExtendedProperties() {
        return this._multiValueExtendedProperties;
    };
    /**
     * Gets the parentFolderId property value. The unique identifier for the message's parent mailFolder.
     * @returns a string
     */
    public get parentFolderId() {
        return this._parentFolderId;
    };
    /**
     * Gets the receivedDateTime property value. The date and time the message was received.  The date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
     * @returns a Date
     */
    public get receivedDateTime() {
        return this._receivedDateTime;
    };
    /**
     * Gets the replyTo property value. The email addresses to use when replying.
     * @returns a recipient
     */
    public get replyTo() {
        return this._replyTo;
    };
    /**
     * Gets the sender property value. 
     * @returns a recipient
     */
    public get sender() {
        return this._sender;
    };
    /**
     * Gets the sentDateTime property value. The date and time the message was sent.  The date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
     * @returns a Date
     */
    public get sentDateTime() {
        return this._sentDateTime;
    };
    /**
     * Gets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the message. Nullable.
     * @returns a singleValueLegacyExtendedProperty
     */
    public get singleValueExtendedProperties() {
        return this._singleValueExtendedProperties;
    };
    /**
     * Gets the subject property value. The subject of the message.
     * @returns a string
     */
    public get subject() {
        return this._subject;
    };
    /**
     * Gets the toRecipients property value. The To: recipients for the message.
     * @returns a recipient
     */
    public get toRecipients() {
        return this._toRecipients;
    };
    /**
     * Gets the uniqueBody property value. 
     * @returns a itemBody
     */
    public get uniqueBody() {
        return this._uniqueBody;
    };
    /**
     * Gets the webLink property value. The URL to open the message in Outlook on the web.You can append an ispopout argument to the end of the URL to change how the message is displayed. If ispopout is not present or if it is set to 1, then the message is shown in a popout window. If ispopout is set to 0, then the browser will show the message in the Outlook on the web review pane.The message will open in the browser if you are logged in to your mailbox via Outlook on the web. You will be prompted to login if you are not already logged in with the browser.This URL cannot be accessed from within an iFrame.
     * @returns a string
     */
    public get webLink() {
        return this._webLink;
    };
    /**
     * The deserialization information for the current model
     * @returns a Map<string, (item: T, node: ParseNode) => void>
     */
    public getFieldDeserializers<T>() : Map<string, (item: T, node: ParseNode) => void> {
        return new Map<string, (item: T, node: ParseNode) => void>([...super.getFieldDeserializers<T>(),
            ["attachments", (o, n) => { (o as unknown as Message).attachments = n.getCollectionOfObjectValues<Attachment>(Attachment); }],
            ["bccRecipients", (o, n) => { (o as unknown as Message).bccRecipients = n.getCollectionOfObjectValues<Recipient>(Recipient); }],
            ["body", (o, n) => { (o as unknown as Message).body = n.getObjectValue<ItemBody>(ItemBody); }],
            ["bodyPreview", (o, n) => { (o as unknown as Message).bodyPreview = n.getStringValue(); }],
            ["ccRecipients", (o, n) => { (o as unknown as Message).ccRecipients = n.getCollectionOfObjectValues<Recipient>(Recipient); }],
            ["conversationId", (o, n) => { (o as unknown as Message).conversationId = n.getStringValue(); }],
            ["conversationIndex", (o, n) => { (o as unknown as Message).conversationIndex = n.getStringValue(); }],
            ["extensions", (o, n) => { (o as unknown as Message).extensions = n.getCollectionOfObjectValues<Extension>(Extension); }],
            ["flag", (o, n) => { (o as unknown as Message).flag = n.getObjectValue<FollowupFlag>(FollowupFlag); }],
            ["from", (o, n) => { (o as unknown as Message).from = n.getObjectValue<Recipient>(Recipient); }],
            ["hasAttachments", (o, n) => { (o as unknown as Message).hasAttachments = n.getBooleanValue(); }],
            ["importance", (o, n) => { (o as unknown as Message).importance = n.getEnumValue<Importance>(Importance); }],
            ["inferenceClassification", (o, n) => { (o as unknown as Message).inferenceClassification = n.getEnumValue<InferenceClassificationType>(InferenceClassificationType); }],
            ["internetMessageHeaders", (o, n) => { (o as unknown as Message).internetMessageHeaders = n.getCollectionOfObjectValues<InternetMessageHeader>(InternetMessageHeader); }],
            ["internetMessageId", (o, n) => { (o as unknown as Message).internetMessageId = n.getStringValue(); }],
            ["isDeliveryReceiptRequested", (o, n) => { (o as unknown as Message).isDeliveryReceiptRequested = n.getBooleanValue(); }],
            ["isDraft", (o, n) => { (o as unknown as Message).isDraft = n.getBooleanValue(); }],
            ["isRead", (o, n) => { (o as unknown as Message).isRead = n.getBooleanValue(); }],
            ["isReadReceiptRequested", (o, n) => { (o as unknown as Message).isReadReceiptRequested = n.getBooleanValue(); }],
            ["multiValueExtendedProperties", (o, n) => { (o as unknown as Message).multiValueExtendedProperties = n.getCollectionOfObjectValues<MultiValueLegacyExtendedProperty>(MultiValueLegacyExtendedProperty); }],
            ["parentFolderId", (o, n) => { (o as unknown as Message).parentFolderId = n.getStringValue(); }],
            ["receivedDateTime", (o, n) => { (o as unknown as Message).receivedDateTime = n.getDateValue(); }],
            ["replyTo", (o, n) => { (o as unknown as Message).replyTo = n.getCollectionOfObjectValues<Recipient>(Recipient); }],
            ["sender", (o, n) => { (o as unknown as Message).sender = n.getObjectValue<Recipient>(Recipient); }],
            ["sentDateTime", (o, n) => { (o as unknown as Message).sentDateTime = n.getDateValue(); }],
            ["singleValueExtendedProperties", (o, n) => { (o as unknown as Message).singleValueExtendedProperties = n.getCollectionOfObjectValues<SingleValueLegacyExtendedProperty>(SingleValueLegacyExtendedProperty); }],
            ["subject", (o, n) => { (o as unknown as Message).subject = n.getStringValue(); }],
            ["toRecipients", (o, n) => { (o as unknown as Message).toRecipients = n.getCollectionOfObjectValues<Recipient>(Recipient); }],
            ["uniqueBody", (o, n) => { (o as unknown as Message).uniqueBody = n.getObjectValue<ItemBody>(ItemBody); }],
            ["webLink", (o, n) => { (o as unknown as Message).webLink = n.getStringValue(); }],
        ]);
    };
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     */
    public serialize(writer: SerializationWriter) : void {
        if(!writer) throw new Error("writer cannot be undefined");
        super.serialize(writer);
        writer.writeCollectionOfObjectValues<Attachment>("attachments", this.attachments);
        writer.writeCollectionOfObjectValues<Recipient>("bccRecipients", this.bccRecipients);
        writer.writeObjectValue<ItemBody>("body", this.body);
        writer.writeStringValue("bodyPreview", this.bodyPreview);
        writer.writeCollectionOfObjectValues<Recipient>("ccRecipients", this.ccRecipients);
        writer.writeStringValue("conversationId", this.conversationId);
        writer.writeStringValue("conversationIndex", this.conversationIndex);
        writer.writeCollectionOfObjectValues<Extension>("extensions", this.extensions);
        writer.writeObjectValue<FollowupFlag>("flag", this.flag);
        writer.writeObjectValue<Recipient>("from", this.from);
        writer.writeBooleanValue("hasAttachments", this.hasAttachments);
        writer.writeEnumValue<Importance>("importance", this.importance);
        writer.writeEnumValue<InferenceClassificationType>("inferenceClassification", this.inferenceClassification);
        writer.writeCollectionOfObjectValues<InternetMessageHeader>("internetMessageHeaders", this.internetMessageHeaders);
        writer.writeStringValue("internetMessageId", this.internetMessageId);
        writer.writeBooleanValue("isDeliveryReceiptRequested", this.isDeliveryReceiptRequested);
        writer.writeBooleanValue("isDraft", this.isDraft);
        writer.writeBooleanValue("isRead", this.isRead);
        writer.writeBooleanValue("isReadReceiptRequested", this.isReadReceiptRequested);
        writer.writeCollectionOfObjectValues<MultiValueLegacyExtendedProperty>("multiValueExtendedProperties", this.multiValueExtendedProperties);
        writer.writeStringValue("parentFolderId", this.parentFolderId);
        writer.writeDateValue("receivedDateTime", this.receivedDateTime);
        writer.writeCollectionOfObjectValues<Recipient>("replyTo", this.replyTo);
        writer.writeObjectValue<Recipient>("sender", this.sender);
        writer.writeDateValue("sentDateTime", this.sentDateTime);
        writer.writeCollectionOfObjectValues<SingleValueLegacyExtendedProperty>("singleValueExtendedProperties", this.singleValueExtendedProperties);
        writer.writeStringValue("subject", this.subject);
        writer.writeCollectionOfObjectValues<Recipient>("toRecipients", this.toRecipients);
        writer.writeObjectValue<ItemBody>("uniqueBody", this.uniqueBody);
        writer.writeStringValue("webLink", this.webLink);
    };
    /**
     * Sets the attachments property value. The fileAttachment and itemAttachment attachments for the message.
     * @param value Value to set for the attachments property.
     */
    public set attachments(value: Attachment[] | undefined) {
        this._attachments = value;
    };
    /**
     * Sets the bccRecipients property value. The Bcc: recipients for the message.
     * @param value Value to set for the bccRecipients property.
     */
    public set bccRecipients(value: Recipient[] | undefined) {
        this._bccRecipients = value;
    };
    /**
     * Sets the body property value. 
     * @param value Value to set for the body property.
     */
    public set body(value: ItemBody | undefined) {
        this._body = value;
    };
    /**
     * Sets the bodyPreview property value. The first 255 characters of the message body. It is in text format. If the message contains instances of mention, this property would contain a concatenation of these mentions as well.
     * @param value Value to set for the bodyPreview property.
     */
    public set bodyPreview(value: string | undefined) {
        this._bodyPreview = value;
    };
    /**
     * Sets the ccRecipients property value. The Cc: recipients for the message.
     * @param value Value to set for the ccRecipients property.
     */
    public set ccRecipients(value: Recipient[] | undefined) {
        this._ccRecipients = value;
    };
    /**
     * Sets the conversationId property value. The ID of the conversation the email belongs to.
     * @param value Value to set for the conversationId property.
     */
    public set conversationId(value: string | undefined) {
        this._conversationId = value;
    };
    /**
     * Sets the conversationIndex property value. Indicates the position of the message within the conversation.
     * @param value Value to set for the conversationIndex property.
     */
    public set conversationIndex(value: string | undefined) {
        this._conversationIndex = value;
    };
    /**
     * Sets the extensions property value. The collection of open extensions defined for the message. Nullable.
     * @param value Value to set for the extensions property.
     */
    public set extensions(value: Extension[] | undefined) {
        this._extensions = value;
    };
    /**
     * Sets the flag property value. 
     * @param value Value to set for the flag property.
     */
    public set flag(value: FollowupFlag | undefined) {
        this._flag = value;
    };
    /**
     * Sets the from property value. 
     * @param value Value to set for the from property.
     */
    public set from(value: Recipient | undefined) {
        this._from = value;
    };
    /**
     * Sets the hasAttachments property value. Indicates whether the message has attachments. This property doesn't include inline attachments, so if a message contains only inline attachments, this property is false. To verify the existence of inline attachments, parse the body property to look for a src attribute, such as <IMG src='cid:image001.jpg@01D26CD8.6C05F070'>.
     * @param value Value to set for the hasAttachments property.
     */
    public set hasAttachments(value: boolean | undefined) {
        this._hasAttachments = value;
    };
    /**
     * Sets the importance property value. 
     * @param value Value to set for the importance property.
     */
    public set importance(value: Importance | undefined) {
        this._importance = value;
    };
    /**
     * Sets the inferenceClassification property value. 
     * @param value Value to set for the inferenceClassification property.
     */
    public set inferenceClassification(value: InferenceClassificationType | undefined) {
        this._inferenceClassification = value;
    };
    /**
     * Sets the internetMessageHeaders property value. A collection of message headers defined by RFC5322. The set includes message headers indicating the network path taken by a message from the sender to the recipient. It can also contain custom message headers that hold app data for the message.  Returned only on applying a $select query option. Read-only.
     * @param value Value to set for the internetMessageHeaders property.
     */
    public set internetMessageHeaders(value: InternetMessageHeader[] | undefined) {
        this._internetMessageHeaders = value;
    };
    /**
     * Sets the internetMessageId property value. The message ID in the format specified by RFC2822.
     * @param value Value to set for the internetMessageId property.
     */
    public set internetMessageId(value: string | undefined) {
        this._internetMessageId = value;
    };
    /**
     * Sets the isDeliveryReceiptRequested property value. Indicates whether a read receipt is requested for the message.
     * @param value Value to set for the isDeliveryReceiptRequested property.
     */
    public set isDeliveryReceiptRequested(value: boolean | undefined) {
        this._isDeliveryReceiptRequested = value;
    };
    /**
     * Sets the isDraft property value. Indicates whether the message is a draft. A message is a draft if it hasn't been sent yet.
     * @param value Value to set for the isDraft property.
     */
    public set isDraft(value: boolean | undefined) {
        this._isDraft = value;
    };
    /**
     * Sets the isRead property value. Indicates whether the message has been read.
     * @param value Value to set for the isRead property.
     */
    public set isRead(value: boolean | undefined) {
        this._isRead = value;
    };
    /**
     * Sets the isReadReceiptRequested property value. Indicates whether a read receipt is requested for the message.
     * @param value Value to set for the isReadReceiptRequested property.
     */
    public set isReadReceiptRequested(value: boolean | undefined) {
        this._isReadReceiptRequested = value;
    };
    /**
     * Sets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the message. Nullable.
     * @param value Value to set for the multiValueExtendedProperties property.
     */
    public set multiValueExtendedProperties(value: MultiValueLegacyExtendedProperty[] | undefined) {
        this._multiValueExtendedProperties = value;
    };
    /**
     * Sets the parentFolderId property value. The unique identifier for the message's parent mailFolder.
     * @param value Value to set for the parentFolderId property.
     */
    public set parentFolderId(value: string | undefined) {
        this._parentFolderId = value;
    };
    /**
     * Sets the receivedDateTime property value. The date and time the message was received.  The date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
     * @param value Value to set for the receivedDateTime property.
     */
    public set receivedDateTime(value: Date | undefined) {
        this._receivedDateTime = value;
    };
    /**
     * Sets the replyTo property value. The email addresses to use when replying.
     * @param value Value to set for the replyTo property.
     */
    public set replyTo(value: Recipient[] | undefined) {
        this._replyTo = value;
    };
    /**
     * Sets the sender property value. 
     * @param value Value to set for the sender property.
     */
    public set sender(value: Recipient | undefined) {
        this._sender = value;
    };
    /**
     * Sets the sentDateTime property value. The date and time the message was sent.  The date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
     * @param value Value to set for the sentDateTime property.
     */
    public set sentDateTime(value: Date | undefined) {
        this._sentDateTime = value;
    };
    /**
     * Sets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the message. Nullable.
     * @param value Value to set for the singleValueExtendedProperties property.
     */
    public set singleValueExtendedProperties(value: SingleValueLegacyExtendedProperty[] | undefined) {
        this._singleValueExtendedProperties = value;
    };
    /**
     * Sets the subject property value. The subject of the message.
     * @param value Value to set for the subject property.
     */
    public set subject(value: string | undefined) {
        this._subject = value;
    };
    /**
     * Sets the toRecipients property value. The To: recipients for the message.
     * @param value Value to set for the toRecipients property.
     */
    public set toRecipients(value: Recipient[] | undefined) {
        this._toRecipients = value;
    };
    /**
     * Sets the uniqueBody property value. 
     * @param value Value to set for the uniqueBody property.
     */
    public set uniqueBody(value: ItemBody | undefined) {
        this._uniqueBody = value;
    };
    /**
     * Sets the webLink property value. The URL to open the message in Outlook on the web.You can append an ispopout argument to the end of the URL to change how the message is displayed. If ispopout is not present or if it is set to 1, then the message is shown in a popout window. If ispopout is set to 0, then the browser will show the message in the Outlook on the web review pane.The message will open in the browser if you are logged in to your mailbox via Outlook on the web. You will be prompted to login if you are not already logged in with the browser.This URL cannot be accessed from within an iFrame.
     * @param value Value to set for the webLink property.
     */
    public set webLink(value: string | undefined) {
        this._webLink = value;
    };
}
