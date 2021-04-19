import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';

export class BodyType implements Parsable<BodyType> {
    public serialize (writer: SerializationWriter) : void {
    };
    public deserializeFields () : Map<string, (item: BodyType, node: ParseNode) => void> {
        return new Map<string, (item: BodyType, node: ParseNode) => void>([
        ]);
    };
}
