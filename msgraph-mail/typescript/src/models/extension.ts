import { deserializeIntoEntity, serializeEntity, type Entity } from './entity';
import { type Parsable, type ParseNode, type SerializationWriter } from '@microsoft/kiota-abstractions';

export function createExtensionFromDiscriminatorValue(parseNode: ParseNode | undefined) {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return deserializeIntoExtension;
}
export function deserializeIntoExtension(extension: Extension | undefined = {} as Extension) : Record<string, (node: ParseNode) => void> {
    return {
        ...deserializeIntoEntity(extension),
    }
}
export interface Extension extends Entity, Parsable {
}
export function serializeExtension(writer: SerializationWriter, extension: Extension | undefined = {} as Extension) : void {
        serializeEntity(writer, extension)
}
