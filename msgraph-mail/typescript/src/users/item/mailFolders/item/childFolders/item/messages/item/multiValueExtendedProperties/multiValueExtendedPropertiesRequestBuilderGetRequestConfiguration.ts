import {MultiValueExtendedPropertiesRequestBuilderGetQueryParameters} from './multiValueExtendedPropertiesRequestBuilderGetQueryParameters';
import {RequestOption} from '@microsoft/kiota-abstractions';

<<<<<<< HEAD
export interface MultiValueExtendedPropertiesRequestBuilderGetRequestConfiguration {
    /** Request headers */
    headers?: Record<string, string> | undefined;
    /** Request options */
    options?: RequestOption[] | undefined;
    /** Request query parameters */
    queryParameters?: MultiValueExtendedPropertiesRequestBuilderGetQueryParameters | undefined;
=======
/** Configuration for the request such as headers, query parameters, and middleware options. */
export class MultiValueExtendedPropertiesRequestBuilderGetRequestConfiguration {
    /** Request headers */
    public headers?: Record<string, string> | undefined;
    /** Request options */
    public options?: RequestOption[] | undefined;
    /** Request query parameters */
    public queryParameters?: MultiValueExtendedPropertiesRequestBuilderGetQueryParameters | undefined;
>>>>>>> main
}
