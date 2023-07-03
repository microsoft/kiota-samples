import {deserializeIntoEntity} from './deserializeIntoEntity';
import {SingleValueLegacyExtendedProperty} from './singleValueLegacyExtendedProperty';
import {Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function deserializeIntoSingleValueLegacyExtendedProperty(singleValueLegacyExtendedProperty: SingleValueLegacyExtendedProperty | undefined = {} as SingleValueLegacyExtendedProperty) : Record<string, (node: ParseNode) => void> {
    return {
        ...deserializeIntoEntity(singleValueLegacyExtendedProperty),
        "value": n => { singleValueLegacyExtendedProperty.value = n.getStringValue(); },
    }
}
