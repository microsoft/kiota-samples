import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';

export class Importance implements Parsable<Importance> {
    public serialize (writer: SerializationWriter) : void {
    };
    public deserializeFields () : Map<string, (item: Importance, node: ParseNode) => void> {
        return new Map<string, (item: Importance, node: ParseNode) => void>([
        ]);
    };
}
