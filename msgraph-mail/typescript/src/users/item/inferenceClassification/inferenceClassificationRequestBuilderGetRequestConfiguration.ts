import {InferenceClassificationRequestBuilderGetQueryParameters} from './inferenceClassificationRequestBuilderGetQueryParameters';
import {RequestOption} from '@microsoft/kiota-abstractions';

export interface InferenceClassificationRequestBuilderGetRequestConfiguration {
    /**
     * Request headers
     */
    headers?: Record<string, string[]> | undefined;
    /**
     * Request options
     */
    options?: RequestOption[] | undefined;
    /**
     * Request query parameters
     */
    queryParameters?: InferenceClassificationRequestBuilderGetQueryParameters | undefined;
}
