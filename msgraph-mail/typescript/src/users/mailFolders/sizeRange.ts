import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';

export class SizeRange implements Parsable<SizeRange> {
    public maximumSize?: number | undefined;
    public minimumSize?: number | undefined;
    public serialize (writer: SerializationWriter) : void {
        writer.writeNumberValue("maximumSize", this.maximumSize);
        writer.writeNumberValue("minimumSize", this.minimumSize);
    };
    public deserializeFields () : Map<string, (item: SizeRange, node: ParseNode) => void> {
        return new Map<string, (item: SizeRange, node: ParseNode) => void>([
            ["maximumSize", (o, n) => { o.maximumSize = n.getNumberValue(); }],
            ["minimumSize", (o, n) => { o.minimumSize = n.getNumberValue(); }],
        ]);
    };
}
