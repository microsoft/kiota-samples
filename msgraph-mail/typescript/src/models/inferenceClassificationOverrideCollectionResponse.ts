import type {InferenceClassificationOverride} from './inferenceClassificationOverride';
import type {AdditionalDataHolder, Parsable} from '@microsoft/kiota-abstractions';

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
