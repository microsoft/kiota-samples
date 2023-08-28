import type {InferenceClassification} from './inferenceClassification';
import type {InferenceClassificationOverride} from './inferenceClassificationOverride';
import {serializeEntity} from './serializeEntity';
import {serializeInferenceClassificationOverride} from './serializeInferenceClassificationOverride';
import type {Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function serializeInferenceClassification(writer: SerializationWriter, inferenceClassification: InferenceClassification | undefined = {} as InferenceClassification) : void {
        serializeEntity(writer, inferenceClassification)
        writer.writeCollectionOfObjectValues<InferenceClassificationOverride>("overrides", inferenceClassification.overrides, serializeInferenceClassificationOverride);
}
