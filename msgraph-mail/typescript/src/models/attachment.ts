import type {Entity} from './entity';
import type {Parsable} from '@microsoft/kiota-abstractions';

export interface Attachment extends Entity, Parsable {
    /**
     * The MIME type.
     */
    contentType?: string | undefined;
    /**
     * true if the attachment is an inline attachment; otherwise, false.
     */
    isInline?: boolean | undefined;
    /**
     * The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
     */
    lastModifiedDateTime?: Date | undefined;
    /**
     * The attachment's file name.
     */
    name?: string | undefined;
    /**
     * The length of the attachment in bytes.
     */
    size?: number | undefined;
}
