import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';

export class EmailAddress implements Parsable<EmailAddress> {
    public address?: string | undefined;
    public name?: string | undefined;
    public serialize (writer: SerializationWriter) : void {
        writer.writeStringValue("address", this.address);
        writer.writeStringValue("name", this.name);
    };
    public deserializeFields () : Map<string, (item: EmailAddress, node: ParseNode) => void> {
        return new Map<string, (item: EmailAddress, node: ParseNode) => void>([
            ["address", (o, n) => { o.address = n.getStringValue(); }],
            ["name", (o, n) => { o.name = n.getStringValue(); }],
        ]);
    };
}
