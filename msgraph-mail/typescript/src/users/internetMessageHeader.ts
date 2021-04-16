import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';

export class InternetMessageHeader implements Parsable<InternetMessageHeader> {
    /** Represents the key in a key-value pair.  */
    public name?: string | undefined;
    /** The value in a key-value pair.  */
    public value?: string | undefined;
    /**
     * Serialiazes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @returns a void
     */
    public serialize (writer: SerializationWriter) : void {
        writer.writeStringValue("name", this.name);
        writer.writeStringValue("value", this.value);
    };
    /**
     * The serialization information for the current model
     * @returns a Map<string, (item: InternetMessageHeader, node: ParseNode) => void>
     */
    public deserializeFields () : Map<string, (item: InternetMessageHeader, node: ParseNode) => void> {
        return new Map<string, (item: InternetMessageHeader, node: ParseNode) => void>([
            ["name", (o, n) => { o.name = n.getStringValue(); }],
            ["value", (o, n) => { o.value = n.getStringValue(); }],
        ]);
    };
}
