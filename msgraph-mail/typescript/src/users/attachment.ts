import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';
import {Entity} from './entity';

export class Attachment extends Entity implements Parsable {
    /** The MIME type.  */
    public contentType?: string | undefined;
    /** true if the attachment is an inline attachment; otherwise, false.  */
    public isInline?: boolean | undefined;
    /** The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z  */
    public lastModifiedDateTime?: Date | undefined;
    /** The display name of the attachment. This does not need to be the actual file name.  */
    public name?: string | undefined;
    /** The length of the attachment in bytes.  */
    public size?: number | undefined;
    /**
     * The deserialization information for the current model
     * @returns a Map<string, (item: T, node: ParseNode) => void>
     */
    public getFieldDeserializers<T> () : Map<string, (item: T, node: ParseNode) => void> {
        return new Map<string, (item: T, node: ParseNode) => void>([...super.getFieldDeserializers<T>(),
            ["contentType", (o, n) => { (o as unknown as Attachment).contentType = n.getStringValue(); }],
            ["isInline", (o, n) => { (o as unknown as Attachment).isInline = n.getBooleanValue(); }],
            ["lastModifiedDateTime", (o, n) => { (o as unknown as Attachment).lastModifiedDateTime = n.getDateValue(); }],
            ["name", (o, n) => { (o as unknown as Attachment).name = n.getStringValue(); }],
            ["size", (o, n) => { (o as unknown as Attachment).size = n.getNumberValue(); }],
        ]);
    };
    /**
     * Serialiazes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @returns a void
     */
    public serialize (writer: SerializationWriter) : void {
        super.serialize(writer);
        writer.writeStringValue("contentType", this.contentType);
        writer.writeBooleanValue("isInline", this.isInline);
        writer.writeDateValue("lastModifiedDateTime", this.lastModifiedDateTime);
        writer.writeStringValue("name", this.name);
        writer.writeNumberValue("size", this.size);
    };
}
