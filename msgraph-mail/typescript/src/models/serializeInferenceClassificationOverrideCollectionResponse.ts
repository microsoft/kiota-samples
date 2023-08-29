import type {InferenceClassificationOverride} from './inferenceClassificationOverride';
import type {InferenceClassificationOverrideCollectionResponse} from './inferenceClassificationOverrideCollectionResponse';
import {serializeInferenceClassificationOverride} from './serializeInferenceClassificationOverride';
import type {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function serializeInferenceClassificationOverrideCollectionResponse(writer: SerializationWriter, inferenceClassificationOverrideCollectionResponse: InferenceClassificationOverrideCollectionResponse | undefined = {} as InferenceClassificationOverrideCollectionResponse) : void {
        writer.writeStringValue("@odata.nextLink", inferenceClassificationOverrideCollectionResponse.odataNextLink);
        writer.writeCollectionOfObjectValues<InferenceClassificationOverride>("value", inferenceClassificationOverrideCollectionResponse.value, serializeInferenceClassificationOverride);
        writer.writeAdditionalData(inferenceClassificationOverrideCollectionResponse.additionalData);
}
