import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';
import {Entity} from './entity';

export class Extension extends Entity implements Parsable<Extension> {
    public serialize (writer: SerializationWriter) : void {
        super.serialize(writer);
    };
    public deserializeFields () : Map<string, (item: Extension, node: ParseNode) => void> {
        return new Map<string, (item: Extension, node: ParseNode) => void>([...super.DeserializeFields(),
        ]);
    };
}
