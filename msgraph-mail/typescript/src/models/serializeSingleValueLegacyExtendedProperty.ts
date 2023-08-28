import {serializeEntity} from './serializeEntity';
import type {SingleValueLegacyExtendedProperty} from './singleValueLegacyExtendedProperty';
import type {Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function serializeSingleValueLegacyExtendedProperty(writer: SerializationWriter, singleValueLegacyExtendedProperty: SingleValueLegacyExtendedProperty | undefined = {} as SingleValueLegacyExtendedProperty) : void {
        serializeEntity(writer, singleValueLegacyExtendedProperty)
        writer.writeStringValue("value", singleValueLegacyExtendedProperty.value);
}
