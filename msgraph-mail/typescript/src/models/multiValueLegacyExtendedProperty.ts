import { deserializeIntoEntity, serializeEntity, type Entity } from './entity';
import { type Parsable, type ParseNode, type SerializationWriter } from '@microsoft/kiota-abstractions';

export function createMultiValueLegacyExtendedPropertyFromDiscriminatorValue(parseNode: ParseNode | undefined) {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return deserializeIntoMultiValueLegacyExtendedProperty;
}
export function deserializeIntoMultiValueLegacyExtendedProperty(multiValueLegacyExtendedProperty: MultiValueLegacyExtendedProperty | undefined = {} as MultiValueLegacyExtendedProperty) : Record<string, (node: ParseNode) => void> {
    return {
        ...deserializeIntoEntity(multiValueLegacyExtendedProperty),
        "value": n => { multiValueLegacyExtendedProperty.value = n.getCollectionOfPrimitiveValues<string>(); },
    }
}
export interface MultiValueLegacyExtendedProperty extends Entity, Parsable {
    /**
     * A collection of property values.
     */
    value?: string[] | undefined;
}
export function serializeMultiValueLegacyExtendedProperty(writer: SerializationWriter, multiValueLegacyExtendedProperty: MultiValueLegacyExtendedProperty | undefined = {} as MultiValueLegacyExtendedProperty) : void {
        serializeEntity(writer, multiValueLegacyExtendedProperty)
        writer.writeCollectionOfPrimitiveValues<string>("value", multiValueLegacyExtendedProperty.value);
}
