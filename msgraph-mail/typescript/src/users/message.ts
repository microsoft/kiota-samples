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
    public bccRecipients?: Recipient[] | undefined;
    public body?: ItemBody | undefined;
    public bodyPreview?: string | undefined;
    public ccRecipients?: Recipient[] | undefined;
    public conversationId?: string | undefined;
    public conversationIndex?: string | undefined;
    public flag?: FollowupFlag | undefined;
    public from?: Recipient | undefined;
    public hasAttachments?: boolean | undefined;
    public importance?: Importance | undefined;
    public inferenceClassification?: InferenceClassificationType | undefined;
    public internetMessageHeaders?: InternetMessageHeader[] | undefined;
    public internetMessageId?: string | undefined;
    public isDeliveryReceiptRequested?: boolean | undefined;
    public isDraft?: boolean | undefined;
    public isRead?: boolean | undefined;
    public isReadReceiptRequested?: boolean | undefined;
    public parentFolderId?: string | undefined;
    public receivedDateTime?: string | undefined;
    public replyTo?: Recipient[] | undefined;
    public sender?: Recipient | undefined;
    public sentDateTime?: string | undefined;
    public subject?: string | undefined;
    public toRecipients?: Recipient[] | undefined;
    public uniqueBody?: ItemBody | undefined;
    public webLink?: string | undefined;
    public attachments?: Attachment[] | undefined;
    public extensions?: Extension[] | undefined;
    public multiValueExtendedProperties?: MultiValueLegacyExtendedProperty[] | undefined;
    public singleValueExtendedProperties?: SingleValueLegacyExtendedProperty[] | undefined;
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
        writer.writeObjectValue<Importance>("importance", this.importance);
        writer.writeObjectValue<InferenceClassificationType>("inferenceClassification", this.inferenceClassification);
        writer.writeCollectionOfObjectValues<InternetMessageHeader>("internetMessageHeaders", this.internetMessageHeaders);
        writer.writeStringValue("internetMessageId", this.internetMessageId);
        writer.writeBooleanValue("isDeliveryReceiptRequested", this.isDeliveryReceiptRequested);
        writer.writeBooleanValue("isDraft", this.isDraft);
        writer.writeBooleanValue("isRead", this.isRead);
        writer.writeBooleanValue("isReadReceiptRequested", this.isReadReceiptRequested);
        writer.writeStringValue("parentFolderId", this.parentFolderId);
        writer.writeStringValue("receivedDateTime", this.receivedDateTime);
        writer.writeCollectionOfObjectValues<Recipient>("replyTo", this.replyTo);
        writer.writeObjectValue<Recipient>("sender", this.sender);
        writer.writeStringValue("sentDateTime", this.sentDateTime);
        writer.writeStringValue("subject", this.subject);
        writer.writeCollectionOfObjectValues<Recipient>("toRecipients", this.toRecipients);
        writer.writeObjectValue<ItemBody>("uniqueBody", this.uniqueBody);
        writer.writeStringValue("webLink", this.webLink);
        writer.writeCollectionOfObjectValues<Attachment>("attachments", this.attachments);
        writer.writeCollectionOfObjectValues<Extension>("extensions", this.extensions);
        writer.writeCollectionOfObjectValues<MultiValueLegacyExtendedProperty>("multiValueExtendedProperties", this.multiValueExtendedProperties);
        writer.writeCollectionOfObjectValues<SingleValueLegacyExtendedProperty>("singleValueExtendedProperties", this.singleValueExtendedProperties);
    };
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
            ["importance", (o, n) => { o.importance = n.getObjectValue<Importance>(Importance); }],
            ["inferenceClassification", (o, n) => { o.inferenceClassification = n.getObjectValue<InferenceClassificationType>(InferenceClassificationType); }],
            ["internetMessageHeaders", (o, n) => { o.internetMessageHeaders = n.getCollectionOfObjectValues<InternetMessageHeader>(InternetMessageHeader); }],
            ["internetMessageId", (o, n) => { o.internetMessageId = n.getStringValue(); }],
            ["isDeliveryReceiptRequested", (o, n) => { o.isDeliveryReceiptRequested = n.getBooleanValue(); }],
            ["isDraft", (o, n) => { o.isDraft = n.getBooleanValue(); }],
            ["isRead", (o, n) => { o.isRead = n.getBooleanValue(); }],
            ["isReadReceiptRequested", (o, n) => { o.isReadReceiptRequested = n.getBooleanValue(); }],
            ["parentFolderId", (o, n) => { o.parentFolderId = n.getStringValue(); }],
            ["receivedDateTime", (o, n) => { o.receivedDateTime = n.getStringValue(); }],
            ["replyTo", (o, n) => { o.replyTo = n.getCollectionOfObjectValues<Recipient>(Recipient); }],
            ["sender", (o, n) => { o.sender = n.getObjectValue<Recipient>(Recipient); }],
            ["sentDateTime", (o, n) => { o.sentDateTime = n.getStringValue(); }],
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
