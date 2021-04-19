import {SerializationWriter, ParseNode, Parsable} from '@microsoft/kiota-abstractions';
import {Entity} from './entity';

export class Extension extends Entity implements Parsable<Extension> {
    /**
     * Serialiazes information the current object
     * @param writer Serialization writer to use to serialize this model
     * @returns a void
     */
    public serialize (writer: SerializationWriter) : void {
        super.serialize(writer);
    };
    /**
     * The serialization information for the current model
     * @returns a Map<string, (item: Extension, node: ParseNode) => void>
     */
    public deserializeFields () : Map<string, (item: Extension, node: ParseNode) => void> {
        return new Map<string, (item: Extension, node: ParseNode) => void>([...super.deserializeFields(),
        ]);
    };
}
