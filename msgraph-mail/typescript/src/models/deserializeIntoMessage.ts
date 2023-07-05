import {Attachment} from './attachment';
import {createAttachmentFromDiscriminatorValue} from './createAttachmentFromDiscriminatorValue';
import {createExtensionFromDiscriminatorValue} from './createExtensionFromDiscriminatorValue';
import {createFollowupFlagFromDiscriminatorValue} from './createFollowupFlagFromDiscriminatorValue';
import {createInternetMessageHeaderFromDiscriminatorValue} from './createInternetMessageHeaderFromDiscriminatorValue';
import {createItemBodyFromDiscriminatorValue} from './createItemBodyFromDiscriminatorValue';
import {createMultiValueLegacyExtendedPropertyFromDiscriminatorValue} from './createMultiValueLegacyExtendedPropertyFromDiscriminatorValue';
import {createRecipientFromDiscriminatorValue} from './createRecipientFromDiscriminatorValue';
import {createSingleValueLegacyExtendedPropertyFromDiscriminatorValue} from './createSingleValueLegacyExtendedPropertyFromDiscriminatorValue';
import {deserializeIntoOutlookItem} from './deserializeIntoOutlookItem';
import {Extension} from './extension';
import {FollowupFlag} from './followupFlag';
import {Importance} from './importance';
import {InferenceClassificationType} from './inferenceClassificationType';
import {InternetMessageHeader} from './internetMessageHeader';
import {ItemBody} from './itemBody';
import {Message} from './message';
import {MultiValueLegacyExtendedProperty} from './multiValueLegacyExtendedProperty';
import {Recipient} from './recipient';
import {serializeAttachment} from './serializeAttachment';
import {serializeExtension} from './serializeExtension';
import {serializeFollowupFlag} from './serializeFollowupFlag';
import {serializeInternetMessageHeader} from './serializeInternetMessageHeader';
import {serializeItemBody} from './serializeItemBody';
import {serializeMultiValueLegacyExtendedProperty} from './serializeMultiValueLegacyExtendedProperty';
import {serializeRecipient} from './serializeRecipient';
import {serializeSingleValueLegacyExtendedProperty} from './serializeSingleValueLegacyExtendedProperty';
import {SingleValueLegacyExtendedProperty} from './singleValueLegacyExtendedProperty';
import {Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

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
