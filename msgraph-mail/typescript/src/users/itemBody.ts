import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';
import {BodyType} from './bodyType';

export class ItemBody implements Parsable {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.  */
    public readonly additionalData: Map<string, unknown> = new Map<string, unknown>();
    /** The content of the item.  */
    public content?: string | undefined;
    public contentType?: BodyType | undefined;
    /**
     * The deserialization information for the current model
     * @returns a Map<string, (item: T, node: ParseNode) => void>
     */
    public getFieldDeserializers<T> () : Map<string, (item: T, node: ParseNode) => void> {
        return new Map<string, (item: T, node: ParseNode) => void>([
            ["content", (o, n) => { (o as unknown as ItemBody).content = n.getStringValue(); }],
            ["contentType", (o, n) => { (o as unknown as ItemBody).contentType = n.getEnumValue<BodyType>(BodyType); }],
        ]);
    };
    /**
     * Serialiazes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @returns a void
     */
    public serialize (writer: SerializationWriter) : void {
        writer.writeStringValue("content", this.content);
        writer.writeEnumValue<BodyType>("contentType", this.contentType);
        writer.writeAdditionalData(this.additionalData);
    };
}
