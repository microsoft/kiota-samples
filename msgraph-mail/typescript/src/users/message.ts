import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';
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

export class Message extends OutlookItem implements Parsable<Message> {
    /** The Bcc: recipients for the message.  */
    public bccRecipients?: Recipient[] | undefined;
    public body?: ItemBody | undefined;
    /** The first 255 characters of the message body. It is in text format. If the message contains instances of mention, this property would contain a concatenation of these mentions as well.  */
    public bodyPreview?: string | undefined;
    /** The Cc: recipients for the message.  */
    public ccRecipients?: Recipient[] | undefined;
    /** The ID of the conversation the email belongs to.  */
    public conversationId?: string | undefined;
    /** Indicates the position of the message within the conversation.  */
    public conversationIndex?: string | undefined;
    public flag?: FollowupFlag | undefined;
    public from?: Recipient | undefined;
    /** Indicates whether the message has attachments. This property doesn't include inline attachments, so if a message contains only inline attachments, this property is false. To verify the existence of inline attachments, parse the body property to look for a src attribute, such as <IMG src='cid:image001.jpg@01D26CD8.6C05F070'>.  */
    public hasAttachments?: boolean | undefined;
    public importance?: Importance | undefined;
    public inferenceClassification?: InferenceClassificationType | undefined;
    /** A collection of message headers defined by RFC5322. The set includes message headers indicating the network path taken by a message from the sender to the recipient. It can also contain custom message headers that hold app data for the message.  Returned only on applying a $select query option. Read-only.  */
    public internetMessageHeaders?: InternetMessageHeader[] | undefined;
    /** The message ID in the format specified by RFC2822.  */
    public internetMessageId?: string | undefined;
    /** Indicates whether a read receipt is requested for the message.  */
    public isDeliveryReceiptRequested?: boolean | undefined;
    /** Indicates whether the message is a draft. A message is a draft if it hasn't been sent yet.  */
    public isDraft?: boolean | undefined;
    /** Indicates whether the message has been read.  */
    public isRead?: boolean | undefined;
    /** Indicates whether a read receipt is requested for the message.  */
    public isReadReceiptRequested?: boolean | undefined;
    /** The unique identifier for the message's parent mailFolder.  */
    public parentFolderId?: string | undefined;
    /** The date and time the message was received.  The date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.  */
    public receivedDateTime?: Date | undefined;
    /** The email addresses to use when replying.  */
    public replyTo?: Recipient[] | undefined;
    public sender?: Recipient | undefined;
    /** The date and time the message was sent.  The date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.  */
    public sentDateTime?: Date | undefined;
    /** The subject of the message.  */
    public subject?: string | undefined;
    /** The To: recipients for the message.  */
    public toRecipients?: Recipient[] | undefined;
    public uniqueBody?: ItemBody | undefined;
    /** The URL to open the message in Outlook on the web.You can append an ispopout argument to the end of the URL to change how the message is displayed. If ispopout is not present or if it is set to 1, then the message is shown in a popout window. If ispopout is set to 0, then the browser will show the message in the Outlook on the web review pane.The message will open in the browser if you are logged in to your mailbox via Outlook on the web. You will be prompted to login if you are not already logged in with the browser.This URL cannot be accessed from within an iFrame.  */
    public webLink?: string | undefined;
    /** The fileAttachment and itemAttachment attachments for the message.  */
    public attachments?: Attachment[] | undefined;
    /** The collection of open extensions defined for the message. Nullable.  */
    public extensions?: Extension[] | undefined;
    /** The collection of multi-value extended properties defined for the message. Nullable.  */
    public multiValueExtendedProperties?: MultiValueLegacyExtendedProperty[] | undefined;
    /** The collection of single-value extended properties defined for the message. Nullable.  */
    public singleValueExtendedProperties?: SingleValueLegacyExtendedProperty[] | undefined;
    /**
     * Serialiazes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @returns a void
     */
    public serialize (writer: SerializationWriter) : void {
        super.serialize(writer);
        writer.writeCollectionOfObjectValues<Recipient>("bccRecipients", this.bccRecipients);
        writer.writeObjectValue<ItemBody>("body", this.body);
        writer.writeStringValue("bodyPreview", this.bodyPreview);
        writer.writeCollectionOfObjectValues<Recipient>("ccRecipients", this.ccRecipients);
        writer.writeStringValue("conversationId", this.conversationId);
        writer.writeStringValue("conversationIndex", this.conversationIndex);
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
        writer.writeStringValue("parentFolderId", this.parentFolderId);
        writer.writeDateValue("receivedDateTime", this.receivedDateTime);
        writer.writeCollectionOfObjectValues<Recipient>("replyTo", this.replyTo);
        writer.writeObjectValue<Recipient>("sender", this.sender);
        writer.writeDateValue("sentDateTime", this.sentDateTime);
        writer.writeStringValue("subject", this.subject);
        writer.writeCollectionOfObjectValues<Recipient>("toRecipients", this.toRecipients);
        writer.writeObjectValue<ItemBody>("uniqueBody", this.uniqueBody);
        writer.writeStringValue("webLink", this.webLink);
        writer.writeCollectionOfObjectValues<Attachment>("attachments", this.attachments);
        writer.writeCollectionOfObjectValues<Extension>("extensions", this.extensions);
        writer.writeCollectionOfObjectValues<MultiValueLegacyExtendedProperty>("multiValueExtendedProperties", this.multiValueExtendedProperties);
        writer.writeCollectionOfObjectValues<SingleValueLegacyExtendedProperty>("singleValueExtendedProperties", this.singleValueExtendedProperties);
    };
    /**
     * The serialization information for the current model
     * @returns a Map<string, (item: Message, node: ParseNode) => void>
     */
    public deserializeFields () : Map<string, (item: Message, node: ParseNode) => void> {
        return new Map<string, (item: Message, node: ParseNode) => void>([...super.deserializeFields(),
            ["bccRecipients", (o, n) => { o.bccRecipients = n.getCollectionOfObjectValues<Recipient>(Recipient); }],
            ["body", (o, n) => { o.body = n.getObjectValue<ItemBody>(ItemBody); }],
            ["bodyPreview", (o, n) => { o.bodyPreview = n.getStringValue(); }],
            ["ccRecipients", (o, n) => { o.ccRecipients = n.getCollectionOfObjectValues<Recipient>(Recipient); }],
            ["conversationId", (o, n) => { o.conversationId = n.getStringValue(); }],
            ["conversationIndex", (o, n) => { o.conversationIndex = n.getStringValue(); }],
            ["flag", (o, n) => { o.flag = n.getObjectValue<FollowupFlag>(FollowupFlag); }],
            ["from", (o, n) => { o.from = n.getObjectValue<Recipient>(Recipient); }],
            ["hasAttachments", (o, n) => { o.hasAttachments = n.getBooleanValue(); }],
            ["importance", (o, n) => { o.importance = n.getEnumValue<Importance>(Importance); }],
            ["inferenceClassification", (o, n) => { o.inferenceClassification = n.getEnumValue<InferenceClassificationType>(InferenceClassificationType); }],
            ["internetMessageHeaders", (o, n) => { o.internetMessageHeaders = n.getCollectionOfObjectValues<InternetMessageHeader>(InternetMessageHeader); }],
            ["internetMessageId", (o, n) => { o.internetMessageId = n.getStringValue(); }],
            ["isDeliveryReceiptRequested", (o, n) => { o.isDeliveryReceiptRequested = n.getBooleanValue(); }],
            ["isDraft", (o, n) => { o.isDraft = n.getBooleanValue(); }],
            ["isRead", (o, n) => { o.isRead = n.getBooleanValue(); }],
            ["isReadReceiptRequested", (o, n) => { o.isReadReceiptRequested = n.getBooleanValue(); }],
            ["parentFolderId", (o, n) => { o.parentFolderId = n.getStringValue(); }],
            ["receivedDateTime", (o, n) => { o.receivedDateTime = n.getDateValue(); }],
            ["replyTo", (o, n) => { o.replyTo = n.getCollectionOfObjectValues<Recipient>(Recipient); }],
            ["sender", (o, n) => { o.sender = n.getObjectValue<Recipient>(Recipient); }],
            ["sentDateTime", (o, n) => { o.sentDateTime = n.getDateValue(); }],
            ["subject", (o, n) => { o.subject = n.getStringValue(); }],
            ["toRecipients", (o, n) => { o.toRecipients = n.getCollectionOfObjectValues<Recipient>(Recipient); }],
            ["uniqueBody", (o, n) => { o.uniqueBody = n.getObjectValue<ItemBody>(ItemBody); }],
            ["webLink", (o, n) => { o.webLink = n.getStringValue(); }],
            ["attachments", (o, n) => { o.attachments = n.getCollectionOfObjectValues<Attachment>(Attachment); }],
            ["extensions", (o, n) => { o.extensions = n.getCollectionOfObjectValues<Extension>(Extension); }],
            ["multiValueExtendedProperties", (o, n) => { o.multiValueExtendedProperties = n.getCollectionOfObjectValues<MultiValueLegacyExtendedProperty>(MultiValueLegacyExtendedProperty); }],
            ["singleValueExtendedProperties", (o, n) => { o.singleValueExtendedProperties = n.getCollectionOfObjectValues<SingleValueLegacyExtendedProperty>(SingleValueLegacyExtendedProperty); }],
        ]);
    };
}
