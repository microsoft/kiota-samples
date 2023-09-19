import { deserializeIntoEntity, serializeEntity, type Entity } from './entity';
import { type Parsable, type ParseNode, type SerializationWriter } from '@microsoft/kiota-abstractions';

export function createSingleValueLegacyExtendedPropertyFromDiscriminatorValue(parseNode: ParseNode | undefined) {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return deserializeIntoSingleValueLegacyExtendedProperty;
}
export function deserializeIntoSingleValueLegacyExtendedProperty(singleValueLegacyExtendedProperty: SingleValueLegacyExtendedProperty | undefined = {} as SingleValueLegacyExtendedProperty) : Record<string, (node: ParseNode) => void> {
    return {
        ...deserializeIntoEntity(singleValueLegacyExtendedProperty),
        "value": n => { singleValueLegacyExtendedProperty.value = n.getStringValue(); },
    }
}
export function serializeSingleValueLegacyExtendedProperty(writer: SerializationWriter, singleValueLegacyExtendedProperty: SingleValueLegacyExtendedProperty | undefined = {} as SingleValueLegacyExtendedProperty) : void {
        serializeEntity(writer, singleValueLegacyExtendedProperty)
        writer.writeStringValue("value", singleValueLegacyExtendedProperty.value);
}
export interface SingleValueLegacyExtendedProperty extends Entity, Parsable {
    /**
     * A property value.
     */
    value?: string | undefined;
}
