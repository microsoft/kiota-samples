import type {Entity} from './entity';
import type {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function deserializeIntoEntity(entity: Entity | undefined = {} as Entity) : Record<string, (node: ParseNode) => void> {
    return {
        "id": n => { entity.id = n.getStringValue(); },
    }
}
