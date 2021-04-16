import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';
import {Entity} from './entity';

export class Attachment extends Entity implements Parsable<Attachment> {
    /** The MIME type.  */
    public contentType?: string | undefined;
    /** true if the attachment is an inline attachment; otherwise, false.  */
    public isInline?: boolean | undefined;
    /** The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z  */
    public lastModifiedDateTime?: string | undefined;
    /** The display name of the attachment. This does not need to be the actual file name.  */
    public name?: string | undefined;
    /** The length of the attachment in bytes.  */
    public size?: number | undefined;
    /**
     * Serialiazes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @returns a void
     */
    public serialize (writer: SerializationWriter) : void {
        super.serialize(writer);
        writer.writeStringValue("contentType", this.contentType);
        writer.writeBooleanValue("isInline", this.isInline);
        writer.writeStringValue("lastModifiedDateTime", this.lastModifiedDateTime);
        writer.writeStringValue("name", this.name);
        writer.writeNumberValue("size", this.size);
    };
    /**
     * The serialization information for the current model
     * @returns a Map<string, (item: Attachment, node: ParseNode) => void>
     */
    public deserializeFields () : Map<string, (item: Attachment, node: ParseNode) => void> {
        return new Map<string, (item: Attachment, node: ParseNode) => void>([...super.deserializeFields(),
            ["contentType", (o, n) => { o.contentType = n.getStringValue(); }],
            ["isInline", (o, n) => { o.isInline = n.getBooleanValue(); }],
            ["lastModifiedDateTime", (o, n) => { o.lastModifiedDateTime = n.getStringValue(); }],
            ["name", (o, n) => { o.name = n.getStringValue(); }],
            ["size", (o, n) => { o.size = n.getNumberValue(); }],
        ]);
    };
}
