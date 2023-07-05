import {MailFolder} from './mailFolder';
import {Message} from './message';
import {MessageRule} from './messageRule';
import {MultiValueLegacyExtendedProperty} from './multiValueLegacyExtendedProperty';
import {serializeEntity} from './serializeEntity';
import {serializeMessage} from './serializeMessage';
import {serializeMessageRule} from './serializeMessageRule';
import {serializeMultiValueLegacyExtendedProperty} from './serializeMultiValueLegacyExtendedProperty';
import {serializeSingleValueLegacyExtendedProperty} from './serializeSingleValueLegacyExtendedProperty';
import {SingleValueLegacyExtendedProperty} from './singleValueLegacyExtendedProperty';
import {Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function serializeMailFolder(writer: SerializationWriter, mailFolder: MailFolder | undefined = {} as MailFolder) : void {
        serializeEntity(writer, mailFolder)
        writer.writeNumberValue("childFolderCount", mailFolder.childFolderCount);
        writer.writeCollectionOfObjectValues<MailFolder>("childFolders", mailFolder.childFolders, serializeMailFolder);
        writer.writeStringValue("displayName", mailFolder.displayName);
        writer.writeBooleanValue("isHidden", mailFolder.isHidden);
        writer.writeCollectionOfObjectValues<MessageRule>("messageRules", mailFolder.messageRules, serializeMessageRule);
        writer.writeCollectionOfObjectValues<Message>("messages", mailFolder.messages, serializeMessage);
        writer.writeCollectionOfObjectValues<MultiValueLegacyExtendedProperty>("multiValueExtendedProperties", mailFolder.multiValueExtendedProperties, serializeMultiValueLegacyExtendedProperty);
        writer.writeStringValue("parentFolderId", mailFolder.parentFolderId);
        writer.writeCollectionOfObjectValues<SingleValueLegacyExtendedProperty>("singleValueExtendedProperties", mailFolder.singleValueExtendedProperties, serializeSingleValueLegacyExtendedProperty);
        writer.writeNumberValue("totalItemCount", mailFolder.totalItemCount);
        writer.writeNumberValue("unreadItemCount", mailFolder.unreadItemCount);
}
