import {SingleValueLegacyExtendedPropertyItemRequestBuilderGetQueryParameters} from './singleValueLegacyExtendedPropertyItemRequestBuilderGetQueryParameters';
import {RequestOption} from '@microsoft/kiota-abstractions';

<<<<<<< HEAD
export interface SingleValueLegacyExtendedPropertyItemRequestBuilderGetRequestConfiguration {
    /** Request headers */
    headers?: Record<string, string> | undefined;
    /** Request options */
    options?: RequestOption[] | undefined;
    /** Request query parameters */
    queryParameters?: SingleValueLegacyExtendedPropertyItemRequestBuilderGetQueryParameters | undefined;
=======
/** Configuration for the request such as headers, query parameters, and middleware options. */
export class SingleValueLegacyExtendedPropertyItemRequestBuilderGetRequestConfiguration {
    /** Request headers */
    public headers?: Record<string, string> | undefined;
    /** Request options */
    public options?: RequestOption[] | undefined;
    /** Request query parameters */
    public queryParameters?: SingleValueLegacyExtendedPropertyItemRequestBuilderGetQueryParameters | undefined;
>>>>>>> main
}
