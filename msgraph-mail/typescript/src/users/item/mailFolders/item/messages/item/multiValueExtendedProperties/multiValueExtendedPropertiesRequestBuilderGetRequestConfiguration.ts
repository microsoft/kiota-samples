import {MultiValueExtendedPropertiesRequestBuilderGetQueryParameters} from './multiValueExtendedPropertiesRequestBuilderGetQueryParameters';
import {RequestOption} from '@microsoft/kiota-abstractions';

export interface MultiValueExtendedPropertiesRequestBuilderGetRequestConfiguration {
    /** Request headers */
    headers?: Record<string, string> | undefined;
    /** Request options */
    options?: RequestOption[] | undefined;
    /** Request query parameters */
    queryParameters?: MultiValueExtendedPropertiesRequestBuilderGetQueryParameters | undefined;
}
