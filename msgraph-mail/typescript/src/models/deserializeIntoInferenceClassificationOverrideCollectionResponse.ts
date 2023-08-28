import {createInferenceClassificationOverrideFromDiscriminatorValue} from './createInferenceClassificationOverrideFromDiscriminatorValue';
import type {InferenceClassificationOverride} from './inferenceClassificationOverride';
import type {InferenceClassificationOverrideCollectionResponse} from './inferenceClassificationOverrideCollectionResponse';
import {serializeInferenceClassificationOverride} from './serializeInferenceClassificationOverride';
import type {AdditionalDataHolder, Parsable, ParseNode, SerializationWriter} from '@microsoft/kiota-abstractions';

export function deserializeIntoInferenceClassificationOverrideCollectionResponse(inferenceClassificationOverrideCollectionResponse: InferenceClassificationOverrideCollectionResponse | undefined = {} as InferenceClassificationOverrideCollectionResponse) : Record<string, (node: ParseNode) => void> {
    return {
        "@odata.nextLink": n => { inferenceClassificationOverrideCollectionResponse.odataNextLink = n.getStringValue(); },
        "value": n => { inferenceClassificationOverrideCollectionResponse.value = n.getCollectionOfObjectValues<InferenceClassificationOverride>(createInferenceClassificationOverrideFromDiscriminatorValue); },
    }
}
