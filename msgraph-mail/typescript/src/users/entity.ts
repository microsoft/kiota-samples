import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';

export class Entity implements Parsable<Entity> {
    /** Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.  */
    public readonly additionalData: Map<string, unknown> = new Map<string, unknown>();
    /** Read-only.  */
    public id?: string | undefined;
    /**
     * The deserialization information for the current model
     * @returns a Map<string, (item: Entity, node: ParseNode) => void>
     */
    public getFieldDeserializers () : Map<string, (item: Entity, node: ParseNode) => void> {
        return new Map<string, (item: Entity, node: ParseNode) => void>([
            ["id", (o, n) => { o.id = n.getStringValue(); }],
        ]);
    };
    /**
     * Serialiazes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @returns a void
     */
    public serialize (writer: SerializationWriter) : void {
        writer.writeStringValue("id", this.id);
        writer.writeAdditionalData(this.additionalData);
    };
}
