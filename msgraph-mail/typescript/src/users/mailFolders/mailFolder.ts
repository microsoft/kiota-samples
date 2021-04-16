import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';
import {Entity} from '../entity';
import {Message} from '../message';
import {MultiValueLegacyExtendedProperty} from '../multiValueLegacyExtendedProperty';
import {SingleValueLegacyExtendedProperty} from '../singleValueLegacyExtendedProperty';
import {MessageRule} from './messageRule';

export class MailFolder extends Entity implements Parsable<MailFolder> {
    /** The number of immediate child mailFolders in the current mailFolder.  */
    public childFolderCount?: number | undefined;
    /** The mailFolder's display name.  */
    public displayName?: string | undefined;
    /** The unique identifier for the mailFolder's parent mailFolder.  */
    public parentFolderId?: string | undefined;
    /** The number of items in the mailFolder.  */
    public totalItemCount?: number | undefined;
    /** The number of items in the mailFolder marked as unread.  */
    public unreadItemCount?: number | undefined;
    /** The collection of child folders in the mailFolder.  */
    public childFolders?: MailFolder[] | undefined;
    /** The collection of rules that apply to the user's Inbox folder.  */
    public messageRules?: MessageRule[] | undefined;
    /** The collection of messages in the mailFolder.  */
    public messages?: Message[] | undefined;
    /** The collection of multi-value extended properties defined for the mailFolder. Read-only. Nullable.  */
    public multiValueExtendedProperties?: MultiValueLegacyExtendedProperty[] | undefined;
    /** The collection of single-value extended properties defined for the mailFolder. Read-only. Nullable.  */
    public singleValueExtendedProperties?: SingleValueLegacyExtendedProperty[] | undefined;
    /**
     * Serialiazes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @returns a void
     */
    public serialize (writer: SerializationWriter) : void {
        super.serialize(writer);
        writer.writeNumberValue("childFolderCount", this.childFolderCount);
        writer.writeStringValue("displayName", this.displayName);
        writer.writeStringValue("parentFolderId", this.parentFolderId);
        writer.writeNumberValue("totalItemCount", this.totalItemCount);
        writer.writeNumberValue("unreadItemCount", this.unreadItemCount);
        writer.writeCollectionOfObjectValues<MailFolder>("childFolders", this.childFolders);
        writer.writeCollectionOfObjectValues<MessageRule>("messageRules", this.messageRules);
        writer.writeCollectionOfObjectValues<Message>("messages", this.messages);
        writer.writeCollectionOfObjectValues<MultiValueLegacyExtendedProperty>("multiValueExtendedProperties", this.multiValueExtendedProperties);
        writer.writeCollectionOfObjectValues<SingleValueLegacyExtendedProperty>("singleValueExtendedProperties", this.singleValueExtendedProperties);
    };
    /**
     * The serialization information for the current model
     * @returns a Map<string, (item: MailFolder, node: ParseNode) => void>
     */
    public deserializeFields () : Map<string, (item: MailFolder, node: ParseNode) => void> {
        return new Map<string, (item: MailFolder, node: ParseNode) => void>([...super.deserializeFields(),
            ["childFolderCount", (o, n) => { o.childFolderCount = n.getNumberValue(); }],
            ["displayName", (o, n) => { o.displayName = n.getStringValue(); }],
            ["parentFolderId", (o, n) => { o.parentFolderId = n.getStringValue(); }],
            ["totalItemCount", (o, n) => { o.totalItemCount = n.getNumberValue(); }],
            ["unreadItemCount", (o, n) => { o.unreadItemCount = n.getNumberValue(); }],
            ["childFolders", (o, n) => { o.childFolders = n.getCollectionOfObjectValues<MailFolder>(MailFolder); }],
            ["messageRules", (o, n) => { o.messageRules = n.getCollectionOfObjectValues<MessageRule>(MessageRule); }],
            ["messages", (o, n) => { o.messages = n.getCollectionOfObjectValues<Message>(Message); }],
            ["multiValueExtendedProperties", (o, n) => { o.multiValueExtendedProperties = n.getCollectionOfObjectValues<MultiValueLegacyExtendedProperty>(MultiValueLegacyExtendedProperty); }],
            ["singleValueExtendedProperties", (o, n) => { o.singleValueExtendedProperties = n.getCollectionOfObjectValues<SingleValueLegacyExtendedProperty>(SingleValueLegacyExtendedProperty); }],
        ]);
    };
}
