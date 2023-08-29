import {InferenceClassificationRequestBuilderGetQueryParameters} from './inferenceClassificationRequestBuilderGetQueryParameters';
import type {RequestOption} from '@microsoft/kiota-abstractions';

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
