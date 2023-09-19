import { deserializeIntoAttachment } from './attachment';
import { deserializeIntoExtension } from './extension';
import { deserializeIntoInferenceClassification } from './inferenceClassification';
import { deserializeIntoInferenceClassificationOverride } from './inferenceClassificationOverride';
import { deserializeIntoMailFolder } from './mailFolder';
import { deserializeIntoMessage } from './message';
import { deserializeIntoMessageRule } from './messageRule';
import { deserializeIntoMultiValueLegacyExtendedProperty } from './multiValueLegacyExtendedProperty';
import { deserializeIntoOutlookItem } from './outlookItem';
import { deserializeIntoSingleValueLegacyExtendedProperty } from './singleValueLegacyExtendedProperty';
import { type AdditionalDataHolder, type Parsable, type ParseNode, type SerializationWriter } from '@microsoft/kiota-abstractions';

export function createEntityFromDiscriminatorValue(parseNode: ParseNode | undefined) {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return deserializeIntoEntity;
}
export function deserializeIntoEntity(entity: Entity | undefined = {} as Entity) : Record<string, (node: ParseNode) => void> {
    return {
        "id": n => { entity.id = n.getStringValue(); },
    }
}
export interface Entity extends AdditionalDataHolder, Parsable {
    /**
     * Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     */
    additionalData?: Record<string, unknown>;
    /**
     * The unique identifier for an entity. Read-only.
     */
    id?: string | undefined;
}
export function serializeEntity(writer: SerializationWriter, entity: Entity | undefined = {} as Entity) : void {
        writer.writeStringValue("id", entity.id);
        writer.writeAdditionalData(entity.additionalData);
}
