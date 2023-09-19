import { deserializeIntoEntity, serializeEntity, type Entity } from './entity';
import { createInferenceClassificationOverrideFromDiscriminatorValue, serializeInferenceClassificationOverride, type InferenceClassificationOverride } from './inferenceClassificationOverride';
import { type Parsable, type ParseNode, type SerializationWriter } from '@microsoft/kiota-abstractions';

export function createInferenceClassificationFromDiscriminatorValue(parseNode: ParseNode | undefined) {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return deserializeIntoInferenceClassification;
}
export function deserializeIntoInferenceClassification(inferenceClassification: InferenceClassification | undefined = {} as InferenceClassification) : Record<string, (node: ParseNode) => void> {
    return {
        ...deserializeIntoEntity(inferenceClassification),
        "overrides": n => { inferenceClassification.overrides = n.getCollectionOfObjectValues<InferenceClassificationOverride>(createInferenceClassificationOverrideFromDiscriminatorValue); },
    }
}
export interface InferenceClassification extends Entity, Parsable {
    /**
     * A set of overrides for a user to always classify messages from specific senders in certain ways: focused, or other. Read-only. Nullable.
     */
    overrides?: InferenceClassificationOverride[] | undefined;
}
export function serializeInferenceClassification(writer: SerializationWriter, inferenceClassification: InferenceClassification | undefined = {} as InferenceClassification) : void {
        serializeEntity(writer, inferenceClassification)
        writer.writeCollectionOfObjectValues<InferenceClassificationOverride>("overrides", inferenceClassification.overrides, );
}
