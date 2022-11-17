import {MultiValueLegacyExtendedPropertyItemRequestBuilderGetQueryParameters} from './multiValueLegacyExtendedPropertyItemRequestBuilderGetQueryParameters';
import {RequestOption} from '@microsoft/kiota-abstractions';

<<<<<<< HEAD
export interface MultiValueLegacyExtendedPropertyItemRequestBuilderGetRequestConfiguration {
    /** Request headers */
    headers?: Record<string, string> | undefined;
    /** Request options */
    options?: RequestOption[] | undefined;
    /** Request query parameters */
    queryParameters?: MultiValueLegacyExtendedPropertyItemRequestBuilderGetQueryParameters | undefined;
=======
/** Configuration for the request such as headers, query parameters, and middleware options. */
export class MultiValueLegacyExtendedPropertyItemRequestBuilderGetRequestConfiguration {
    /** Request headers */
    public headers?: Record<string, string> | undefined;
    /** Request options */
    public options?: RequestOption[] | undefined;
    /** Request query parameters */
    public queryParameters?: MultiValueLegacyExtendedPropertyItemRequestBuilderGetQueryParameters | undefined;
>>>>>>> main
}
