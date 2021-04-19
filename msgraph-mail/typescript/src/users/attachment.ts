import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';
import {Entity} from './entity';

export class Attachment extends Entity implements Parsable<Attachment> {
    public contentType?: string | undefined;
    public isInline?: boolean | undefined;
    public lastModifiedDateTime?: string | undefined;
    public name?: string | undefined;
    public size?: number | undefined;
    public serialize (writer: SerializationWriter) : void {
        super.serialize(writer);
        writer.writeStringValue("contentType", this.contentType);
        writer.writeBooleanValue("isInline", this.isInline);
        writer.writeStringValue("lastModifiedDateTime", this.lastModifiedDateTime);
        writer.writeStringValue("name", this.name);
        writer.writeNumberValue("size", this.size);
    };
    public deserializeFields () : Map<string, (item: Attachment, node: ParseNode) => void> {
        return new Map<string, (item: Attachment, node: ParseNode) => void>([...super.DeserializeFields(),
            ["contentType", (o, n) => { o.contentType = n.getStringValue(); }],
            ["isInline", (o, n) => { o.isInline = n.getBooleanValue(); }],
            ["lastModifiedDateTime", (o, n) => { o.lastModifiedDateTime = n.getStringValue(); }],
            ["name", (o, n) => { o.name = n.getStringValue(); }],
            ["size", (o, n) => { o.size = n.getNumberValue(); }],
        ]);
    };
}
