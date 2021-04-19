import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';

export class Entity implements Parsable<Entity> {
    public id?: string | undefined;
    public serialize (writer: SerializationWriter) : void {
        writer.writeStringValue("id", this.id);
    };
    public deserializeFields () : Map<string, (item: Entity, node: ParseNode) => void> {
        return new Map<string, (item: Entity, node: ParseNode) => void>([
            ["id", (o, n) => { o.id = n.getStringValue(); }],
        ]);
    };
}
