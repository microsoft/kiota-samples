import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';
import {BodyType} from './bodyType';

export class ItemBody implements Parsable<ItemBody> {
    public content?: string | undefined;
    public contentType?: BodyType | undefined;
    public serialize (writer: SerializationWriter) : void {
        writer.writeStringValue("content", this.content);
        writer.writeEnumValue<BodyType>("contentType", this.contentType);
    };
    public deserializeFields () : Map<string, (item: ItemBody, node: ParseNode) => void> {
        return new Map<string, (item: ItemBody, node: ParseNode) => void>([
            ["content", (o, n) => { o.content = n.getStringValue(); }],
            ["contentType", (o, n) => { o.contentType = n.getEnumValue<BodyType>(BodyType); }],
        ]);
    };
}
