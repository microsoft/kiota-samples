import { createAttachmentFromDiscriminatorValue, serializeAttachment, type Attachment } from './attachment';
import { createExtensionFromDiscriminatorValue, serializeExtension, type Extension } from './extension';
import { createFollowupFlagFromDiscriminatorValue, serializeFollowupFlag, type FollowupFlag } from './followupFlag';
import { Importance } from './importance';
import { InferenceClassificationType } from './inferenceClassificationType';
import { createInternetMessageHeaderFromDiscriminatorValue, serializeInternetMessageHeader, type InternetMessageHeader } from './internetMessageHeader';
import { createItemBodyFromDiscriminatorValue, serializeItemBody, type ItemBody } from './itemBody';
import { createMultiValueLegacyExtendedPropertyFromDiscriminatorValue, serializeMultiValueLegacyExtendedProperty, type MultiValueLegacyExtendedProperty } from './multiValueLegacyExtendedProperty';
import { deserializeIntoOutlookItem, serializeOutlookItem, type OutlookItem } from './outlookItem';
import { createRecipientFromDiscriminatorValue, serializeRecipient, type Recipient } from './recipient';
import { createSingleValueLegacyExtendedPropertyFromDiscriminatorValue, serializeSingleValueLegacyExtendedProperty, type SingleValueLegacyExtendedProperty } from './singleValueLegacyExtendedProperty';
import { type Parsable, type ParseNode, type SerializationWriter } from '@microsoft/kiota-abstractions';

export function createMessageFromDiscriminatorValue(parseNode: ParseNode | undefined) {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return deserializeIntoMessage;
}
export function deserializeIntoMessage(message: Message | undefined = {} as Message) : Record<string, (node: ParseNode) => void> {
    return {
        ...deserializeIntoOutlookItem(message),
        "attachments": n => { message.attachments = n.getCollectionOfObjectValues<Attachment>(createAttachmentFromDiscriminatorValue); },
        "bccRecipients": n => { message.bccRecipients = n.getCollectionOfObjectValues<Recipient>(createRecipientFromDiscriminatorValue); },
        "body": n => { message.body = n.getObjectValue<ItemBody>(createItemBodyFromDiscriminatorValue); },
        "bodyPreview": n => { message.bodyPreview = n.getStringValue(); },
        "ccRecipients": n => { message.ccRecipients = n.getCollectionOfObjectValues<Recipient>(createRecipientFromDiscriminatorValue); },
        "conversationId": n => { message.conversationId = n.getStringValue(); },
        "conversationIndex": n => { message.conversationIndex = n.getStringValue(); },
        "extensions": n => { message.extensions = n.getCollectionOfObjectValues<Extension>(createExtensionFromDiscriminatorValue); },
        "flag": n => { message.flag = n.getObjectValue<FollowupFlag>(createFollowupFlagFromDiscriminatorValue); },
        "from": n => { message.from = n.getObjectValue<Recipient>(createRecipientFromDiscriminatorValue); },
        "hasAttachments": n => { message.hasAttachments = n.getBooleanValue(); },
        "importance": n => { message.importance = n.getEnumValue<Importance>(Importance); },
        "inferenceClassification": n => { message.inferenceClassification = n.getEnumValue<InferenceClassificationType>(InferenceClassificationType); },
        "internetMessageHeaders": n => { message.internetMessageHeaders = n.getCollectionOfObjectValues<InternetMessageHeader>(createInternetMessageHeaderFromDiscriminatorValue); },
        "internetMessageId": n => { message.internetMessageId = n.getStringValue(); },
        "isDeliveryReceiptRequested": n => { message.isDeliveryReceiptRequested = n.getBooleanValue(); },
        "isDraft": n => { message.isDraft = n.getBooleanValue(); },
        "isRead": n => { message.isRead = n.getBooleanValue(); },
        "isReadReceiptRequested": n => { message.isReadReceiptRequested = n.getBooleanValue(); },
        "multiValueExtendedProperties": n => { message.multiValueExtendedProperties = n.getCollectionOfObjectValues<MultiValueLegacyExtendedProperty>(createMultiValueLegacyExtendedPropertyFromDiscriminatorValue); },
        "parentFolderId": n => { message.parentFolderId = n.getStringValue(); },
        "receivedDateTime": n => { message.receivedDateTime = n.getDateValue(); },
        "replyTo": n => { message.replyTo = n.getCollectionOfObjectValues<Recipient>(createRecipientFromDiscriminatorValue); },
        "sender": n => { message.sender = n.getObjectValue<Recipient>(createRecipientFromDiscriminatorValue); },
        "sentDateTime": n => { message.sentDateTime = n.getDateValue(); },
        "singleValueExtendedProperties": n => { message.singleValueExtendedProperties = n.getCollectionOfObjectValues<SingleValueLegacyExtendedProperty>(createSingleValueLegacyExtendedPropertyFromDiscriminatorValue); },
        "subject": n => { message.subject = n.getStringValue(); },
        "toRecipients": n => { message.toRecipients = n.getCollectionOfObjectValues<Recipient>(createRecipientFromDiscriminatorValue); },
        "uniqueBody": n => { message.uniqueBody = n.getObjectValue<ItemBody>(createItemBodyFromDiscriminatorValue); },
        "webLink": n => { message.webLink = n.getStringValue(); },
    }
}
export interface Message extends OutlookItem, Parsable {
    /**
     * The fileAttachment and itemAttachment attachments for the message.
     */
    attachments?: Attachment[] | undefined;
    /**
     * The Bcc: recipients for the message.
     */
    bccRecipients?: Recipient[] | undefined;
    /**
     * The body property
     */
    body?: ItemBody | undefined;
    /**
     * The first 255 characters of the message body. It is in text format.
     */
    bodyPreview?: string | undefined;
    /**
     * The Cc: recipients for the message.
     */
    ccRecipients?: Recipient[] | undefined;
    /**
     * The ID of the conversation the email belongs to.
     */
    conversationId?: string | undefined;
    /**
     * Indicates the position of the message within the conversation.
     */
    conversationIndex?: string | undefined;
    /**
     * The collection of open extensions defined for the message. Nullable.
     */
    extensions?: Extension[] | undefined;
    /**
     * The flag property
     */
    flag?: FollowupFlag | undefined;
    /**
     * The from property
     */
    from?: Recipient | undefined;
    /**
     * Indicates whether the message has attachments. This property doesn't include inline attachments, so if a message contains only inline attachments, this property is false. To verify the existence of inline attachments, parse the body property to look for a src attribute, such as <IMG src='cid:image001.jpg@01D26CD8.6C05F070'>.
     */
    hasAttachments?: boolean | undefined;
    /**
     * The importance property
     */
    importance?: Importance | undefined;
    /**
     * The inferenceClassification property
     */
    inferenceClassification?: InferenceClassificationType | undefined;
    /**
     * The internetMessageHeaders property
     */
    internetMessageHeaders?: InternetMessageHeader[] | undefined;
    /**
     * The internetMessageId property
     */
    internetMessageId?: string | undefined;
    /**
     * The isDeliveryReceiptRequested property
     */
    isDeliveryReceiptRequested?: boolean | undefined;
    /**
     * The isDraft property
     */
    isDraft?: boolean | undefined;
    /**
     * The isRead property
     */
    isRead?: boolean | undefined;
    /**
     * The isReadReceiptRequested property
     */
    isReadReceiptRequested?: boolean | undefined;
    /**
     * The collection of multi-value extended properties defined for the message. Nullable.
     */
    multiValueExtendedProperties?: MultiValueLegacyExtendedProperty[] | undefined;
    /**
     * The parentFolderId property
     */
    parentFolderId?: string | undefined;
    /**
     * The receivedDateTime property
     */
    receivedDateTime?: Date | undefined;
    /**
     * The replyTo property
     */
    replyTo?: Recipient[] | undefined;
    /**
     * The sender property
     */
    sender?: Recipient | undefined;
    /**
     * The sentDateTime property
     */
    sentDateTime?: Date | undefined;
    /**
     * The collection of single-value extended properties defined for the message. Nullable.
     */
    singleValueExtendedProperties?: SingleValueLegacyExtendedProperty[] | undefined;
    /**
     * The subject property
     */
    subject?: string | undefined;
    /**
     * The toRecipients property
     */
    toRecipients?: Recipient[] | undefined;
    /**
     * The uniqueBody property
     */
    uniqueBody?: ItemBody | undefined;
    /**
     * The webLink property
     */
    webLink?: string | undefined;
}
export function serializeMessage(writer: SerializationWriter, message: Message | undefined = {} as Message) : void {
        serializeOutlookItem(writer, message)
        writer.writeCollectionOfObjectValues<Attachment>("attachments", message.attachments, );
        writer.writeCollectionOfObjectValues<Recipient>("bccRecipients", message.bccRecipients, );
        writer.writeObjectValue<ItemBody>("body", message.body, );
        writer.writeStringValue("bodyPreview", message.bodyPreview);
        writer.writeCollectionOfObjectValues<Recipient>("ccRecipients", message.ccRecipients, );
        writer.writeStringValue("conversationId", message.conversationId);
        writer.writeStringValue("conversationIndex", message.conversationIndex);
        writer.writeCollectionOfObjectValues<Extension>("extensions", message.extensions, );
        writer.writeObjectValue<FollowupFlag>("flag", message.flag, );
        writer.writeObjectValue<Recipient>("from", message.from, );
        writer.writeBooleanValue("hasAttachments", message.hasAttachments);
        writer.writeEnumValue<Importance>("importance", message.importance);
        writer.writeEnumValue<InferenceClassificationType>("inferenceClassification", message.inferenceClassification);
        writer.writeCollectionOfObjectValues<InternetMessageHeader>("internetMessageHeaders", message.internetMessageHeaders, );
        writer.writeStringValue("internetMessageId", message.internetMessageId);
        writer.writeBooleanValue("isDeliveryReceiptRequested", message.isDeliveryReceiptRequested);
        writer.writeBooleanValue("isDraft", message.isDraft);
        writer.writeBooleanValue("isRead", message.isRead);
        writer.writeBooleanValue("isReadReceiptRequested", message.isReadReceiptRequested);
        writer.writeCollectionOfObjectValues<MultiValueLegacyExtendedProperty>("multiValueExtendedProperties", message.multiValueExtendedProperties, );
        writer.writeStringValue("parentFolderId", message.parentFolderId);
        writer.writeDateValue("receivedDateTime", message.receivedDateTime);
        writer.writeCollectionOfObjectValues<Recipient>("replyTo", message.replyTo, );
        writer.writeObjectValue<Recipient>("sender", message.sender, );
        writer.writeDateValue("sentDateTime", message.sentDateTime);
        writer.writeCollectionOfObjectValues<SingleValueLegacyExtendedProperty>("singleValueExtendedProperties", message.singleValueExtendedProperties, );
        writer.writeStringValue("subject", message.subject);
        writer.writeCollectionOfObjectValues<Recipient>("toRecipients", message.toRecipients, );
        writer.writeObjectValue<ItemBody>("uniqueBody", message.uniqueBody, );
        writer.writeStringValue("webLink", message.webLink);
}
