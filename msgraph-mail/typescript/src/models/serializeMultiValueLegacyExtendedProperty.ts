import type {MultiValueLegacyExtendedProperty} from './multiValueLegacyExtendedProperty';
import {serializeEntity} from './serializeEntity';
import type {Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function serializeMultiValueLegacyExtendedProperty(writer: SerializationWriter, multiValueLegacyExtendedProperty: MultiValueLegacyExtendedProperty | undefined = {} as MultiValueLegacyExtendedProperty) : void {
        serializeEntity(writer, multiValueLegacyExtendedProperty)
        writer.writeCollectionOfPrimitiveValues<string>("value", multiValueLegacyExtendedProperty.value);
}
