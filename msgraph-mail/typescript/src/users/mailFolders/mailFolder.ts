import {Entity} from '../entity';
import {Message} from '../message';
import {MultiValueLegacyExtendedProperty} from '../multiValueLegacyExtendedProperty';
import {SingleValueLegacyExtendedProperty} from '../singleValueLegacyExtendedProperty';
import {MessageRule} from './messageRule';

export class MailFolder extends Entity {
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
}
