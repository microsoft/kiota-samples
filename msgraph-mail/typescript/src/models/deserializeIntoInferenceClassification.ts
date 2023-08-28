import {createInferenceClassificationOverrideFromDiscriminatorValue} from './createInferenceClassificationOverrideFromDiscriminatorValue';
import {deserializeIntoEntity} from './deserializeIntoEntity';
import type {InferenceClassification} from './inferenceClassification';
import type {InferenceClassificationOverride} from './inferenceClassificationOverride';
import {serializeInferenceClassificationOverride} from './serializeInferenceClassificationOverride';
import type {Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function deserializeIntoInferenceClassification(inferenceClassification: InferenceClassification | undefined = {} as InferenceClassification) : Record<string, (node: ParseNode) => void> {
    return {
        ...deserializeIntoEntity(inferenceClassification),
        "overrides": n => { inferenceClassification.overrides = n.getCollectionOfObjectValues<InferenceClassificationOverride>(createInferenceClassificationOverrideFromDiscriminatorValue); },
    }
}
