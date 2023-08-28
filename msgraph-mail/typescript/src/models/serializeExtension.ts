import type {Extension} from './extension';
import {serializeEntity} from './serializeEntity';
import type {Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function serializeExtension(writer: SerializationWriter, extension: Extension | undefined = {} as Extension) : void {
        serializeEntity(writer, extension)
}
