import { createInferenceClassificationOverrideFromDiscriminatorValue, serializeInferenceClassificationOverride, type InferenceClassificationOverride } from './inferenceClassificationOverride';
import { type AdditionalDataHolder, type Parsable, type ParseNode, type SerializationWriter } from '@microsoft/kiota-abstractions';

export function createInferenceClassificationOverrideCollectionResponseFromDiscriminatorValue(parseNode: ParseNode | undefined) {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return deserializeIntoInferenceClassificationOverrideCollectionResponse;
}
export function deserializeIntoInferenceClassificationOverrideCollectionResponse(inferenceClassificationOverrideCollectionResponse: InferenceClassificationOverrideCollectionResponse | undefined = {} as InferenceClassificationOverrideCollectionResponse) : Record<string, (node: ParseNode) => void> {
    return {
        "@odata.nextLink": n => { inferenceClassificationOverrideCollectionResponse.odataNextLink = n.getStringValue(); },
        "value": n => { inferenceClassificationOverrideCollectionResponse.value = n.getCollectionOfObjectValues<InferenceClassificationOverride>(createInferenceClassificationOverrideFromDiscriminatorValue); },
    }
}
export interface InferenceClassificationOverrideCollectionResponse extends AdditionalDataHolder, Parsable {
    /**
     * Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     */
    additionalData?: Record<string, unknown>;
    /**
     * The OdataNextLink property
     */
    odataNextLink?: string | undefined;
    /**
     * The value property
     */
    value?: InferenceClassificationOverride[] | undefined;
}
export function serializeInferenceClassificationOverrideCollectionResponse(writer: SerializationWriter, inferenceClassificationOverrideCollectionResponse: InferenceClassificationOverrideCollectionResponse | undefined = {} as InferenceClassificationOverrideCollectionResponse) : void {
        writer.writeStringValue("@odata.nextLink", inferenceClassificationOverrideCollectionResponse.odataNextLink);
        writer.writeCollectionOfObjectValues<InferenceClassificationOverride>("value", inferenceClassificationOverrideCollectionResponse.value, );
        writer.writeAdditionalData(inferenceClassificationOverrideCollectionResponse.additionalData);
}
