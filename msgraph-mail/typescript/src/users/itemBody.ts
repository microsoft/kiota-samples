import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';
import {BodyType} from './bodyType';

export class ItemBody implements Parsable<ItemBody> {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.  */
    public readonly additionalData: Map<string, unknown> = new Map<string, unknown>();
    /** The content of the item.  */
    public content?: string | undefined;
    public contentType?: BodyType | undefined;
    /**
     * The serialization information for the current model
     * @returns a Map<string, (item: ItemBody, node: ParseNode) => void>
     */
    public deserializeFields () : Map<string, (item: ItemBody, node: ParseNode) => void> {
        return new Map<string, (item: ItemBody, node: ParseNode) => void>([
            ["content", (o, n) => { o.content = n.getStringValue(); }],
            ["contentType", (o, n) => { o.contentType = n.getEnumValue<BodyType>(BodyType); }],
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
