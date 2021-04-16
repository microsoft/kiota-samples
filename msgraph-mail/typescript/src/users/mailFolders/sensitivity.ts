import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';

export class Sensitivity implements Parsable<Sensitivity> {
    public serialize (writer: SerializationWriter) : void {
    };
    public deserializeFields () : Map<string, (item: Sensitivity, node: ParseNode) => void> {
        return new Map<string, (item: Sensitivity, node: ParseNode) => void>([
        ]);
    };
}
