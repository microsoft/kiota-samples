import {createMailFolderFromDiscriminatorValue} from './createMailFolderFromDiscriminatorValue';
import {createMessageFromDiscriminatorValue} from './createMessageFromDiscriminatorValue';
import {createMessageRuleFromDiscriminatorValue} from './createMessageRuleFromDiscriminatorValue';
import {createMultiValueLegacyExtendedPropertyFromDiscriminatorValue} from './createMultiValueLegacyExtendedPropertyFromDiscriminatorValue';
import {createSingleValueLegacyExtendedPropertyFromDiscriminatorValue} from './createSingleValueLegacyExtendedPropertyFromDiscriminatorValue';
import {deserializeIntoEntity} from './deserializeIntoEntity';
import type {MailFolder} from './mailFolder';
import type {Message} from './message';
import type {MessageRule} from './messageRule';
import type {MultiValueLegacyExtendedProperty} from './multiValueLegacyExtendedProperty';
import {serializeMailFolder} from './serializeMailFolder';
import {serializeMessage} from './serializeMessage';
import {serializeMessageRule} from './serializeMessageRule';
import {serializeMultiValueLegacyExtendedProperty} from './serializeMultiValueLegacyExtendedProperty';
import {serializeSingleValueLegacyExtendedProperty} from './serializeSingleValueLegacyExtendedProperty';
import type {SingleValueLegacyExtendedProperty} from './singleValueLegacyExtendedProperty';
import type {Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function deserializeIntoMailFolder(mailFolder: MailFolder | undefined = {} as MailFolder) : Record<string, (node: ParseNode) => void> {
    return {
        ...deserializeIntoEntity(mailFolder),
        "childFolderCount": n => { mailFolder.childFolderCount = n.getNumberValue(); },
        "childFolders": n => { mailFolder.childFolders = n.getCollectionOfObjectValues<MailFolder>(createMailFolderFromDiscriminatorValue); },
        "displayName": n => { mailFolder.displayName = n.getStringValue(); },
        "isHidden": n => { mailFolder.isHidden = n.getBooleanValue(); },
        "messageRules": n => { mailFolder.messageRules = n.getCollectionOfObjectValues<MessageRule>(createMessageRuleFromDiscriminatorValue); },
        "messages": n => { mailFolder.messages = n.getCollectionOfObjectValues<Message>(createMessageFromDiscriminatorValue); },
        "multiValueExtendedProperties": n => { mailFolder.multiValueExtendedProperties = n.getCollectionOfObjectValues<MultiValueLegacyExtendedProperty>(createMultiValueLegacyExtendedPropertyFromDiscriminatorValue); },
        "parentFolderId": n => { mailFolder.parentFolderId = n.getStringValue(); },
        "singleValueExtendedProperties": n => { mailFolder.singleValueExtendedProperties = n.getCollectionOfObjectValues<SingleValueLegacyExtendedProperty>(createSingleValueLegacyExtendedPropertyFromDiscriminatorValue); },
        "totalItemCount": n => { mailFolder.totalItemCount = n.getNumberValue(); },
        "unreadItemCount": n => { mailFolder.unreadItemCount = n.getNumberValue(); },
    }
}
