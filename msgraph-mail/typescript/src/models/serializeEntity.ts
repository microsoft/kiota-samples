import type {Entity} from './entity';
import type {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function serializeEntity(writer: SerializationWriter, entity: Entity | undefined = {} as Entity) : void {
        writer.writeStringValue("id", entity.id);
        writer.writeAdditionalData(entity.additionalData);
}
