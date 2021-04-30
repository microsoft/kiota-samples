import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';

export class EmailAddress implements Parsable<EmailAddress> {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.  */
    public readonly additionalData: Map<string, unknown> = new Map<string, unknown>();
    /** The email address of the person or entity.  */
    public address?: string | undefined;
    /** The display name of the person or entity.  */
    public name?: string | undefined;
    /**
     * The serialization information for the current model
     * @returns a Map<string, (item: EmailAddress, node: ParseNode) => void>
     */
    public deserializeFields () : Map<string, (item: EmailAddress, node: ParseNode) => void> {
        return new Map<string, (item: EmailAddress, node: ParseNode) => void>([
            ["address", (o, n) => { o.address = n.getStringValue(); }],
            ["name", (o, n) => { o.name = n.getStringValue(); }],
        ]);
    };
    /**
     * Serialiazes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @returns a void
     */
    public serialize (writer: SerializationWriter) : void {
        writer.writeStringValue("address", this.address);
        writer.writeStringValue("name", this.name);
        writer.writeAdditionalData(this.additionalData);
    };
}
