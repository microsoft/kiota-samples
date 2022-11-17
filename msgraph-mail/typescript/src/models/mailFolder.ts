import {Entity} from './entity';
import {Message} from './message';
import {MessageRule} from './messageRule';
import {MultiValueLegacyExtendedProperty} from './multiValueLegacyExtendedProperty';
import {SingleValueLegacyExtendedProperty} from './singleValueLegacyExtendedProperty';
import {Parsable} from '@microsoft/kiota-abstractions';

export interface MailFolder extends Entity, Partial<Parsable> {
    /** The number of immediate child mailFolders in the current mailFolder. */
    childFolderCount?: number | undefined;
    /** The collection of child folders in the mailFolder. */
    childFolders?: MailFolder[] | undefined;
    /** The mailFolder's display name. */
    displayName?: string | undefined;
    /** Indicates whether the mailFolder is hidden. This property can be set only when creating the folder. Find more information in Hidden mail folders. */
    isHidden?: boolean | undefined;
    /** The collection of rules that apply to the user's Inbox folder. */
    messageRules?: MessageRule[] | undefined;
    /** The collection of messages in the mailFolder. */
    messages?: Message[] | undefined;
    /** The collection of multi-value extended properties defined for the mailFolder. Read-only. Nullable. */
    multiValueExtendedProperties?: MultiValueLegacyExtendedProperty[] | undefined;
    /** The unique identifier for the mailFolder's parent mailFolder. */
    parentFolderId?: string | undefined;
    /** The collection of single-value extended properties defined for the mailFolder. Read-only. Nullable. */
    singleValueExtendedProperties?: SingleValueLegacyExtendedProperty[] | undefined;
    /** The number of items in the mailFolder. */
    totalItemCount?: number | undefined;
    /** The number of items in the mailFolder marked as unread. */
    unreadItemCount?: number | undefined;
}
