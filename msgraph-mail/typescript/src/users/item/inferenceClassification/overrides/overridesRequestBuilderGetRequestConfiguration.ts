import {OverridesRequestBuilderGetQueryParameters} from './overridesRequestBuilderGetQueryParameters';
import type {RequestOption} from '@microsoft/kiota-abstractions';

export interface OverridesRequestBuilderGetRequestConfiguration {
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
    queryParameters?: OverridesRequestBuilderGetQueryParameters | undefined;
}
