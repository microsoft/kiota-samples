import {SingleValueExtendedPropertiesRequestBuilderGetQueryParameters} from './singleValueExtendedPropertiesRequestBuilderGetQueryParameters';
import {RequestOption} from '@microsoft/kiota-abstractions';

export interface SingleValueExtendedPropertiesRequestBuilderGetRequestConfiguration {
    /** Request headers */
    headers?: Record<string, string> | undefined;
    /** Request options */
    options?: RequestOption[] | undefined;
    /** Request query parameters */
    queryParameters?: SingleValueExtendedPropertiesRequestBuilderGetQueryParameters | undefined;
}
