import {ExtensionsRequestBuilderGetQueryParameters} from './extensionsRequestBuilderGetQueryParameters';
import {RequestOption} from '@microsoft/kiota-abstractions';

<<<<<<< HEAD
export interface ExtensionsRequestBuilderGetRequestConfiguration {
    /** Request headers */
    headers?: Record<string, string> | undefined;
    /** Request options */
    options?: RequestOption[] | undefined;
    /** Request query parameters */
    queryParameters?: ExtensionsRequestBuilderGetQueryParameters | undefined;
=======
/** Configuration for the request such as headers, query parameters, and middleware options. */
export class ExtensionsRequestBuilderGetRequestConfiguration {
    /** Request headers */
    public headers?: Record<string, string> | undefined;
    /** Request options */
    public options?: RequestOption[] | undefined;
    /** Request query parameters */
    public queryParameters?: ExtensionsRequestBuilderGetQueryParameters | undefined;
>>>>>>> main
}
