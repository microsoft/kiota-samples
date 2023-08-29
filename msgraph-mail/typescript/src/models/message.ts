import type {Attachment} from './attachment';
import type {Extension} from './extension';
import type {FollowupFlag} from './followupFlag';
import {Importance} from './importance';
import {InferenceClassificationType} from './inferenceClassificationType';
import type {InternetMessageHeader} from './internetMessageHeader';
import type {ItemBody} from './itemBody';
import type {MultiValueLegacyExtendedProperty} from './multiValueLegacyExtendedProperty';
import type {OutlookItem} from './outlookItem';
import type {Recipient} from './recipient';
import type {SingleValueLegacyExtendedProperty} from './singleValueLegacyExtendedProperty';
import type {Parsable} from '@microsoft/kiota-abstractions';

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
