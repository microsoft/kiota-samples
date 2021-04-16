import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';

export class InternetMessageHeader implements Parsable<InternetMessageHeader> {
    public name?: string | undefined;
    public value?: string | undefined;
    public serialize (writer: SerializationWriter) : void {
        writer.writeStringValue("name", this.name);
        writer.writeStringValue("value", this.value);
    };
    public deserializeFields () : Map<string, (item: InternetMessageHeader, node: ParseNode) => void> {
        return new Map<string, (item: InternetMessageHeader, node: ParseNode) => void>([
            ["name", (o, n) => { o.name = n.getStringValue(); }],
            ["value", (o, n) => { o.value = n.getStringValue(); }],
        ]);
    };
}
