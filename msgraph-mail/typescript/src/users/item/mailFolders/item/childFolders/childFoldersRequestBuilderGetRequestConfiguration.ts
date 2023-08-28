import {ChildFoldersRequestBuilderGetQueryParameters} from './childFoldersRequestBuilderGetQueryParameters';
import type {RequestOption} from '@microsoft/kiota-abstractions';

export interface ChildFoldersRequestBuilderGetRequestConfiguration {
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
    queryParameters?: ChildFoldersRequestBuilderGetQueryParameters | undefined;
}
