import {CountRequestBuilderGetQueryParameters} from './countRequestBuilderGetQueryParameters';
import type {RequestOption} from '@microsoft/kiota-abstractions';

export interface CountRequestBuilderGetRequestConfiguration {
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
    queryParameters?: CountRequestBuilderGetQueryParameters | undefined;
}
