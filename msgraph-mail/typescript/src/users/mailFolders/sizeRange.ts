import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';

export class SizeRange implements Parsable<SizeRange> {
    /** The maximum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.  */
    public maximumSize?: number | undefined;
    /** The minimum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.  */
    public minimumSize?: number | undefined;
    /**
     * Serialiazes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @returns a void
     */
    public serialize (writer: SerializationWriter) : void {
        writer.writeNumberValue("maximumSize", this.maximumSize);
        writer.writeNumberValue("minimumSize", this.minimumSize);
    };
    /**
     * The serialization information for the current model
     * @returns a Map<string, (item: SizeRange, node: ParseNode) => void>
     */
    public deserializeFields () : Map<string, (item: SizeRange, node: ParseNode) => void> {
        return new Map<string, (item: SizeRange, node: ParseNode) => void>([
            ["maximumSize", (o, n) => { o.maximumSize = n.getNumberValue(); }],
            ["minimumSize", (o, n) => { o.minimumSize = n.getNumberValue(); }],
        ]);
    };
}
