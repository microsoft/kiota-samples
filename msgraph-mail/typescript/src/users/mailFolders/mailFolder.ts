import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';
import {Entity} from '../entity';
import {Message} from '../message';
import {MultiValueLegacyExtendedProperty} from '../multiValueLegacyExtendedProperty';
import {SingleValueLegacyExtendedProperty} from '../singleValueLegacyExtendedProperty';
import {MessageRule} from './messageRule';

export class MailFolder extends Entity implements Parsable<MailFolder> {
    public childFolderCount?: number | undefined;
    public displayName?: string | undefined;
    public parentFolderId?: string | undefined;
    public totalItemCount?: number | undefined;
    public unreadItemCount?: number | undefined;
    public childFolders?: MailFolder[] | undefined;
    public messageRules?: MessageRule[] | undefined;
    public messages?: Message[] | undefined;
    public multiValueExtendedProperties?: MultiValueLegacyExtendedProperty[] | undefined;
    public singleValueExtendedProperties?: SingleValueLegacyExtendedProperty[] | undefined;
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
