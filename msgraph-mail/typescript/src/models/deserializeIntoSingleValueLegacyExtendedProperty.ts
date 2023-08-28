import {deserializeIntoEntity} from './deserializeIntoEntity';
import type {SingleValueLegacyExtendedProperty} from './singleValueLegacyExtendedProperty';
import type {Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function deserializeIntoSingleValueLegacyExtendedProperty(singleValueLegacyExtendedProperty: SingleValueLegacyExtendedProperty | undefined = {} as SingleValueLegacyExtendedProperty) : Record<string, (node: ParseNode) => void> {
    return {
        ...deserializeIntoEntity(singleValueLegacyExtendedProperty),
        "value": n => { singleValueLegacyExtendedProperty.value = n.getStringValue(); },
    }
}
