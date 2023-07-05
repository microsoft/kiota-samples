import {Entity} from './entity';
import {Parsable} from '@microsoft/kiota-abstractions';

export interface OutlookItem extends Entity, Parsable {
    /**
     * The categories associated with the item
     */
    categories?: string[] | undefined;
    /**
     * Identifies the version of the item. Every time the item is changed, changeKey changes as well. This allows Exchange to apply changes to the correct version of the object. Read-only.
     */
    changeKey?: string | undefined;
    /**
     * The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
     */
    createdDateTime?: Date | undefined;
    /**
     * The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
     */
    lastModifiedDateTime?: Date | undefined;
}
