import { deserializeIntoEntity, serializeEntity, type Entity } from './entity';
import { createMessageFromDiscriminatorValue, serializeMessage, type Message } from './message';
import { createMessageRuleFromDiscriminatorValue, serializeMessageRule, type MessageRule } from './messageRule';
import { createMultiValueLegacyExtendedPropertyFromDiscriminatorValue, serializeMultiValueLegacyExtendedProperty, type MultiValueLegacyExtendedProperty } from './multiValueLegacyExtendedProperty';
import { createSingleValueLegacyExtendedPropertyFromDiscriminatorValue, serializeSingleValueLegacyExtendedProperty, type SingleValueLegacyExtendedProperty } from './singleValueLegacyExtendedProperty';
import { type Parsable, type ParseNode, type SerializationWriter } from '@microsoft/kiota-abstractions';

export function createMailFolderFromDiscriminatorValue(parseNode: ParseNode | undefined) {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return deserializeIntoMailFolder;
}
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
export interface MailFolder extends Entity, Parsable {
    /**
     * The number of immediate child mailFolders in the current mailFolder.
     */
    childFolderCount?: number | undefined;
    /**
     * The collection of child folders in the mailFolder.
     */
    childFolders?: MailFolder[] | undefined;
    /**
     * The mailFolder's display name.
     */
    displayName?: string | undefined;
    /**
     * Indicates whether the mailFolder is hidden. This property can be set only when creating the folder. Find more information in Hidden mail folders.
     */
    isHidden?: boolean | undefined;
    /**
     * The collection of rules that apply to the user's Inbox folder.
     */
    messageRules?: MessageRule[] | undefined;
    /**
     * The collection of messages in the mailFolder.
     */
    messages?: Message[] | undefined;
    /**
     * The collection of multi-value extended properties defined for the mailFolder. Read-only. Nullable.
     */
    multiValueExtendedProperties?: MultiValueLegacyExtendedProperty[] | undefined;
    /**
     * The unique identifier for the mailFolder's parent mailFolder.
     */
    parentFolderId?: string | undefined;
    /**
     * The collection of single-value extended properties defined for the mailFolder. Read-only. Nullable.
     */
    singleValueExtendedProperties?: SingleValueLegacyExtendedProperty[] | undefined;
    /**
     * The number of items in the mailFolder.
     */
    totalItemCount?: number | undefined;
    /**
     * The number of items in the mailFolder marked as unread.
     */
    unreadItemCount?: number | undefined;
}
export function serializeMailFolder(writer: SerializationWriter, mailFolder: MailFolder | undefined = {} as MailFolder) : void {
        serializeEntity(writer, mailFolder)
        writer.writeNumberValue("childFolderCount", mailFolder.childFolderCount);
        writer.writeCollectionOfObjectValues<MailFolder>("childFolders", mailFolder.childFolders, );
        writer.writeStringValue("displayName", mailFolder.displayName);
        writer.writeBooleanValue("isHidden", mailFolder.isHidden);
        writer.writeCollectionOfObjectValues<MessageRule>("messageRules", mailFolder.messageRules, );
        writer.writeCollectionOfObjectValues<Message>("messages", mailFolder.messages, );
        writer.writeCollectionOfObjectValues<MultiValueLegacyExtendedProperty>("multiValueExtendedProperties", mailFolder.multiValueExtendedProperties, );
        writer.writeStringValue("parentFolderId", mailFolder.parentFolderId);
        writer.writeCollectionOfObjectValues<SingleValueLegacyExtendedProperty>("singleValueExtendedProperties", mailFolder.singleValueExtendedProperties, );
        writer.writeNumberValue("totalItemCount", mailFolder.totalItemCount);
        writer.writeNumberValue("unreadItemCount", mailFolder.unreadItemCount);
}
