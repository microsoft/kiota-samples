import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';

export class SizeRange implements Parsable {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.  */
    public readonly additionalData: Map<string, unknown> = new Map<string, unknown>();
    /** The maximum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.  */
    public maximumSize?: number | undefined;
    /** The minimum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.  */
    public minimumSize?: number | undefined;
    /**
     * The deserialization information for the current model
     * @returns a Map<string, (item: T, node: ParseNode) => void>
     */
    public getFieldDeserializers<T> () : Map<string, (item: T, node: ParseNode) => void> {
        return new Map<string, (item: T, node: ParseNode) => void>([
            ["maximumSize", (o, n) => { (o as unknown as SizeRange).maximumSize = n.getNumberValue(); }],
            ["minimumSize", (o, n) => { (o as unknown as SizeRange).minimumSize = n.getNumberValue(); }],
        ]);
    };
    /**
     * Serialiazes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @returns a void
     */
    public serialize (writer: SerializationWriter) : void {
        writer.writeNumberValue("maximumSize", this.maximumSize);
        writer.writeNumberValue("minimumSize", this.minimumSize);
        writer.writeAdditionalData(this.additionalData);
    };
}
