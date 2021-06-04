import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';
import {Entity} from './entity';
import {Message} from './message';
import {MessageRule} from './messageRule';
import {MultiValueLegacyExtendedProperty} from './multiValueLegacyExtendedProperty';
import {SingleValueLegacyExtendedProperty} from './singleValueLegacyExtendedProperty';

export class MailFolder extends Entity implements Parsable {
    /** The number of immediate child mailFolders in the current mailFolder.  */
    public childFolderCount?: number | undefined;
    /** The collection of child folders in the mailFolder.  */
    public childFolders?: MailFolder[] | undefined;
    /** The mailFolder's display name.  */
    public displayName?: string | undefined;
    /** Indicates whether the mailFolder is hidden. This property can be set only when creating the folder. Find more information in Hidden mail folders.  */
    public isHidden?: boolean | undefined;
    /** The collection of rules that apply to the user's Inbox folder.  */
    public messageRules?: MessageRule[] | undefined;
    /** The collection of messages in the mailFolder.  */
    public messages?: Message[] | undefined;
    /** The collection of multi-value extended properties defined for the mailFolder. Read-only. Nullable.  */
    public multiValueExtendedProperties?: MultiValueLegacyExtendedProperty[] | undefined;
    /** The unique identifier for the mailFolder's parent mailFolder.  */
    public parentFolderId?: string | undefined;
    /** The collection of single-value extended properties defined for the mailFolder. Read-only. Nullable.  */
    public singleValueExtendedProperties?: SingleValueLegacyExtendedProperty[] | undefined;
    /** The number of items in the mailFolder.  */
    public totalItemCount?: number | undefined;
    /** The number of items in the mailFolder marked as unread.  */
    public unreadItemCount?: number | undefined;
    /**
     * The deserialization information for the current model
     * @returns a Map<string, (item: T, node: ParseNode) => void>
     */
    public getFieldDeserializers<T> () : Map<string, (item: T, node: ParseNode) => void> {
        return new Map<string, (item: T, node: ParseNode) => void>([...super.getFieldDeserializers<T>(),
            ["childFolderCount", (o, n) => { (o as unknown as MailFolder).childFolderCount = n.getNumberValue(); }],
            ["childFolders", (o, n) => { (o as unknown as MailFolder).childFolders = n.getCollectionOfObjectValues<MailFolder>(MailFolder); }],
            ["displayName", (o, n) => { (o as unknown as MailFolder).displayName = n.getStringValue(); }],
            ["isHidden", (o, n) => { (o as unknown as MailFolder).isHidden = n.getBooleanValue(); }],
            ["messageRules", (o, n) => { (o as unknown as MailFolder).messageRules = n.getCollectionOfObjectValues<MessageRule>(MessageRule); }],
            ["messages", (o, n) => { (o as unknown as MailFolder).messages = n.getCollectionOfObjectValues<Message>(Message); }],
            ["multiValueExtendedProperties", (o, n) => { (o as unknown as MailFolder).multiValueExtendedProperties = n.getCollectionOfObjectValues<MultiValueLegacyExtendedProperty>(MultiValueLegacyExtendedProperty); }],
            ["parentFolderId", (o, n) => { (o as unknown as MailFolder).parentFolderId = n.getStringValue(); }],
            ["singleValueExtendedProperties", (o, n) => { (o as unknown as MailFolder).singleValueExtendedProperties = n.getCollectionOfObjectValues<SingleValueLegacyExtendedProperty>(SingleValueLegacyExtendedProperty); }],
            ["totalItemCount", (o, n) => { (o as unknown as MailFolder).totalItemCount = n.getNumberValue(); }],
            ["unreadItemCount", (o, n) => { (o as unknown as MailFolder).unreadItemCount = n.getNumberValue(); }],
        ]);
    };
    /**
     * Serialiazes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @returns a void
     */
    public serialize (writer: SerializationWriter) : void {
        super.serialize(writer);
        writer.writeNumberValue("childFolderCount", this.childFolderCount);
        writer.writeCollectionOfObjectValues<MailFolder>("childFolders", this.childFolders);
        writer.writeStringValue("displayName", this.displayName);
        writer.writeBooleanValue("isHidden", this.isHidden);
        writer.writeCollectionOfObjectValues<MessageRule>("messageRules", this.messageRules);
        writer.writeCollectionOfObjectValues<Message>("messages", this.messages);
        writer.writeCollectionOfObjectValues<MultiValueLegacyExtendedProperty>("multiValueExtendedProperties", this.multiValueExtendedProperties);
        writer.writeStringValue("parentFolderId", this.parentFolderId);
        writer.writeCollectionOfObjectValues<SingleValueLegacyExtendedProperty>("singleValueExtendedProperties", this.singleValueExtendedProperties);
        writer.writeNumberValue("totalItemCount", this.totalItemCount);
        writer.writeNumberValue("unreadItemCount", this.unreadItemCount);
    };
}
